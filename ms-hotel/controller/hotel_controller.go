package controller

import (
	"context"
	"ms-hotel/pb"
	"ms-hotel/repository"
)

type HotelController struct {
	pb.UnimplementedHotelServiceServer
	HotelRepository repository.HotelRepository
}

func NewHotelController(HotelRepository repository.HotelRepository) *HotelController {
	return &HotelController{
		HotelRepository: HotelRepository,
	}
}

func (h *HotelController) CreateHotel(ctx context.Context, in *pb.CreateHotelRequest) (*pb.CreateHotelResponse, error) {
	hotelID, err := h.HotelRepository.CreateHotel(in.Name, in.Location, in.Description)
	if err != nil {
		return nil, err
	}

	return &pb.CreateHotelResponse{
		Id: hotelID,
	}, nil
}

func (h *HotelController) UpdateHotel(ctx context.Context, in *pb.UpdateHotelRequest) (*pb.UpdateHotelResponse, error) {
	err := h.HotelRepository.UpdateHotel(in.Id, in.Name, in.Location, in.Description)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateHotelResponse{
		Id: in.Id,
	}, nil
}

func (h *HotelController) DeleteHotel(ctx context.Context, in *pb.DeleteHotelRequest) (*pb.DeleteHotelResponse, error) {
	err := h.HotelRepository.DeleteHotel(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteHotelResponse{
		Id: in.Id,
	}, nil
}

func (h *HotelController) GetHotel(ctx context.Context, in *pb.GetHotelRequest) (*pb.GetHotelResponse, error) {
	hotel, err := h.HotelRepository.GetHotel(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetHotelResponse{
		Hotel: &pb.Hotel{
			Id:          int32(hotel.Id),
			Name:        hotel.Name,
			Location:    hotel.Location,
			Description: hotel.Description,
			CreatedAt:   hotel.CreatedAt,
			UpdatedAt:   hotel.UpdatedAt,
		},
	}, nil
}

func (h *HotelController) ListHotels(ctx context.Context, in *pb.ListHotelsRequest) (*pb.ListHotelsResponse, error) {
	hotels, err := h.HotelRepository.ListHotels()
	if err != nil {
		return nil, err
	}

	var pbHotels []*pb.Hotel
	for _, hotel := range hotels {
		pbHotels = append(pbHotels, &pb.Hotel{
			Id:          int32(hotel.Id),
			Name:        hotel.Name,
			Location:    hotel.Location,
			Description: hotel.Description,
			CreatedAt:   hotel.CreatedAt,
			UpdatedAt:   hotel.UpdatedAt,
		})
	}

	return &pb.ListHotelsResponse{
		Hotels: pbHotels,
	}, nil
}

func (h *HotelController) CreateRoom(ctx context.Context, in *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	roomID, err := h.HotelRepository.CreateRoom(int(in.HotelId), int(in.Capacity), int(in.Price))
	if err != nil {
		return nil, err
	}

	return &pb.CreateRoomResponse{
		Id: int32(roomID),
	}, nil
}

func (h *HotelController) UpdateRoom(ctx context.Context, in *pb.UpdateRoomRequest) (*pb.UpdateRoomResponse, error) {
	err := h.HotelRepository.UpdateRoom(int(in.Id), int(in.HotelId), int(in.Capacity), int(in.Price))
	if err != nil {
		return nil, err
	}

	return &pb.UpdateRoomResponse{
		Id: in.Id,
	}, nil
}

func (h *HotelController) DeleteRoom(ctx context.Context, in *pb.DeleteRoomRequest) (*pb.DeleteRoomResponse, error) {
	err := h.HotelRepository.DeleteRoom(int(in.Id))
	if err != nil {
		return nil, err
	}

	return &pb.DeleteRoomResponse{
		Id: in.Id,
	}, nil
}

func (h *HotelController) GetRoom(ctx context.Context, in *pb.GetRoomRequest) (*pb.GetRoomResponse, error) {
	room, err := h.HotelRepository.GetRoom(int(in.Id))
	if err != nil {
		return nil, err
	}

	return &pb.GetRoomResponse{
		Room: &pb.Room{
			Id:        int32(room.Id),
			HotelId:   int32(room.HotelID),
			Capacity:  int32(room.Capacity),
			Price:     int32(room.Price),
			CreatedAt: room.CreatedAt,
			UpdatedAt: room.UpdatedAt,
		},
	}, nil
}

func (h *HotelController) ListRooms(ctx context.Context, in *pb.ListRoomsRequest) (*pb.ListRoomsResponse, error) {
	rooms, err := h.HotelRepository.ListRooms(int(in.HotelId))
	if err != nil {
		return nil, err
	}

	var pbRooms []*pb.Room
	for _, room := range rooms {
		pbRooms = append(pbRooms, &pb.Room{
			Id:        int32(room.Id),
			HotelId:   int32(room.HotelID),
			Capacity:  int32(room.Capacity),
			Price:     int32(room.Price),
			CreatedAt: room.CreatedAt,
			UpdatedAt: room.UpdatedAt,
		})
	}

	return &pb.ListRoomsResponse{
		Rooms: pbRooms,
	}, nil
}
