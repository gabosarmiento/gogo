package execute_transaction

import "github.com/gabosarmiento/gogo/entity"

type ExecuteTransaction struct {
	Repository entity.TransactionRepository
}

func NewExecuteTransaction(repository entity.TransactionRepository) *ExecuteTransaction {
	return &ExecuteTransaction{
		Repository: repository,
	}
}

func (p *ExecuteTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {

	transaction := entity.NewTransaction()
	transaction.Amount = input.Amount
	transaction.AccountID = input.AccountID
	transaction.ID = input.ID
	invalidtransaction := transaction.IsValid()

	if invalidtransaction != nil {
		err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "rejected", invalidtransaction.Error())

		if err != nil {
			return TransactionDtoOutput{}, err
		}

		output := TransactionDtoOutput{
			ID:            transaction.ID,
			Status:        "rejected",
			ErrrorMessage: invalidtransaction.Error(),
		}

		return output, nil
	}

	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "approved", "")

	if err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		ID:            transaction.ID,
		Status:        "approved",
		ErrrorMessage: "",
	}

	return output, nil
}
