package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetCommentTypeTest(t *testing.T) {
	val1 := getCommentType("bname")
	val2 := getCommentType("hoge")
	assert.NotEqual(val1, val2)
}
