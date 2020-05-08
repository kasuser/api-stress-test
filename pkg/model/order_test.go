package model_test

import (
	"stresstest/pkg/model"
	"testing"
)

func TestInitializeOrders(t *testing.T) {
	for _, o := range model.Orders() {
		if o.Code != "" {
			t.Error("orders should be not initialized at start")
		}
	}

	model.InitializeOrders()

	for _, o := range model.Orders() {
		if o.Code == "" {
			t.Error("orders not initialized")
		}
	}
}

func TestUpdateActualOrders(t *testing.T) {
	model.InitializeOrders()

	oldOrds := *model.Orders()

	model.UpdateActualOrders()

	updated := false
	for i, o := range model.Orders() {
		if oldOrds[i].Code != o.Code {
			updated = true
		}
	}

	if updated == false {
		t.Error("orders not initialized")
	}
}