package databases

import (
	"docker-db-management/types"
	"fmt"
)

type MariaDB struct {
	LatestImage  bool
	Password     string
	DatabaseName string
}

var _ DBHandler = &MariaDB{}

func (m *MariaDB) SetConfig(config types.Config) error {
	m.LatestImage = config.LatestImage
	m.Password = config.Password
	m.DatabaseName = config.DatabaseName
	return nil
}

func (m MariaDB) Create() error {
	fmt.Printf("Creating MariaDB database: %s\n", m.DatabaseName)
	return nil
}

func (m MariaDB) Remove(name string) error {
	fmt.Printf("Removing MariaDB database: %s\n", name)
	return nil
}
