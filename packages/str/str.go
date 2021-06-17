package str

func ReverseString(str string) string  {
	var reverseStr []byte;
	// convert the original string into a byte
	original := []byte(str);
	for i:= len(str) - 1; i >= 0; i-- {
		reverseStr = append(reverseStr, original[i]);
	}
	return string(reverseStr);
}