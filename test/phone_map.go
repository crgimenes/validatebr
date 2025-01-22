package main

var invalidDDD = map[string]bool{
	"00": true, "01": true, "02": true, "03": true, "04": true, "05": true, "06": true, "07": true, "08": true, "09": true, "10": true,
	"20": true, "23": true, "25": true, "26": true, "29": true, "30": true, "36": true, "39": true, "40": true, "50": true, "52": true,
	"56": true, "57": true, "58": true, "59": true, "60": true, "70": true, "72": true, "76": true, "78": true, "80": true, "90": true,
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

	if invalidDDD[phone[:2]] {
		return false
	}

	return true
}
