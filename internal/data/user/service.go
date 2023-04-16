package data_user

import (
	data_user_protocols "github.com/claudiomozer/muxgrpc/internal/data/user/protocols"
	domain "github.com/claudiomozer/muxgrpc/internal/domain/user"
)

type Service struct {
	create_user_repository data_user_protocols.CreateUserRepository
	find_user_repository   data_user_protocols.FindUserByDocumentRepository
}

func New(
	create_user_repository data_user_protocols.CreateUserRepository,
	find_user_repository data_user_protocols.FindUserByDocumentRepository,
) *Service {
	return &Service{
		create_user_repository: create_user_repository,
		find_user_repository:   find_user_repository,
	}
}

func (s *Service) Create(dto domain.CreateUserUseCaseDTO) error {
	return s.create_user_repository.Create(dto)
}

func (s *Service) FindByDocument(document string) (*domain.FindUserByDocumentUseCaseDTO, error) {
	return s.find_user_repository.FindByDocument(document)
}
