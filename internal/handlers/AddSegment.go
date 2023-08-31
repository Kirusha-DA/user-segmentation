package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Kirusha-DA/user-segmentation/internal/models"
	jsonschemes "github.com/Kirusha-DA/user-segmentation/internal/models/json_schemes"
	"github.com/Kirusha-DA/user-segmentation/internal/repositories/segments"
)

func (h handler) AddSegment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var modelSegments []models.Segment
	if err := json.Unmarshal(body, &modelSegments); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	SegmentsRepo := segments.NewReporepository(h.DB)

	var jsonItems jsonschemes.Items
	for _, modelSegment := range modelSegments {
		item, err := SegmentsRepo.Create(modelSegment)
		if err != nil {
			jsonItems = append(jsonItems, jsonschemes.Item{
				Slug:    modelSegment.Slug,
				Message: err.Error(),
			})
		} else {
			jsonItems = append(jsonItems, jsonschemes.Item{
				Slug:    item.Slug,
				Message: "OK",
			})
		}
	}
	dataJson, _ := json.MarshalIndent(jsonItems, "", " ")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}
