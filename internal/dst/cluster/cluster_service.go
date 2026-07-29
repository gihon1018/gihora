package cluster

import "gihora/internal/dst/cluster/model"

type Service struct {
	manager *Manager
}

func NewService(manager *Manager) *Service {
	s := &Service{
		manager: manager,
	}
	return s
}

func (s *Service) ListClusters() []model.Cluster {
	return s.manager.ListClusters()
}
