package main

import (
	"fmt"
	"testing"
)

// test db connection
func TestDb(t *testing.T) {
	dbConnect()
	err := db.Connect()
	if err != nil {
		t.Errorf("Database connection error.")
	}
}

// deduct stock
func PurchaseTest(t *testing.T) {
	deductStock(false, 20)
	logAction(false, fmt.Sprintf("%v bought", "test-animal"), 1)
}

// is stock sufficient (1/2)
func CatTest(t *testing.T) {
	countStock()
		if 1 > cats {
			t.Errorf("cat stock insuffient.")
		}
}

// is stock sufficient (2/2)
func DogTest(t *testing.T) {
	countStock()
		if 1 > dogs {
			t.Errorf("dog stock insuffient.")
		}
}
