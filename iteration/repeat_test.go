package iteration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {

	t.Run("repeat one char", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assert.Equal(t, want, got)
	})

	t.Run("repeat two char", func(t *testing.T) {
		got := Repeat("aa", 4)
		want := "aaaaaaaa"
		assert.Equal(t, want, got)
	})

	t.Run("repeat three char", func(t *testing.T) {
		got := Repeat("aaa", 2)
		want := "aaaaaa"
		assert.Equal(t, want, got)
	})
}

func BenchmarkRepeatSize5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkRepeatSize15(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 15)
	}
}

func BenchmarkRepeatSize55(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 55)
	}
}
