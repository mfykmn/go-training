package fee

type Yen int

type Fee interface {
	fee() Yen
	label() string
}

type AdultFee struct{}

func (_ AdultFee) fee() Yen {
	return Yen(100)
}

func (_ AdultFee) label() string {
	return "大人"
}

type ChildFee struct{}

func (_ ChildFee) fee() Yen {
	return Yen(50)
}

func (_ ChildFee) label() string {
	return "子供"
}

type SeniorFee struct{}

func (_ SeniorFee) fee() Yen {
	return Yen(80)
}

func (_ SeniorFee) label() string {
	return "老人"
}

//go:generate stringer -type=feeType
type feeType int

const (
	_ feeType = iota
	Adult
	Child
	Senior
)

func Factory(feeType feeType) Fee {
	switch feeType {
	case Adult:
		return &AdultFee{}
	case Child:
		return &ChildFee{}
	case Senior:
		return &ChildFee{}
	}
	return nil
}
