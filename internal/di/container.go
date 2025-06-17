package di

import (
	"go.uber.org/dig"
	"gorm.io/gorm"
	"log"
	"sakshyahere/tuko/internal/controller"
	"sakshyahere/tuko/internal/db"
	"sakshyahere/tuko/internal/repository"
	"sakshyahere/tuko/internal/service"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	//DB
	container.Provide(db.ConnectDb)
	container.Provide(db.NewMigrationService)

	// Repositories
	container.Provide(repository.NewUserRepository)

	// Services
	container.Provide(service.NewUserService)

	// Controllers
	container.Provide(controller.NewUserController)

	return container
}

func InitializeApp(container *dig.Container) error {
	return container.Invoke(func(database *gorm.DB, migrationService *db.MigrationService) {
		if err := migrationService.Migrate(); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		log.Println("Database and migrations initialized successfully")
	})
}
