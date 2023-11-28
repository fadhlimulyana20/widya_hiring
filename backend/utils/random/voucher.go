package random

import (
	"errors"
	"math/rand"
)

func GenerateRandomVoucher(length int) (string, error) {
	if length < 4 {
		return "", errors.New("voucher length should be greater than 4")
	}

	charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	voucher := ""

	for i := 0; i < length; i++ {
		randNum := rand.Intn(len(charset))
		voucher += string(charset[randNum])
	}

	return voucher, nil
}
