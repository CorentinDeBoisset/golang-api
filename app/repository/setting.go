package repository;

// SettingRepository handles recipe manipulations in the database
type SettingRepository struct {
}

// ProvideSettingRepository is the provider for SettingRepository
func ProvideSettingRepository() (*SettingRepository, error) {
	return &SettingRepository{}, nil
}
