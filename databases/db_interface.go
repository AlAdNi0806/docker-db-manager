package databases

import "docker-db-management/types"

type DBHandler interface {
	SetConfig(config types.Config) error
	Create() error
	Remove(name string) error
}
