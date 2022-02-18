package cekgocek

import "regexp"

func CheckValidateEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

func CheckValidateIndonesiaPhone(phone string) bool {
	re := regexp.MustCompile(`^(\+62|62|0)8[1-9][0-9]{6,9}$`)
	return re.MatchString(phone)
}
