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

// response for swagger
type RegisterResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type InputRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddressRequest struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type InfoCustomerResponse struct {
	User    RegisterResponse `json:"user"`
	Address AddressRequest   `json:"address"`
}
