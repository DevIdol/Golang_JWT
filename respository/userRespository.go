package respository

import (
	"errors"

	"github.com/DevIdol/Golang_JWT/data/request"
	"github.com/DevIdol/Golang_JWT/helpers"
	"github.com/DevIdol/Golang_JWT/models"
	"gorm.io/gorm"
)

type UserRespositoryImpl struct {
	Db *gorm.DB
}

func NewUserRespositoryImpl(Db *gorm.DB) UserRespository {
	return &UserRespositoryImpl{Db: Db}
}

// Delete implements UserRespository
func (user *UserRespositoryImpl) Delete(userId int) {
	var users models.User
	result := user.Db.Where("id = ?", userId).Delete(&users)
	helpers.ErrorPanic(result.Error)
}

// FindAll implements UserRespository
func (user *UserRespositoryImpl) FindAll() []models.User {
	var users []models.User
	result := user.Db.Find(&users)
	helpers.ErrorPanic(result.Error)
	return users
}

// FindById implements UserRespository
func (user *UserRespositoryImpl) FindById(userId int) (models.User, error) {
	var users models.User
	result := user.Db.Find(&users, userId)
	if result != nil {
		return users, nil
	} else {
		return users, errors.New("user not found")
	}
}

// FindByUsername implements UserRespository
func (user *UserRespositoryImpl) FindByUsername(username string) (models.User, error) {
	var users models.User
	result := user.Db.First(&users, "username = ?", username)

	if result.Error != nil {
		return users, errors.New("invalid username or password")
	}
	return users, nil
}

// Save implements UserRespository
func (user *UserRespositoryImpl) Save(users models.User) {
	result := user.Db.Create(&users)
	helpers.ErrorPanic(result.Error)
}

// Update implements UserRespository
func (user *UserRespositoryImpl) Update(users models.User) {
	var updateUsers = request.UpdateUserRequest{
		Id:       users.Id,
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}

	result := user.Db.Model(&users).Updates(updateUsers)
	helpers.ErrorPanic(result.Error)
}
