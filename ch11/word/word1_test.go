package word

import (
	"testing"
	"math/rand"
	"time"
)

func TestPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want bool
	} {
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste", true},
		{"Palindrome", false},
		{"dessers", false},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

// generates a random palindrome
func randomPalindrome(rnd *rand.Rand) string {
	n := rnd.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		rn := rune(rnd.Intn(0x1000))
		runes[i] = rn
		runes[n - i - 1] = rn
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Seed: %d", seed)
	rnd := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		word := randomPalindrome(rnd)
		if !IsPalindrome(word) {
			t.Errorf("IsPalindrome(%q) = false", word)
		}
	}

}
