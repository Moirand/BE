package roles

import (
	"github.com/sev-2/raiden"
)

type Users struct {
	raiden.RoleBase
}

func (r *Users) Name() string {
	return "users"
}

func (r *Users) CanLogin() bool {
	return true
}

