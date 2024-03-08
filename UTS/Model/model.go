package Model

type Accounts struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
type Games struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Max_player int    `json:"max_player"`
}
type Rooms struct {
	Id        int    `json:"id"`
	Room_name string `json:"room_name"`
	Id_game   Games
}
type Participants struct {
	Id         int   `json:"id"`
	Id_room    Rooms `json:"Room"`
	Id_account Accounts
}

type DetailRooms struct {
	Id      int   `json:"id"`
	Id_room Rooms `json:"Room"`
}

//response
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Rooms  `json:"data"`
}
type RoomsResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Rooms `json:"data"`
}
type DetailRoomsResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    []Participants `json:"data"`
}
