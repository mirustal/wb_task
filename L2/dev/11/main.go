package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Event struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Title  string    `json:"title"`
	Detail string    `json:"detail"`
	Date   time.Time `json:"date"`
}



func SerializeEvent(event Event) ([]byte, error) {
	return json.Marshal(event)
}

func DeserializeEvent(data []byte) (Event, error) {
	var event Event
	err := json.Unmarshal(data, &event)
	return event, err
}


func ParseQueryParams(r *http.Request) (int, time.Time, error) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		return 0, time.Time{}, err
	}
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		return 0, time.Time{}, err
	}
	return userID, date, nil
}


func createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var event Event
	err := decoder.Decode(&event)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := createEventInStore(event); err != nil {
		http.Error(w, `{"error":"business logic failure"}`, http.StatusServiceUnavailable)
		return
	}
	resp, _ := json.Marshal(map[string]string{"result": "event created"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received %s request for %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}


func main() {
	http.HandleFunc("/create_event", createEvent)
	http.HandleFunc("/update_event", updateEvent)
	http.HandleFunc("/delete_event", deleteEvent)
	http.HandleFunc("/events_for_day", eventsForDay)
	http.HandleFunc("/events_for_week", eventsForWeek)
	http.HandleFunc("/events_for_month", eventsForMonth)

	// Оборачиваем все хэндлеры логгирующим middleware
	http.Handle("/", logRequest(http.DefaultServeMux))

	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
