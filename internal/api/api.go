package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stpiech/gosolve-task/internal/search"
)

type httpError struct {
	message string
	code    int
}

func RegisterSearchValueEndpoint(values []int) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/endpoint/{value}", func(w http.ResponseWriter, r *http.Request) {
		value := chi.URLParam(r, "value")

		responseJson, err := searchValueResponse(value, values)

		if err != (httpError{}) {
			http.Error(w, err.message, err.code)
			return
		}

		w.Write(responseJson)
	})

	http.ListenAndServe(":3001", r)

	return nil
}

func searchValueResponse(value string, values []int) ([]byte, httpError) {
	parsedValue, err := transformValueParam(value)

	if err != nil {
		return []byte{}, httpError{message: err.Error(), code: http.StatusBadRequest}
	}

	response, err := search.FindIndexOrClosest(values, parsedValue)

	if err != nil {
		return []byte{}, httpError{message: err.Error(), code: http.StatusNotFound}
	}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return []byte{}, httpError{message: err.Error(), code: http.StatusBadRequest}
	}

	return jsonResponse, httpError{}
}

func transformValueParam(valueStr string) (int, error) {
	if valueStr == "" {
		return 0, errors.New("Value is required")
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, errors.New("Invalid value")
	}

	return value, nil
}
