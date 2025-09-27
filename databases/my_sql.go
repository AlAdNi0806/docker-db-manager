package databases

import "fmt"

// MySQL implements the DBHandler interface for MySQL databases.
type MySQL struct {
	// You can add fields such as connection parameters here.
	// For this simple example we only keep a name for demonstration.
	Name string
}

// Ensure MySQL satisfies the DBHandler interface.
var _ DBHandler = MySQL{}

// Create creates a new MySQL database. In a real implementation this would
// run the appropriate command (e.g., `mysqladmin create`). Here we just
// print a message.
func (m MySQL) Create(name string) error {
	fmt.Printf("Creating MySQL database: %s\n", name)
	return nil
}

// Remove deletes an existing MySQL database. A real implementation would
// execute a command such as `mysqladmin drop`.
func (m MySQL) Remove(name string) error {
	fmt.Printf("Removing MySQL database: %s\n", name)
	return nil
}
