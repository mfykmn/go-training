package main

import (
	"fmt"

	"github.com/mafuyuk/go-training/book/principles-of-the-systems-architecture/chapter2/fee"
)

func main() {

	fmt.Println(fee.NewCharge(fee.AdultFee{}).Yen())
	fmt.Println(fee.NewCharge(fee.ChildFee{}).Yen())
}
