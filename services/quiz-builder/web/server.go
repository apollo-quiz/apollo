package web

import (
	"log"
	"net/http"
	"quiz-builder/web/quiz"
)

func StartServer() error {
	log.Println("starting server")
	mux := http.NewServeMux()

	mux.HandleFunc("/quiz/search", quiz.SearchQuizzes)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return err
	}

	return nil
}
