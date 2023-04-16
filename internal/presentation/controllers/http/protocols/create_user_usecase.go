package controllers_protocols

import domain_user "github.com/claudiomozer/muxgrpc/internal/domain/user"

type CreateUserUseCase interface {
	Create(dto domain_user.CreateUserUseCaseDTO) error
}
