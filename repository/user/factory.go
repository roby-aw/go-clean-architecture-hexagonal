package user

import (
	"github.com/roby-aw/go-clean-architecture-hexagonal/business/user"
	"github.com/roby-aw/go-clean-architecture-hexagonal/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) user.Repository {
	adminRepo := NewMongoRepository(dbCon.MongoDB)
	return adminRepo
}
