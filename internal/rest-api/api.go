package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	icalculator "gitlab.com/lmoz25/kheiron-technical-test/internal/infix-calculator"
	pcalculator "gitlab.com/lmoz25/kheiron-technical-test/internal/prefix-calculator"
)

// SumRequest is the struct for making sum requests to the API, for both calculators
type SumRequest struct {
	Sum string `json:"sum"`
}

// SumResponse is the struct in which responses to sum requests will return, for both calculators
type SumResponse struct {
	Answer float32 `json:"answer" validate:"required"`
}

// SetupRouter sets up the router for handling requests to both calculators
func SetupRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/prefix", UsePrefixCalculator).Methods("POST")
	router.HandleFunc("/infix", UseInfixCalculator).Methods("POST")
	return router
}

// HandleRequests handles requests to the two calculators using a mux router
func HandleRequests() {
	router := SetupRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}

// UsePrefixCalculator handles requests to the /prefix endpoint
func UsePrefixCalculator(w http.ResponseWriter, r *http.Request) {
	var calc pcalculator.PrefixCalculator

	reqBody, _ := ioutil.ReadAll(r.Body)
	var requestSum SumRequest
	json.Unmarshal(reqBody, &requestSum)

	err := calc.ParseInput(requestSum.Sum)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	result, err := calc.Result()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var responseSum = SumResponse{
		Answer: result,
	}

	json.NewEncoder(w).Encode(responseSum)
}

// UseInfixCalculator handles requests to the /infix endpoint
func UseInfixCalculator(w http.ResponseWriter, r *http.Request) {
	var calc icalculator.InfixCalculator

	reqBody, _ := ioutil.ReadAll(r.Body)
	var requestSum SumRequest
	json.Unmarshal(reqBody, &requestSum)

	err := calc.ParseInput(requestSum.Sum)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	result, err := calc.Result()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var responseSum = SumResponse{
		Answer: result,
	}

	json.NewEncoder(w).Encode(responseSum)
}
