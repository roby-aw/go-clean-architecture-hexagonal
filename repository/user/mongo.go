package user

import (
	"errors"
	"time"

	businessUser "github.com/roby-aw/go-clean-architecture-hexagonal/business/user"
	"github.com/roby-aw/go-clean-architecture-hexagonal/repository"
	"github.com/roby-aw/go-clean-architecture-hexagonal/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type MongoDBRepository struct {
	colUser *mongo.Collection
}

func NewMongoRepository(col *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		colUser: col.Collection("user"),
	}
}

func (repo *MongoDBRepository) FindUserByEmail(email string) (businessUser.User, error) {
	var user repository.User
	var userBusiness businessUser.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	queryFilter := repository.NewFilterQuery()
	queryFilter.SetEmail(email)

	err := repo.colUser.FindOne(ctx, queryFilter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return userBusiness, errors.New("wrong email")
		}
		return userBusiness, err
	}
	userBusiness.ID = user.ID.Hex()
	userBusiness.Email = user.Email
	userBusiness.Password = user.Password

	return userBusiness, nil
}

func (repo *MongoDBRepository) CreateUser(data businessUser.Register) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	passwd, err := utils.Hash(data.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}
	_, err = repo.colUser.InsertOne(ctx, repository.RegisterUser{ID: primitive.NewObjectID(), Email: data.Email, Password: string(passwd), CreatedAt: time.Now()})
	if err != nil {
		return err
	}
	return nil
}

func (repo *MongoDBRepository) FindUserByID(id string) (businessUser.User, error) {
	var user repository.User
	var userBusiness businessUser.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return userBusiness, errors.New("invalid id")
	}

	queryFilter := repository.NewFilterQuery()
	queryFilter.SetID(objID)

	err = repo.colUser.FindOne(ctx, queryFilter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return userBusiness, errors.New("wrong id")
		}
		return userBusiness, err
	}
	userBusiness.ID = user.ID.Hex()
	userBusiness.Email = user.Email
	userBusiness.Password = user.Password
	return userBusiness, nil
}
