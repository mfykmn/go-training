package fee

import "errors"

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

//go:generate enumer -type=FeeType -transform=kebab
type FeeType int

const (
	Adult FeeType = iota
	Child
)

func Factory(feeType string) (Fee, error) {
	target, err := FeeTypeString(feeType)
	if err != nil {
		return nil, err
	}
	switch target {
	case Adult:
		return &AdultFee{}, nil
	case Child:
		return &ChildFee{}, nil
	default:
		return nil, errors.New("unknown fee type")
	}
}
