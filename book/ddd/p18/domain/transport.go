package domain

import "errors"

type Transporter interface {
	Booking(cargo *Cargo) error
}

// Cargo 貨物を表す
type Cargo struct {
	size float32
}

func NewCargo(size float32) *Cargo {
	return &Cargo{
		size: size,
	}
}

// Voyage 公開を表す
type Voyage struct {
	cargoList []*Cargo
	capacity float32
}

func NewVoyage(capacity float32) *Voyage {
	return &Voyage{
		capacity: capacity,
	}
}

func (v *Voyage) bookedCargoSize() float32 {
	var size float32
	for _, v := range v.cargoList {
		size += v.size
	}
	return size
}

func (v *Voyage) Booking(cargo *Cargo) error {
	maxBooking := v.capacity * 1.1
	if (v.bookedCargoSize() + cargo.size) > maxBooking {
		return errors.New("Over max booking size")
	}

	v.cargoList = append(v.cargoList, cargo)
	return nil
}