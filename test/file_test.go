package test

import (
	"bayudsatriyo/belajar-golang-restfull-api/simple"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Databases")
	fmt.Println(connection)
	assert.NotNil(t, connection)

	cleanup()
}
