package controller

import (
	"context"
	"ms-gateway/helper"
	"ms-gateway/model"
	"ms-gateway/pb"
	"regexp"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type UserContoller struct {
	userGRPC pb.UserServiceClient
}

func NewUserController(userGRPC pb.UserServiceClient) *UserContoller {
	return &UserContoller{
		userGRPC: userGRPC,
	}
}

func (u *UserContoller) Register(c echo.Context) error {
	var payload model.User

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid request payload",
		})
	}

	if payload.Name == "" || payload.Email == "" || payload.Password == "" {
		return c.JSON(400, helper.Response{
			Message: "name, email, and password are required",
		})
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(payload.Email) {
		return c.JSON(400, helper.Response{
			Message: "invalid email format",
		})
	}

	in := pb.RegisterRequest{
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  payload.Password,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	response, err := u.userGRPC.Register(context.TODO(), &in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to register user",
			Detail:  err.Error(),
		})
	}

	return c.JSON(201, response)
}

func (u *UserContoller) Login(c echo.Context) error {
	var loginRequest model.User
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid login request payload",
		})
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		return c.JSON(400, helper.Response{
			Message: "email and password are required",
		})
	}

	response, err := u.userGRPC.Login(context.TODO(), &pb.LoginRequest{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	})

	if err != nil {
		return c.JSON(401, helper.Response{
			Message: "invalid email or password",
		})
	}

	userIDClaims := strconv.Itoa(int(response.Id))

	claims := jwt.MapClaims{
		"id":   userIDClaims,
		"role": response.Role,
	}

	token, err := helper.GenerateJWTTokenWithClaims(claims)
	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to generate JWT token",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "login success",
		"token":   token,
	})
}

func (u *UserContoller) GetInfoCustomer(c echo.Context) error {
	userIDConv, ok := c.Get("id").(string)
	if !ok {
		return c.JSON(400, helper.Response{
			Message: "invalid user id",
		})
	}

	userID, err := strconv.Atoi(userIDConv)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid user id",
		})
	}

	response1, err := u.userGRPC.GetCustomer(context.TODO(), &pb.GetCustomerRequest{
		Id: int32(userID),
	})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to get user information",
		})
	}

	response2, err := u.userGRPC.GetCustomerAddress(context.TODO(), &pb.GetCustomerAddressRequest{
		CustomerId: int32(userID),
	})

	return c.JSON(200, echo.Map{
		"user":    response1,
		"address": response2,
	})
}

func (u *UserContoller) UpdateCustomer(c echo.Context) error {
	userIDConv, ok := c.Get("id").(string)
	if !ok {
		return c.JSON(400, helper.Response{
			Message: "invalid user id",
		})
	}

	userID, err := strconv.Atoi(userIDConv)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid user id",
		})
	}

	var updateRequest pb.UpdateCustomerRequest
	if err := c.Bind(&updateRequest); err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid update request payload",
		})
	}

	if updateRequest.Email != "" {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(updateRequest.Email) {
			return c.JSON(400, helper.Response{
				Message: "invalid email format",
			})
		}
	}

	response, err := u.userGRPC.UpdateCustomer(context.TODO(), &pb.UpdateCustomerRequest{
		Id:        int32(userID),
		Name:      updateRequest.Name,
		Email:     updateRequest.Email,
		Password:  updateRequest.Password,
		UpdatedAt: time.Now().Format(time.RFC3339),
	})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to update user",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "customer updated successfully",
		"user":    response,
	})
}

func (u *UserContoller) DeleteCustomer(c echo.Context) error {
	userIDConv, ok := c.Get("id").(string)
	if !ok {
		return c.JSON(400, helper.Response{
			Message: "invalid user id",
		})
	}

	userID, err := strconv.Atoi(userIDConv)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid user id",
		})
	}

	response, err := u.userGRPC.DeleteCustomer(context.TODO(), &pb.DeleteCustomerRequest{
		Id: int32(userID),
	})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to delete user",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "user deleted successfully",
		"user":    response,
	})
}

func (u *UserContoller) GetCustomerAdmin(c echo.Context) error {
	userIDConv := c.Param("id")

	userID, err := strconv.Atoi(userIDConv)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid user id",
		})
	}

	response, err := u.userGRPC.GetCustomerAdmin(context.TODO(), &pb.GetCustomerAdminRequest{
		UserId: int32(userID),
	})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "user not found",
		})
	}

	return c.JSON(200, helper.Response{
		Message: "successfully get user info",
		Detail:  response,
	})
}

func (u *UserContoller) GetAllCustomerAdmin(c echo.Context) error {
	response, err := u.userGRPC.GetAllCustomerAdmin(context.TODO(), &pb.Empty{})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to get all user",
		})
	}

	return c.JSON(200, helper.Response{
		Message: "successfully get all user info",
		Detail:  response,
	})
}

func (u *UserContoller) UpdateCustomerAdmin(c echo.Context) error {
	userIDConv := c.Param("id")

	userID, err := strconv.Atoi(userIDConv)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid user id",
		})
	}

	var updateRequest pb.UpdateCustomerRequest
	if err := c.Bind(&updateRequest); err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid update request payload",
		})
	}

	response, err := u.userGRPC.UpdateCustomer(context.TODO(), &pb.UpdateCustomerRequest{
		Id:        int32(userID),
		Name:      updateRequest.Name,
		Email:     updateRequest.Email,
		Password:  updateRequest.Password,
		UpdatedAt: time.Now().Format(time.RFC3339),
	})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to update user",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "customer updated successfully",
		"user":    response,
	})
}

func (u *UserContoller) DeleteCustomerAdmin(c echo.Context) error {
	userIDConv := c.Param("id")

	userID, err := strconv.Atoi(userIDConv)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid user id",
		})
	}

	response, err := u.userGRPC.DeleteCustomer(context.TODO(), &pb.DeleteCustomerRequest{
		Id: int32(userID),
	})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to delete user",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "user deleted successfully",
		"user":    response,
	})
}

func (u *UserContoller) AddAddress(c echo.Context) error {
	userIDConv, ok := c.Get("id").(string)
	if !ok {
		return c.JSON(400, helper.Response{
			Message: "invalid user id",
		})
	}

	userID, err := strconv.Atoi(userIDConv)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid user id",
		})
	}

	_, err = u.userGRPC.GetCustomerAddress(context.TODO(), &pb.GetCustomerAddressRequest{
		CustomerId: int32(userID),
	})
	if err == nil {
		return c.JSON(400, helper.Response{
			Message: "address already exist",
		})
	}

	var addAddressRequest pb.AddAddressCustomerRequest
	if err := c.Bind(&addAddressRequest); err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid update request payload",
		})
	}

	if addAddressRequest.Country == "" || addAddressRequest.City == "" {
		return c.JSON(400, helper.Response{
			Message: "country and city are required",
		})
	}

	response, err := u.userGRPC.AddAddressCustomer(context.TODO(), &pb.AddAddressCustomerRequest{
		CustomerId: int32(userID),
		Country:    addAddressRequest.Country,
		City:       addAddressRequest.City,
		UpdatedAt:  time.Now().Format(time.RFC3339),
	})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to add address",
			Detail:  err.Error(),
		})
	}

	return c.JSON(201, echo.Map{
		"message": "address created successfully",
		"address": response,
	})
}

func (u *UserContoller) UpdateAddress(c echo.Context) error {
	userIDConv, ok := c.Get("id").(string)
	if !ok {
		return c.JSON(400, helper.Response{
			Message: "invalid user id",
		})
	}

	userID, err := strconv.Atoi(userIDConv)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid user id",
		})
	}

	var updateRequest pb.UpdateAddressCustomerRequest
	if err := c.Bind(&updateRequest); err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid update request payload",
		})
	}

	if updateRequest.Country == "" || updateRequest.City == "" {
		return c.JSON(400, helper.Response{
			Message: "country and city are required",
		})
	}

	response, err := u.userGRPC.UpdateAddressCustomer(context.TODO(), &pb.UpdateAddressCustomerRequest{
		CustomerId: int32(userID),
		Country:    updateRequest.Country,
		City:       updateRequest.City,
		UpdatedAt:  time.Now().Format(time.RFC3339),
	})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to update address",
			Detail:  err.Error(),
		})
	}

	return c.JSON(200, helper.Response{
		Message: "address updated successfully",
		Detail:  response,
	})
}
