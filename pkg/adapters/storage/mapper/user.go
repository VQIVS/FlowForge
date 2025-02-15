package mapper

import (
	"github.com/VQIVS/FlowForge/internal/user/domain"
	"github.com/VQIVS/FlowForge/pkg/adapters/storage/types"
	"gorm.io/gorm"
)

func UserDomain2Storage(userDomain domain.User) *types.User {
	return &types.User{
		Model: gorm.Model{
			ID:        uint(userDomain.ID),
			CreatedAt: userDomain.CreatedAt,
		},
		FirstName: userDomain.FirstName,
		LastName:  userDomain.LastName,
		Phone:     string(userDomain.Phone),
		Password:  userDomain.Password,
	}
}
func UserStorage2Domain(userStorage types.User) *domain.User {
	return &domain.User{
		ID:        domain.UserID(userStorage.ID),
		CreatedAt: userStorage.CreatedAt,
		FirstName: userStorage.FirstName,
		LastName:  userStorage.LastName,
		Phone:     domain.Phone(userStorage.Phone),
		Password:  userStorage.Password,
	}
}
