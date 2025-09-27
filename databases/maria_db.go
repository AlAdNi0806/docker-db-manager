package databases

import "fmt"

// MariaDB implements the DBHandler interface for MariaDB databases.
type MariaDB struct {
	// You can add fields such as connection parameters here.
	// For this simple example we only keep a name for demonstration.
	Name string
}

// Ensure MariaDB satisfies the DBHandler interface.
var _ DBHandler = MariaDB{}

// Create creates a new MariaDB database. In a real implementation this would
// run the appropriate command. Here we just print a message.
func (m MariaDB) Create(name string) error {
	fmt.Printf("Creating MariaDB database: %s\n", name)
	return nil
}

// Remove deletes an existing MariaDB database. In a real implementation this would
// run the appropriate command. Here we just print a message.
func (m MariaDB) Remove(name string) error {
	fmt.Printf("Removing MariaDB database: %s\n", name)
	return nil
}
