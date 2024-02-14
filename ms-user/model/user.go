package model

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	RoleID    int    `json:"role_id"`
	AddressID int    `json:"address_id"`
	CreatedAt string `json:"create_at"`
	UpdatedAt string `json:"update_at"`
}

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"role"`
}

type Address struct {
	Id      int    `json:"id"`
	Country string `json:"country"`
	City    string `json:"city"`
}
