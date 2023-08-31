package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kirusha-DA/user-segmentation/internal/models"
	jsonschemes "github.com/Kirusha-DA/user-segmentation/internal/models/json_schemes"
	"github.com/Kirusha-DA/user-segmentation/internal/repositories/segments"
	"github.com/gorilla/mux"
)

func (h handler) DeleteUserSegments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var modelSegments []models.Segment
	if err := json.Unmarshal(body, &modelSegments); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	usersSegmentRepo := segments.NewReporepository(h.DB)
	modelsDeleted := usersSegmentRepo.DeleteClients(modelSegments, id)

	segmentsRepo := segments.NewReporepository(h.DB)
	var slugIds []int
	for _, value := range modelsDeleted {
		slugIds = append(slugIds, value.Id)
	}
	modelsWithSlugs := segmentsRepo.GetSlugssById(slugIds)
	var jsonItems jsonschemes.Items
	for _, value := range modelsWithSlugs {
		jsonItems = append(jsonItems, jsonschemes.Item{
			Slug:    value.Slug,
			Message: "Ok",
		})
	}

	jsonData, _ := json.MarshalIndent(jsonItems, "", " ")
	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}
