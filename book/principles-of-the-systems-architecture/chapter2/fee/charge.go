package fee

type Charge struct {
	fee Fee
}

func NewCharge(fee Fee) *Charge {
	return &Charge{
		fee: fee,
	}
}

func (c *Charge) Yen() Yen {
	return c.fee.fee()
}
