package mod

import (
	"fmt"
	"gihora/pkg/util"
	"path/filepath"
	"regexp"
	"sync"
)

type Manager struct {
	ModList []Mod
}

var (
	manager *Manager
	once    sync.Once
	modReg  = regexp.MustCompile(`^ServerModSetup\("([^"]+)"\)\s*(?:--\s*(.*))?$`)
	initErr error
)

func GetManager() (*Manager, error) {
	once.Do(func() {
		manager = &Manager{}
		initErr = initManager(manager)
	})
	if initErr != nil {
		return nil, initErr
	}
	return manager, nil
}

func initManager(manager *Manager) error {
	join := filepath.Join("test", "mods.lua")
	lineSlice, err := util.ReadFileByLine(join)
	if err != nil {
		return fmt.Errorf("初始化模组管理器失败: %w", err)
	}
	for _, line := range lineSlice {
		parts := modReg.FindStringSubmatch(line)
		if parts != nil {
			manager.ModList = append(manager.ModList, Mod{
				Id:     parts[1],
				Remark: parts[2],
			})
		}
	}
	return nil
}
