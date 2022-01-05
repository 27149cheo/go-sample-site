package orm

import (
	"context"
	_ "embed"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"go-sample-site/pkg/core/user/repository/models"
	gormtool "go-sample-site/pkg/tool/gorm"
)

//go:embed user.sql
var userSQL string

type UserColl struct {
	*gorm.DB

	coll string
}

func NewUserColl() *UserColl {
	name := models.User{}.TableName()
	return &UserColl{
		DB:   gormtool.DB("loutest"),
		coll: name,
	}
}

func (c *UserColl) GetCollectionName() string {
	return c.coll
}

func (c *UserColl) EnsureTable(_ context.Context) error {
	tx := c.Exec(userSQL)
	return tx.Error
}

func (c *UserColl) List() ([]*models.User, error) {
	var res []*models.User
	result := c.Find(&res)

	return res, result.Error
}

func (c *UserColl) Get(id string) (*models.User, error) {
	res := &models.User{}
	result := c.First(&res, "id=?", id)

	return res, result.Error
}

func (c *UserColl) Create(obj *models.User) error {
	u := uuid.New()
	obj.ID = u.String()

	result := c.DB.Create(obj)

	return result.Error
}

func (c *UserColl) Delete(id string) error {
	result := c.DB.Delete(&models.User{ID: id})

	return result.Error
}
