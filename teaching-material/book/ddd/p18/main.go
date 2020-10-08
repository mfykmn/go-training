package main

import (
	"fmt"

	"github.com/mfykmn/go-training/teaching-material/book/ddd/p18/domain"
)

func main() {
	voyage := domain.NewVoyage(10)
	cargoList := []*domain.Cargo{
		domain.NewCargo(1),
		domain.NewCargo(4),
		domain.NewCargo(5.5),
		domain.NewCargo(2),
	}

	for _, cargo := range cargoList {
		if err := voyage.Booking(cargo); err != nil {
			fmt.Printf("failed err: %v\n", err)
		} else {
			fmt.Println("success")
		}
	}
}
