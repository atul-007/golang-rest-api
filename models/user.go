package models

type User struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	Name      string `json:"name"`
}

func (*User) GetBaseUrl() string {
	return "https://645e3aae12e0a87ac0eadd29.mockapi.io/api"
}
