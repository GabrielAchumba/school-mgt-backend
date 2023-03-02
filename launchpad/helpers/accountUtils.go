package helpers

import (
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/account-module/models"
)

type AccountUtils struct {
}

func NewAccountUtils() AccountUtils {
	return AccountUtils{}
}

func (impl AccountUtils) FindAccountByCategoryIdId(
	acoounts []models.Account, categoryId string) models.Account {

	var account = models.Account{}

	for _, item := range acoounts {
		if item.ContributorId == categoryId {
			account = item
			return account
		}
	}
	return account
}
