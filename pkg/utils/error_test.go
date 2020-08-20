package utils

import "testing"

func TestCheckErr(t *testing.T) {
	CheckErr("Failed!", nil)
}

func TestCheckTestingErr(t *testing.T) {
	CheckTestingErr(t, nil)
}
