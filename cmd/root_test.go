package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommentType(t *testing.T) {
	val1 := getCommentType("bname")
	val2 := getCommentType("hoge")
	assert.NotEqual(t, val1, val2)
}
