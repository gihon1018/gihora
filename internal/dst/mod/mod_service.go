package mod

type Service struct {
	manager *Manager
}

func NewService(manager *Manager) *Service {
	s := &Service{
		manager: manager,
	}
	return s
}

func (s *Service) ListMods() []Mod {
	return s.manager.ListMods()
}

func (s *Service) SubMod(id, remark string) error {
	if err := s.manager.SubMod(id, remark); err != nil {
		return err
	}

	return nil
}

func (s *Service) UnsubMod(id string) error {
	if err := s.manager.UnsubMod(id); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateModRemark(id, remark string) error {
	if err := s.manager.UpdateModRemark(id, remark); err != nil {
		return err
	}

	return nil
}
