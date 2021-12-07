package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/quadrosh/user-list/user"
)

// Handler is interface of API handler
type Handler interface {
	Create(http.ResponseWriter, *http.Request)
	Find(http.ResponseWriter, *http.Request)
	Filter(http.ResponseWriter, *http.Request)
}

type handler struct {
	userService user.ServiceInterface
}

// NewHandler returns the type, which implements Handler interface
func NewHandler(service user.ServiceInterface) Handler {
	return &handler{userService: service}
}

// Greate handles create request and calls the Create function of user service interface
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var user user.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(err)
		return
	}
	err = h.userService.Create(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	log.Println("Successfully added")

}

// propValueRequest is type of property/value request
type propValueRequest struct {
	Property string      `json:"property"`
	Value    interface{} `json:"value"`
}

// Find handles request to /find address, calls Find function of user service interface
func (h *handler) Find(w http.ResponseWriter, r *http.Request) {
	var req propValueRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(err)
		return
	}

	users, err := h.userService.Find(req.Property, req.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	responseBody, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	setupResponse(w, responseBody, http.StatusOK)
}

type filterRequest struct {
	DateFrom int `json:"recording_date_from"`
	DateTo   int `json:"recording_date_to"`
	AgeFrom  int `json:"age_from"`
	AgeTo    int `json:"age_to"`
}
type filterResponse struct {
	Users []user.User `json:"users"`
	Sum   int         `json:"sum"`
}

// Filter handles reguest for users by specific age and recording date diapazons, calls Filter function of user service interface
func (h *handler) Filter(w http.ResponseWriter, r *http.Request) {
	var req filterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(err)
		return
	}

	users, length, err := h.userService.Filter(req.DateFrom, req.DateTo, req.AgeFrom, req.AgeTo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	res := filterResponse{
		Users: users,
		Sum:   length,
	}
	responseBody, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	setupResponse(w, responseBody, http.StatusOK)
}

func setupResponse(w http.ResponseWriter, body []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}
