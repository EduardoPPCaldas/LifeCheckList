package usecase

import (
	"errors"

	"github.com/EduardoPPCaldas/LifeCheckList/application/dto"
	"github.com/EduardoPPCaldas/LifeCheckList/application/validator"
	"github.com/EduardoPPCaldas/LifeCheckList/domain/entity"
	"github.com/EduardoPPCaldas/LifeCheckList/infrastructure/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUCHandler interface {
	SignUp(dto.CreateUserDto) (*dto.UserDto, error)
}

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (uc UserUseCase) SignUp(userDto dto.CreateUserDto) (*dto.UserDto, error) {
	// Check if email is valid
	if !validator.IsEmailValid(userDto.Email) {
		return nil, errors.New("Email invalid")
	}

	// Create hash of password
	hash, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), 10)

	if err != nil {
		return nil, err
	}

	// Create new user
	user := entity.NewUser(userDto.Email, userDto.Name, string(hash))
	newUser, err := uc.repository.Create(*user)
	if err != nil {
		return nil, err
	}

	result := &dto.UserDto{
		ID: newUser.ID,
		Name: newUser.Name,
		Email: newUser.Email,
	}

	return result, nil
}
