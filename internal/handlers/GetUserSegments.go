package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kirusha-DA/user-segmentation/internal/repositories/segments"
	"github.com/gorilla/mux"
)

func (h handler) GetUserSegments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatoin/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	segmentsRepo := segments.NewReporepository(h.DB)
	segments, _ := segmentsRepo.ReadClients(id)

	json, _ := json.MarshalIndent(segments, "", " ")
	if len(segments) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}
