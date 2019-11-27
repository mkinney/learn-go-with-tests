package iteration

import (
	"testing"
	"fmt"
)

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("repeat 'a' four times", func(t *testing.T) {
		got := Repeat("a", 4)
		want := "aaaa"
		assertCorrectMessage(t, got, want)
	})
	t.Run("repeat 'ab' three times", func(t *testing.T) {
		got := Repeat("ab", 3)
		want := "ababab"
		assertCorrectMessage(t, got, want)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	output := Repeat("ab", 3)
	fmt.Println(output)
	// Output: ababab
}
