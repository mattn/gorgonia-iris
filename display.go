package main

import (
	"bytes"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/mattn/go-sixel"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func plotData(x []float64, a []float64) {
	p := plot.New()
	p.Title.Text = "sepal length & width"
	p.X.Label.Text = "length"
	p.Y.Label.Text = "width"
	p.Add(plotter.NewGrid())

	l := len(x) / len(a)
	for k := 1; k <= 3; k++ {
		data0 := make(plotter.XYs, 0)
		for i := 0; i < len(a); i++ {
			if k != int(a[i]) {
				continue
			}
			x1 := x[i*l+0] // sepal_length
			y1 := x[i*l+1] // sepal_width
			data0 = append(data0, plotter.XY{X: x1, Y: y1})
		}
		data, err := plotter.NewScatter(data0)
		if err != nil {
			log.Fatal(err)
		}
		data.GlyphStyle.Color = plotutil.Color(k - 1)
		data.Shape = &draw.PyramidGlyph{}
		p.Add(data)
		p.Legend.Add(fmt.Sprint(k), data)
	}

	w, err := p.WriterTo(4*vg.Inch, 4*vg.Inch, "png")
	if err != nil {
		panic(err)
	}
	var b bytes.Buffer
	w.WriteTo(&b)
	if false {
		img, err := png.Decode(&b)
		if err != nil {
			println("foO")
			panic(err)
		}
		sixel.NewEncoder(os.Stdout).Encode(img)
	} else {
		ioutil.WriteFile("out.png", b.Bytes(), 0644)
	}
}
