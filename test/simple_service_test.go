package test

import (
	"bayudsatriyo/belajar-golang-restfull-api/simple"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializeSimpleService(false)
	fmt.Println(simpleService)
	fmt.Println(err)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}
