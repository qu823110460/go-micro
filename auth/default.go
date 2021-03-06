package auth

import (
	"github.com/google/uuid"
)

var (
	DefaultAuth = NewAuth()
)

func NewAuth(opts ...Option) Auth {
	return &noop{}
}

type noop struct {
	opts Options
}

// String returns the name of the implementation
func (n *noop) String() string {
	return "noop"
}

// Init the auth
func (n *noop) Init(opts ...Option) {
	for _, o := range opts {
		o(&n.opts)
	}
}

// Options set for auth
func (n *noop) Options() Options {
	return n.opts
}

// Generate a new account
func (n *noop) Generate(id string, opts ...GenerateOption) (*Account, error) {
	options := NewGenerateOptions(opts...)

	return &Account{
		ID:       id,
		Roles:    options.Roles,
		Metadata: options.Metadata,
		Secret:   uuid.New().String(),
	}, nil
}

// Grant access to a resource
func (n *noop) Grant(role string, res *Resource) error {
	return nil
}

// Revoke access to a resource
func (n *noop) Revoke(role string, res *Resource) error {
	return nil
}

// Verify an account has access to a resource
func (n *noop) Verify(acc *Account, res *Resource) error {
	return nil
}

// Inspect a token
func (n *noop) Inspect(token string) (*Account, error) {
	return &Account{
		ID: uuid.New().String(),
	}, nil
}

// Token generation using an account id and secret
func (n *noop) Token(id, secret string, opts ...TokenOption) (*Token, error) {
	return &Token{}, nil
}
