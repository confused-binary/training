package mascot_test

import (
	"testing"

	"github.com/confused-binary/training/golang/001_vscode_setup/mascot"
)

func testMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal(("Wrong mascot! :("))
	}
}
