package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile.txt")
	if err != nil {
		panic(err)
	}

	ParseProfile(contents)
}
