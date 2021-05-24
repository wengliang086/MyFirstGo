package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	for i := 0; i < 10; i++ {
		Fetch(fmt.Sprintf("test-%d", i))
	}
}
