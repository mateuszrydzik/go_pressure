package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PressureData struct {
	Station  string `json:"stacja"`
	Pressure string `json:"cisnienie"`
	Date     string `json:"data_pomiaru"`
	Hour     string `json:"godzina_pomiaru"`
}

func main() {
	// loading .env variables
	path, err := filepath.Abs("../../.env")
	if err != nil {
		fmt.Println("Cannot access .env:", err)
		os.Exit(1)
	}
	
	err = godotenv.Load(path)
	if err != nil {
		fmt.Println("Cannot load .env:", err)
		os.Exit(1)
	}
	
	// downloading imgw data
 	station_id := os.Getenv("STATION_ID")
	url := fmt.Sprintf("https://danepubliczne.imgw.pl/api/data/synop/id/%s", station_id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error when sending HTTP request:", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Serwer did not respond with 200 OK:", resp.Status)
		return
	}

	// changing json to PressureData type
	respData, errResp := io.ReadAll(resp.Body)
	if errResp != nil {
		fmt.Println("Cannot read response", errResp)
		return
	}

	var pressure PressureData
	errPressure := json.Unmarshal(respData, &pressure)
	if errPressure != nil {
		fmt.Println("Cannot unmarshal JSON:", errPressure)
		return
	}

	// converting string to float
	pressureFloat, errConv := strconv.ParseFloat(pressure.Pressure, 64)
	if errConv != nil {
		fmt.Println("Cannot converte pressure from string to float:", errConv)
		return
	}

	pressure.Pressure = fmt.Sprintf("%.2f", pressureFloat)

	// connecting to the database
	dbuser := os.Getenv("DBUSER")
	dbpassword := os.Getenv("DBPWD")
	dbname := os.Getenv("DBNAME")
	dbhost := os.Getenv("DBHOST")
	dbport, err := strconv.Atoi(os.Getenv("DBPORT"))

	if err != nil {
		fmt.Println("Cannot converte db_port from string to integer", err)
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

	// sending data to database
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
