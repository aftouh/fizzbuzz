package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aftouh/fizzbuzz/models"
	"github.com/google/go-cmp/cmp"
)

func fizzbuzz16() []string {
	return []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16"}
}

func TestSuccess(t *testing.T) {
	tt := []struct {
		name               string
		request            *http.Request
		expectedStatusCode int
		expectedResponseKO *models.ErrorResponse
		expectedResponseOK *models.FizzbuzzResponse
	}{
		{
			"success",
			httptest.NewRequest("GET", "http://test.com?int1=3&int2=5&limit=16&str1=fizz&str2=buzz", nil),
			http.StatusOK,
			nil,
			&models.FizzbuzzResponse{Result: fizzbuzz16(), Status: "Success"},
		},
		{
			"error missing parameter",
			httptest.NewRequest("GET", "ttp://test.com?int2=5&limit=16&str1=fizz&str2=buzz", nil),
			http.StatusBadRequest,
			&models.ErrorResponse{Message: "Failed to parse query parameters: int1 is empty", Status: "Bad request"},
			nil,
		},
		{
			"error converting int parameter",
			httptest.NewRequest("GET", "http://test.com?int1=s&int2=5&limit=16&str1=fizz&str2=buzz", nil),
			http.StatusBadRequest,
			&models.ErrorResponse{Message: "Failed to parse query parameters: schema: error converting value for \"int1\"", Status: "Bad request"},
			nil,
		},
		{
			"error calculating fizzbuzz",
			httptest.NewRequest("GET", "http://test.com?int1=3&int2=5&limit=0&str1=fizz&str2=buzz", nil),
			http.StatusBadRequest,
			&models.ErrorResponse{Message: "Failed to calculate fizzbuzz result: limit must be great or equal to 1", Status: "Bad request"},
			nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			Fizzbuzz(w, tc.request)

			// Check http status code
			if tc.expectedStatusCode != w.Code {
				t.Errorf("Expected status code: %d, got: %d", tc.expectedStatusCode, w.Code)
			}

			// Check bad request result
			if tc.expectedResponseKO != nil {
				var got models.ErrorResponse
				b, _ := ioutil.ReadAll(w.Body)
				if err := json.Unmarshal(b, &got); err != nil {
					t.Errorf("failed to parse result; %v", err)
				}
				if !cmp.Equal(*tc.expectedResponseKO, got) {
					t.Errorf("Exptected response: %v, got: %v", tc.expectedResponseKO, got)
				}
			}

			// Check success request result
			if tc.expectedResponseOK != nil {
				var got models.FizzbuzzResponse
				b, _ := ioutil.ReadAll(w.Body)
				if err := json.Unmarshal(b, &got); err != nil {
					t.Errorf("failed to parse result; %v", err)
				}
				if !cmp.Equal(*tc.expectedResponseOK, got) {
					t.Errorf("Exptected response: %v, got: %v", tc.expectedResponseOK, got)
				}
			}
		})
	}

	// w := httptest.NewRecorder()
	// r := httptest.NewRequest("GET", "http://test.com?int1=3&int2=5&limit=16&str1=fizz&str2=buzz", nil)

	// Fizzbuzz(w, r)
	// if w.Code != http.StatusOK {
	// 	t.Errorf("ko")
	// }

	// expected := models.FizzbuzzResponse{
	// 	Result: fizzbuzz16(),
	// 	Status: "Success",
	// }

	// var got models.FizzbuzzResponse

	// b, _ := ioutil.ReadAll(w.Body)
	// if err := json.Unmarshal(b, &got); err != nil {
	// 	t.Errorf("failed to parse result; %v", err)
	// }
	// if !cmp.Equal(expected, got) {
	// 	t.Errorf("Exptected response: %v, got: %v", expected, got)
	// }
}
