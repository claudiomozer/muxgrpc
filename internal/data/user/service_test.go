package data_user

import (
	"errors"
	"testing"

	mock_data_user_protocols "github.com/claudiomozer/muxgrpc/internal/data/user/protocols/mocks"
	domain_user "github.com/claudiomozer/muxgrpc/internal/domain/user"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type UserServiceSuite struct {
	suite.Suite
	ctrl                   *gomock.Controller
	create_user_repository *mock_data_user_protocols.MockCreateUserRepository
	find_user_repository   *mock_data_user_protocols.MockFindUserByDocumentRepository
	sut                    *Service
}

func (suite *UserServiceSuite) SetupTest() {
	suite.create_user_repository = mock_data_user_protocols.NewMockCreateUserRepository(suite.ctrl)
	suite.find_user_repository = mock_data_user_protocols.NewMockFindUserByDocumentRepository(suite.ctrl)
	suite.sut = New(suite.create_user_repository, suite.find_user_repository)
}

func (suite *UserServiceSuite) TearDowntest() {
	suite.ctrl.Finish()
}

func (suite *UserServiceSuite) TestShouldReturnAnErrorIfCreateUserRepositoryReturnsError() {
	dto := domain_user.CreateUserUseCaseDTO{}
	mockError := errors.New("error")
	suite.create_user_repository.EXPECT().Create(dto).Return(mockError)
	err := suite.sut.Create(dto)

	suite.NotNil(err, "should return an error if CreateUserRepository returns error")
}

func (suite *UserServiceSuite) TestShouldNotReturnAnErrorIfCreateUserRepositorySucceed() {
	dto := domain_user.CreateUserUseCaseDTO{}
	suite.create_user_repository.EXPECT().Create(dto).Return(nil)
	err := suite.sut.Create(dto)

	suite.Nil(err, "should not return an error if CreateUserRepository succeeds")
}

func (suite *UserServiceSuite) TestShouldReturnAnErrorIfFindUserByDocumentRepositoryReturnsError() {
	document := "123123123"
	mockError := errors.New("error")
	suite.find_user_repository.EXPECT().FindByDocument(document).Return(nil, mockError)
	dto, err := suite.sut.FindByDocument(document)

	suite.NotNil(err, "should return an error if CreateUserRepository returns error")
	suite.Nil(dto, "should not return a valid dto")
}

func (suite *UserServiceSuite) TestShouldNotReturnAnErrorIfFindUserByDocumentRepositorySucceed() {
	document := "123123123"
	dto := &domain_user.FindUserByDocumentUseCaseDTO{Document: document}
	suite.find_user_repository.EXPECT().FindByDocument(document).Return(dto, nil)
	findDtoResult, err := suite.sut.FindByDocument(document)

	suite.Nil(err, "should not return an error if CreateUserRepository succeeds")
	suite.Equal(dto, findDtoResult, "should be returned the mocked dto")
}

func TestSuite(t *testing.T) {
	testSuite := new(UserServiceSuite)
	testSuite.ctrl = gomock.NewController(t)
	suite.Run(t, testSuite)
}
