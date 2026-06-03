package enum2str

// supported numeric types, the tilde(~) means set of all types whose underlying type is `Numeric`
type Numeric interface {
	~uint8 | ~int8 | ~int
}

// String - returns string value of an input from a map of string list
func String[T Numeric](input T, strList ...string) (str string) {
	if len(strList) == 0 {
		return ""
	}

	i := int(input)

	if i < 0 || i >= len(strList) {
		i = 0
	}

	return strList[i]
}
