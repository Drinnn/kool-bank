package usecase

import (
	"github.com/Drinnn/kool-bank/domain"
	"github.com/Drinnn/kool-bank/dto"
)

type TransactionUseCase struct {
	TransactionRepository domain.TransactionRepository
}

func NewTransactionUseCase(transactionRepository domain.TransactionRepository) *TransactionUseCase {
	return &TransactionUseCase{
		TransactionRepository: transactionRepository,
	}
}

func (uc *TransactionUseCase) ProcessTransaction(transactionDto dto.TransactionDto) (*domain.Transaction, error) {
	creditCard := uc.hydrateCreditCard(transactionDto)

	ccBalanceAndLimit, err := uc.TransactionRepository.GetCreditCard(creditCard)
	if err != nil {
		return &domain.Transaction{}, err
	}

	creditCard.ID = ccBalanceAndLimit.ID
	creditCard.Limit = ccBalanceAndLimit.Limit
	creditCard.Balance = ccBalanceAndLimit.Balance

	transaction := uc.newTransaction(transactionDto, creditCard)
	transaction.ProcessAndValidate(creditCard)

	err = uc.TransactionRepository.SaveTransaction(transaction, creditCard)
	if err != nil {
		return &domain.Transaction{}, err
	}

	return transaction, nil
}

func (uc *TransactionUseCase) hydrateCreditCard(transactionDto dto.TransactionDto) *domain.CreditCard {
	creditCard := domain.NewCreditCard()
	creditCard.Name = transactionDto.Name
	creditCard.Number = transactionDto.Number
	creditCard.ExpirationMonth = transactionDto.ExpirationMonth
	creditCard.ExpirationYear = transactionDto.ExpirationYear
	creditCard.CVV = transactionDto.CVV

	return creditCard
}

func (uc *TransactionUseCase) newTransaction(transactionDto dto.TransactionDto, cc *domain.CreditCard) *domain.Transaction {
	transaction := domain.NewTransaction()
	transaction.CreditCardId = cc.ID
	transaction.Amount = transactionDto.Amount
	transaction.Store = transactionDto.Store
	transaction.Description = transactionDto.Description

	return transaction
}
