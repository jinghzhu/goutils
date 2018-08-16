package data

import (
	"math"
	"testing"

	"github.com/jinghzhu/goutils/converter"
)

func TestStrToUint64(t *testing.T) {
	validUints := []uint64{0, 1, math.MaxUint8, math.MaxUint16, math.MaxUint32, math.MaxUint64}
	invalidUints := []string{"1.233", "a", "false"}

	for _, f := range validUints {
		_, err := StrToUint64(converter.Uint64ToStr(f))
		if err != nil {
			t.Errorf("Should pass for %+v but encounter error %s\n", f, err.Error())
		}
	}
	for _, f := range invalidUints {
		_, err := StrToUint64(f)
		if err == nil {
			t.Errorf("Should encounter error but pass for %+v\n", f)
		}
	}
}
