package di

import (
	"go.uber.org/dig"
	"gorm.io/gorm"
	"log"
	"sakshyahere/tuko/internal/controller"
	authController "sakshyahere/tuko/internal/controller/auth"
	"sakshyahere/tuko/internal/db"
	"sakshyahere/tuko/internal/repository"
	authRepository "sakshyahere/tuko/internal/repository/auth"
	"sakshyahere/tuko/internal/service"
	authService "sakshyahere/tuko/internal/service/auth"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	//DB
	container.Provide(db.ConnectDb)
	container.Provide(db.NewMigrationService)

	// Repositories
	container.Provide(repository.NewUserRepository)
	container.Provide(authRepository.NewAuthRepository)

	// Services
	container.Provide(service.NewUserService)
	container.Provide(authService.NewAuthService)

	// Controllers
	container.Provide(controller.NewUserController)
	container.Provide(authController.NewAuthController)

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
