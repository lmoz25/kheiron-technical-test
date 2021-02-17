package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	api "gitlab.com/lmoz25/kheiron-technical-test/internal/rest-api"
)

// type APITests struct {
// 	httpExpect *httpexpect.Expect
// }

// func (s *APITests) Setup(t *testing.T) {
// 	router := api.SetupRouter()
// 	s.httpExpect = httpexpect.
// }

func testMain(t *testing.T, request *api.SumRequest, expectedResponse *api.SumResponse, endpoint string, calculator http.HandlerFunc) {
	requestBody, err := json.Marshal(request)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(calculator)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response = api.SumResponse{}

	responseBody, _ := ioutil.ReadAll(rr.Body)
	json.Unmarshal(responseBody, &response)
	if response != *expectedResponse {
		t.Errorf("handler returned unexpected body: got %f want %f",
			response.Answer, expectedResponse.Answer)
	}
}

func TestInfixAPI(t *testing.T) {
	for _, tc := range InfixAPITestData {
		testName := fmt.Sprintf("Infix Calculator API: %s", tc.TestDescription)
		t.Run(testName, func(t *testing.T) {
			testMain(t, &tc.Sum, &tc.ExpectedResult, "/infix", api.UseInfixCalculator)
		})
	}
}

func TestPrefixAPI(t *testing.T) {
	for _, tc := range PrefixAPITestData {
		testName := fmt.Sprintf("Prefix Calculator API: %s", tc.TestDescription)
		t.Run(testName, func(t *testing.T) {
			testMain(t, &tc.Sum, &tc.ExpectedResult, "/prefix", api.UsePrefixCalculator)
		})
	}
}

var InfixAPITestData = []struct {
	TestDescription string
	Sum             api.SumRequest
	ExpectedResult  api.SumResponse
}{
	{
		"Add two numbers",
		api.SumRequest{
			Sum: "( 1 + 2 )",
		},
		api.SumResponse{
			Answer: 3,
		},
	},
	{
		"Subtract two numbers",
		api.SumRequest{
			Sum: "( 1 - 2 )",
		},
		api.SumResponse{
			Answer: -1,
		},
	},
	{
		"Multiply two numbers",
		api.SumRequest{
			Sum: "( 3 *  4 )",
		},
		api.SumResponse{
			Answer: 12,
		},
	},
	{
		"Divide two numbers",
		api.SumRequest{
			Sum: "( 3 / 4 )",
		},
		api.SumResponse{
			Answer: 0.75,
		},
	},
	{
		"Combine two operations",
		api.SumRequest{
			Sum: "( 1 + ( 2 * 3 ) )",
		},
		api.SumResponse{
			Answer: 7,
		},
	},
	{
		"Combine to operations again",
		api.SumRequest{
			Sum: "( ( 1 * 2 ) + 3 )",
		},
		api.SumResponse{
			Answer: 5,
		},
	},
	{
		"Combine more operations",
		api.SumRequest{
			Sum: "( ( ( 1 + 1 ) / 10 ) - ( 1 * 2 ) )",
		},
		api.SumResponse{
			Answer: -1.8,
		},
	},
}

var PrefixAPITestData = []struct {
	TestDescription string
	Sum             api.SumRequest
	ExpectedResult  api.SumResponse
}{
	{
		"Add two numbers",
		api.SumRequest{
			Sum: "+ 3 4",
		},
		api.SumResponse{
			Answer: 7,
		},
	},
	{
		"Subtract two numbers",
		api.SumRequest{
			Sum: "- 3 4",
		},
		api.SumResponse{
			Answer: -1,
		},
	},
	{
		"Multiply two numbers",
		api.SumRequest{
			Sum: "* 3 4",
		},
		api.SumResponse{
			Answer: 12,
		},
	},
	{
		"Divide two numbers",
		api.SumRequest{
			Sum: "/ 3 4",
		},
		api.SumResponse{
			Answer: 0.75,
		},
	},
	{
		"Combine two operations",
		api.SumRequest{
			Sum: "+ 1 * 2 3",
		},
		api.SumResponse{
			Answer: 7,
		},
	},
	{
		"Combine more operations",
		api.SumRequest{
			Sum: "- / 10 + 1 1 * 1 2",
		},
		api.SumResponse{
			Answer: 3,
		},
	},
}
