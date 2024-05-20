package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	assert.Equal(t,
		true,
		HashIsEqual(
			"unzip.go",
			"6eb8f2e194d3dae061acf60267eff31599a11efd",
		),
	)
}
