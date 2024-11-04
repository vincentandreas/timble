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

func (repo *Implementation) FindUserById(id uint64, ctx context.Context) (*models.User, error) {
	var user *models.User
	err := repo.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
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

func (repo *Implementation) UpdatePremium(id uint64, ctx context.Context) error{
	var user models.User
    err := repo.db.WithContext(ctx).First(&user, id).Error
    if err!= nil {
        return err
    }
	now := time.Now()
    user.PremiumAccount = true
	user.PremiumExpired =  now.AddDate(0, 1, 0) //todo change to more custom
    err = repo.db.Save(&user).Error
	return err
}