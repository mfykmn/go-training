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
