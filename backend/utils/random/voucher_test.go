package random

import "testing"

func TestGenerenateVoucherNormal(t *testing.T) {
	vocuher, err := GenerateRandomVoucher(12)
	if err != nil {
		t.Error(err)
	}

	t.Log(vocuher)
}

func TestGenerenateVoucherTooShort(t *testing.T) {
	_, err := GenerateRandomVoucher(2)
	if err == nil {
		t.Error(err)
	}
}
