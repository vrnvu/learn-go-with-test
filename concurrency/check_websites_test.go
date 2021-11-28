package concurrency

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func mock(url string) bool {
	if !strings.HasPrefix(url, "http://") {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	tests := []struct {
		web  string
		want bool
	}{
		{web: "http://google.com", want: true},
		{web: "http://blog.dave.com", want: true},
		{web: "invalid://invalid.com", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.web, func(t *testing.T) {
			got := CheckWebsites(mock, []string{tt.web})
			assert.Equal(t, tt.want, got[tt.web])
		})
	}
}

func slowMock(wc WebsiteChecker) WebsiteChecker {
	time.Sleep(50 * time.Millisecond)
	return wc
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 500)
	for i := 0; i < len(urls); i++ {
		urls[i] = fmt.Sprintf("http://url%d.com", i)
	}

	got := CheckWebsites(slowMock(mock), urls)
	for _, v := range got {
		assert.True(b, v)
	}
}
