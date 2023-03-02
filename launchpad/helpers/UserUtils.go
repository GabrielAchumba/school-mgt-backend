package helpers

import (
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/user-module/dtos"
)

type UserUtils struct {
}

func NewUserUtils() UserUtils {
	return UserUtils{}
}

func (impl UserUtils) FindUserById(users []dtos.UserResponse, id string) dtos.UserResponse {

	var user dtos.UserResponse = dtos.UserResponse{}

	for _, value := range users {
		if value.ContributorId == id {
			return value
		}
	}

	return user
}
