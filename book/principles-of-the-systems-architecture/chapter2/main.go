package main

import (
	"fmt"

	"github.com/mafuyuk/go-training/book/principles-of-the-systems-architecture/chapter2/fee"
)

func main() {

	aFee := fee.Factory(fee.Adult)
	sFee := fee.Factory(fee.Senior)
	cFee := fee.Factory(fee.Child)

	r := fee.NewReservation()
	r.AddFee(aFee)
	r.AddFee(sFee)
	r.AddFee(cFee)
	fmt.Println(r.FeeTotal())
}
