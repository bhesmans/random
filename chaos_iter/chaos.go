package main

// Choas and iterations
// REF : Mathematics : From the birth of numbers, Jan Gullberg, p348
// (...)
// Non linear fun (1 - cx^2) applied to itself n times with -1 < x_0 < 1
//
// There is stable point up to c = 3/4
// Then there is the first bifurcation
// Then one more,
// Then it become chaos!

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type t struct {
	n int
	c float64
}

func f(c, x float64) float64 {
	return 1 - c*x*x
}

func nth_f(n int, c, x float64) float64 {
	res := x
	for i := 0; i < n; i++ {
		res = f(c, res)
	}
	return res
}

func (t t) Len() int {
	return 300
}

func (t t) XY(i int) (float64, float64) {
	x0 := -1 + (2/float64(t.Len()))*float64(i)
	return float64(t.c), nth_f(t.n, t.c, x0)
}

func main() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Iteration and chaos"
	p.X.Label.Text = "C"
	p.Y.Label.Text = "Results"

	for i := float64(0); i < 2; i += 0.01 {
		err = plotutil.AddScatters(p, t{1000, i})

		if err != nil {
			panic(err)
		}
	}

	if err := p.Save(10*vg.Inch, 10*vg.Inch, "chaos.png"); err != nil {
		panic(err)
	}
}
