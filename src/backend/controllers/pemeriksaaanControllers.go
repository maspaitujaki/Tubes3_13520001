package controllers

import (
	"dna-matching-api/database"
	"dna-matching-api/entity"
	"encoding/json"
	"net/http"
)

// GetAllPemeriksaan returns all pemeriksaan
func GetAllPemeriksaan(w http.ResponseWriter, r *http.Request) {
	var pemeriksaans []entity.Pemeriksaan
	database.Connector.Find(&pemeriksaans)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pemeriksaans)
}

// CreatePemeriksaan creates pemeriksaan
func CreatePemeriksaan(w http.ResponseWriter, r *http.Request) {
	// TODO : Implement CreatePemeriksaan dengan KMP dan Boyer-Moore
}
