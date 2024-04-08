// sumber dataset: oleh ALLENA VENKATA SAI ABY di kaggle, Licensi: CC0: Public Domain

package main

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// membaca file dataset
	f, err := os.Open("salary_dataset.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// membuat data untuk model regresi
	xs := make([]float64, len(records))
	ys := make([]float64, len(records))

	for i, record := range records {
		if len(record[0]) == 0 || len(record[1]) == 0 {
			log.Printf("Skipping record %d because it has an empty value", i)
			continue
		}

		x, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Skipping record %d because it has an invalid value: %v", i, err)
			continue
		}

		y, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Skipping record %d because it has an invalid value: %v", i, err)
			continue
		}

		xs[i] = x
		ys[i] = y
	}


	//regresi linear
	alpha, beta := stat.LinearRegression(xs, ys, nil, false)

	log.Printf("alpha = %v, beta = %v\n", alpha, beta)

	// menghitung Mean Absolute Error (MAE)
	mae := MAE(xs, ys)
	log.Printf("Mean Absolute Error = %v\n", mae)

	// membuat plot
	p := plot.New()
	p.Title.Text = "===Scatter plot dengan garis yang paling sesuai==="
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	pts := make(plotter.XYs, len(xs))

	for i := range pts {
		pts[i].X = xs[i]
		pts[i].Y = ys[i]
	}

	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Panic(err)
	}

	line := plotter.NewFunction(func(x float64) float64 {
		return alpha + beta*x
	})

	p.Add(s, line)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "gambar_plot.png"); err != nil {
		log.Panic(err)
	}
}

func MAE(yTrue, yPred []float64) float64 {
	sum := 0.0
	for i := range yTrue {
		sum += math.Abs(yTrue[i] - yPred[i])
	}
	return sum / float64(len(yTrue))
}
