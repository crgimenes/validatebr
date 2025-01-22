package main

func PhoneWithBrazilianAreaCodeNoMap(phone string) bool { // DDD
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

	invalidDDDs := []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "20", "23", "25", "26", "29", "30", "36", "39", "40", "50", "52", "56", "57", "58", "59", "60", "70", "72", "76", "78", "80", "90"}

	for _, ddd := range invalidDDDs {
		if phone[:2] == ddd {
			return false
		}
	}

	return true
}
