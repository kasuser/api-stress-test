package model

import (
	"math/rand"
	"sync"
)

type order struct {
	Code   string
	Usages uint
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz")
var orders [50]order
var OrdMutex sync.Mutex

// Orders returns slice with orders
func Orders() *[50]order {
	return &orders
}

// InitializeOrders initializes orders and save at slice
func InitializeOrders() {
	for i := 0; i < 50; i++ {
		orders[i] = order{Code: getRandCode(), Usages: 0}
	}
}

// UpdateActualOrders substitutes a new order instead one of existing. It
// should be start by timer
func UpdateActualOrders() {
	OrdMutex.Lock()
	defer OrdMutex.Unlock()

	orders[rand.Int63()%int64(len(letters))] = order{Code: getRandCode(), Usages: 0}
}

func getRandCode() string {
	c := make([]rune, 2)
	for i := range c {
		c[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(c)
}
