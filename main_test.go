package main

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

func Test_toCSVRecord(t *testing.T) {
	for _, tc := range []struct {
		rec  record
		want []string
	}{
		{rec1, []string{"Foo", "1", "2.1", "false", ""}},
		{rec2, []string{"Bar", "4", "6.75", "true", ""}},
		{rec3, []string{"Baz", "2", "2.5", "false", ptr2}},
	} {
		if got := tc.rec.toCSVRecord(); !slices.Equal(got, tc.want) {
			t.Errorf("toCSVRecord(%v)=%q; want %q", tc.rec, got, tc.want)
		}
	}
}

func Test_toJTTRow(t *testing.T) {
	for _, tc := range []struct {
		rec  record
		want string
	}{
		{rec1, `"Foo"	1	2.1	false	""`},
		{rec2, `"Bar"	4	6.75	true	null`},
		{rec3, `"Baz"	2	2.5	false	"wants to be promoted to \"Boss\", and\nget a raise"`},
	} {
		if got := tc.rec.toJTTRow(); got != tc.want {
			t.Errorf("toJTTRow(%v);\n got: %v\nwant: %v", tc.rec, pretty(got), pretty(tc.want))
		}
	}
}

func pretty(s string) string {
	s = strings.ReplaceAll(fmt.Sprintf("%q", s), "\\\"", "|")
	s = strings.ReplaceAll(s, "\"", "")
	s = strings.ReplaceAll(s, "|", "\"")
	s = strings.ReplaceAll(s, "\\t", "|")
	return s
}
