package validatebr

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	cpfP1  = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	cpfP2  = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	cnpjP1 = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	cnpjP2 = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	cnpjRegex  = regexp.MustCompile(`^[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$`)
	cpfRegex   = regexp.MustCompile(`^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2}$`)
	emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)

func IsEmailValid(e string) bool {
	s := strings.ToLower(e)
	return emailRegex.MatchString(s)
}

func IsCNPJ(e string) bool {
	return cnpjRegex.MatchString(e)
}

func IsCPF(e string) bool {
	return cpfRegex.MatchString(e)
}

func RemoveNonDigits(s string) string {
	var r string
	for _, c := range s {
		if unicode.IsDigit(c) {
			r += string(c)
		}
	}
	return r
}

func IsRepetitive(s string) bool {
	c := s[0]
	for _, v := range []byte(s) {
		if c != v {
			return false
		}
	}
	return true
}

func sum(s string, table []int) int {
	r := 0

	for i, v := range table {
		c := s[i]
		d := int(c - '0')
		r += v * d
	}

	return r
}

func CPF(cpf string) bool {
	cpf = RemoveNonDigits(cpf)
	if len(cpf) != 11 {
		return false
	}

	if IsRepetitive(cpf) {
		return false
	}

	p1 := cpf[:9]
	p2 := cpf[9:]

	s := sum(p1, cpfP1)
	r1 := s % 11
	d1 := 0

	if r1 >= 2 {
		d1 = 11 - r1
	}

	s2 := p1 + strconv.Itoa(d1)
	dsum := sum(s2, cpfP2)
	r2 := dsum % 11

	d2 := 0
	if r2 >= 2 {
		d2 = 11 - r2
	}

	p1aux := int(p2[0] - '0')
	p2aux := int(p2[1] - '0')

	return byte(d1) == byte(p1aux) && byte(d2) == byte(p2aux)
}

func CNPJ(cnpj string) bool {
	cnpj = RemoveNonDigits(cnpj)
	if len(cnpj) != 14 {
		return false
	}

	if IsRepetitive(cnpj) {
		return false
	}

	p1 := cnpj[:12]
	p2 := cnpj[12:]

	s := sum(p1, cnpjP1)
	r1 := s % 11

	d1 := 0
	if r1 >= 2 {
		d1 = 11 - r1
	}

	s2 := p1 + strconv.Itoa(d1)
	dsum := sum(s2, cnpjP2)
	r2 := dsum % 11

	d2 := 0
	if r2 >= 2 {
		d2 = 11 - r2
	}

	p1aux := int(p2[0] - '0')
	p2aux := int(p2[1] - '0')

	return byte(d1) == byte(p1aux) && byte(d2) == byte(p2aux)
}

func PhoneWithBrazilianAreaCode(phone string) bool { // DDD
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

	invalidDDD := []string{
		"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10",
		"20", "23", "25", "26", "29", "30", "36", "39", "40", "50", "52",
		"56", "57", "58", "59", "60", "70", "72", "76", "78", "80", "90",
	}

	for _, v := range invalidDDD {
		if v == phone[:2] {
			return false
		}
	}

	return true
}

func PixKeyType(pixkey string) ([]string, error) {
	types := map[string]bool{
		"CNPJ":  false,
		"CPF":   false,
		"EMAIL": false,
		"EVP":   false,
		"PHONE": false,
	}

	if IsEmailValid(pixkey) {
		types["EMAIL"] = true
	}
	if IsCNPJ(pixkey) && CNPJ(pixkey) {
		types["CNPJ"] = true
	}
	if IsCPF(pixkey) && CPF(pixkey) {
		types["CPF"] = true
	}
	if len(pixkey) == 36 {
		types["EVP"] = true
	}
	if PhoneWithBrazilianAreaCode(pixkey) {
		types["PHONE"] = true
	}

	var ret []string
	for k, v := range types {
		if v {
			ret = append(ret, k)
		}
	}

	return ret, nil
}
