package repository

import (
	"database/sql"

	"github.com/Drinnn/kool-bank/domain"
)

type TransactionRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{
		db,
	}
}

func (r *TransactionRepositoryDb) SaveTransaction(transaction *domain.Transaction, creditCard *domain.CreditCard) error {
	stmt, err := r.db.Prepare(`insert into transactions(id, amount, status, description, store,
		credit_card_id, created_at)
		values($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		transaction.ID,
		transaction.Amount,
		transaction.Status,
		transaction.Description,
		transaction.Store,
		transaction.CreditCardId,
		transaction.CreatedAt,
	)
	if err != nil {
		return err
	}

	if transaction.Status == "approved" {
		err = r.updateCreditCardBalance(creditCard)
		if err != nil {
			return err
		}
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepositoryDb) updateCreditCardBalance(creditCard *domain.CreditCard) error {
	_, err := r.db.Exec("update credit_cards set balance = $1 where id = $2",
		creditCard.Balance, creditCard.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepositoryDb) CreateCreditCard(creditCard *domain.CreditCard) error {
	stmt, err := r.db.Prepare(`insert into credit_cards(id, name, number, expiration_month,
		expiration_year, cvv, balance, limit, created_at)
		values($1, $2, $3, $4, $5, $6, $7, $8, $9)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		creditCard.ID,
		creditCard.Name,
		creditCard.Number,
		creditCard.ExpirationMonth,
		creditCard.ExpirationYear,
		creditCard.CVV,
		creditCard.Balance,
		creditCard.Limit,
		creditCard.CreatedAt,
	)
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}
