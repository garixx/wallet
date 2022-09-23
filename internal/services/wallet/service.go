package wallet

type Wallet struct {
	UserID  int
	Balance float64
}

type Repository interface {
}

// Service holds calendar business logic and works with repository
type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}
