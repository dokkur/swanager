package drivers

import (
	"fmt"
	"strings"

	"github.com/da4nik/swanager/core/entities"
	"github.com/da4nik/swanager/core/entities/drivers/mongo"
)

// DBDriver interface for backend db drivers
type DBDriver interface {
	GetUser(emailOrID string) *entities.User
}

// GetDBDriver return selected db driver
func GetDBDriver(driver string) (DBDriver, error) {
	switch strings.ToLower(driver) {
	case "mongo":
		return new(mongo.Mongo), nil
	default:
		return nil, fmt.Errorf("Unknown driver %s.", driver)
	}
}