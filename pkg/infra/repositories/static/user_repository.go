package static

import domain_user "github.com/claudiomozer/muxgrpc/internal/domain/user"

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) Create(dto domain_user.CreateUserUseCaseDTO) error {
	return nil
}

func (repo *UserRepository) FindByDocument(document string) (*domain_user.FindUserByDocumentUseCaseDTO, error) {
	return &domain_user.FindUserByDocumentUseCaseDTO{
		Document: document,
		Name:     "Test",
	}, nil
}
