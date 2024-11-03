package repo

import (
	"context"
	"errors"
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

	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		today := time.Now() 
		
		var user models.User
		err := tx.First(&user, userId).Error
		if err!= nil {
            return err
        }

		curQuota := user.Quota
		curQuotaDate := user.QuotaLatestUpdate

		if utilities.DateEqual(curQuotaDate, today) {
			if curQuota == 0 {
				return errors.New("noquota")
			}
			curQuota--
		}else{
			curQuota = 10
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