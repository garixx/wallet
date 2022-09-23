package validator

import (
	"fmt"
	"wallet/internal/services/transaction"
	"wallet/internal/services/user"
	"wallet/internal/services/wallet"
)

type validationError struct {
	message string
}

func (v validationError) Error() string {
	return v.message
}

func Error(message string) error {
	return validationError{message: fmt.Sprintf("validation error: %s", message)}
}

func WalletError(message string) error {
	return validationError{message: fmt.Sprintf("wallet validation error: %s", message)}
}

func TransactionError(message string) error {
	return validationError{message: fmt.Sprintf("transaction validation error: %s", message)}
}

func UserError(message string) error {
	return validationError{message: fmt.Sprintf("user validation error: %s", message)}
}

// Service validates structures
type Service struct {
}

func (s *Service) Validate(entity any) error {
	switch e := entity.(type) {
	case transaction.Transaction:
		return validateTransaction(e)
	case user.User:
		return validateUser(e)
	case wallet.Wallet:
		return validateWallet(e)
	default:
		return Error(fmt.Sprintf("unsupported type %T", e))
	}
}

func validateTransaction(ta transaction.Transaction) error {
	if ta.Amount < 0 {
		return TransactionError("amount less than zero")
	}
	return nil
}

func validateUser(u user.User) error {
	if len(u.Name) < 5 {
		return UserError("name length less than 5")
	}
	if len(u.Name) > 20 {
		return UserError("name length more than 20")
	}
	return nil
}

func validateWallet(w wallet.Wallet) error {
	if w.Balance < 0.0 {
		return WalletError("balance less than zero")
	}
	return nil
}
