package models

//TODO search info about using dates on structs

type date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

//Game is the struc of the table called "games"
type Game struct {
	ID            uint64  `json:"id" gorm:"primaryKey"`
	GameName      string  `json:"game_name"`
	GameImage     string  `json:"game_image"`
	Rating        float64 `json:"rating"`
	IDOwner       uint    `json:"id_owner"`
	EntryDate     string  `json:"entry_date"`
	Disponibility bool    `json:"disponibility"`
	Comments      string  `json:"comments" gorm:"size:200"`
}
