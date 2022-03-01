package usecase

import (
	"testing"

	"github.com/anandawira/anandapay/domain"
	"github.com/anandawira/anandapay/pkg/wallet/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WalletUsecaseTestSuite struct {
	suite.Suite

	mockRepo *repo.MockWalletRepo
	usecase  domain.WalletUsecase
}

func (ts *WalletUsecaseTestSuite) SetupSuite() {
	ts.mockRepo = new(repo.MockWalletRepo)
	ts.usecase = NewWalletUsecase(ts.mockRepo)
}

func (ts *WalletUsecaseTestSuite) TestGetBalance() {
	ts.T().Run("It should return balance if wallet found", func(t *testing.T) {
		ts.mockRepo.On(
			"GetBalance",
			mock.AnythingOfType("string"),
		).Return(12, nil).Once()

		balance, err := ts.usecase.GetBalance("walletId1")
		assert.NoError(t, err)
		assert.Equal(t, uint64(12), balance)
	})

	ts.T().Run("It should return error if wallet not found", func(t *testing.T) {
		ts.mockRepo.On(
			"GetBalance",
			mock.AnythingOfType("string"),
		).Return(0, domain.ErrWalletNotFound).Once()

		_, err := ts.usecase.GetBalance("walletId1")
		assert.Error(t, err)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(WalletUsecaseTestSuite))
}