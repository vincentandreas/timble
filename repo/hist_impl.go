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

		err = tx.Create(&hist).Error
		if err!= nil {
            return err
        }
		return nil
	})
	return err
}

func (repo *Implementation) SeenUserIds(userId uint64, ctx context.Context) ([]uint64, error) {

	var watchedUserIDs []uint64

    // Query for the watched user IDs
    err := repo.db.Model(&models.History{}).Where("user_id = ?", userId).Pluck("watched_user_id", &watchedUserIDs).Error
	fmt.Println(watchedUserIDs)
	return watchedUserIDs, err
}