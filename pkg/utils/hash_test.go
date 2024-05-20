package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	assert.Equal(t,
		true,
		HashIsEqual(
			"ichor.go",
			"3ff67992e20ee313e4145e8a00a093cd2e2e31d3",
		),
	)
}
