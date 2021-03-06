package controllers

import (
	"encoding/json"
	"fmt"
	"fs5-gestiondesocios-backend/api/models"
	"fs5-gestiondesocios-backend/api/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//TODO Create generic controllers

//GetGame retrieves a game by id
func GetGame(writer http.ResponseWriter, r *http.Request) {
	game := models.Game{}
	id := mux.Vars(r)["id"]
	if !initDB {
		db = utils.GetConnection()
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error clossing the DB")
	} else {
		defer sqlDB.Close()
	}
	db.Find(&game, id)
	if game.ID > 0 {
		jGame, _ := json.Marshal(game)
		utils.SendResponse(writer, http.StatusOK, jGame)
	} else {
		utils.SendError(writer, http.StatusNotFound, "Game not found")
	}

}

//GetGames retrieves a game by id
func GetGames(writer http.ResponseWriter, r *http.Request) {
	games := []models.Game{}

	if !initDB {
		db = utils.GetConnection()
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error clossing the DB")
	} else {
		defer sqlDB.Close()
	}
	db.Find(&games)
	if len(games) > 0 {
		jGames, _ := json.Marshal(games)
		utils.SendResponse(writer, http.StatusOK, jGames)
	} else {
		utils.SendError(writer, http.StatusNotFound, "Games not found")
	}

}

//InsertGame inserts a game into the "games" table
func InsertGame(writer http.ResponseWriter, r *http.Request) {
	game := models.Game{}
	if !initDB {
		db = utils.GetConnection()
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error clossing the DB")
	} else {
		defer sqlDB.Close()
	}
	errDB := json.NewDecoder(r.Body).Decode(&game)
	if errDB != nil {
		fmt.Println(errDB)
		utils.SendError(writer, http.StatusBadRequest, "Error formating game")
		return
	}

	errDB = db.Create(&game).Error
	if errDB != nil {
		fmt.Println(errDB)
		utils.SendError(writer, http.StatusBadRequest, "Error inserting game")
		return
	}

	j, _ := json.Marshal(game)
	utils.SendResponse(writer, http.StatusCreated, j)
}

//DeleteGame deletes a game given the ID
func DeleteGame(writer http.ResponseWriter, r *http.Request) {
	game := models.Game{}
	id := mux.Vars(r)["id"]
	if !initDB {
		db = utils.GetConnection()
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error clossing the DB")
	} else {
		defer sqlDB.Close()
	}
	db.Find(&game, id)
	if game.ID > 0 {
		db.Delete(game)
		utils.SendResponse(writer, http.StatusOK, []byte(`Game removed successfully`))
	} else {
		utils.SendError(writer, http.StatusNotFound, "Error removing the game, might not exist")
	}
}

//UpdateGame updates a game given the ID
func UpdateGame(writer http.ResponseWriter, r *http.Request) {
	gameFind := models.Game{}
	gameNewData := models.Game{}

	id := mux.Vars(r)["id"]
	if !initDB {
		db = utils.GetConnection()
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error clossing the DB")
	} else {
		defer sqlDB.Close()
	}
	db.Find(&gameFind, id)

	if gameFind.ID > 0 {
		err := json.NewDecoder(r.Body).Decode(&gameNewData)
		if err != nil {
			utils.SendError(writer, http.StatusBadRequest, "Invalid data")
			return
		}
		db.Model(&gameFind).Updates(gameNewData)
		j, _ := json.Marshal(gameFind)
		utils.SendResponse(writer, http.StatusOK, j)
	} else {
		utils.SendError(writer, http.StatusNotFound, "Game not found")
	}
}
