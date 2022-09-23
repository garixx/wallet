package validator

import (
	"fmt"
	"wallet/internal/services/transaction"
	"wallet/internal/services/user"
	"wallet/internal/services/wallet"
)

type ErrUnsupportedType struct {
	message string
}

func (e ErrUnsupportedType) Error() string {
	return fmt.Sprintf("validation error: unsupported type %s", e.message)
}

type ErrUserValidation struct {
	message string
}

func (e ErrUserValidation) Error() string {
	return fmt.Sprintf("user validation error: %s", e.message)
}

type ErrWalletValidation struct {
	message string
}

func (e ErrWalletValidation) Error() string {
	return fmt.Sprintf("wallet validation error: %s", e.message)
}

type ErrTransactionValidation struct {
	message string
}

func (e ErrTransactionValidation) Error() string {
	return fmt.Sprintf("transaction validation error: %s", e.message)
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
		return ErrUnsupportedType{message: fmt.Sprintf("%T", e)}
	}
}

func validateTransaction(ta transaction.Transaction) error {
	if ta.Amount < 0 {
		return ErrTransactionValidation{message: "amount less than zero"}
	}
	return nil
}

func validateUser(u user.User) error {
	if len(u.Name) < 5 {
		return ErrUserValidation{message: "name length less than 5"}
	}
	if len(u.Name) > 20 {
		return ErrUserValidation{message: "name length more than 20"}
	}
	return nil
}

func validateWallet(w wallet.Wallet) error {
	if w.Balance < 0.0 {
		return ErrWalletValidation{message: "balance less than zero"}
	}
	return nil
}
