package fee

type Reservation struct {
	fees []Fee
}

func NewReservation() *Reservation {
	return &Reservation{}
}

func (r *Reservation) AddFee(fee Fee) {
	r.fees = append(r.fees, fee)
}

func (r *Reservation) FeeTotal() Yen {
	var total int
	for _, v := range r.fees {
		total += int(v.fee())
	}
	return Yen(total)
}
