package validatebr

var (
	// invalid DDD codes (true is invalid)
	invalidDDD = []bool{
		true,  // "00",
		true,  // "01",
		true,  // "02",
		true,  // "03",
		true,  // "04",
		true,  // "05",
		true,  // "06",
		true,  // "07",
		true,  // "08",
		true,  // "09",
		true,  // "10",
		false, // "11",
		false, // "12",
		false, // "13",
		false, // "14",
		false, // "15",
		false, // "16",
		false, // "17",
		false, // "18",
		false, // "19",
		true,  // "20",
		false, // "21",
		false, // "22",
		true,  // "23",
		false, // "24",
		true,  // "25",
		true,  // "26",
		false, // "27",
		false, // "28",
		true,  // "29",
		true,  // "30",
		false, // "31",
		false, // "32",
		false, // "33",
		false, // "34",
		false, // "35",
		true,  // "36",
		false, // "37",
		false, // "38",
		true,  // "39",
		true,  // "40",
		false, // "41",
		false, // "42",
		false, // "43",
		false, // "44",
		false, // "45",
		false, // "46",
		false, // "47",
		false, // "48",
		false, // "49",
		true,  // "50",
		false, // "51",
		true,  // "52",
		false, // "53",
		false, // "54",
		false, // "55",
		true,  // "56",
		true,  // "57",
		true,  // "58",
		true,  // "59",
		true,  // "60",
		false, // "61",
		false, // "62",
		false, // "63",
		false, // "64",
		false, // "65",
		false, // "66",
		false, // "67",
		false, // "68",
		false, // "69",
		true,  // "70",
		false, // "71",
		true,  // "72",
		false, // "73",
		false, // "74",
		false, // "75",
		true,  // "76",
		false, // "77",
		true,  // "78",
		false, // "79",
		true,  // "80",
		false, // "81",
		false, // "82",
		false, // "83",
		false, // "84",
		false, // "85",
		false, // "86",
		false, // "87",
		false, // "88",
		false, // "89",
		true,  // "90",
		false, // "91",
		false, // "92",
		false, // "93",
		false, // "94",
		false, // "95",
		false, // "96",
		false, // "97",
		false, // "98",
		false, // "99",
	}
)

func fastAtoi(b [2]byte) int {
	return int(b[0]-'0')*10 + int(b[1]-'0')
}

func PhoneWithBrazilianAreaCodeArray(phone string) bool { // DDD
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

	ddd := fastAtoi([2]byte{phone[0], phone[1]})
	return invalidDDD[ddd]
}
