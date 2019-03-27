package storage

import (
	"strings"
	"testing"
)

func TestCheckQuota(t *testing.T) {
	var wasCalled bool
	var actualMsg string
	var actualUser string

	tests := []struct{
		username string
		bytesUsed int64
		msgContains string
		wantNotify bool
	} {
		{"john", 910000000, "91%", true},
		{"mary", 890000000, "89%", false},
	}

	// override notify with custom test
	notifyBkp := notify
	notify = func(msg string, username string) {
		wasCalled = true
		actualMsg = msg
		actualUser = username
	}
	
	for _, test := range tests {
		wasCalled = false
		usage[test.username] = test.bytesUsed

		CheckQuota(test.username)

		if (wasCalled != test.wantNotify) {
			t.Errorf("notify was called: %t, want %t", wasCalled, test.wantNotify)
		}

		if (!test.wantNotify) {
			continue // no further checks
		}

		if (!strings.Contains(actualMsg, test.msgContains)) {
			t.Errorf("message was %q, want it to contain %q", actualMsg, test.msgContains)
		}
		if (actualUser != test.username) {
			t.Errorf("username was %q, want %q", actualUser, test.username)
		}
	}

	notify = notifyBkp
}
