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

// вводимые данные
type InputValues struct {
	L, L1, m, k, beta, phi1, phi2, T float64
}

// вспомагательная функция для графиков
func minFromSlice(array []float64) float64 {
	res := array[0]
	for i := range array {
		if array[i] < res {
			res = array[i]
		}
	}
	return res
}

// вспомагательная функция для графиков
func maxFromSlice(array []float64) float64 {
	res := array[0]
	for i := range array {
		if array[i] > res {
			res = array[i]
		}
	}
	return res
}

// вычисление углов и скорости
func calculation(inputValues InputValues) (t, phi1T, phi2T, v1T, v2T []float64) {
	const g = 9.82

	omega1 := math.Sqrt(g / inputValues.L)
	omega2 := math.Sqrt(g/inputValues.L + 2*inputValues.k*inputValues.L1*inputValues.L1/(inputValues.m*inputValues.L*inputValues.L))

	𝜉1 := (inputValues.phi1 + inputValues.phi2) / 2
	𝜉2 := (inputValues.phi1 - inputValues.phi2) / 2

	dt := 0.001

	t = make([]float64, 0, int(inputValues.T/dt)+1)
	phi1T = make([]float64, 0, int(inputValues.T/dt)+1)
	phi2T = make([]float64, 0, int(inputValues.T/dt)+1)
	v1T = make([]float64, 0, int(inputValues.T/dt)+1)
	v2T = make([]float64, 0, int(inputValues.T/dt)+1)

	for i := 0.0; i <= inputValues.T/dt; i++ {
		time := i * dt

		phi12 := (𝜉1*math.Cos(omega1*time) + 𝜉2*math.Cos(omega2*time)) * math.Exp(-inputValues.beta*time)
		phi22 := (𝜉1*math.Cos(omega1*time) - 𝜉2*math.Cos(omega2*time)) * math.Exp(-inputValues.beta*time)
		v1 := (𝜉1*omega1*(-math.Sin(omega1*time)) + 𝜉2*omega2*(-math.Sin(omega2*time))) * math.Exp((-inputValues.beta)*time)
		v2 := (𝜉1*omega1*(-math.Sin(omega1*time)) - 𝜉2*omega2*(-math.Sin(omega2*time))) * math.Exp((-inputValues.beta)*time)

		t = append(t, time)
		phi1T = append(phi1T, phi12)
		phi2T = append(phi2T, phi22)
		v1T = append(v1T, v1)
		v2T = append(v2T, v2)
	}

	fmt.Printf("Нормальная частота omega1: %0.6f Герц\n", omega1)
	fmt.Printf("Нормальная частота omega2: %0.6f Герц\n", omega2)

	return t, phi1T, phi2T, v1T, v2T
}

// Зависимость скорости от времени для каждого маятника
func VFromTPlot2(t, v1T, v2T []float64) {
	p := plot.New()

	p.Title.Text = "Зависимость скорости от времени для каждого маятника V(t)"
	p.X.Label.Text = "Время (сек)"
	p.Y.Label.Text = "Скорость (рад/сек)"
	width := vg.Points(0.5)

	first := make(plotter.XYs, len(t))
	second := make(plotter.XYs, len(t))

	for i := range t {
		first[i].X = t[i]
		first[i].Y = v1T[i]

		second[i].X = t[i]
		second[i].Y = v2T[i]
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

	p.Add(firstLine, secondLine)

	p.Legend.Add("Первый", firstLine)
	p.Legend.Add("Второй", secondLine)

	p.X.Min = minFromSlice(t)
	p.X.Max = maxFromSlice(t)
	p.Y.Min = min(minFromSlice(v1T), minFromSlice(v2T))
	p.Y.Max = max(maxFromSlice(v1T), maxFromSlice(v2T))

	filename := "v(t).png"

	if err := p.Save(5*vg.Inch, 3*vg.Inch, filename); err != nil {
		log.Fatal("failed to save p", "err", err)
	}
}

// Зависимость угла от времени для каждого маятника
func PhiFromTPlot2(t, phi1T, phi2T []float64) {
	p := plot.New()

	p.Title.Text = "Зависимость угла от времени для каждого маятника phi(t)"
	p.X.Label.Text = "Время (сек)"
	p.Y.Label.Text = "Угол (рад)"
	width := vg.Points(0.5)

	first := make(plotter.XYs, len(t))
	second := make(plotter.XYs, len(t))

	for i := range t {
		first[i].X = t[i]
		first[i].Y = phi1T[i]

		second[i].X = t[i]
		second[i].Y = phi2T[i]

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

	p.Add(firstLine, secondLine)

	p.Legend.Add("Первый", firstLine)
	p.Legend.Add("Второй", secondLine)

	p.X.Min = minFromSlice(t)
	p.X.Max = maxFromSlice(t)
	p.Y.Min = min(minFromSlice(phi1T), minFromSlice(phi2T))
	p.Y.Max = max(maxFromSlice(phi1T), maxFromSlice(phi2T))

	filename := "phi(t).png"

	if err := p.Save(5*vg.Inch, 3*vg.Inch, filename); err != nil {
		log.Fatal("failed to save p", "err", err)
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

	// ввод и проверка на валидность
	var phi1, phi2 float64
	_, err := fmt.Scan(&inputValues.L)
	if err != nil {
		return
	}
	if inputValues.L < 0 {
		fmt.Println("L can't be negative, try again")
		return
	}
	_, err = fmt.Scan(&inputValues.L1)
	if err != nil {
		return
	}
	if inputValues.L1 < 0 {
		fmt.Println("L can't be negative, try again")
		return
	}
	_, err = fmt.Scan(&inputValues.m)
	if err != nil {
		return
	}
	if inputValues.m < 0 {
		fmt.Println("m can't be negative, try again")
		return
	}
	_, err = fmt.Scan(&inputValues.k)
	if err != nil {
		return
	}
	if inputValues.k < 0 {
		fmt.Println("k can't be negative, try again")
		return
	}
	_, err = fmt.Scan(&inputValues.beta)
	if err != nil {
		return
	}
	if inputValues.beta < 0 {
		fmt.Println("beta can't be negative, try again")
		return
	}
	_, err = fmt.Scan(&phi1)
	if err != nil {
		return
	}
	if phi1 < 0 {
		fmt.Println("phi1 can't be negative, try again")
		return
	}
	inputValues.phi1 = math.Pi / 180 * phi1
	_, err = fmt.Scan(&phi2)
	if err != nil {
		return
	}
	if phi1 < 0 {
		fmt.Println("phi2 can't be negative, try again")
		return
	}
	inputValues.phi2 = math.Pi / 180 * phi2
	_, err = fmt.Scan(&inputValues.T)
	if err != nil {
		return
	}
	if inputValues.T < 0 {
		fmt.Println("T can't be negative, try again")
		return
	}
	t, phi1T, phi2T, v1T, v2T := calculation(inputValues)
	PhiFromTPlot2(t, phi1T, phi2T)
	VFromTPlot2(t, v1T, v2T)

}

/*

example input
1.0
1.0
1.0
0.1
0.01
10
20
100

*/
