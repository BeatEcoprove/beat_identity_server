package repositories

import (
	interfaces "github.com/BeatEcoprove/identityService/pkg/adapters"
)

type (
	ProfileRepository struct {
		interfaces.RepositoryBase
	}

	IProfileRepository interface {
		interfaces.Repository
	}
)

func NewProfileRepository(database interfaces.Database) *ProfileRepository {
	return &ProfileRepository{
		RepositoryBase: *interfaces.NewRepositoryBase(database),
	}
}