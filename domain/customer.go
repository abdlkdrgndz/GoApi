package domain

type Customer struct {
	Name     string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
