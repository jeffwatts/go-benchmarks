package go_benchmarks

import (
	"regexp"
	"fmt"
)

type CompileOnceObject struct {
	compiledRegexp *regexp.Regexp
}

func (compileOnceObject CompileOnceObject) FindMatchIndexes(text string) [][]int {
	return compileOnceObject.compiledRegexp.FindAllStringIndex(text, -1)
}

func NewCanonicalCompileOnceObject() CompileOnceObject {
	return CompileOnceObject{compileMasterPattern()}
}

func OnDemandFindMatchIndexes(text string) [][]int {
	regex := compileMasterPattern()
	return regex.FindAllStringIndex(text, -1)
}

func compileMasterPattern() *regexp.Regexp {
	visaPattern := `4\d{3}(?:[\s-]*)?\d{4}(?:[\s-]*)?\d{4}(?:[\s-]*)?\d{4}`
	mastercardPattern := `5[1-5]\d{2}(?:[\s-]*)?\d{4}(?:[\s-]*)?\d{4}(?:[\s-]*)?\d{4}`
	amexPattern := `3[47]\d{2}(?:[\s-]*)?\d{6}(?:[\s-]*)?\d{5}`
	discoverPattern := `6(?:011|22[1-9]|4[4-9]\d|5\d{2})(?:[\s-]*)?\d{4}(?:[\s-]*)?\d{4}(?:[\s-]*)?\d{4}`
	jcbPattern := `35(?:2[89]|[3-8]\d)(?:[\s-]*)?\d{4}(?:[\s-]*)?\d{4}(?:[\s-]*)?\d{4}`
	maestroPattern := `(?:5[0678]\d{2}|6\d{3})(?:[\s-]*)?\d{4}(?:[\s-]*)?\d{4}(?:[\s-]*)?\d{1,7}?`
	return regexp.MustCompile(fmt.Sprintf("\\b(?:(%s)|(%s)|(%s)|(%s)|(%s)|(%s))\\b", visaPattern,
		mastercardPattern, amexPattern, discoverPattern, jcbPattern, maestroPattern))
}