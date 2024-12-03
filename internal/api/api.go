package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/stpiech/gosolve-task/internal/logger"
	"github.com/stpiech/gosolve-task/internal/search"
)

type httpError struct {
	message string
	code    int
}

func RegisterSearchValueEndpoint(values []int, port int) error {
	r := chi.NewRouter()
	r.Use(LoggerMiddleware)

	r.Get("/endpoint/{value}", func(w http.ResponseWriter, r *http.Request) {
		value := chi.URLParam(r, "value")

		responseJson, err := searchValueResponse(value, values)

		if err != (httpError{}) {
			logger.ErrorLogger(err.message)
			http.Error(w, err.message, err.code)
			return
		}

		w.Write(responseJson)
	})

	logger.InfoLogger(fmt.Sprintf("Server is starting at port %d", port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)

	if err != nil {
		return err
	}

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
