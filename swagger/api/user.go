// Package api defines the user model.
package api

// User represents body of User request and response.
type User struct {
	// User's name.
	// Required: true
	Name string `json:"name"`

	// User's nickname.
	// Required: true
	Nickname string `json:"nickname"`

	// User's address.
	Address string `json:"address"`

	// User's email.
	Email string `json:"email"`
}
