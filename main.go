package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Struct for storing time records from the database
type TimeRecord struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

// Initialize the database connection
func initDB() {
	var err error
	dsn := "root:rishwanth@17@tcp(127.0.0.1:3306)/time_api" // Update with your MySQL credentials
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening MySQL database: ", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging MySQL database: ", err)
	}
	fmt.Println("Connected to MySQL database")
}

// Handler for the /current-time endpoint
func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the current time in UTC
	currentTime := time.Now().UTC()

	// Convert UTC to Toronto time zone (EST/EDT)
	torontoTime := currentTime.In(time.FixedZone("EST", -5*60*60)) // For EST (UTC-5)

	// Insert the time into the database (in Toronto local time)
	_, err := db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", torontoTime)
	if err != nil {
		http.Error(w, "Error inserting time into database", http.StatusInternalServerError)
		return
	}

	// Create a response with the current Toronto time in JSON format
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"current_time": torontoTime.Format(time.RFC3339),
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

// Handler for the /all-times endpoint
func allTimesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, timestamp FROM time_log")
	if err != nil {
		http.Error(w, "Error retrieving times from database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var times []TimeRecord
	for rows.Next() {
		var t TimeRecord
		var timestamp []byte // Scan the timestamp as a byte slice
		if err := rows.Scan(&t.ID, &timestamp); err != nil {
			log.Printf("Error scanning time record: %v", err)
			http.Error(w, "Error scanning time record", http.StatusInternalServerError)
			return
		}

		// Convert the byte slice to time.Time
		t.Timestamp, err = time.Parse("2006-01-02 15:04:05", string(timestamp))
		if err != nil {
			log.Printf("Error parsing timestamp: %v", err)
			http.Error(w, "Error parsing timestamp", http.StatusInternalServerError)
			return
		}

		// Convert the UTC timestamp to Toronto time zone
		torontoTime := t.Timestamp.In(time.FixedZone("EST", -5*60*60)) // For EST (UTC-5)
		t.Timestamp = torontoTime

		times = append(times, t)
	}

	// Respond with the times in JSON format
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"times": times}); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func main() {
	// Initialize database connection
	initDB()
	defer db.Close()

	// Setup routes
	http.HandleFunc("/current-time", currentTimeHandler)
	http.HandleFunc("/all-times", allTimesHandler)

	// Start the server
	port := ":8080"
	fmt.Printf("Starting server on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
