package repository;

// UserRepository handles user manipulations in the database
type UserRepository struct {
}

// ProvideUserRepository is the provider for UserRepository
func ProvideUserRepository() (*UserRepository, error) {
	return &UserRepository{}, nil
}
