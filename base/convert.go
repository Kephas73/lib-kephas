package base

import (
	"strconv"
	"strings"
)

// StringToInt32 func;
func StringToInt32(s string) (int32, error) {
	i, err := strconv.Atoi(s)
	return int32(i), err
}

// StringToUint32 func;
func StringToUint32(s string) (uint32, error) {
	i, err := strconv.Atoi(s)
	return uint32(i), err
}

// StringToInt64 func;
func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// StringToUint64 func;
func StringToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

// StringToBool func;
func StringToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// Int64ToString func;
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// Int32ToString func;
func Int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

// BoolToInt8 func;
func BoolToInt8(b bool) int8 {
	if b {
		return 1
	} else {
		return 0
	}
}

// Int8ToBool func;
func Int8ToBool(b int8) bool {
	if b == 1 {
		return true
	} else {
		return false
	}
}

// HexToInt64 func;
func HexToInt64(hexString string) int64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexString, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return int64(result)
}

// Int64ToBin func;
func Int64ToBin(value int64) string {
	return strconv.FormatInt(value, 2) // base 2 for binary
}

// BinToInt64 func;
func BinToInt64(binString string) int64 {
	// base 2 for binary
	result, _ := strconv.ParseInt(binString, 2, 64)
	return result
}

// BinTo2sCompleteInt32 func;
func BinTo2sCompleteInt32(bin string) int32 {
	if strings.HasPrefix(bin, "0") {
		return int32(BinToInt64(bin))
	}
	myBin := ""
	for _, r := range bin {
		myBin += string(r)
	}
	return int32(BinToInt64(myBin))
}

// HexTo2sCompleteInt32 func;
func HexTo2sCompleteInt32(hexString string) int32 {
	return BinTo2sCompleteInt32(Int64ToBin(HexToInt64(hexString)))
}