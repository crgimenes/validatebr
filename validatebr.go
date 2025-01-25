package validatebr

import (
	"errors"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

var (
	cpfP1  = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	cpfP2  = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	cnpjP1 = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	cnpjP2 = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	cnpjRegex = regexp.MustCompile(
		`^[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$`)
	cnpjAlphaRegex = regexp.MustCompile(
		`^[0-9A-Z]{2}\.?[0-9A-Z]{3}\.?[0-9A-Z]{3}/?[0-9A-Z]{4}-?[0-9]{2}$`)
	cpfRegex   = regexp.MustCompile(`^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2}$`)
	emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)

	ErrInvalidPixType   = errors.New("invalid pix type")
	ErrInvalidCharacter = errors.New("invalid character")
)

func IsEmailValid(e string) bool {
	s := strings.ToLower(e)
	return emailRegex.MatchString(s)
}

func IsCNPJ(e string) bool {
	return cnpjRegex.MatchString(e)
}

func IsCNPJAlpha(e string) bool {
	return cnpjAlphaRegex.MatchString(strings.ToUpper(e))
}

func IsCPF(e string) bool {
	return cpfRegex.MatchString(e)
}

func RemoveNonDigits(s string) string {
	var b strings.Builder
	for _, c := range s {
		if unicode.IsDigit(c) {
			b.WriteRune(c)
		}
	}
	return b.String()
}

func RemoveNonAlphaNum(s string) string {
	var b strings.Builder
	for _, c := range s {
		if unicode.IsDigit(c) ||
			unicode.IsLetter(c) {
			b.WriteRune(c)
		}
	}
	return b.String()
}

func IsRepetitive(s string) bool {
	if len(s) == 0 {
		return false
	}
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

func getAlphanumericValue(r rune) (int, error) {
	r = unicode.ToUpper(r)
	if unicode.IsDigit(r) ||
		unicode.IsLetter(r) {
		return int(r) - 48, nil
	}
	return 0, ErrInvalidCharacter
}

func sumAlpha(s string, table []int) (int, error) {
	total := 0
	for i, v := range table {
		val, err := getAlphanumericValue(rune(s[i]))
		if err != nil {
			return -1, err
		}
		total += val * v
	}
	return total, nil
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

// CNPJAlphanumeric is a experimental function to validate CNPJ with alphanumeric characters use with caution.
func CNPJAlphanumeric(cnpj string) bool {
	cnpj = RemoveNonAlphaNum(cnpj)
	if len(cnpj) != 14 {
		return false
	}

	if IsRepetitive(cnpj) {
		return false
	}

	for i := 12; i < 14; i++ {
		if cnpj[i] < '0' || cnpj[i] > '9' {
			return false
		}
	}

	p1 := cnpj[:12]
	p2 := cnpj[12:]

	s, err := sumAlpha(p1, cnpjP1)
	if err != nil {
		return false
	}
	if s < 0 {
		return false
	}
	r1 := s % 11
	d1 := 0
	if r1 >= 2 {
		d1 = 11 - r1
	}

	p1ComDV := p1 + strconv.Itoa(d1)

	s2, err := sumAlpha(p1ComDV, cnpjP2)
	if err != nil {
		return false
	}
	if s2 < 0 {
		return false
	}
	r2 := s2 % 11
	d2 := 0
	if r2 >= 2 {
		d2 = 11 - r2
	}

	dv1, _ := strconv.Atoi(string(p2[0]))
	dv2, _ := strconv.Atoi(string(p2[1]))

	return d1 == dv1 && d2 == dv2
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

	if len(ret) == 0 {
		return nil, ErrInvalidPixType
	}

	slices.Sort(ret)

	return ret, nil
}
