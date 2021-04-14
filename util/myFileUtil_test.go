package util

import (
	"os"
	"strings"
	"testing"
)

func TestExits(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		expect   bool
	}{
		{
			name:     "exits",
			filename: "BNB",
			expect:   true,
		},
		{
			name:     "notExits",
			filename: "HELLO",
			expect:   false,
		},
		{
			name:     "exits",
			filename: "Aave",
			expect:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if b := Exits(strings.Join([]string{"..", GetCoinDir(tt.filename)}, string(os.PathSeparator))); b != tt.expect {
				t.Errorf("writeSourceCode() result = %v, expect %v", b, tt.expect)
			}
		})
	}
}
