package validatebr

func IsValidDDD(ddd int) bool {
	const (
		invalidDDDBitmask1 = 0b0001111100010100000000011001000001100110100100000000011111111111
		invalidDDDBitmask2 = 0b0000000000000000000000000000000000000100000000010101000101000000
	)

	if ddd < 0 || ddd > 99 {
		return false
	}
	if ddd < 64 {
		return (invalidDDDBitmask1 & (1 << ddd)) == 0
	}
	dddAdjusted := ddd - 64
	return (invalidDDDBitmask2 & (1 << dddAdjusted)) == 0
}

func PhoneWithBrazilianAreaCode(phone string) bool {
	var digits [13]byte
	digitCount := 0

	// remove all non-digits
	for i := 0; i < len(phone) && digitCount < 13; i++ {
		c := phone[i]
		if c >= '0' && c <= '9' {
			digits[digitCount] = c
			digitCount++
		}
	}

	// if the phone number is 13 digits long and starts with 55, then remove the 55
	if digitCount == 13 && digits[0] == '5' && digits[1] == '5' {
		copy(digits[0:], digits[2:13])
		digitCount -= 2
	}

	// validate fone number length
	if digitCount != 11 {
		return false
	}

	// validate if all digits are the same
	allSame := true
	for i := 1; i < 11; i++ {
		if digits[i] != digits[0] {
			allSame = false
			break
		}
	}
	if allSame {
		return false
	}

	// validate DDD
	ddd := int(digits[0]-'0')*10 + int(digits[1]-'0')
	return IsValidDDD(ddd)
}

func PhoneWithBrazilianAreaCodeMap(phone string) bool { // DDD
	phone = RemoveNonDigits(phone)
	if len(phone) == 13 {
		phone = phone[2:]
	}

	if len(phone) != 11 {
		return false
	}

	if IsRepetitive(phone) {
		return false
	}

	invalidDDD := map[string]struct{}{
		"00": {}, "01": {}, "02": {}, "03": {}, "04": {}, "05": {}, "06": {}, "07": {}, "08": {}, "09": {},
		"10": {}, "20": {}, "23": {}, "25": {}, "26": {}, "29": {}, "30": {}, "36": {}, "39": {}, "40": {},
		"50": {}, "52": {}, "56": {}, "57": {}, "58": {}, "59": {}, "60": {}, "70": {}, "72": {}, "76": {},
		"78": {}, "80": {}, "90": {},
	}

	if _, ok := invalidDDD[phone[:2]]; ok {
		return false
	}

	return true
}
