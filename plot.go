package main

import (
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
)

// Plotter は検証結果のプロットをする
type Plotter struct {
	Plot  *plot.Plot
	Lines []Line
}

// Line は、プロットしてできる直線を表す
type Line struct {
	XYs  plotter.XYs
	Name string
}

func (p *Plotter) plot() error {
	err := plotutil.AddLinePoints(p.Plot, p.Lines[0].Name, p.Lines[0].XYs, p.Lines[1].Name, p.Lines[1].XYs)
	if err != nil {
		return err
	}

	if err := p.Plot.Save(640, 480, "result.png"); err != nil {
		return err
	}

	return nil
}
