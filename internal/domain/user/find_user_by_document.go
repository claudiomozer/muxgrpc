package domain_user

type FindUserByDocumentUseCase interface {
	FindByDocument(document string) (*FindUserByDocumentUseCaseDTO, error)
}

type FindUserByDocumentUseCaseDTO struct {
	Name     string `json:"name"`
	Document string `json:"document"`
}
