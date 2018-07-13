package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Just an example showing stretchr's testify package
func TestHelloWorld(t *testing.T) {
	t.Parallel()

	assert.Nil(t, nil)
}
