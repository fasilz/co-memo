package models

import (
	"fmt"
	"regexp"

	"github.com/revel/revel"
)

const (
	MAXUSERNAMESIZE = 15
	MINUSERNAMESIZE = 4
	MAXPASSWORDSIZE = 30
	MINPASSWORDSIZE = 8
)

// User ...
type User struct {
	UserID             uint32
	Name               string
	Username, Password string
	HashedPassword     []byte
	Privileged         bool
	TeamID             uint32
	DepartmentID       uint32
	OrganizationID     uint32
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Username)
}

var userRegex = regexp.MustCompile("^\\w*$")

func (user *User) Validate(v *revel.Validation) {
	v.Check(user.Username,
		revel.Required{},
		revel.MaxSize{MAXUSERNAMESIZE},
		revel.MinSize{MINUSERNAMESIZE},
		revel.Match{userRegex},
	)

	ValidatePassword(v, user.Password).
		Key("user.Password")

	v.Check(user.Name,
		revel.Required{},
		revel.MaxSize{100},
	)
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{MAXPASSWORDSIZE},
		revel.MinSize{MINPASSWORDSIZE},
	)
}

func (user *User) IsPreviledged() bool {
	return user.Privileged
}
