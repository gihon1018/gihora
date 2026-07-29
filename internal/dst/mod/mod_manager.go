package mod

import (
	"fmt"
	"gihora/internal/dst/mod/model"
	"gihora/pkg/apperr"
	"gihora/pkg/util"
	"log"
	"regexp"
	"slices"
	"strings"
	"sync"
)

var modReg = regexp.MustCompile(`^ServerModSetup\("([^"]+)"\)\s*(?:--\s*(.*))?$`)

type Manager struct {
	ModList []model.Mod
	cfg     Config
	mu      sync.RWMutex
}

func NewManager() (*Manager, error) {
	m := &Manager{}

	if err := m.init(); err != nil {
		log.Printf("[ERROR] 模组管理器新建失败: %v\n", err)
		return nil, err
	}

	return m, nil
}

func (m *Manager) init() error {
	var err error
	if m.cfg, err = InitConfig(); err != nil {
		log.Printf("[ERROR] 模组管理器配置初始化失败: %v\n", err)
		return err
	}

	lineSlice, err := util.ReadFileByLine(m.cfg.ModsSetupPath)
	if err != nil {
		log.Printf("[ERROR] 模组管理器初始化失败: %v\n", err)
		return err
	}

	for _, line := range lineSlice {
		parts := modReg.FindStringSubmatch(line)
		if parts != nil {
			id := parts[1]
			remark := parts[2]
			m.ModList = append(m.ModList, model.NewMod(id, remark))
		}
	}

	return nil
}

func (m *Manager) save() error {
	sb := strings.Builder{}
	for _, mod := range m.ModList {
		if _, err := fmt.Fprintf(&sb, "ServerModSetup(\"%s\") -- %s\n", mod.Id, mod.Remark); err != nil {
			log.Printf("[ERROR] 模组列表遍历失败: %v\n", err)
			return err
		}
	}

	if err := util.WriteToFile(m.cfg.ModsSetupPath, sb.String()); err != nil {
		log.Printf("[ERROR] 模组列表保存失败: %v\n", err)
		return err
	}

	return nil
}

func (m *Manager) hasMod(id string) bool {
	for _, mod := range m.ModList {
		if mod.Id == id {
			return true
		}
	}

	return false
}

func (m *Manager) ListMods() []model.Mod {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return slices.Clone(m.ModList)
}

func (m *Manager) SubMod(id, remark string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if hasMod := m.hasMod(id); hasMod {
		return apperr.ModAlreadySubbed.WithMsg(fmt.Sprintf("模组 %s 已订阅", id))
	}

	m.ModList = append(m.ModList, model.NewMod(id, remark))

	if err := m.save(); err != nil {
		log.Printf("[ERROR] 模组 %s 订阅失败: %v\n", id, err)
		return apperr.ModSubFail
	}

	return nil
}

func (m *Manager) UnsubMod(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if hasMod := m.hasMod(id); !hasMod {
		return apperr.ModNotSubbed.WithMsg(fmt.Sprintf("模组 %s 未订阅", id))
	}

	for i, mod := range m.ModList {
		if mod.Id == id {
			m.ModList = append(m.ModList[:i], m.ModList[i+1:]...)
		}
	}

	if err := m.save(); err != nil {
		log.Printf("[ERROR] 模组 %s 取消订阅失败: %v\n", id, err)
		return apperr.ModUnsubFail
	}

	return nil
}

func (m *Manager) UpdateModRemark(id, remark string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if hasMod := m.hasMod(id); !hasMod {
		return apperr.ModNotSubbed.WithMsg(fmt.Sprintf("模组 %s 未订阅", id))
	}

	for i := range m.ModList {
		if m.ModList[i].Id == id {
			m.ModList[i].Remark = remark
		}
	}

	if err := m.save(); err != nil {
		log.Printf("[ERROR] 模组 %s 修改备注失败: %v\n", id, err)
		return apperr.ModUpdateRemarkFail
	}

	return nil
}
