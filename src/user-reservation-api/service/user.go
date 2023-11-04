package service

import (
	"errors"
	"user-reservation-api/client"
	"user-reservation-api/dto"
	"user-reservation-api/model"

	"golang.org/x/crypto/bcrypt"
)

type userService struct{}

type userServiceInterface interface {
	InsertUser(userDto dto.UserDto) (dto.UserDto, error)
	GetUserById(id int) (dto.UserDto, error)
	GetUsers() (dto.UsersDto, error)
	UserLogin(loginDto dto.UserDto) (dto.UserDto, error)
}

var UserService userServiceInterface

func init() {
	UserService = &userService{}
}

func (s *userService) InsertUser(userDto dto.UserDto) (dto.UserDto, error) {
	var user model.User

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)

	if err != nil {
		return userDto, err
	}

	user.Name = userDto.Name
	user.LastName = userDto.LastName
	user.Dni = userDto.Dni
	user.Email = userDto.Email
	user.Password = string(encryptedPassword)
	user.Role = "Customer"

	user = client.InsertUser(user)

	userDto.Id = user.Id
	userDto.Role = user.Role
	userDto.Password = user.Password

	if user.Id == 0 {
		return userDto, errors.New("error creating user")
	}

	return userDto, nil
}

func (s *userService) GetUserById(id int) (dto.UserDto, error) {

	var user model.User = client.GetUserById(id)

	var userDto dto.UserDto

	if user.Id == 0 {
		return userDto, errors.New("user not found")
	}

	userDto.Id = user.Id
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.Dni = user.Dni
	userDto.Email = user.Email
	userDto.Role = user.Role

	return userDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, error) {
	var users model.Users = client.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Id = user.Id
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.Dni = user.Dni
		userDto.Email = user.Email
		userDto.Role = user.Role

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) UserLogin(loginDto dto.UserDto) (dto.UserDto, error) {

	var user = client.GetUserByEmail(loginDto.Email)

	if user.Id == 0 {
		return loginDto, errors.New("user not registered")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password))
	if err != nil {
		// Passwords don't match
		return loginDto, errors.New("incorrect password")
	}

	var userDto dto.UserDto

	userDto.Id = user.Id
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.Dni = user.Dni
	userDto.Email = user.Email
	userDto.Role = user.Role
	return userDto, nil
}
