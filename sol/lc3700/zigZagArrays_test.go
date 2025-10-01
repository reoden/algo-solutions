package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestZigZagArrays(t *testing.T) {
	assert.Equal(t, 2, zigZagArrays(3, 4, 5))
	assert.Equal(t, 10, zigZagArrays(3, 1, 3))
}
