package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

type Data struct {
	Profit1  string
	Penalty1 string
	Loss1    string
	Profit2  string
	Penalty2 string
	Loss2    string
}

func calculate(power, cost, sigma float64) (float64, float64, float64, string) {
	delta := power * 0.05
	b1 := power - delta
	b2 := power + delta
	step := 0.001

	energyShare := 0.0
	for p := b1; p < b2; p += step {
		pd := (1 / (sigma * math.Sqrt(2*math.Pi))) * math.Exp(-((p-power)*(p-power))/(2*sigma*sigma))
		energyShare += pd * step
	}

	energyWithoutImbalance := math.Round(power * 24 * energyShare)
	profit := energyWithoutImbalance * cost * 1000
	energyWithImbalance := math.Round(power * 24 * (1 - energyShare))
	penalty := energyWithImbalance * cost * 1000
	loss := profit - penalty
	choice := " (збиток)"
	if loss > 0 {
		choice = " (прибуток)"
	}
	return profit, penalty, loss, choice
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	if r.Method == http.MethodPost {
		power, _ := strconv.ParseFloat(r.FormValue("power"), 64)
		cost, _ := strconv.ParseFloat(r.FormValue("cost"), 64)
		sigma1, _ := strconv.ParseFloat(r.FormValue("sigma1"), 64)
		sigma2, _ := strconv.ParseFloat(r.FormValue("sigma2"), 64)
		profit1, penalty1, loss1, choice1 := calculate(power, cost, sigma1)
		profit2, penalty2, loss2, choice2 := calculate(power, cost, sigma2)

		data := Data{
			Profit1:  fmt.Sprintf("%.2f грн", profit1),
			Penalty1: fmt.Sprintf("%.2f грн", penalty1),
			Loss1:    fmt.Sprintf("%.2f грн %s", loss1, choice1),
			Profit2:  fmt.Sprintf("%.2f грн", profit2),
			Penalty2: fmt.Sprintf("%.2f грн", penalty2),
			Loss2:    fmt.Sprintf("%.2f грн %s", loss2, choice2),
		}
		tmpl.Execute(w, data)
	} else {
		tmpl.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", handler)
	port := ":8080"
	fmt.Println("Server running on http://localhost" + port)
	http.ListenAndServe(port, nil)
}
