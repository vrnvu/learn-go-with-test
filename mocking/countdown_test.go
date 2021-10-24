package mocking

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockSleeper struct{}

func (m MockSleeper) Sleep() {}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	mockSleeper := MockSleeper{}

	Countdown(buffer, mockSleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	assert.Equal(t, want, got)
}
