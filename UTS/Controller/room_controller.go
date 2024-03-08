package Controller

import (
	m "UTS/Model"

	// "database/sql"
	// "encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	params := mux.Vars(r)
	gameID, err := strconv.Atoi(params["id"])
	if err != nil {
		SendErrorResponse(w, 505, "Failed to convert")
		return
	}

	query := fmt.Sprintf(`SELECT r.id, r.room_name FROM rooms r WHERE r.id_game = %d`, gameID)

	rows, err := db.Query(query)
	if err != nil {
		SendErrorResponse(w, 500, "Failed to execute.")
		return
	}

	var room m.Rooms
	var rooms []m.Rooms
	for rows.Next() {
		if err := rows.Scan(&room.Id, &room.Room_name); err != nil {
			SendErrorResponse(w, 404, "Data not found")
			return
		} else {
			rooms = append(rooms, room)
		}
	}

	SendSuccessAllRoomsResponse(w, 200, "Get Room Data successfully!", rooms)
}

func GetDetailRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	param := mux.Vars(r)
	roomId, err := strconv.Atoi(param["id"])
	if err != nil {
		SendErrorResponse(w, 505, "Failed to convert")
		return
	}

	query := fmt.Sprintf(`SELECT r.id, r.room_name, p.id, p.id_account, a.username
	FROM rooms r
	JOIN participants p ON p.id_room = r.id
	JOIN accounts a ON p.id_account = a.id
	WHERE r.id = %d`, roomId)

	getDetailRoomRow, err := db.Query(query)
	if err != nil {
		SendErrorResponse(w, 500, "Failed to execute.")
		return
	}
	defer getDetailRoomRow.Close()

	var detailRooms []m.Participants
	for getDetailRoomRow.Next() {
		var room m.Rooms
		var account m.Accounts
		var participant m.Participants

		if err := getDetailRoomRow.Scan(
			&room.Id, &room.Room_name, &participant.Id, &account.Id, &account.Username); err != nil {
			SendErrorResponse(w, 500, "Couldn't get room details")
			return
		}

		participant.Id_room = room
		participant.Id_account = account

		detailRooms = append(detailRooms, participant)
	}

	SendSuccessDetailRoomsResponse(w, 200, "Get room data successfully!", detailRooms)
}

func InsertRoom(w http.ResponseWriter, r *http.Request) {

}
