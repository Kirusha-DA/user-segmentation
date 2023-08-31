package handlers

import (
	"encoding/json"
	"net/http"

	jsonschemes "github.com/Kirusha-DA/user-segmentation/internal/models/json_schemes"
	"github.com/Kirusha-DA/user-segmentation/internal/repositories/segments"
	"github.com/gorilla/mux"
)

func (h handler) DeleteSegment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	slug := vars["slug"]

	segmentsRepo := segments.NewReporepository(h.DB)
	ok, err := segmentsRepo.Delete(slug)

	var jsonData jsonschemes.Item
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		jsonData = jsonschemes.Item{
			Slug:    slug,
			Message: err.Error(),
		}
	} else {
		w.WriteHeader(http.StatusOK)
		jsonData = jsonschemes.Item{
			Slug:    slug,
			Message: "OK",
		}
	}

	json, _ := json.MarshalIndent(jsonData, "", " ")
	w.Write(json)
}
