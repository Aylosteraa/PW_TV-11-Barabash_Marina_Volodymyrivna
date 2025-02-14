package calculators

import (
	// "fmt"
	"html/template"
	// "log"
	"net/http"
	"strconv"
)

// Структура для збереження даних форми та результатів
type Calculation struct {
	Carbon          float64
	Hydrogen        float64
	Oxygen          float64
	Sulfur          float64
	OilHeat         float64
	FuelMoisture    float64
	Ash             float64
	Vanadium        float64
	CarbonWM        float64
	HydrogenWM      float64
	OxygenWM        float64
	SulfurWM        float64
	AshWM           float64
	VanadiumWM      float64
	LowerHeatResult float64
	Calculated      bool
}

// func calculator2() {
// 	http.HandleFunc("/", CalculatorHandler2)
// 	http.ListenAndServe(":8080", nil)
// }

// Головна сторінка
func CalculatorHandler2(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/calculator2.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Ініціалізація структури
	data := Calculation{}

	if r.Method == http.MethodPost {
		// Отримання даних з форми
		data.Carbon = parseInput(r.FormValue("carbon"))
		data.Hydrogen = parseInput(r.FormValue("hydrogen"))
		data.Oxygen = parseInput(r.FormValue("oxygen"))
		data.Sulfur = parseInput(r.FormValue("sulfur"))
		data.OilHeat = parseInput(r.FormValue("oilHeat"))
		data.FuelMoisture = parseInput(r.FormValue("fuelMoisture"))
		data.Ash = parseInput(r.FormValue("ash"))
		data.Vanadium = parseInput(r.FormValue("vanadium"))

		// Розрахунки
		factor1 := (100 - data.FuelMoisture - data.Ash) / 100
		factor2 := (100 - data.FuelMoisture/10 - data.Ash/10) / 100
		factor3 := (100 - data.FuelMoisture) / 100

		data.CarbonWM = formatValue(data.Carbon * factor1)
		data.HydrogenWM = formatValue(data.Hydrogen * factor1)
		data.OxygenWM = formatValue(data.Oxygen * factor2)
		data.SulfurWM = formatValue(data.Sulfur * factor1)
		data.AshWM = formatValue(data.Ash * factor3)
		data.VanadiumWM = formatValue(data.Vanadium * factor3)
		data.LowerHeatResult = formatValue(data.OilHeat*factor1 - 0.025*data.FuelMoisture)

		data.Calculated = true
	}

	tmpl.Execute(w, data)
}

// Функція для парсингу значень
func parseInput(value string) float64 {
	if val, err := strconv.ParseFloat(value, 64); err == nil {
		return val
	}
	return 0.0
}

// Функція для форматування значень
func formatValue(value float64) float64 {
	return float64(int(value*100)) / 100
}
