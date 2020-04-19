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
	n          int
	c          float64
	minx, maxx float64
	minc, maxc float64
	f          f
	fileName   string
	fun        string
}

type f func(c, x float64) float64

func f1(c, x float64) float64 {
	return 1 - c*x*x
}

func f2(c, x float64) float64 {
	return c * x * (1 - x)
}

func nth_f(n int, c, x float64, f f) float64 {
	res := x
	for i := 0; i < n; i++ {
		res = f(c, res)
	}
	return res
}

func (t *t) Len() int {
	return 300
}

func (t *t) XY(i int) (float64, float64) {
	x0 := t.minx + ((t.maxx-t.minx)/float64(t.Len()+2))*float64(i+1)
	return float64(t.c), nth_f(t.n, t.c, x0, t.f)
}

func chaosGraph(t *t) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = t.fun
	p.X.Label.Text = "C"
	p.Y.Label.Text = "Results"

	for i := t.minc; i < t.maxc; i += 0.01 {
		t.c = i
		err = plotutil.AddScatters(p, t)

		if err != nil {
			panic(err)
		}
	}

	if err := p.Save(10*vg.Inch, 10*vg.Inch, t.fileName); err != nil {
		panic(err)
	}
}

func main() {
	chaosGraph(&t{1000, 0, -1, 1, 0, 2, f1, "chaos.png", "1 - c * x^2"})
	chaosGraph(&t{1000, 0, 0, 1, 2.4, 4, f2, "chaos2.png", "c * x * (1 - x)"})
}
