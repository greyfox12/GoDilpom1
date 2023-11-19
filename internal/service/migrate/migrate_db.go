package migrate

import (
	"fmt"
)

func (s *Service) Execute() error {

	err := s.myRepo.MigrateSchema()

	if err != nil {
		return fmt.Errorf("execute migrsate requers: %w", err) // внутренняя ошибка сервера
	}

	return nil

}
