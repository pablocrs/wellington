package sprite_sass_test

import (
	"testing"

	. "github.com/drewwells/sprite_sass"
)

func TestParser(t *testing.T) {
	p := Parser{}
	p.Start("test/_var.scss")
	t.Errorf("")
}

func TestImporter(t *testing.T) {
	//p := Parser{}
	//p.Start("test/import.scss")

}
