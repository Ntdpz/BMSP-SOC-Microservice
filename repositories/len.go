package repositories

import "bmsp-backend-service/models"

func (r Repositories) GetLen(customerName string) (int64, int64) {
	var countSent int64
	r.db.Model(&models.Document{}).Where("customer = ? AND status = ?", customerName, "sent").Count(&countSent)

	var countWaiting int64
	r.db.Model(&models.Document{}).Where("customer = ? AND status = ?", customerName, "waiting").Count(&countWaiting)
	return countSent, countWaiting
}
