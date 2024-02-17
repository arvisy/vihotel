package controller

import (
	"context"
	"ms-gateway/helper"
	"ms-gateway/model"
	"ms-gateway/pb"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HotelController struct {
	hotelGRPC pb.HotelServiceClient
}

func NewHotelController(hotelGRPC pb.HotelServiceClient) *HotelController {
	return &HotelController{
		hotelGRPC: hotelGRPC,
	}
}

// @Summary      Private
// @Description  Create Hotel
// @Tags         Hotel
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param		 data body model.CreateHotelRequest true "The input user struct"
// @Success      201  {object}  helper.Message
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.MessageDetails
// @Router       /api/v1/hotel [Post]
func (h *HotelController) CreateHotel(c echo.Context) error {
	var payload model.Hotel

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid request payload",
		})
	}

	if payload.Name == "" || payload.Location == "" || payload.Description == "" {
		return c.JSON(400, helper.Response{
			Message: "name, location, and description are required",
		})
	}

	in := &pb.CreateHotelRequest{
		Name:        payload.Name,
		Location:    payload.Location,
		Description: payload.Description,
	}

	response, err := h.hotelGRPC.CreateHotel(context.TODO(), in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to create hotel",
			Detail:  err.Error(),
		})
	}

	return c.JSON(201, response)
}

// @Summary      Private
// @Description  Update Hotel
// @Tags         Hotel
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param ID path int true "Hotel ID"
// @Param		 data body model.CreateHotelRequest true "The input user struct"
// @Success      201  {object}  helper.Message
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.MessageDetails
// @Router       /api/v1/hotel/:id [Put]
func (h *HotelController) UpdateHotel(c echo.Context) error {
	var payload model.Hotel

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid request payload",
		})
	}

	if payload.Id == 0 {
		return c.JSON(400, helper.Response{
			Message: "hotel ID is required",
		})
	}

	in := &pb.UpdateHotelRequest{
		Id:          int32(payload.Id),
		Name:        payload.Name,
		Location:    payload.Location,
		Description: payload.Description,
	}

	response, err := h.hotelGRPC.UpdateHotel(context.TODO(), in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to update hotel",
			Detail:  err.Error(),
		})
	}

	return c.JSON(200, response)
}

// @Summary      Private
// @Description  Delete Hotel
// @Tags         Hotel
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param ID path int true "Hotel ID"
// @Success      201  {object}  helper.Message
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.MessageDetails
// @Router       /api/v1/hotel/:id [Delete]
func (h *HotelController) DeleteHotel(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(400, helper.Response{
			Message: "hotel ID is required",
		})
	}

	idConv, _ := strconv.Atoi(id)

	in := &pb.DeleteHotelRequest{
		Id: int32(idConv),
	}

	response, err := h.hotelGRPC.DeleteHotel(context.TODO(), in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to delete hotel",
			Detail:  err.Error(),
		})
	}

	return c.JSON(200, response)
}

// @Summary      Public
// @Description  Get info Hotel
// @Tags         User
// @Accept       json
// @Produce      json
// @Param ID path int true "Hotel ID"
// @Success      201  {object}  model.Hotel
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.MessageDetails
// @Router       /hotel/:id [Get]
func (h *HotelController) GetHotel(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(400, helper.Response{
			Message: "hotel ID is required",
		})
	}

	idConv, _ := strconv.Atoi(id)

	in := &pb.GetHotelRequest{
		Id: int32(idConv),
	}

	response, err := h.hotelGRPC.GetHotel(context.TODO(), in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to get hotel details",
			Detail:  err.Error(),
		})
	}

	return c.JSON(200, response)
}

// @Summary      Public
// @Description  Get all Hotel
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      201  {object}  model.Hotel
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.MessageDetails
// @Router       /hotel [Get]
func (h *HotelController) ListHotels(c echo.Context) error {
	in := &pb.ListHotelsRequest{}

	response, err := h.hotelGRPC.ListHotels(context.TODO(), in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to list hotels",
			Detail:  err.Error(),
		})
	}

	return c.JSON(200, response)
}

// @Summary      Private
// @Description  Create room
// @Tags         Hotel
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param		 Data body model.CreateRoomRequest true "The input user struct"
// @Success      201  {object}  helper.Message
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.MessageDetails
// @Router       /api/v1/hotel/room [Post]
func (h *HotelController) CreateRoom(c echo.Context) error {
	var payload model.Room

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid request payload",
		})
	}

	if payload.HotelID == 0 || payload.Capacity == 0 || payload.Price == 0 {
		return c.JSON(400, helper.Response{
			Message: "hotel ID, capacity, and price are required",
		})
	}

	in := &pb.CreateRoomRequest{
		HotelId:  int32(payload.HotelID),
		Capacity: int32(payload.Capacity),
		Price:    int32(payload.Price),
	}

	response, err := h.hotelGRPC.CreateRoom(context.TODO(), in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to create room",
			Detail:  err.Error(),
		})
	}

	return c.JSON(201, response)
}

// @Summary      Private
// @Description  Update room
// @Tags         Hotel
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param		 Data body model.UpdateRoomRequest true "The input user struct"
// @Success      201  {object}  helper.Message
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.MessageDetails
// @Router       /api/v1/hotel/room [Put]
func (h *HotelController) UpdateRoom(c echo.Context) error {
	var payload model.Room

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid request payload",
		})
	}

	if payload.Id == 0 || payload.HotelID == 0 || payload.Capacity == 0 || payload.Price == 0 {
		return c.JSON(400, helper.Response{
			Message: "room ID, hotel ID, type, capacity, and price are required",
		})
	}

	in := &pb.UpdateRoomRequest{
		Id:       int32(payload.Id),
		HotelId:  int32(payload.HotelID),
		Capacity: int32(payload.Capacity),
		Price:    int32(payload.Price),
	}

	response, err := h.hotelGRPC.UpdateRoom(context.TODO(), in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to update room",
			Detail:  err.Error(),
		})
	}

	return c.JSON(200, response)
}

// @Summary      Private
// @Description  Delete room
// @Tags         Hotel
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param ID path int true "Room ID"
// @Success      201  {object}  helper.Message
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.MessageDetails
// @Router       /api/v1/hotel/room/:id [Delete]
func (h *HotelController) DeleteRoom(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(400, helper.Response{
			Message: "room ID is required",
		})
	}

	idConv, _ := strconv.Atoi(id)

	in := &pb.DeleteRoomRequest{
		Id: int32(idConv),
	}

	response, err := h.hotelGRPC.DeleteRoom(context.TODO(), in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to delete room",
			Detail:  err.Error(),
		})
	}

	return c.JSON(200, response)
}

// @Summary      Public
// @Description  Get info room
// @Tags         User
// @Accept       json
// @Produce      json
// @Param ID path int true "Room ID"
// @Success      201  {object}  model.Room
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.MessageDetails
// @Router       /hotel/room/:id [Get]
func (h *HotelController) GetRoom(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(400, helper.Response{
			Message: "room ID is required",
		})
	}

	idConv, _ := strconv.Atoi(id)

	in := &pb.GetRoomRequest{
		Id: int32(idConv),
	}

	response, err := h.hotelGRPC.GetRoom(context.TODO(), in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to get room details",
			Detail:  err.Error(),
		})
	}

	return c.JSON(200, response)
}

// @Summary      Public
// @Description  Get all room
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      201  {object}  model.Room
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.MessageDetails
// @Router       /hotel/room [Get]
func (h *HotelController) ListRooms(c echo.Context) error {
	hotelID := c.QueryParam("hotel_id")

	hotelIDConv, _ := strconv.Atoi(hotelID)

	in := &pb.ListRoomsRequest{
		HotelId: int32(hotelIDConv),
	}

	response, err := h.hotelGRPC.ListRooms(context.TODO(), in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to list rooms",
			Detail:  err.Error(),
		})
	}

	return c.JSON(200, response)
}
