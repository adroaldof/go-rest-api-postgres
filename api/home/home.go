package home

import (
	"encoding/json"
	"net/http"
	"time"
)

type Home struct {
	Now string `json:"Now"`
}

func HomePage(response http.ResponseWriter, request *http.Request) {
	now := Home{
		Now: time.Now().String(),
	}

	json.NewEncoder(response).Encode(now)
}
