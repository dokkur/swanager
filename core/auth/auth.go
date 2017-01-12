package auth
import (
	"github.com/da4nik/swanager/core/entities"
)

// WithToken authenticates with token
func WithToken(token string) bool {
    _, _ = entities.FindUserByToken(token)
    // TODO: Auth
	if token == "" {
		return false
	}
	return true
}
