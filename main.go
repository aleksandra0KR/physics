package main

import (
	"fmt"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"gonum.org/v1/plot"
	"image/color"
	"log"
	"math"
)

const (
	dataFile = "demo.dat"
	gifFile  = "demo.gif"
	numPts   = 100
)

func minFromSlice(array []float64) float64 {
	res := array[0]
	for i := range array {
		if array[i] < res {
			res = array[i]
		}
	}
	return res
}

func maxFromSlice(array []float64) float64 {
	res := array[0]
	for i := range array {
		if array[i] > res {
			res = array[i]
		}
	}
	return res
}

type InputValues struct {
	T float64 // период решетки
	N float64 // общее число штрихов
	d float64 // расстояние (м)
}

func check(input float64) {
	if input < float64(0) {
		fmt.Errorf("input is negative, should be positive")
	}
}

func inputValues(input *InputValues) {
	fmt.Println("Enter the periods")
	fmt.Scan(&input.T)
	check(input.T)

	fmt.Println("Enter the total number of strokes")
	fmt.Scan(&input.N)
	check(input.N)

	fmt.Println("Enter the total number of strokes")
	fmt.Scan(&input.d)
	check(input.d)

}

func intensity() {

}

func main() {
	var input InputValues
	//inputValues(&input)
	input.T = 1.0
	input.N = 1000.0
	input.d = 3.0
	//theta := make([]float64, 180)
	//for i := 0; i < 180; i++ {
	//	theta[i] = -math.Pi/2 + float64(i)*(math.Pi/2)/180.0
	//	}

	theta := []float64{-1.57079633, -1.53906307, -1.50732981, -1.47559655, -1.44386329,
		-1.41213003, -1.38039677, -1.34866351, -1.31693025, -1.28519699,
		-1.25346374, -1.22173048, -1.18999722, -1.15826396, -1.1265307,
		-1.09479744, -1.06306418, -1.03133092, -0.99959766, -0.9678644,
		-0.93613114, -0.90439789, -0.87266463, -0.84093137, -0.80919811,
		-0.77746485, -0.74573159, -0.71399833, -0.68226507, -0.65053181,
		-0.61879855, -0.58706529, -0.55533203, -0.52359878, -0.49186552,
		-0.46013226, -0.428399, -0.39666574, -0.36493248, -0.33319922,
		-0.30146596, -0.2697327, -0.23799944, -0.20626618, -0.17453293,
		-0.14279967, -0.11106641, -0.07933315, -0.04759989, -0.01586663,
		0.01586663, 0.04759989, 0.07933315, 0.11106641, 0.14279967,
		0.17453293, 0.20626618, 0.23799944, 0.2697327, 0.30146596,
		0.33319922, 0.36493248, 0.39666574, 0.428399, 0.46013226,
		0.49186552, 0.52359878, 0.55533203, 0.58706529, 0.61879855,
		0.65053181, 0.68226507, 0.71399833, 0.74573159, 0.77746485,
		0.80919811, 0.84093137, 0.87266463, 0.90439789, 0.93613114,
		0.9678644, 0.99959766, 1.03133092, 1.06306418, 1.09479744,
		1.1265307, 1.15826396, 1.18999722, 1.22173048, 1.25346374,
		1.28519699, 1.31693025, 1.34866351, 1.38039677, 1.41213003,
		1.44386329, 1.47559655, 1.50732981, 1.53906307, 1.57079633}
	wavelength := make([]float64, 5000)
	for i := 0; i < 5000; i++ {
		wavelength[i] = 400 + float64(i)*(750-400)/5000.0
	}

	Theta := make([][]float64, len(theta))
	Wavelength := make([][]float64, len(theta))

	for i := range Theta {
		Theta[i] = make([]float64, len(wavelength))
		Wavelength[i] = make([]float64, len(wavelength))
		for j := range Theta[i] {
			Theta[i][j] = theta[i]
			Wavelength[i][j] = wavelength[j]
		}
	}

	Intensity := make([][]float64, len(theta))
	for i := range Intensity {
		Intensity[i] = make([]float64, len(wavelength))
		for j := range Intensity[i] {
			Intensity[i][j] = (math.Sin(math.Pi*input.T*math.Sin(Theta[i][j])/Wavelength[i][j]) *
				math.Sin(math.Pi*input.T*math.Sin(Theta[i][j])/Wavelength[i][j]) /
				((math.Pi * input.T * math.Sin(Theta[i][j]) / Wavelength[i][j]) *
					(math.Pi * input.T * math.Sin(Theta[i][j]) / Wavelength[i][j]))) * (math.Sin(math.Pi*input.N*input.d*math.Sin(Theta[i][j])/Wavelength[i][j]) *
				math.Sin(math.Pi*input.N*input.d*math.Sin(Theta[i][j])/Wavelength[i][j]) /
				(math.Sin(math.Pi*input.d*math.Sin(Theta[i][j])/Wavelength[i][j]) *
					math.Sin(math.Pi*input.d*math.Sin(Theta[i][j])/Wavelength[i][j])))
		}
	}

	intens := make([]float64, 180)
	wave := 400.0
	for i := 0; i < 100; i++ {
		intens[i] = (math.Sin(math.Pi*input.T*math.Sin(theta[i])/wave) *
			math.Sin(math.Pi*input.T*math.Sin(theta[i])/wave) /
			((math.Pi * input.T * math.Sin(theta[i]) / wave) *
				(math.Pi * input.T * math.Sin(theta[i]) / wave))) * (math.Sin(math.Pi*input.N*input.d*math.Sin(theta[i])/wave) *
			math.Sin(math.Pi*input.N*input.d*math.Sin(theta[i])/wave) /
			(math.Sin(math.Pi*input.d*math.Sin(theta[i])/wave) *
				math.Sin(math.Pi*input.d*math.Sin(theta[i])/wave)))
	}

	VFromTPlot2(theta, intens)
}

func VFromTPlot2(t, v1T []float64) {
	p := plot.New()

	p.Title.Text = "Зависимость скорости от времени для каждого маятника V(t)"
	p.X.Label.Text = "Время (сек)"
	p.Y.Label.Text = "Скорость (рад/сек)"
	width := vg.Points(0.5)

	first := make(plotter.XYs, len(t))

	for i := range t {
		first[i].X = t[i]
		first[i].Y = v1T[i]

	}

	firstLine, err := plotter.NewLine(first)
	if err != nil {
		log.Fatal("failed to get total line", "err", err)
	}
	firstLine.Width = width
	firstLine.Color = color.RGBA{R: 255, G: 192, B: 203, A: 255}

	p.Add(firstLine)

	p.Legend.Add("Первый", firstLine)

	p.X.Min = minFromSlice(t)
	p.X.Max = maxFromSlice(t)
	fmt.Println(p.X.Max)
	filename := "v(t).png"

	if err := p.Save(10*vg.Inch, 5*vg.Inch, filename); err != nil {
		log.Fatal("failed to save p", "err", err)
	}
}
