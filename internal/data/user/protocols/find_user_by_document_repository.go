package data_user_protocols

import domain_user "github.com/claudiomozer/muxgrpc/internal/domain/user"

type FindUserByDocumentRepository interface {
	FindByDocument(document string) (*domain_user.FindUserByDocumentUseCaseDTO, error)
}
