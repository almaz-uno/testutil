package testutil

import "github.com/pmezard/go-difflib/difflib"

func DiffForAssert(expected, actual string) string {
	diff, err := difflib.GetContextDiffString(difflib.ContextDiff{
		A:        difflib.SplitLines(expected),
		FromFile: "expected",
		B:        difflib.SplitLines(actual),
		ToFile:   "actual",
	})
	if err != nil {
		panic(err)
	}
	return "\n" + diff
}
