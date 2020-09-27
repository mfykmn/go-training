// Code generated by "stringer -type=feeType"; DO NOT EDIT.

package fee

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Adult-1]
	_ = x[Child-2]
	_ = x[Senior-3]
}

const _feeType_name = "AdultChildSenior"

var _feeType_index = [...]uint8{0, 5, 10, 16}

func (i feeType) String() string {
	i -= 1
	if i < 0 || i >= feeType(len(_feeType_index)-1) {
		return "feeType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _feeType_name[_feeType_index[i]:_feeType_index[i+1]]
}