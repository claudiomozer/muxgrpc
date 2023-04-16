package data_user_protocols

import domain_user "github.com/claudiomozer/muxgrpc/internal/domain/user"

type CreateUserRepository interface {
	Create(dto domain_user.CreateUserUseCaseDTO) error
}
