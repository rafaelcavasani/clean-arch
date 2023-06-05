package repository

import (
	"clean-arch/core/domain"
	"clean-arch/core/repository"
	"clean-arch/infrastructure/logger"
	"context"
)

type UserEntity struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepositoryDynamoDB struct {
	db NoSQL
}

var _ repository.UserRepository = (*UserRepositoryDynamoDB)(nil)

func NewUserRepositoryDynamoDB(db NoSQL) UserRepositoryDynamoDB {
	return UserRepositoryDynamoDB{
		db: db,
	}
}

func (repository UserRepositoryDynamoDB) FindById(ctx context.Context, userId string) (domain.User, error) {
	logger.Infof("M=FindById, stage=init, userId=%s", userId)
	out, err := repository.db.FindById(ctx, userId)
	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{Id: out.Id, Name: out.Name, Email: out.Email}
	if err != nil {
		return domain.User{}, err
	}
	logger.Infof("M=FindById, stage=finish, user=%s", user)
	return user, nil
}

func (repository UserRepositoryDynamoDB) PutItem(ctx context.Context, item domain.User) (domain.User, error) {
	logger.Infof("M=PutItem, stage=init, user=%s", item)
	userEntity := UserEntity{Id: item.Id, Name: item.Name, Email: item.Email}
	out, err := repository.db.PutItem(ctx, userEntity)
	if err != nil {
		return domain.User{}, err
	}

	user, err := domain.NewUser(out.Id, out.Name, out.Email)
	if err != nil {
		return domain.User{}, err
	}
	logger.Infof("M=PutItem, stage=finish, user=%s", user)
	return user, nil
}

func (repository UserRepositoryDynamoDB) UpdateItem(ctx context.Context, item domain.User) (domain.User, error) {
	logger.Infof("M=UpdateItem, stage=init, user=%s", item)
	userEntity := UserEntity{Id: item.Id, Name: item.Name, Email: item.Email}
	out, err := repository.db.UpdateItem(ctx, userEntity)
	if err != nil {
		return domain.User{}, err
	}

	user, err := domain.NewUser(out.Id, out.Name, out.Email)
	if err != nil {
		return domain.User{}, err
	}
	logger.Infof("M=UpdateItem, stage=finish, user=%s", user)
	return user, nil
}

func (repository UserRepositoryDynamoDB) DeleteItem(ctx context.Context, id string) error {
	logger.Infof("M=DeleteItem, stage=init, id=%s", id)
	err := repository.db.DeleteItem(ctx, id)
	if err != nil {
		return err
	}
	logger.Infof("M=DeleteItem, stage=finish, id=%s", id)
	return nil
}
