package db

import (
	"fmt"
	"gorm.io/gorm"
	"sakshyahere/tuko/internal/model"
)

type MigrationService struct {
	db *gorm.DB
}

func NewMigrationService(db *gorm.DB) *MigrationService {
	return &MigrationService{
		db: db,
	}
}

func (s *MigrationService) Migrate() error {
	err := s.db.AutoMigrate(&model.User{})
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	fmt.Println("Migration Success")
	return nil
}
