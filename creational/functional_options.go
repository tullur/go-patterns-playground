package creational

import "errors"

var (
	ErrEmptyId        = errors.New("id cannot be empty")
	ErrEmptyEmail     = errors.New("email cannot be empty")
	ErrInvalidEmail   = errors.New("email is invalid")
	ErrEmptyFirstName = errors.New("first name cannot be empty")
	ErrEmptyLastName  = errors.New("last lane cannot be empty")
)

type Customer struct {
	id            string
	email         string
	firstName     *string
	lastName      *string
	loyaltyPoints int
}

type Option func(c *Customer)

func WithName(firstName, lastName *string) Option {
	return func(c *Customer) {
		c.firstName = firstName
		c.lastName = lastName
	}
}

func WithLoyaltyPoints(points int) Option {
	return func(c *Customer) {
		c.loyaltyPoints = points
	}
}

func NewCustomer(id, email string, opts ...Option) (*Customer, error) {
	c := &Customer{id: id, email: email, loyaltyPoints: 100}
	for _, opt := range opts {
		opt(c)
	}

	if err := c.validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Customer) validate() error {
	if c.id == "" {
		return ErrEmptyId
	}

	if c.email == "" {
		return ErrEmptyEmail
	}

	if c.firstName != nil && *c.firstName == "" {
		return ErrEmptyFirstName
	}
	if c.lastName != nil && *c.lastName == "" {
		return ErrEmptyLastName
	}

	return nil
}
