package main

import (
	"fmt"

	"github.com/mafuyuk/go-training/book/principles-of-the-systems-architecture/chapter2/fee"
)

func main() {

	a1Fee, err := fee.Factory(fee.Adult.String())
	if err != nil {
		fmt.Println(err)
		return
	}
	a2Fee, err := fee.Factory(fee.Adult.String())
	if err != nil {
		fmt.Println(err)
		return
	}
	c1Fee, err := fee.Factory(fee.Child.String())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fee.NewCharge(a1Fee).Yen())
	fmt.Println(fee.NewCharge(c1Fee).Yen())

	r := fee.NewReservation()
	r.AddFee(a1Fee)
	r.AddFee(a2Fee)
	r.AddFee(c1Fee)
	fmt.Println(r.FeeTotal())
}
