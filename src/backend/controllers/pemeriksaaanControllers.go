package controllers

import (
	"dna-matching-api/database"
	"dna-matching-api/entity"
	"encoding/json"
	"io/ioutil"
	"time"

	// "fmt"
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

func KmpMatch(text string, pattern string) int {
	var n int = len(text)
	var m int = len(pattern)
	var fail []int = ComputeFail(pattern)

	var i int = 0
	var j int = 0

	for i < n {
		if text[i] == pattern[j] {
			if j == m-1 {
				return i - m + 1
			}
			i++
			j++
		} else if j > 0 {
			j = fail[j-1]
		} else {
			i++
		}
	}
	return -1
}

func ComputeFail(pattern string) []int {
	var m int = len(pattern)
	var fail []int = make([]int, m)
	fail[0] = 0

	var j int = 0
	var i int = 1
	for x := 1; x < m; x++ {
		if pattern[i] == pattern[j] {
			fail[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = fail[j-1]
		} else {
			fail[i] = 0
			i++
		}
	}
	return fail
}

// CreatePemeriksaan creates pemeriksaan
func CreatePemeriksaan(w http.ResponseWriter, r *http.Request) {
	// TODO : Implement CreatePemeriksaan dengan KMP dan Boyer-Moore
	requestBody, _ := ioutil.ReadAll(r.Body)
	var pemeriksaan entity.Pemeriksaan
	json.Unmarshal(requestBody, &pemeriksaan)

	// TODO: Cek penyakit ada nggak
	var penyakit entity.Penyakit
	database.Connector.Where("nama = ?", pemeriksaan.Penyakit).First(&penyakit)

	// // kalo ada lanjut ke bawah

	//TODO : Pemeriksaan rantai
	var cek int = KmpMatch(pemeriksaan.Rantai, penyakit.Rantai)
	if cek == -1 {
		pemeriksaan.Hasil = false
		pemeriksaan.Tanggal = time.Now()
	} else {
		pemeriksaan.Hasil = true
		pemeriksaan.Tanggal = time.Now()
	}

	// TODO : Assign Hasil dan tanggal()
	// pemeriksaan.Hasil =
	// pemeriksaan.Tanggal = time.Now()

	// Penambahan ke database
	if result := database.Connector.Create(&pemeriksaan); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pemeriksaan)
}
