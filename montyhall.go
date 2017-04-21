package main

import (
	"math/rand"
	"time"

	set "github.com/deckarep/golang-set"
)

// MontyHall はモンティ・ホール問題をシミュレートする
type MontyHall struct {
	N     int
	Doors []interface{} // monty.Doors = []int{1, 2, 3}
}

func (monty *MontyHall) simulate() (picked []interface{}, switched []interface{}) {
	n := monty.N

	pickedV := make([]interface{}, n)
	switchedV := make([]interface{}, n)

	for i := 0; i < n; i++ {
		// 当たりのドアと選ぶドアをランダムに決める
		nDoors := len(monty.Doors)
		rand.Seed(time.Now().UnixNano())
		car := monty.Doors[rand.Intn(nDoors)]
		picked := monty.Doors[rand.Intn(nDoors)]

		// 最初に選んだドア, 全てのドアのうち最初に選んだドア以外のドアの１つ(ただし、その１つ以外のドアは必ず全てハズレ)の中から１つ選ぶ
		doorSet := set.NewSetFromSlice(monty.Doors)
		carAndPicked := []interface{}{car, picked}
		goats := doorSet.Difference(set.NewSetFromSlice(carAndPicked))
		two := doorSet.Difference(goats).ToSlice()
		switched := two[rand.Intn(len(two))]

		// 選び直さず正解
		if picked == car {
			pickedV[i] = 1
		}

		// 選びなおして正解
		if switched == car {
			switchedV[i] = 1
		}

	}

	return pickedV, switchedV
}
