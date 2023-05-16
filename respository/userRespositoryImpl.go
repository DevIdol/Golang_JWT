package respository

import "github.com/DevIdol/Golang_JWT/models"

type UserRespository interface {
	Save(user models.User)
	Update(user models.User)
	Delete(userId int)
	FindById(userId int) (models.User, error)
	FindAll() []models.User
	FindByUsername(username string) (models.User, error)
}
