package mod

import (
	"fmt"
	"gihora/pkg/util"
	"log"
	"regexp"
	"slices"
	"strings"
	"sync"
)

var modReg = regexp.MustCompile(`^ServerModSetup\("([^"]+)"\)\s*(?:--\s*(.*))?$`)

type Manager struct {
	ModList []Mod
	cfg     Config
	mu      sync.RWMutex
}

func NewManager() (*Manager, error) {
	m := &Manager{}

	if err := m.init(); err != nil {
		log.Println("[ERROR] 新建模组管理器失败")
		return nil, err
	}

	return m, nil
}

func (m *Manager) init() error {
	var err error
	if m.cfg, err = InitConfig(); err != nil {
		log.Println("[ERROR] 初始化模组配置失败")
		return err
	}

	lineSlice, err := util.ReadFileByLine(m.cfg.ModFilePath)
	if err != nil {
		log.Println("[ERROR] 初始化模组管理器失败")
		return err
	}

	for _, line := range lineSlice {
		parts := modReg.FindStringSubmatch(line)
		if parts != nil {
			id := parts[1]
			remark := parts[2]
			m.ModList = append(m.ModList, NewMod(id, remark))
		}
	}

	return nil
}

func (m *Manager) save() error {
	sb := strings.Builder{}
	for _, mod := range m.ModList {
		fmt.Fprintf(&sb, "ServerModSetup(\"%s\") -- %s\n", mod.Id, mod.Remark)
	}

	if err := util.WriteToFile(m.cfg.ModFilePath, sb.String()); err != nil {
		log.Println("[ERROR] 保存模组列表失败")
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

func (m *Manager) ListMods() []Mod {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return slices.Clone(m.ModList)
}

func (m *Manager) SubMod(id, remark string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if hasMod := m.hasMod(id); hasMod {
		return fmt.Errorf("模组 %s 已订阅", id)
	}

	m.ModList = append(m.ModList, NewMod(id, remark))

	if err := m.save(); err != nil {
		log.Printf("[ERROR] 模组 %s 订阅失败\n", id)
		return err
	}

	return nil
}

func (m *Manager) UnsubMod(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if hasMod := m.hasMod(id); !hasMod {
		return fmt.Errorf("模组 %s 未订阅", id)
	}

	for i, mod := range m.ModList {
		if mod.Id == id {
			m.ModList = append(m.ModList[:i], m.ModList[i+1:]...)
		}
	}

	if err := m.save(); err != nil {
		log.Printf("[ERROR] 模组 %s 取消订阅失败\n", id)
		return err
	}

	return nil
}

func (m *Manager) UpdateModRemark(id, remark string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if hasMod := m.hasMod(id); !hasMod {
		return fmt.Errorf("模组 %s 未订阅", id)
	}

	for i := range m.ModList {
		if m.ModList[i].Id == id {
			m.ModList[i].Remark = remark
		}
	}

	if err := m.save(); err != nil {
		log.Printf("[ERROR] 模组 %s 修改备注失败\n", id)
		return err
	}

	return nil
}
