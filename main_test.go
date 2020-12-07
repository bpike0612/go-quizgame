package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_main(t *testing.T) {
	a := assert.New(t)
	a.NotNil(main)
}
