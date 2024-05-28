package main
import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"log"
)

//  Tried to implement as much as I can without any framework like GIN, Bigo, Fiber
// Go has good support for Http and handle requests but does not provides better ways to get request parameters
// We need to make reg expressions hence we install gorilla mux

func WriteJSON(w  http.ResponseWriter, status int, v any) error{
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
} 

// ApiFunc is the function signature of function we are using 
type apiFunc func(http.ResponseWriter, *http.Request) error 

type ApiError struct{
	Error string
}

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if err:= f(w, r); err!=nil{
			WriteJSON(w, http.StatusBadRequest, ApiError{Error : err.Error()})
		}
	}
}

type ApiServer struct {
	listenAddr string 
}

func NewApiServer(listenAddr string) *ApiServer{
	return &ApiServer{
		listenAddr : listenAddr,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHttpHandleFunc(s.handleAccount))

	router.HandleFunc("/account/{id}", makeHttpHandleFunc(s.handleGetAccount))

	log.Println("JSON API server running on port : ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

// create account, get account, delete account, transfer ->
// common practice to prefix with handle (community practice)
func (s *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleTransfer(w, r)
	}

	return fmt.Errorf("Method not allowed %s", r.Method)
}

// We can also handle error inside the function, we need to convert below function into Http handler
func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	//account := NewAccount("Harshit", "Awasthi")
	// db.get(id)
	fmt.Println(id)

	return WriteJSON(w, http.StatusOK, &Account{})
}

func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Making a mock, using postgres as database