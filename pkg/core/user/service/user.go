package service

import (
	"go.uber.org/zap"

	"go-sample-site/pkg/core/user/repository/models"
	"go-sample-site/pkg/core/user/repository/orm"
)

type User struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Alias            string `json:"alias"`
}

func ListUsers(logger *zap.SugaredLogger) ([]*User, error) {
	users, err := orm.NewUserColl().List()
	if err != nil {
		logger.Errorf("Failed to list users, err: %s", err)
		return nil, err
	}

	var res []*User
	for _, u := range users {
		res = append(res, &User{
			ID:     u.ID,
			Name:   u.Name,
			Alias: u.Alias,
		})
	}

	return res, nil
}

func GetUser(id string, logger *zap.SugaredLogger) (*User, error) {
	u, err := orm.NewUserColl().Get(id)
	if err != nil {
		logger.Errorf("Failed to get user %s, err: %s", id, err)
		return nil, err
	}

	return &User{
		ID:     u.ID,
		Name:   u.Name,
		Alias: u.Alias,
	}, nil

}

func DeleteUser(id string, _ *zap.SugaredLogger) error {
	return orm.NewUserColl().Delete(id)
}

func CreateUser(u *User, _ *zap.SugaredLogger) error {
	obj := &models.User{
		ID:     u.ID,
		Name:   u.Name,
		Alias: u.Alias,
	}

	return orm.NewUserColl().Create(obj)
}
