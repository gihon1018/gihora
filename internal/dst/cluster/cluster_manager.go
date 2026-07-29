package cluster

import (
	"gihora/internal/dst/cluster/model"
	"gihora/pkg/util"
	"log"
	"sync"
)

type Manager struct {
	clusterList []*model.Cluster
	cfg         Config
	mu          sync.RWMutex
}

func NewManager() (*Manager, error) {
	m := &Manager{}

	if err := m.init(); err != nil {
		log.Println("[ERROR] 集群管理器初始化失败")
		return nil, err
	}

	return m, nil
}

func (m *Manager) init() error {
	var err error
	if m.cfg, err = InitConfig(); err != nil {
		log.Println("[ERROR] 集群管理器配置初始化失败")
		return err
	}

	subDirPaths, err := util.ListSubDirPaths(m.cfg.SaveDirPath)
	if err != nil {
		log.Println("[ERROR] 获取集群子目录失败")
		return err
	}

	for _, subDirPath := range subDirPaths {
		if !model.IsValidClusterDirPath(subDirPath) {
			continue
		}
		cluster, err := model.NewCluster(subDirPath)
		if err != nil {
			log.Println("[ERROR] 新建集群对象失败")
		}
		m.clusterList = append(m.clusterList, cluster)
	}

	return nil
}

func (m *Manager) ListClusters() []model.Cluster {
	m.mu.RLock()
	defer m.mu.RUnlock()

	res := make([]model.Cluster, 0, len(m.clusterList))
	for _, ptr := range m.clusterList {
		if ptr == nil {
			continue
		}
		res = append(res, *ptr)
	}
	return res
}
