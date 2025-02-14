package calculators

import (
	"html/template"
	"net/http"
)

// Структура для збереження даних форми та результатів
type Calculation1 struct {
	Carbon              float64
	Hydrogen            float64
	Oxygen              float64
	Sulfur              float64
	Moisture            float64
	Ash                 float64
	CoefficientWtoD     float64
	CoefficientWtoC     float64
	HeatWorkingMass     float64
	HeatDryMass         float64
	HeatCombustibleMass float64
	Calculated          bool
}

// Сторінка
func CalculatorHandler1(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/calculator1.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Ініціалізація структури
	data := Calculation1{}

	if r.Method == http.MethodPost {
		// Отримання даних з форми
		data.Carbon = parseInput(r.FormValue("carbon"))
		data.Hydrogen = parseInput(r.FormValue("hydrogen"))
		data.Oxygen = parseInput(r.FormValue("oxygen"))
		data.Sulfur = parseInput(r.FormValue("sulfur"))
		data.Ash = parseInput(r.FormValue("ash"))
		data.Moisture = parseInput(r.FormValue("moisture"))

		// Розрахунки
		data.CoefficientWtoD = formatValue(100 / (100 - data.Moisture))
		data.CoefficientWtoC = formatValue(100 / (100 - data.Moisture - data.Ash))

		data.HeatWorkingMass = formatValue((339*data.Carbon + 1030*data.Hydrogen - 108.8*(data.Oxygen-data.Sulfur) - 25*data.Moisture) / 1000)
		data.HeatDryMass = formatValue((data.HeatWorkingMass + 0.025*data.Moisture) * 100 / (100 - data.Moisture))
		data.HeatCombustibleMass = formatValue((data.HeatWorkingMass + 0.025*data.Moisture) * 100 / (100 - data.Moisture - data.Ash))

		data.Calculated = true
	}

	tmpl.Execute(w, data)
}
