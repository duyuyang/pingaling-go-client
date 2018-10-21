// +build integration

package main_test

import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmdVersion(t *testing.T) {

	cmd := exec.Command("./pingaling", "version")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	assert.Nil(t, err)
	assert.Equal(t, "0.5.0\n", out.String())

}
