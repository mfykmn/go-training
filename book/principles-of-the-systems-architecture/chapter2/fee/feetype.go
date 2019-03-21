package fee

//go:generate stringer -type=feeType
type feeType int

const (
	_ feeType = iota
	Adult
	Child
	Senior
)
