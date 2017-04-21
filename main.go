package main

import (
	"log"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
)

func makeXYsWithSlices(x []float64, y []float64) plotter.XYs {
	// xとyは同じ数あるか
	if len(x) != len(y) {
		log.Fatalln("Error: Two slices' length are difference")
	}

	xys := make(plotter.XYs, len(x))

	for i := 0; i < len(x); i++ {
		xys[i].X = x[i]
		xys[i].Y = y[i]
	}

	return xys
}

func convertISliceToF64S(arr []interface{}) (converted []float64) {
	conv := make([]float64, len(arr))
	for i, v := range arr {
		// TODO: ちゃんとnilチェックする
		if v != nil {
			conv[i] = float64(v.(int))
		} else {
			conv[i] = float64(0)
		}
	}
	return conv
}

func cumsum(arr []float64) []float64 {
	if len(arr) <= 0 {
		panic(arr)
	}

	r := []float64{}
	var sum float64
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		r = append(r, sum)
	}
	return r
}

func main() {
	// MontyHall問題を、変更した場合・しなかった場合それぞれ10000回試行する
	monty := MontyHall{
		N:     10000,
		Doors: []interface{}{1, 2, 3},
	}
	pickedWin, switchedWin := monty.simulate()

	// 結果を描画
	// プロット用データの整形
	x := make([]interface{}, monty.N)
	for i := 0; i < monty.N; i++ {
		x[i] = i
	}

	pickedLine := Line{
		XYs:  makeXYsWithSlices(convertISliceToF64S(x), cumsum(convertISliceToF64S(pickedWin))),
		Name: "No swithing",
	}

	switchedLine := Line{
		XYs:  makeXYsWithSlices(convertISliceToF64S(x), cumsum(convertISliceToF64S(switchedWin))),
		Name: "Swithing",
	}

	// タイトル, ラベルの設定
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Monty Hall Problem"
	p.X.Label.Text = "Challenges" // "How many times She did"
	p.Y.Label.Text = "Cars"       // "Total Score"

	// プロット
	plotter := Plotter{
		Plot:  p,
		Lines: []Line{pickedLine, switchedLine},
	}
	if err := plotter.plot(); err != nil {
		log.Fatalln("Error while plotting")
	}
}
