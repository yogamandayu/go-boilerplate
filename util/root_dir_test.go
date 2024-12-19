package util_test

import (
	"testing"

	"github.com/yogamandayu/go-boilerplate/util"

	"github.com/stretchr/testify/assert"
)

func TestRootDir(t *testing.T) {
	rootDir := util.RootDir()
	assert.NotEmpty(t, rootDir)
}
