package model

import (
	"math/rand"
)

type order struct {
	Code   string
	Usages uint
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz")
var orders [50]order

func Orders() *[50]order {
	return &orders
}

func InitializeOrders() {
	for i := 0; i < 50; i++ {
		orders[i] = order{Code: getRandCode(), Usages: 0}
	}
}

func UpdateActualOrders() {
	orders[rand.Int63() % int64(len(letters))] = order{Code: getRandCode(), Usages: 0}
}

func getRandCode() string {
	c := make([]rune, 2)
	for i := range c {
		c[i] = letters[rand.Int63() % int64(len(letters))]
	}
	return string(c)
}
