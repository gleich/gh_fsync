package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckExistence(t *testing.T) {
	err := os.Chdir("../..")
	checkError(t, err)

	// yml file
	testCheckExistenceInstance(t, 0)
	testCheckExistenceInstance(t, 1)
}

func testCheckExistenceInstance(t *testing.T, index int) {
	// Creating file
	f, err := os.Create(validLocations[0])
	checkError(t, err)
	err = f.Close()
	checkError(t, err)
	// Checking if the file exists
	yml_instance := checkExistence()
	assert.Equal(t, yml_instance, validLocations[0])
	// Removing the file
	err = os.Remove(validLocations[0])
	checkError(t, err)
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
