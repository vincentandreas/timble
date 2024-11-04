package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/vincentandreas/dealls/models"
	"github.com/vincentandreas/dealls/utilities"
	"gorm.io/gorm"
)

func (repo *Implementation) SaveHistory(userId uint64, feel models.Feeling, ctx context.Context) error {
	hist := &models.History{
		UserID: userId,
        WatchedUserID: feel.WatchedUserID,
        FeelMatch: feel.FeelMatch,
	}

	if userId == feel.WatchedUserID {
		return errors.New("user cant value themselves")
	}

	// check existing history
	histWatchedIds, err := repo.SeenUserIds(userId, ctx)
	utilities.CheckError(err)

	for _, id := range histWatchedIds {
		if feel.WatchedUserID == id {
			return errors.New("you already decide for this user")
		}
	}


	err = repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		today := time.Now() 
		
		var user models.User
		err := tx.First(&user, userId).Error
		if err!= nil {
            return err
        }

		curQuota := user.Quota
		curQuotaDate := user.QuotaLatestUpdate
		if utilities.DateBefore(user.PremiumExpired,today) {
			user.PremiumAccount = false
		}

		if !user.PremiumAccount{
			if utilities.DateEqual(curQuotaDate, today) {
				if curQuota <= 0 {
					return errors.New("noquota")
				}
				curQuota--
			}else{
				curQuota = 9
				curQuotaDate = today
			}
			user.QuotaLatestUpdate = curQuotaDate
			user.Quota = curQuota

			err = tx.Save(user).Error
			if err!= nil {
				return err
			}
		}

		err = tx.Create(&hist).Error
		if err!= nil {
            return err
        }
		return nil
	})
	return err
}

func (repo *Implementation) GetInterestWithUser(userId uint64, ctx context.Context) ([]models.Recommendation, error) {
	var hists []models.History
	
	err := repo.db.Preload("User").Preload("User.UserPreference").Where("watched_user_id = ?", userId).Where("feel_match = true").Find(&hists)
	recoms := make([]models.Recommendation,0)
	for _, data := range hists {
		recoms = append(recoms, models.Recommendation{
			UserId: data.UserID,
			Hobby: data.User.UserPreference.Hobby,
			Name: data.User.Name,
		})
	}
	return recoms, err.Error
}

func (repo *Implementation) SeenUserIds(userId uint64, ctx context.Context) ([]uint64, error) {

	var watchedUserIDs []uint64
	currentDate := time.Now().Format("2006-01-02")

    err := repo.db.Model(&models.History{}).Where("user_id = ?", userId).Where("DATE(created_at) = ?", currentDate).Pluck("watched_user_id", &watchedUserIDs).Error
	fmt.Println(watchedUserIDs)
	return watchedUserIDs, err
}