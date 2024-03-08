package Controller

import (
	"encoding/json"
	"net/http"

	m "UTS/Model"
)

func SendSuccessResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response m.SuccessResponse
	response.Status = code
	response.Message = message

	json.NewEncoder(w).Encode(response)
}

func SendErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response m.ErrorResponse
	response.Status = code
	response.Message = message

	json.NewEncoder(w).Encode(response)
}

func SendSuccessAllRoomsResponse(w http.ResponseWriter, code int, message string, rooms []m.Rooms) {
	w.Header().Set("Content-Type", "application/json")
	var response m.RoomsResponse
	response.Status = code
	response.Message = message
	response.Data = rooms
	json.NewEncoder(w).Encode(response)
}

func SendSuccessDetailRoomsResponse(w http.ResponseWriter, code int, message string, detailRooms []m.Participants) {
	w.Header().Set("Content-Type", "application/json")
	var response m.DetailRoomsResponse
	response.Status = code
	response.Message = message
	response.Data = detailRooms
	json.NewEncoder(w).Encode(response)
}
