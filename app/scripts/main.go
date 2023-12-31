package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"io"
	"net/http"
	"os"
	"strconv"
)

type PressureData struct {
	Station  string `json:"stacja"`
	Pressure string `json:"cisnienie"`
	Date     string `json:"data_pomiaru"`
	Hour     string `json:"godzina_pomiaru"`
}

func main() {
	// wczytanie zmiennych z pliku .env
	godotenv.Load(".env")

	// pobranie danych imgw
	station_id := os.Getenv("STATION_ID")
	url := fmt.Sprintf("https://danepubliczne.imgw.pl/api/data/synop/id/%s", station_id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Błąd podczas wykonywania żądania HTTP:", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Serwer zwrócił status inny niż 200 OK:", resp.Status)
		return
	}

	// zamiana json na PressureData
	respData, errResp := io.ReadAll(resp.Body)
	if errResp != nil {
		fmt.Println("Błąd podczas odczytywania ciała odpowiedzi:", errResp)
		return
	}

	var pressure PressureData

	errPressure := json.Unmarshal(respData, &pressure)
	if errPressure != nil {
		fmt.Println("Błąd podczas unmarshalowania JSON:", errPressure)
		return
	}

	// konwersja stringa na float64
	pressureFloat, errConv := strconv.ParseFloat(pressure.Pressure, 64)
	if errConv != nil {
		fmt.Println("Błąd podczas konwersji ciśnienia:", errConv)
		return
	}

	pressure.Pressure = fmt.Sprintf("%.2f", pressureFloat) // Formatujemy float do dwóch miejsc po przecinku

	// połączenie z bazą
	dbuser := os.Getenv("DBUSER")
	dbpassword := os.Getenv("DBPWD")
	dbname := os.Getenv("DBNAME")
	dbhost := os.Getenv("DBHOST")
	dbport, err := strconv.Atoi(os.Getenv("DBPORT"))

	if err != nil {
		fmt.Println("Błąd podczas konwersji portu:", err)
		return
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbuser, dbpassword, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// dodanie danych do bazy
	sqlStatement := `
	INSERT INTO pressure (station, pressure, date, hour)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0

	err = db.QueryRow(sqlStatement, pressure.Station, pressure.Pressure, pressure.Date, pressure.Hour).Scan(&id)
	if err != nil {
		panic(err)
	}
}
