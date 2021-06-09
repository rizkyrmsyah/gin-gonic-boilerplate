package usecase

import (
	"github.com/rizkyrmsyah/gin-gonic-boilerplate/repository"
)

type Example struct {
	repo repository.ExampleRepositoryI
}

func NewExample(repo repository.ExampleRepositoryI) ExampleUseCaseI {
	return &Example{
		repo,
	}
}

// function implemented from interface
