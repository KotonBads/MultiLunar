package instance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateInstance(t *testing.T) {
	err := CreateInstance("test instance", "../../temp/instances", "1.8.9")
	assert.Equal(t, nil, err)
}
