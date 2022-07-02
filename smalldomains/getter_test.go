package smalldomains

import (
	"testing"
)

func TestSuccessCode(t *testing.T) {
	if !isSuccessStatusCode(200) {
		t.Errorf("200 was incorrectly detected as a failed HTTP Status Code")
	}
}

func TestFailCode(t *testing.T) {
	statusCodes := []int{300, 400, 500}

	for _, c := range statusCodes {
		if isSuccessStatusCode(c) {
			t.Errorf("%d was incorrectly detected as a successful HTTP Status Code", c)
		}
	}
}
