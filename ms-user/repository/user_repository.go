package repository

import (
	"database/sql"
	"errors"
	"ms-user/model"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (u *UserRepository) AddUser(user model.User) error {
	_, err := u.DB.Exec("INSERT INTO users(name, email, password, role_id, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)", user.Name, user.Email, user.Password, user.RoleID, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetCustomer(user *model.User) error {
	query := "SELECT name, email FROM users WHERE id = $1"
	err := u.DB.QueryRow(query, user.Id).Scan(&user.Name, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	return nil
}

func (u *UserRepository) UpdateCustomer(userID int, customer model.User) error {
	query := "UPDATE users SET name=$1, email=$2, password=$3, updated_at=$4 WHERE id=$5"
	_, err := u.DB.Exec(query, customer.Name, customer.Email, customer.Password, customer.UpdatedAt, userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) DeleteCustomer(userID int) error {
	_, err := u.DB.Exec("DELETE FROM users WHERE id=$1", userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) FindByCredentials(email, password string) (*model.User, error) {
	var user model.User

	err := u.DB.QueryRow("SELECT id, name, email, password, role_id FROM users WHERE email=$1", email).
		Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.RoleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetUserAdmin(userID int) (*model.User, error) {
	query := "SELECT id, name, email, role_id FROM users WHERE id=$1"
	row := u.DB.QueryRow(query, userID)

	var user model.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.RoleID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetAllCustomerAdmin() ([]*model.User, error) {
	query := "SELECT id, name, email, role_id FROM users"
	rows, err := u.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.RoleID)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) UpdateCustomerAdmin(userID int, customer model.User) error {
	query := "UPDATE users SET name=$1, email=$2, password=$3, updated_at=$4 WHERE id=$5"
	_, err := u.DB.Exec(query, customer.Name, customer.Email, customer.Password, customer.UpdatedAt, userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) UserExists(userID int) (bool, error) {
	var exists bool

	query := "SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)"
	err := u.DB.QueryRow(query, userID).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return exists, nil
}

func (u *UserRepository) AddAddress(address model.Address) (int, error) {
	var id int
	err := u.DB.QueryRow("INSERT INTO address(country, city) VALUES($1, $2) RETURNING id", address.Country, address.City).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UserRepository) SetAddressCustomer(user model.User, addressID int) error {
	query := "UPDATE users SET address_id=$1, updated_at=$2 WHERE id=$3"
	_, err := u.DB.Exec(query, addressID, user.UpdatedAt, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetAddressID(userID int) (int, error) {
	var addressID int
	err := u.DB.QueryRow("SELECT address_id FROM users WHERE id=$1", userID).Scan(&addressID)
	if err != nil {
		return 0, err
	}

	return addressID, nil
}

func (u *UserRepository) SetUpdatedAtAfterUpdateAddress(user model.User) error {
	query := "UPDATE users SET updated_at=$1 WHERE id=$2"
	_, err := u.DB.Exec(query, user.UpdatedAt, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) UpdateAddress(addressID int, address model.Address) error {
	query := "UPDATE address SET country=$1, city=$2 WHERE id=$3"
	_, err := u.DB.Exec(query, address.Country, address.City, addressID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetAddress(address *model.Address) (*model.Address, error) {
	query := "SELECT * FROM address WHERE id = $1"
	err := u.DB.QueryRow(query, address.Id).Scan(&address.Id, &address.Country, &address.City)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return address, nil
}
