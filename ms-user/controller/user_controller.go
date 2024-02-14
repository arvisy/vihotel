package controller

import (
	"context"
	"errors"
	"fmt"
	"ms-user/model"
	"ms-user/pb"
	"ms-user/repository"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	pb.UnimplementedUserServiceServer
	UserRepository repository.UserRepository
}

func NewUserController(UserRepository repository.UserRepository) *UserController {
	return &UserController{
		UserRepository: UserRepository,
	}
}

func (u *UserController) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	user := model.User{
		Name:      in.Name,
		Email:     in.Email,
		Password:  string(hashedPassword),
		RoleID:    2,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}

	err = u.UserRepository.AddUser(user)
	if err != nil {
		fmt.Println(err)
		return &pb.RegisterResponse{}, err
	}

	return &pb.RegisterResponse{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *UserController) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := u.UserRepository.FindByCredentials(in.Email, in.Password)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	intRoleConv := strconv.Itoa(user.RoleID)

	return &pb.LoginResponse{
		Id:    int32(user.Id),
		Role:  intRoleConv,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *UserController) GetCustomer(ctx context.Context, in *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	userID := in.Id

	exists, err := u.UserRepository.UserExists(int(userID))
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("user not found")
	}

	user := model.User{
		Id: int(userID),
	}

	err = u.UserRepository.GetCustomer(&user)
	if err != nil {
		return nil, err
	}

	if user.Id == 0 {
		return nil, errors.New("user not found")
	}

	return &pb.GetCustomerResponse{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *UserController) UpdateCustomer(ctx context.Context, in *pb.UpdateCustomerRequest) (*pb.UpdateCustomerResponse, error) {
	userID := in.Id

	exists, err := u.UserRepository.UserExists(int(userID))
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("user not found")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Id:        int(userID),
		Name:      in.Name,
		Email:     in.Email,
		Password:  string(hashedPassword),
		UpdatedAt: in.UpdatedAt,
	}

	err = u.UserRepository.UpdateCustomer(user.Id, user)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateCustomerResponse{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *UserController) DeleteCustomer(ctx context.Context, in *pb.DeleteCustomerRequest) (*pb.Empty, error) {
	userID := in.Id

	exists, err := u.UserRepository.UserExists(int(userID))
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("user not found")
	}

	err = u.UserRepository.DeleteCustomer(int(userID))
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (u *UserController) GetCustomerAdmin(ctx context.Context, in *pb.GetCustomerAdminRequest) (*pb.GetCustomerAdminResponse, error) {
	userID := in.UserId

	exists, err := u.UserRepository.UserExists(int(userID))
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("user not found")
	}

	user, err := u.UserRepository.GetUserAdmin(int(userID))
	if err != nil {
		return nil, err
	}

	return &pb.GetCustomerAdminResponse{
		UserId: in.UserId,
		Name:   user.Name,
		Email:  user.Email,
	}, nil
}

func (u *UserController) GetAllCustomerAdmin(ctx context.Context, in *pb.Empty) (*pb.GetAllCustomerAdminResponse, error) {
	users, err := u.UserRepository.GetAllCustomerAdmin()
	if err != nil {
		return nil, err
	}

	var customerResponses []*pb.CustomerResponse
	for _, user := range users {
		userID := user.Id

		customerResponses = append(customerResponses, &pb.CustomerResponse{
			UserId: int32(userID),
			Name:   user.Name,
			Email:  user.Email,
		})
	}

	return &pb.GetAllCustomerAdminResponse{
		Customers: customerResponses,
	}, nil
}

func (u *UserController) UpdateCustomerAdmin(ctx context.Context, in *pb.UpdateCustomerAdminRequest) (*pb.UpdateCustomerAdminResponse, error) {
	userID := in.Id

	exists, err := u.UserRepository.UserExists(int(userID))
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("user not found")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Id:        int(userID),
		Name:      in.Name,
		Email:     in.Email,
		Password:  string(hashedPassword),
		UpdatedAt: in.UpdatedAt,
	}

	err = u.UserRepository.UpdateCustomer(user.Id, user)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateCustomerAdminResponse{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *UserController) DeleteCustomerAdmin(ctx context.Context, in *pb.DeleteCustomerAdminRequest) (*pb.Empty, error) {
	userID := in.Id

	exists, err := u.UserRepository.UserExists(int(userID))
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("user not found")
	}

	user := model.User{
		Id: int(userID),
	}

	err = u.UserRepository.DeleteCustomer(user.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (u *UserController) AddAddressCustomer(ctx context.Context, in *pb.AddAddressCustomerRequest) (*pb.AddAddressCustomerResponse, error) {
	userID := in.CustomerId

	exists, err := u.UserRepository.UserExists(int(userID))
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("user not found")
	}

	customer := model.User{
		Id:        int(userID),
		UpdatedAt: in.UpdatedAt,
	}

	address := model.Address{
		Country: in.Country,
		City:    in.City,
	}

	addressID, err := u.UserRepository.AddAddress(address)
	if err != nil {
		return nil, err
	}

	err = u.UserRepository.SetAddressCustomer(customer, addressID)
	if err != nil {
		return nil, err
	}

	return &pb.AddAddressCustomerResponse{
		Country: address.Country,
		City:    address.City,
	}, nil
}

func (u *UserController) UpdateAddressCustomer(ctx context.Context, in *pb.UpdateAddressCustomerRequest) (*pb.UpdateAddressCustomerResponse, error) {
	userID := in.CustomerId

	exists, err := u.UserRepository.UserExists(int(userID))
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("user not found")
	}

	customer := model.User{
		Id:        int(userID),
		UpdatedAt: in.UpdatedAt,
	}

	address := model.Address{
		Country: in.Country,
		City:    in.City,
	}

	addressID, err := u.UserRepository.GetAddressID(customer.Id)
	if err != nil {
		return nil, err
	}

	err = u.UserRepository.UpdateAddress(addressID, address)
	if err != nil {
		return nil, err
	}

	err = u.UserRepository.SetUpdatedAtAfterUpdateAddress(customer)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateAddressCustomerResponse{
		Country: address.Country,
		City:    address.City,
	}, nil
}

func (u *UserController) GetCustomerAddress(ctx context.Context, in *pb.GetCustomerAddressRequest) (*pb.GetCustomerAddressResponse, error) {
	userID := in.CustomerId

	exists, err := u.UserRepository.UserExists(int(userID))
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("user not found")
	}

	addID, err := u.UserRepository.GetAddressID(int(userID))
	if err != nil {
		return nil, err
	}

	address := model.Address{
		Id: addID,
	}

	response, err := u.UserRepository.GetAddress(&address)
	if err != nil {
		return nil, err
	}

	addCon := strconv.Itoa(response.Id)

	response2 := pb.GetCustomerAddressResponse{
		AddressId: addCon,
		Country:   response.Country,
		City:      response.City,
	}

	return &response2, nil
}
