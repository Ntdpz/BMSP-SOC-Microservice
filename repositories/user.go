package repositories

import "bmsp-backend-service/models"

func (r Repositories) FindUser(username string) (models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
