package iconv

import (
	"testing"

	"golang.org/x/text/encoding/charmap"
)

const (
	src      = `Jaime De ArmiÃ±Ã¡n y HÃ©ctor Alterio`
	expected = `Jaime De Armiñán y Héctor Alterio`
)

func TestCharmapWindows1252Conversion(t *testing.T) {
	dst, err := charmap.Windows1252.NewDecoder().Bytes([]byte(src))
	if err != nil {
		t.Fatal(err)
	}

	if string(dst) != expected {
		t.Fatalf("expected %q but got %q instead", expected, string(dst))
	}
}

func TestIconvWindows1252Conversion(t *testing.T) {
	cd, err := Open("WINDOWS-1252", "UTF-8")
	if err != nil {
		t.Fatalf("Error on opening: %s", err)
	}

	str, err := cd.Conv(src)
	if err != nil {
		t.Fatalf("Error on conversion: %s", err)
	}

	if str != expected {
		t.Errorf("Unexpected value: %#v (expected %#v)", str, expected)
	}

	err = cd.Close()
	if err != nil {
		t.Fatalf("Error on close: %s", err)
	}
}
