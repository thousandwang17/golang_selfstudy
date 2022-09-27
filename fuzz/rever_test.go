package main

import (
	"testing"
	"unicode/utf8"
)

func TestRever(t *testing.T) {
	for _, tc := range []struct {
		in, want string
	}{
		{"hello , word!", "!drow , olleh"},
		{" ", " "},
		{"!12345", "54321!"},
	} {
		rev, err := rever(tc.in)
		if err != nil {
			return
		}
		if rev != tc.want {
			t.Errorf("rever: %q , want %q", rev, tc.want)
		}
	}
}

func FuzzRever(f *testing.F) {
	for _, v := range []string{
		"hello , word!",
		" ",
		"!12345",
	} {
		f.Add(v)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := rever(orig)

		if err1 != nil {
			return
		}

		doubleRev, err2 := rever(rev)

		if err2 != nil {
			return
		}

		if doubleRev != orig {
			t.Errorf("rever Rec : %q, Want : %q", orig, doubleRev)
		}

		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("rever invaild utf8 fails rev : %q", rev)
		}
	})
}
