package model

type Hotel struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	Description string `json:"description"`
	CreatedAt   string `json:"create_at"`
	UpdatedAt   string `json:"update_at"`
}

type Room struct {
	Id        int    `json:"id"`
	HotelID   int    `json:"hotel_id"`
	Capacity  int    `json:"capacity"`
	Price     int    `json:"price"`
	CreatedAt string `json:"create_at"`
	UpdatedAt string `json:"update_at"`
}

// for swagger
type CreateHotelRequest struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

type CreateRoomRequest struct {
	HotelID  int32 `json:"hotel_id"`
	Capacity int32 `json:"capacity"`
	Price    int32 `json:"price"`
}

type UpdateRoomRequest struct {
	ID       int32 `json:"id"`
	HotelID  int32 `json:"hotel_id"`
	Capacity int32 `json:"capacity"`
	Price    int32 `json:"price"`
}
