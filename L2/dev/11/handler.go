package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Server struct {
	Mu    sync.RWMutex
	Cache map[string]Event
}


func makeJsonResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to serialize response:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}


func validateTime(queryParam string, multiplier int) (time.Time, time.Time, error) {
	handledTime, err := strconv.Atoi(queryParam)
	if err != nil {
		log.Println("Invalid time parameter:", err)
		return time.Time{}, time.Time{}, err
	}
	timeFrom := time.Unix(0, 0).Add(time.Duration(handledTime) * time.Duration(multiplier) * 24 * time.Hour)
	timeTo := timeFrom.Add(time.Duration(multiplier) * 24 * time.Hour)
	return timeFrom, timeTo, nil
}

func (s *Server) EventByName(w http.ResponseWriter, r *http.Request) {
	eventName := r.URL.Query().Get("event_name")
	if eventName == "" {
		makeJsonResponse(w, http.StatusBadRequest, jsonError("invalid request: event_name is required"))
		return
	}
	s.Mu.RLock()
	event, ok := s.Cache[eventName]
	s.Mu.RUnlock()
	if !ok {
		makeJsonResponse(w, http.StatusNotFound, jsonError("no event for this name"))
		return
	}
	makeJsonResponse(w, http.StatusOK, event)
}

func (s *Server) EventsForDay(w http.ResponseWriter, r *http.Request) {
	day := r.URL.Query().Get("day")
	if day == "" {
		makeJsonResponse(w, http.StatusBadRequest, jsonError("invalid request: day is required"))
		return
	}
	timeFrom, timeTo, err := validateTime(day, 1)
	if err != nil {
		makeJsonResponse(w, http.StatusBadRequest, jsonError("invalid day parameter"))
		return
	}
	var result []Event
	s.Mu.RLock()
	for _, event := range s.Cache {
		if inTimeSpan(timeFrom, timeTo, event.Time) {
			result = append(result, event)
		}
	}
	s.Mu.RUnlock()
	makeJsonResponse(w, http.StatusOK, result)
}



func jsonError(message string) map[string]string {
	return map[string]string{"error": message}
}


func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func (s *Server) EventByName(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	eventName, ok := m["event_name"]
	if !ok {
		makeJsonRespond(w, 400, jsonError("invalid request"))
		return

	}
	s.Mu.RLock()
	event, ok := s.Cache[eventName[0]]
	s.Mu.RUnlock()
	if !ok {
		makeJsonRespond(w, 500, jsonError("no event for this name"))
		return
	}
	data, err := json.Marshal(event)
	if err != nil {
		log.Println(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	makeJsonRespond(w, 200, jsonResult(string(data)))
}

func (s *Server) EventsForDay(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	day, ok := m["day"]
	if !ok {
		makeJsonRespond(w, 400, jsonError("invalid request"))
		return
	}
	handeledTime, err := strconv.Atoi(day[0])
	if err != nil {
		log.Println(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	result := make([]Event, 0)
	timeFrom := time.Unix(0, 0).Add(time.Duration(handeledTime) * 24 * time.Hour)
	timeTo := timeFrom.AddDate(0, 0, 1)
	s.Mu.RLock()
	for _, event := range s.Cache {
		if inTimeSpan(timeFrom, timeTo, event.Time) {
			result = append(result, event)
		}
	}
	s.Mu.RUnlock()
	data, err := json.Marshal(result)
	if err != nil {
		log.Panicln(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	makeJsonRespond(w, 200, jsonResult(string(data)))
}

func (s *Server) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	week, ok := m["week"]
	if !ok {
		makeJsonRespond(w, 400, jsonError("invalid request"))
		return
	}
	handeledTime, err := strconv.Atoi(week[0])
	if err != nil {
		log.Println(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	result := make([]Event, 0)
	timeFrom := time.Unix(0, 0).Add(time.Duration(handeledTime) * 24 * time.Hour * 7)
	timeTo := timeFrom.AddDate(0, 0, 7)
	s.Mu.RLock()
	for _, event := range s.Cache {
		if inTimeSpan(timeFrom, timeTo, event.Time) {
			result = append(result, event)
		}
	}
	s.Mu.RUnlock()
	data, err := json.Marshal(result)
	if err != nil {
		log.Panicln(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	makeJsonRespond(w, 200, jsonResult(string(data)))
}

func (s *Server) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	month, ok := m["month"]
	if !ok {
		makeJsonRespond(w, 400, jsonError("invalid request"))
		return
	}
	handeledTime, err := strconv.Atoi(month[0])
	if err != nil {
		log.Println(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	result := make([]Event, 0)
	s.Mu.RLock()
	for _, event := range s.Cache {
		if event.Time.Month() == time.Month(handeledTime) {
			result = append(result, event)
		}
	}
	s.Mu.RUnlock()
	data, err := json.Marshal(result)
	if err != nil {
		log.Panicln(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	makeJsonRespond(w, 200, jsonResult(string(data)))
}

