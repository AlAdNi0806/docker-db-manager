package databases

// DBHandler defines the common operations that any database implementation
// must provide. By using this interface the CLI can work with MySQL, MariaDB,
// or any future database types without needing to know their concrete
// implementations.
//
// The methods return an error to allow callers to handle failures gracefully.
type DBHandler interface {
	// Create creates a new database with the given name.
	Create(name string) error

	// Remove deletes the named database.
	Remove(name string) error
}
