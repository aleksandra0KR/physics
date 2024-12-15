package main

import (
	"fmt"
	"math"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

const (
	a = 1.0
	b = 0.8
	c = 2.0
	U = 1.0
)

func getV(x []float64) []float64 {
	V := make([]float64, len(x))

	for i, xi := range x {
		n := math.Floor(xi / c)
		if n*c < xi && xi < n*c+a {
			V[i] = 0
		} else if n*c+a < xi && xi < (n+1)*c {
			V[i] = U
		}
	}

	return V
}

func getFunc(x []float64, P float64) []float64 {
	y := make([]float64, len(x))
	for i, xi := range x {
		if xi == 0 {
			y[i] = math.Cos(xi)
		} else {
			y[i] = P*math.Sin(xi)/xi + math.Cos(xi)
		}
	}
	return y
}

func generatePlotWithPotentialPits() {
	x := linspace(-6, 7, 1000)
	V := getV(x)

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Potential Relief for Electron (Kronig-Penney Model)"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "X"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "V(x)"}),
	)

	points := make([]opts.LineData, len(x))
	for i, _ := range x {
		points[i] = opts.LineData{Value: V[i]}
	}
	line.SetXAxis(x).AddSeries("V(x)", points)

	f, _ := os.Create("potential_pits.html")
	defer f.Close()
	line.Render(f)
}

func generatePlotWithKronigPenney(P float64) {
	x := linspace(-25, 25, 1000)
	y := getFunc(x, P)
	yPlusOne := make([]float64, len(x))
	yMinusOne := make([]float64, len(x))

	for i := range x {
		yPlusOne[i] = 1
		yMinusOne[i] = -1
	}

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Graphical Analysis of Kronig-Penney Equation"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "a * alpha"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "cos(a * alpha) + P * sin(a * alpha) / (a * alpha)"}),
	)

	pointsY := make([]opts.LineData, len(x))
	pointsYPlusOne := make([]opts.LineData, len(x))
	pointsYMinusOne := make([]opts.LineData, len(x))

	for i := range x {
		pointsY[i] = opts.LineData{Value: y[i]}
		pointsYPlusOne[i] = opts.LineData{Value: yPlusOne[i]}
		pointsYMinusOne[i] = opts.LineData{Value: yMinusOne[i]}
	}

	line.SetXAxis(x).
		AddSeries("y", pointsY).
		AddSeries("y+1", pointsYPlusOne).
		AddSeries("y-1", pointsYMinusOne)

	f, _ := os.Create("kronig_penney_P" + fmt.Sprintf("%.2f", P) + ".html")
	defer f.Close()
	line.Render(f)
}

func linspace(start, end float64, num int) []float64 {
	step := (end - start) / float64(num-1)
	result := make([]float64, num)
	for i := 0; i < num; i++ {
		result[i] = start + float64(i)*step
	}
	return result
}

func main() {
	generatePlotWithPotentialPits()
	generatePlotWithKronigPenney(15)
	generatePlotWithKronigPenney(0)
	generatePlotWithKronigPenney(9999999999999)
}
