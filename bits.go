package patricia

var _len8tab = [256]uint8{
	0x00, 0x01, 0x02, 0x02, 0x03, 0x03, 0x03, 0x03, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05,
	0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06,
	0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06,
	0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
	0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
	0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
	0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
}

var _leftMasks32 []uint32
var _leftMasks64 []uint64

func initBuildLeftMasks() {
	_leftMasks32 = make([]uint32, 33)
	for i := uint(1); i < 33; i++ {
		_leftMasks32[i] = uint32(_leftMasks32[i-1] | 1<<(32-i))
	}

	_leftMasks64 = make([]uint64, 65)
	for i := uint(1); i < 65; i++ {
		_leftMasks64[i] = uint64(_leftMasks64[i-1] | 1<<(64-i))
	}
}

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
func LeadingZeros32(x uint32) (n int) {
	if x >= 1<<16 {
		x >>= 16
		n = 16
	}
	if x >= 1<<8 {
		x >>= 8
		n += 8
	}
	return 32 - n - int(_len8tab[x])
}

// LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.
func LeadingZeros64(x uint64) (n int) {
	if x >= 1<<32 {
		x >>= 32
		n = 32
	}
	if x >= 1<<16 {
		x >>= 16
		n += 16
	}
	if x >= 1<<8 {
		x >>= 8
		n += 8
	}
	return 64 - n - int(_len8tab[x])
}

// MergePrefixes32 merges two 32-bit prefixes, returning new prefix, new length
func MergePrefixes32(left uint32, leftLength uint, right uint32, rightLength uint) (uint32, uint) {
	return (left & _leftMasks32[leftLength]) | ((right & _leftMasks32[rightLength]) >> leftLength), (leftLength + rightLength)
}

// MergePrefixes64 merges two pairs of uint64s, returning a new prefix, new length
func MergePrefixes64(leftLeft uint64, leftRight uint64, leftLength uint, rightLeft uint64, rightRight uint64, rightLength uint) (uint64, uint64, uint) {
	// mask the left 128 bits
	if leftLength <= 64 {
		leftLeft &= _leftMasks64[leftLength]
		leftRight = 0
	} else {
		leftRight &= _leftMasks64[leftLength-64]
	}

	// mask the right 128 bits
	if rightLength <= 64 {
		rightLeft &= _leftMasks64[rightLength]
		rightRight = 0
	} else {
		rightRight &= _leftMasks64[rightLength-64]
	}

	// shift the right 128 bits to the right
	rightLeft, rightRight = ShiftRightIPv6(rightLeft, rightRight, leftLength)

	// now merge the two
	return leftLeft | rightLeft, leftRight | rightRight, leftLength + rightLength
}
