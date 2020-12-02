package user

// User is base entity of user
type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Verified bool   `json:"verified"`
}
