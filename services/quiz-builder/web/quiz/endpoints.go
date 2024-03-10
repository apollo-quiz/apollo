package quiz

import (
	"common/quiz"
	"common/util"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
)

func SearchQuizzes(w http.ResponseWriter, r *http.Request) {
	util.AllowCrossOrigin(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}

	myquiz := quiz.Quiz{
		Id:         "my-myquiz",
		Title:      "My Test Quiz",
		SectionIds: []string{"section1", "section2", "section3"},
	}

	bytes, err := protojson.Marshal(&myquiz)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
