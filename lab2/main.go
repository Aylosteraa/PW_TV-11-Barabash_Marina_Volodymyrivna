package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

type FuelData struct {
	BCoal  float64
	BMasut float64
	BGas   float64
	KCoal  string
	ECoal  string
	KMasut string
	EMasut string
	KGas   string
	EGas   string
}

func calculateEmissions(bCoal, bMasut, bGas float64) FuelData {
	gasDensity := 0.273
	bGas *= gasDensity

	kCoal := math.Pow(10, 6) / 20.47 * 0.8 * 25.2 / (100 - 1.5) * (1 - 0.985)
	eCoal := math.Pow(10, -6) * kCoal * 20.47 * bCoal

	kMasut := math.Pow(10, 6) / 39.48 * 1 * 0.15 / (100 - 0) * (1 - 0.985)
	eMasut := math.Pow(10, -6) * kMasut * 39.48 * bMasut

	kGas := math.Pow(10, 6) / 33.08 * 0 * 0 / (100 - 0) * (1 - 0.985)
	eGas := math.Pow(10, -6) * kGas * 33.08 * bGas

	return FuelData{
		BCoal:  bCoal,
		BMasut: bMasut,
		BGas:   bGas,
		KCoal:  fmt.Sprintf("%.2f г/ГДж", kCoal),
		ECoal:  fmt.Sprintf("%.2f т", eCoal),
		KMasut: fmt.Sprintf("%.2f г/ГДж", kMasut),
		EMasut: fmt.Sprintf("%.2f т", eMasut),
		KGas:   fmt.Sprintf("%.2f г/ГДж", kGas),
		EGas:   fmt.Sprintf("%.2f т", eGas),
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	if r.Method == http.MethodPost {
		bCoal, _ := strconv.ParseFloat(r.FormValue("coal"), 64)
		bMasut, _ := strconv.ParseFloat(r.FormValue("masut"), 64)
		bGas, _ := strconv.ParseFloat(r.FormValue("gas"), 64)
		data := calculateEmissions(bCoal, bMasut, bGas)
		tmpl.Execute(w, data)
	} else {
		tmpl.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", handler)
	port := ":8080"
	fmt.Println("Server running on http://localhost" + port)
	// log.Fatal(http.ListenAndServe(port, nil))
	http.ListenAndServe(port, nil)
}
