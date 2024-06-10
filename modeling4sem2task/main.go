package main

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io"
	"math"
	"os"
	"sync"
)

type InputValues struct {
	T float64 // период решетки
	N float64 // общее число штрихов
	d float64 // расстояние (м)
}

// проверка валидности данных
func check(input float64) {
	if input < float64(0) {
		fmt.Errorf("input is negative, should be positive")
	}
}

// ввод данных
func inputValues(input *InputValues) {
	fmt.Println("Enter the periods")
	fmt.Scan(&input.T)
	check(input.T)

	fmt.Println("Enter the total number of strokes")
	fmt.Scan(&input.N)
	check(input.N)

	fmt.Println("Enter the distance")
	fmt.Scan(&input.d)
	check(input.d)

}

// вычисление интенсивности
func intensity(input *InputValues, wave float64, Theta, Wavelength, Intensity *[][]float64, wavelength, theta *[]float64, wg *sync.WaitGroup) {
	defer wg.Done()
	Theta2 := (*Theta)
	if wave != 0 {
		for i := range len(*Intensity) {
			(*Intensity)[i] = make([]float64, len(*theta))
			for j := range len((*Intensity)[i]) {
				(*Intensity)[i][j] = (math.Sin(math.Pi*input.T*math.Sin(Theta2[i][j])/wave) *
					math.Sin(math.Pi*input.T*math.Sin(Theta2[i][j])/wave) /
					((math.Pi * input.T * math.Sin(Theta2[i][j]) / wave) *
						(math.Pi * input.T * math.Sin(Theta2[i][j]) / wave))) * (math.Sin(math.Pi*input.N*input.d*math.Sin(Theta2[i][j])/wave) *
					math.Sin(math.Pi*input.N*input.d*math.Sin(Theta2[i][j])/wave) /
					(math.Sin(math.Pi*input.d*math.Sin(Theta2[i][j])/wave) *
						math.Sin(math.Pi*input.d*math.Sin(Theta2[i][j])/wave)))
			}
		}
		return
	}

	Wavelength2 := *Wavelength
	for i := range len(*Intensity) {
		(*Intensity)[i] = make([]float64, len(*theta))
		for j := range len((*Intensity)[i]) {
			(*Intensity)[i][j] = (math.Sin(math.Pi*input.T*math.Sin(Theta2[i][j])/Wavelength2[i][j]) *
				math.Sin(math.Pi*input.T*math.Sin(Theta2[i][j])/Wavelength2[i][j]) /
				((math.Pi * input.T * math.Sin(Theta2[i][j]) / Wavelength2[i][j]) *
					(math.Pi * input.T * math.Sin(Theta2[i][j]) / Wavelength2[i][j]))) * (math.Sin(math.Pi*input.N*input.d*math.Sin(Theta2[i][j])/Wavelength2[i][j]) *
				math.Sin(math.Pi*input.N*input.d*math.Sin(Theta2[i][j])/Wavelength2[i][j]) /
				(math.Sin(math.Pi*input.d*math.Sin(Theta2[i][j])/Wavelength2[i][j]) *
					math.Sin(math.Pi*input.d*math.Sin(Theta2[i][j])/Wavelength2[i][j])))
		}
	}

}

func main() {
	var input InputValues
	inputValues(&input)

	// углы от -pi/2 до pi/2
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

	// длины возможных волн
	// чем больше значений, тем точнее, но медленее работает программа
	wavelength := make([]float64, 4000)
	for i := 0; i < 4000; i++ {
		wavelength[i] = 400 + float64(i)*(750-400)/4000.0
	}

	// массивы для всех возмодных комбинаций длин волн и углов
	Theta := make([][]float64, len(wavelength))
	Wavelength := make([][]float64, len(wavelength))

	for i := range Theta {
		Theta[i] = make([]float64, len(theta))
		Wavelength[i] = make([]float64, len(theta))
		for j := range Theta[0] {
			Theta[i][j] = theta[j]
			Wavelength[i][j] = wavelength[i]

		}
	}

	// длины близких волн
	// чем больше значений, тем точнее, но медленее работает программа
	wavelength2 := make([]float64, 4000)
	for i := 0; i < 4000; i++ {
		wavelength2[i] = 600 + float64(i)*(685-600)/4000.0
	}
	// массивы для всех возмодных комбинаций длин волн и углов
	Wavelength2 := make([][]float64, len(wavelength))

	for i := range Theta {
		Wavelength2[i] = make([]float64, len(theta))
		for j := range Theta[0] {
			Wavelength2[i][j] = wavelength2[i]

		}
	}

	red := 620.0

	// для параллельных вычислений
	var wg sync.WaitGroup
	wg.Add(2)

	// вычисление интенсивности 1 графика
	Intensity2 := make([][]float64, len(wavelength2))
	go intensity(&input, red, &Theta, &Wavelength2, &Intensity2, &wavelength2, &theta, &wg)

	// вычисление интенсивности 2 графика
	Intensity := make([][]float64, len(wavelength))
	go intensity(&input, 0, &Theta, &Wavelength, &Intensity, &wavelength, &theta, &wg)
	wg.Wait()

	wg.Add(2)

	// построение 3д 1 графика из полученных данных
	go Graphics(&Theta, &Wavelength, &Intensity, &wavelength, &theta, &wg)

	// построение 3д 2 графика из полученных данных
	go Graphics2(&Theta, &Wavelength2, &Intensity2, &wavelength2, &theta, &wg)
	wg.Wait()

}

// массив цветов всех волн
var line3DColor = []string{
	"#8B00FF", "#4B0082", "#0000FF", "#00FF00", "#FFFF00",
	"#FFA500", "#FF0000",
}

// массив цветов близких волн
var line3DColorRed = []string{
	"#FFFF00", "#FFA500", "#FF0000",
}

// преобразование данных для построения 3д графика
func genLine3dData(Theta, Wavelength, Intensity *[][]float64, wavelength, theta *[]float64) []opts.Chart3DData {
	data := make([]opts.Chart3DData, 0, len(*theta)*len(*wavelength))

	for i := range len(*wavelength) {
		for j := range len(*theta) {
			data = append(data, opts.Chart3DData{
				Value: []interface{}{
					(*Theta)[i][j],
					(*Wavelength)[i][j],
					(*Intensity)[i][j],
				},
			})
		}
	}

	return data
}

// настройка графика 1
func line3DBase(Theta, Wavelength, Intensity *[][]float64, wavelength, theta *[]float64) *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Вывод цветного спектра\nи графика зависимости интенсивности от угловой координаты", Subtitle: "X - Угол дифракции\nY - Длина волны\nZ - Интенсивность"}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Max:     1000000,
			InRange: &opts.VisualMapInRange{Color: line3DColor},
		}),
	)

	line3d.AddSeries("", genLine3dData(Theta, Wavelength, Intensity, wavelength, theta))
	return line3d
}

// вывод графика в виде страницы html
func Graphics(Theta, Wavelength, Intensity *[][]float64, wavelength, theta *[]float64, wg *sync.WaitGroup) {
	defer wg.Done()
	page := components.NewPage()
	page.AddCharts(
		line3DBase(Theta, Wavelength, Intensity, wavelength, theta),
	)

	f, err := os.Create("line3d.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}

// настройка графика 2
func line3DBaseNear(Theta, Wavelength, Intensity *[][]float64, wavelength, theta *[]float64) *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Вывод цветного спектра\nи графика зависимости интенсивности от угловой координаты", Subtitle: "X - Угол дифракции\nY - Длина волны\nZ - Интенсивность"}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Max:     1000000,
			InRange: &opts.VisualMapInRange{Color: line3DColorRed},
		}),
	)

	line3d.AddSeries("", genLine3dData(Theta, Wavelength, Intensity, wavelength, theta))
	return line3d
}

// вывод графика в виде страницы html 2
func Graphics2(Theta, Wavelength, Intensity *[][]float64, wavelength, theta *[]float64, wg *sync.WaitGroup) {
	defer wg.Done()
	page := components.NewPage()
	page.AddCharts(
		line3DBaseNear(Theta, Wavelength, Intensity, wavelength, theta),
	)

	f, err := os.Create("line3d2.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}

/*
 example input:
	1.0
	1000.0
	3.0
*/
