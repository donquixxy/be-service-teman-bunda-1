package test

import (
	"strings"
	"testing"
)

func GenerateReferalCode(t *testing.T) {
	name := "I Putu Tensu Qiuwulu"
	splitName := strings.Split(name, "")

	if len(splitName) != 4 {
		t.Errorf("salah	")
	}
}
