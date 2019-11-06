package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Azaradel/mm-slash-command/types"
)

type server struct {
	authToken string
	router    *http.ServeMux
}

func NewServer(authToken string) *server {
	s := &server{
		authToken,
		http.NewServeMux(),
	}
	s.routes()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) rootHandler() http.HandlerFunc {
	type Response struct {
		ResponseType string `json:"response_type"`
		Text         string `json:"text"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("failed to parse form values: %v", err)
			return
		}

		slashCommandPayload := types.NewSlashCommandForm(r.Form)

		res := &Response{
			"in_channel",
			fmt.Sprintf("Hello %s!", slashCommandPayload.Username),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
