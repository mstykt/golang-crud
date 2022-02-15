package service

import (
	"github.com/mashingan/smapping"
	. "golang-crud/pkg/entity"
	. "golang-crud/pkg/model/request"
	. "golang-crud/pkg/model/response"
	. "golang-crud/pkg/repo"
)

func NewUserService(userRepo *UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

type UserService struct {
	userRepo *UserRepo
}

func (userService UserService) CreateUser(request UserRequest) (UserResponse, error) {
	var user User
	fields := smapping.MapFields(&request)
	reqMapErr := smapping.FillStruct(&user, fields)
	if reqMapErr != nil {
		return UserResponse{}, reqMapErr
	}

	savedUser, err := userService.userRepo.Create(user)
	if err != nil {
		return UserResponse{}, err
	}

	savedFields := smapping.MapFields(&savedUser)
	var userResponse UserResponse
	resMapErr := smapping.FillStruct(&userResponse, savedFields)

	if resMapErr != nil {
		return UserResponse{}, resMapErr
	}

	return userResponse, nil
}

func (userService UserService) GetUser(id uint64) (UserResponse, error) {
	user, err := userService.userRepo.GetUser(id)

	if err != nil {
		return UserResponse{}, err
	}
	fields := smapping.MapFields(&user)
	var userResponse UserResponse
	mappingErr := smapping.FillStruct(&userResponse, fields)
	if mappingErr != nil {
		return UserResponse{}, mappingErr
	}
	return userResponse, nil
}

func (userService UserService) GetAllUser() ([]UserResponse, error) {
	users, err := userService.userRepo.GetAllUser()
	if err != nil {
		return []UserResponse{}, err
	}

	var userResponseList []UserResponse
	for _, user := range users {
		fields := smapping.MapFields(&user)
		var userResponse UserResponse
		_ = smapping.FillStruct(&userResponse, fields)
		userResponseList = append(userResponseList, userResponse)
	}

	return userResponseList, nil
}

func (userService UserService) UpdateUser(id uint64, userRequest UserRequest) (UserResponse, error) {
	user, err := userService.userRepo.GetUser(id)
	if err != nil {
		return UserResponse{}, err
	}

	fields := smapping.MapFields(&userRequest)
	mappingErr := smapping.FillStruct(&user, fields)
	if mappingErr != nil {
		return UserResponse{}, mappingErr
	}
	updatedUser := userService.userRepo.Update(user)
	updatedFields := smapping.MapFields(&updatedUser)
	var userResponse UserResponse
	updatedMappingErr := smapping.FillStruct(&userResponse, updatedFields)
	if updatedMappingErr != nil {
		return UserResponse{}, updatedMappingErr
	}

	return userResponse, nil
}

func (userService UserService) DeleteUser(id uint64) error {
	err := userService.userRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
