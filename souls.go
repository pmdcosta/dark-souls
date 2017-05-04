package dark_souls

// SaveService represents a service for managing saves.
type SaveService interface {
	Save() error
	Load() error
	Start() error
	Stop() error
}
