package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	icalculator "gitlab.com/lmoz25/kheiron-technical-test/internal/infix-calculator"
	pcalculator "gitlab.com/lmoz25/kheiron-technical-test/internal/prefix-calculator"
)

type SumRequest struct {
	Sum string `json:"sum"`
}

type SumResponse struct {
	Answer float32 `json:"answer" validate:"required"`
}

func SetupRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/prefix", UsePrefixCalculator).Methods("POST")
	router.HandleFunc("/infix", UseInfixCalculator).Methods("POST")
	return router
}

func HandleRequests() {
	router := SetupRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}

func UsePrefixCalculator(w http.ResponseWriter, r *http.Request) {
	var calc pcalculator.PrefixCalculator

	reqBody, _ := ioutil.ReadAll(r.Body)
	var requestSum SumRequest
	json.Unmarshal(reqBody, &requestSum)

	err := calc.ParseInput(requestSum.Sum)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	result, err := calc.Result()
	if err != nil {
		return
	}

	var responseSum = SumResponse{
		Answer: result,
	}

	json.NewEncoder(w).Encode(responseSum)
}

func UseInfixCalculator(w http.ResponseWriter, r *http.Request) {
	var calc icalculator.InfixCalculator

	reqBody, _ := ioutil.ReadAll(r.Body)
	var requestSum SumRequest
	json.Unmarshal(reqBody, &requestSum)

	err := calc.ParseInput(requestSum.Sum)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	result, err := calc.Result()
	if err != nil {
		return
	}

	var responseSum = SumResponse{
		Answer: result,
	}

	json.NewEncoder(w).Encode(responseSum)
}
