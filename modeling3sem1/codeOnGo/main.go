package main

import (
	"fmt"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"log"
	"math"

	"gonum.org/v1/plot"
)

type InputValues struct {
	L, L1, m, k, beta, phi1, phi2, T float64
}

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

func calculation(inputValues InputValues) (t, phi1_t, phi2_t, v1_t, v2_t []float64) {
	const g = 9.82

	omega1 := math.Sqrt(g / inputValues.L)
	omega2 := math.Sqrt(g/inputValues.L + 2*inputValues.k*inputValues.L1*inputValues.L1/(inputValues.m*inputValues.L*inputValues.L))

	𝜉1 := (inputValues.phi1 + inputValues.phi2) / 2
	𝜉2 := (inputValues.phi1 - inputValues.phi2) / 2

	dt := 0.001

	t = make([]float64, 0, int(inputValues.T/dt)+1)
	phi1_t = make([]float64, 0, int(inputValues.T/dt)+1)
	phi2_t = make([]float64, 0, int(inputValues.T/dt)+1)
	v1_t = make([]float64, 0, int(inputValues.T/dt)+1)
	v2_t = make([]float64, 0, int(inputValues.T/dt)+1)

	for i := 0.0; i <= inputValues.T/dt; i++ {
		time := i * dt

		phi12 := (𝜉1*math.Cos(omega1*time) + 𝜉2*math.Cos(omega2*time)) * math.Exp(-inputValues.beta*time)
		phi22 := (𝜉1*math.Cos(omega1*time) - 𝜉2*math.Cos(omega2*time)) * math.Exp(-inputValues.beta*time)
		v1 := (𝜉1*omega1*(-math.Sin(omega1*time)) + 𝜉2*omega2*(-math.Sin(omega2*time))) * math.Exp((-inputValues.beta)*time)
		v2 := (𝜉1*omega1*(-math.Sin(omega1*time)) - 𝜉2*omega2*(-math.Sin(omega2*time))) * math.Exp((-inputValues.beta)*time)

		t = append(t, time)
		phi1_t = append(phi1_t, phi12)
		phi2_t = append(phi2_t, phi22)
		v1_t = append(v1_t, v1)
		v2_t = append(v2_t, v2)
	}

	fmt.Printf("Нормальная частота omega1: %0.6f Герц\n", omega1)
	fmt.Printf("Нормальная частота omega2: %0.6f Герц\n", omega2)

	return t, phi1_t, phi2_t, v1_t, v2_t
}

func VFromTPlot2(t, v1_t, v2_t []float64) {
	plot := plot.New()

	plot.Title.Text = "Зависимость скорости от времени для каждого маятника V(t)"
	plot.X.Label.Text = "Время (сек)"
	plot.Y.Label.Text = "Скорость (рад/сек)"
	width := vg.Points(0.5)

	first := make(plotter.XYs, len(t))
	second := make(plotter.XYs, len(t))

	for i, _ := range t {
		first[i].X = t[i]
		first[i].Y = v1_t[i]

		second[i].X = t[i]
		second[i].Y = v2_t[i]
	}

	firstLine, err := plotter.NewLine(first)
	if err != nil {
		log.Fatal("failed to get total line", "err", err)
	}
	firstLine.Width = width
	firstLine.Color = color.RGBA{R: 255, G: 192, B: 203, A: 255}

	secondLine, err := plotter.NewLine(second)
	if err != nil {
		log.Fatal("failed to get total line", "err", err)
	}
	secondLine.Width = width
	secondLine.Color = color.RGBA{R: 0, G: 255, B: 255, A: 255}

	plot.Add(firstLine, secondLine)

	plot.Legend.Add("Первый", firstLine)
	plot.Legend.Add("Второй", secondLine)

	plot.X.Min = minFromSlice(t)
	plot.X.Max = maxFromSlice(t)
	plot.Y.Min = min(minFromSlice(v1_t), minFromSlice(v2_t))
	plot.Y.Max = max(maxFromSlice(v1_t), maxFromSlice(v2_t))

	filename := "v(t).png"

	if err := plot.Save(5*vg.Inch, 3*vg.Inch, filename); err != nil {
		log.Fatal("failed to save plot", "err", err)
	}
}

func PhiFromTPlot2(t, phi1_t, phi2_t []float64) {
	plot := plot.New()

	plot.Title.Text = "Зависимость угла от времени для каждого маятника phi(t)"
	plot.X.Label.Text = "Время (сек)"
	plot.Y.Label.Text = "Угол (рад)"
	width := vg.Points(0.5)

	first := make(plotter.XYs, len(t))
	second := make(plotter.XYs, len(t))

	for i, _ := range t {
		first[i].X = t[i]
		first[i].Y = phi1_t[i]

		second[i].X = t[i]
		second[i].Y = phi2_t[i]

	}

	firstLine, err := plotter.NewLine(first)
	if err != nil {
		log.Fatal("failed to get total line", "err", err)
	}
	firstLine.Width = width
	firstLine.Color = color.RGBA{R: 255, G: 192, B: 203, A: 255}

	secondLine, err := plotter.NewLine(second)
	if err != nil {
		log.Fatal("failed to get total line", "err", err)
	}
	secondLine.Width = width
	secondLine.Color = color.RGBA{R: 0, G: 255, B: 255, A: 255}

	plot.Add(firstLine, secondLine)

	plot.Legend.Add("Первый", firstLine)
	plot.Legend.Add("Второй", secondLine)

	plot.X.Min = minFromSlice(t)
	plot.X.Max = maxFromSlice(t)
	plot.Y.Min = min(minFromSlice(phi1_t), minFromSlice(phi2_t))
	plot.Y.Max = max(maxFromSlice(phi1_t), maxFromSlice(phi2_t))

	filename := "phi(t).png"

	if err := plot.Save(5*vg.Inch, 3*vg.Inch, filename); err != nil {
		log.Fatal("failed to save plot", "err", err)
	}
}

func main() {
	inputValues := InputValues{
		L:    1.0,
		L1:   1.0,
		m:    1.0,
		k:    0.1,
		beta: 0.01,
		phi1: math.Pi / 180 * 10.0,
		phi2: math.Pi / 180 * 20.0,
		T:    100.0,
	}

	var phi1, phi2 float64
	fmt.Scan(&inputValues.L)
	fmt.Scan(&inputValues.L1)
	fmt.Scan(&inputValues.m)
	fmt.Scan(&inputValues.k)
	fmt.Scan(&inputValues.beta)
	fmt.Scan(&phi1)
	inputValues.phi1 = math.Pi / 180 * phi1
	fmt.Scan(&phi2)
	inputValues.phi2 = math.Pi / 180 * phi2
	fmt.Scan(&inputValues.T)

	t, phi1_t, phi2_t, v1_t, v2_t := calculation(inputValues)
	PhiFromTPlot2(t, phi1_t, phi2_t)
	VFromTPlot2(t, v1_t, v2_t)

}
