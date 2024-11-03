package repo

import (
	"context"
	"time"

	"github.com/vincentandreas/dealls/models"
	"github.com/vincentandreas/dealls/utilities"
)

func (repo *Implementation) SaveUser(user models.UserRegister, ctx context.Context) error {
	var savedUser models.User
	savedUser.Password = utilities.HashParams(user.Password)
	savedUser.Quota = 10 //todo use redis
	savedUser.QuotaLatestUpdate = time.Now()
	savedUser.Email = user.Email
	savedUser.Name = user.Name
	savedUser.Gender = user.Gender

	savedUser.UserPreference = &models.UserPreference{
		InterestGender:  user.InterestGender,
		Hobby:  user.Hobby,
	}
	err := repo.db.WithContext(ctx).Create(&savedUser).Error
	return err
}

func (repo *Implementation) FindUserById(id uint, ctx context.Context) (*models.User, error) {
	var user *models.User
	selSql := "full_name, user_name, phone_number, role"
	err := repo.db.WithContext(ctx).Select(selSql).Where("id = ?", id).First(&user).Error
	return user, err
}

func (repo *Implementation) Login(req models.UserLogin, ctx context.Context) (*models.User, error) {
	var user *models.User
	hashPasswd := utilities.HashParams(req.Password)
	err := repo.db.WithContext(ctx).Where("email = ? AND password = ?", req.Email, hashPasswd).First(&user).Error

	if err != nil || user == nil {
		return nil, err
	}

	return user, nil
}


