package user

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	uid      string
	name     string
	email    string
	password string
	role     UserRole
}

type UserRole string

func (r UserRole) String() string { return string(r) }

const (
	RoleBasic   = UserRole("basic")
	RoleManager = UserRole("manager")
	RoleAdmin   = UserRole("admin")
)

func NewUser(name string, email string, role UserRole, password string) (*User, error) {
	if name == "" {
		return nil, errors.New("name can't be empty")
	}

	if email == "" {
		return nil, errors.New("email can't be empty")
	}

	if password == "" {
		return nil, errors.New("password can't be empty")
	}

	return &User{
		uid:      uuid.New().String(),
		name:     name,
		email:    email,
		role:     role,
		password: password,
	}, nil
}

func (u User) UID() string                 { return u.uid }
func (u User) Name() string                { return u.name }
func (u User) Email() string               { return u.email }
func (u User) Role() UserRole              { return u.role }
func (u User) PasswordMatch(p string) bool { return p == u.password } // TODO: replace with hash

func (u *User) SetRole(r UserRole) { u.role = r }
