package modules

import (
	"github.com/roby-aw/go-clean-architecture-hexagonal/api"
	userController "github.com/roby-aw/go-clean-architecture-hexagonal/api/user"
	userBusiness "github.com/roby-aw/go-clean-architecture-hexagonal/business/user"
	"github.com/roby-aw/go-clean-architecture-hexagonal/config"
	userRepository "github.com/roby-aw/go-clean-architecture-hexagonal/repository/user"
	"github.com/roby-aw/go-clean-architecture-hexagonal/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, _ *config.AppConfig) api.Controller {
	userPermitRepository := userRepository.RepositoryFactory(dbCon)
	userPermitService := userBusiness.NewService(userPermitRepository)
	userPermitController := userController.NewController(userPermitService)
	// Register controller
	controller := api.Controller{
		UserController: userPermitController,
	}

	return controller
}
