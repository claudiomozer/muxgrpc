package domain_user

type CreateUserUseCase interface {
	Create(dto CreateUserUseCaseDTO) error
}

type CreateUserUseCaseDTO struct {
	Name     string `json:"name"`
	Document string `json:"document"`
	Age      int16  `json:"age"`
}
