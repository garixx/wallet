package validator

import (
	"errors"
	"testing"
	"wallet/internal/services/transaction"
	"wallet/internal/services/user"
	"wallet/internal/services/wallet"
)

func TestService_Validate(t *testing.T) {
	type args struct {
		entity any
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			name: "unsupported type",
			args: args{
				entity: struct{}{},
			},
			expectedErr: ErrUnsupportedType{message: "struct {}"},
		},
		{
			name: "user",
			args: args{
				entity: user.User{Name: "valid"},
			},
		},
		{
			name: "wallet",
			args: args{
				entity: wallet.Wallet{Balance: 3},
			},
		},
		{
			name: "transaction",
			args: args{
				entity: transaction.Transaction{Amount: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			switch {
			case tt.expectedErr == nil:
				if err := s.Validate(tt.args.entity); err != nil {
					t.Errorf("Validate() error = %v", err)
				}
			case tt.expectedErr != nil:
				if err := s.Validate(tt.args.entity); !errors.Is(err, tt.expectedErr) {
					t.Errorf("validate() error.\nactual   = %v,\nexpected = %v", err, tt.expectedErr)
				}
			}
		})
	}
}

func Test_validateUser(t *testing.T) {
	type args struct {
		u user.User
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			name: "too short name",
			args: args{
				u: user.User{
					Name: "xxx",
				},
			},
			expectedErr: ErrUserValidation{
				message: "name length less than 5",
			},
		},
		{
			name: "too long name",
			args: args{
				u: user.User{
					Name: "mytoolongnamelongerthantwenty",
				},
			},
			expectedErr: ErrUserValidation{
				message: "name length more than 20",
			},
			//expectedErr: validationError{
			//	message: "name length more than 20",
			//},
		},
		{
			name: "valid name",
			args: args{
				u: user.User{
					Name: "validName",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			switch {
			case tt.expectedErr == nil:
				if err := s.Validate(tt.args.u); err != nil {
					t.Errorf("validateUser() error = %v", err)
				}
			case tt.expectedErr != nil:
				if err := s.Validate(tt.args.u); !errors.Is(err, tt.expectedErr) {
					t.Errorf("validateUser() error.\nactual   = %v,\nexpected = %v", err, tt.expectedErr)
				}
			}
		})
	}
}

func Test_validateWallet(t *testing.T) {
	type args struct {
		w wallet.Wallet
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			name: "valid balance",
			args: args{
				w: wallet.Wallet{
					Balance: 10.0,
				},
			},
		},
		{
			name: "negative balance",
			args: args{
				w: wallet.Wallet{
					Balance: -10.0,
				},
			},
			expectedErr: ErrWalletValidation{
				message: "balance less than zero",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			switch {
			case tt.expectedErr == nil:
				if err := s.Validate(tt.args.w); err != nil {
					t.Errorf("validateWallet() error = %v", err)
				}
			case tt.expectedErr != nil:
				if err := s.Validate(tt.args.w); !errors.Is(err, tt.expectedErr) {
					t.Errorf("validateUser() error.\nactual   = %v,\nexpected = %v", err, tt.expectedErr)
				}
			}
		})
	}
}

func Test_validateTransaction(t *testing.T) {
	type args struct {
		ta transaction.Transaction
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			name: "valid amount",
			args: args{
				ta: transaction.Transaction{
					Amount: 10.7,
				},
			},
		},
		{
			name: "negative amount",
			args: args{
				ta: transaction.Transaction{
					Amount: -10.0,
				},
			},
			expectedErr: ErrTransactionValidation{
				message: "amount less than zero",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			switch {
			case tt.expectedErr == nil:
				if err := s.Validate(tt.args.ta); err != nil {
					t.Errorf("validateTransaction() error = %v", err)
				}
			case tt.expectedErr != nil:
				if err := s.Validate(tt.args.ta); !errors.Is(err, tt.expectedErr) {
					t.Errorf("validateUser() error.\nactual   = %v,\nexpected = %v", err, tt.expectedErr)
				}
			}
		})
	}
}
