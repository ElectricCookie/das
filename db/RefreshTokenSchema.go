package db

// RefreshToken is a token that can be used to gain access to an access_token
type RefreshToken struct {
	Expires int64  `json:"expires" bson:"expires,omitempty"`
	Token   string `json:"token" bson:"token"`
	UserID  string `json:"userId" bson:"userId"`
}
