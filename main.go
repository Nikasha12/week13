package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	TorontoTime string `json:"toronto_time"`
}

func main() {
	http.HandleFunc("/current-time", currentTimeHandler)
	http.ListenAndServe(":8080", nil)
}

func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	torontoTime := getCurrentTorontoTime()
	err := saveTimeToDatabase(torontoTime)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := TimeResponse{TorontoTime: torontoTime.Format(time.RFC3339)}
	json.NewEncoder(w).Encode(response)
}

func getCurrentTorontoTime() time.Time {
	loc, _ := time.LoadLocation("America/Toronto")
	return time.Now().In(loc)
}

func saveTimeToDatabase(time time.Time) error {
	db, err := sql.Open("goracle", "SYS AS SYSDBA/oracle@localhost:1521/ORCLPDB")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (:1)", time)
	return err
}
