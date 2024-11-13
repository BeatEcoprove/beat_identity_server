package domain

import (
	"errors"

	interfaces "github.com/BeatEcoprove/identityService/pkg/domain"
	"github.com/BeatEcoprove/identityService/pkg/services"
	"gorm.io/gorm"
)

type Role int

const (
	Client Role = iota
	Organization
	Admin
)

var ErrUndefinedRole = errors.New("role not defined")

type IdentityUser struct {
	interfaces.EntityBase
	Email    string
	Password string
	Salt     string `gorm:"column:salt"`
	IsActive bool
	Role     Role
}

func NewIdentityUser(email, password string, role Role) *IdentityUser {
	return &IdentityUser{
		Email:    email,
		Password: password,
		Role:     role,
	}
}

func (b *IdentityUser) TableName() string {
	return "auths"
}

func (u *IdentityUser) BeforeCreate(tx *gorm.DB) error {
	u.GetId()

	salt, err := services.GenerateSalt(services.SALT_COST)

	if err != nil {
		return err
	}

	password, err := services.HashPassword(u.Password, salt)

	if err != nil {
		return err
	}

	u.Salt = salt
	u.Password = password
	u.DeletedAt = nil
	return nil
}

func GetRole(role Role) (string, error) {
	switch role {
	case Client:
		return "client", nil
	case Admin:
		return "admin", nil
	case Organization:
		return "organization", nil
	}

	return "", ErrUndefinedRole
}