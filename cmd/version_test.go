package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionText(t *testing.T) {
	assert.NotEmpty(t, version)
}
