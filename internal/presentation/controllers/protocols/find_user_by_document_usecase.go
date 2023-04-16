package controllers_protocols

import domain_user "github.com/claudiomozer/muxgrpc/internal/domain/user"

type FindUserByDocumentUseCase interface {
	FindByDocument(document string) (*domain_user.FindUserByDocumentUseCaseDTO, error)
}
