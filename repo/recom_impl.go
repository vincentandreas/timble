package repo

import (
	"context"
	"fmt"

	"github.com/vincentandreas/dealls/models"
)

func (repo *Implementation) GetRecommendation(userId uint64, ctx context.Context) (models.Recommendation, error) {
	var user models.User
	var err error

	excludeUserIds, err := repo.SeenUserIds(userId, ctx)
	excludeUserIds = append(excludeUserIds, userId)
	fmt.Println("excludeUserIds: ", excludeUserIds)
	if err!= nil {
        return models.Recommendation{}, err
    }

	trx := repo.db.Preload("UserPreference").Joins("JOIN user_preferences ON user_preferences.user_id = users.id")

	err = trx.Where("users.id NOT IN ?", excludeUserIds).First(&user).Error

	if err!= nil {
        return models.Recommendation{}, err
    }

	recom := &models.Recommendation{
		Name : user.Name,
		Hobby: user.UserPreference.Hobby,
		UserId: uint64(user.ID),
	}
	return *recom, nil
}