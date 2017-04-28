package db

// User describes a user in the system
type User struct {
	ID               string `json:"id"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"LastName"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Email            string `json:"email"`
	Language         string `json:"language"`
	EmailVerified    bool   `json:"emailVerified"`
	EmailVerifyToken string `json:"emailVerifyToken"`
	Salt             string `json:"salt"`
	Registered       int64  `json:"registered"`
	LastLogin        int64  `json:"lastLogin"`
}
