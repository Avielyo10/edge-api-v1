package update

// UpdatesInterface is the interface for the updateable objects.
type UpdatesInterface interface {
	IsSuccessful() bool
	Upgrade() error
	CheckForUpdate() error
	Rollback() error
}
