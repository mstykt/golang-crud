package repo

import (
	"errors"
	"golang-crud/configs"
	. "golang-crud/pkg/entity"
)

type UserRepo struct {
	connection *configs.Connection
}

func NewUserRepo(connection *configs.Connection) *UserRepo {
	return &UserRepo{connection: connection}
}

func (userRepo *UserRepo) Create(user User) (User, error) {
	session := userRepo.connection.GetSession()
	tx := session.Create(&user)
	if tx.Error != nil {
		return User{}, errors.New(tx.Error.Error())
	}
	return user, nil
}

func (userRepo UserRepo) GetUser(id uint64) (User, error) {
	session := userRepo.connection.GetSession()
	var user User
	tx := session.Find(&user, User{ID: id})
	if tx.Error != nil {
		return User{}, errors.New(tx.Error.Error())
	}

	if user.ID == 0 {
		return User{}, errors.New("cannot find user")
	}

	return user, nil
}

func (userRepo UserRepo) GetAllUser() ([]User, error) {
	session := userRepo.connection.GetSession()
	var users []User
	tx := session.Find(&users)
	if tx.Error != nil {
		return []User{}, errors.New(tx.Error.Error())
	}

	return users, nil
}

func (userRepo UserRepo) Update(user User) User {
	session := userRepo.connection.GetSession()
	session.Save(&user)
	return user
}

func (userRepo UserRepo) Delete(id uint64) error {
	session := userRepo.connection.GetSession()
	tx := session.Delete(&User{}, id)

	if tx.Error != nil {
		return errors.New(tx.Error.Error())
	}

	return nil
}
