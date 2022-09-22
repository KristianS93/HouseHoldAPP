package validation

import (
	"testing"
)

const errorFormat string = "\nGot: %v\nWant: %v\nGiven: %v\n"

// reflect.DeepEqual was not appropriate in this scenario
func testSliceEquality(t testing.TB, a, b []string) bool {
	t.Helper()
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestCheckPassword(t *testing.T) {
	table := []struct {
		desc  string
		given string
		want  []string
	}{
		// length
		{
			"too short",
			"a",
			[]string{passwordErrLength},
		},
		{
			"too long",
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			[]string{passwordErrLength},
		},
		// one error message
		{
			"no upper",
			"aaa333###",
			[]string{passwordErrUpper},
		},
		{
			"no lower",
			"AAA333###",
			[]string{passwordErrLower},
		},
		{
			"no number",
			"aaaAAA###",
			[]string{passwordErrNumber},
		},
		{
			"no symbol",
			"aaaAAA333",
			[]string{passwordErrSymbol},
		},
		// two error messages
		{
			"no upper or lower",
			"######333",
			[]string{passwordErrUpper, passwordErrLower},
		},
		{
			"no upper or number",
			"######aaa",
			[]string{passwordErrUpper, passwordErrNumber},
		},
		{
			"no upper or symbol",
			"333333aaa",
			[]string{passwordErrUpper, passwordErrSymbol},
		},
		{
			"no lower or number",
			"######AAA",
			[]string{passwordErrLower, passwordErrNumber},
		},
		{
			"no lower or symbol",
			"333333AAA",
			[]string{passwordErrLower, passwordErrSymbol},
		},
		{
			"no number or symbol",
			"aaaaaaAAA",
			[]string{passwordErrNumber, passwordErrSymbol},
		},
		// three error messages
		{
			"no upper or lower or number",
			"###########",
			[]string{passwordErrUpper, passwordErrLower, passwordErrNumber},
		},
		{
			"no upper or lower or symbol",
			"33333333333",
			[]string{passwordErrUpper, passwordErrLower, passwordErrSymbol},
		},
		{
			"no upper or number or symbol",
			"aaaaaaaaaaa",
			[]string{passwordErrUpper, passwordErrNumber, passwordErrSymbol},
		},
		{
			"no lower or number or symbol",
			"AAAAAAAAAAA",
			[]string{passwordErrLower, passwordErrNumber, passwordErrSymbol},
		},
		// four error messages - 3 normal + 1 invalid
		{
			"no upper or lower or number and invalid chars",
			"########==",
			[]string{passwordErrUpper, passwordErrLower, passwordErrNumber, passwordErrInvalidChar([]string{"=", "="})},
		},
		{
			"no upper or lower or symbol and invalid chars",
			"33333333==",
			[]string{passwordErrUpper, passwordErrLower, passwordErrSymbol, passwordErrInvalidChar([]string{"=", "="})},
		},
		{
			"no upper or number or symbol and invalid chars",
			"aaaaaaaa==",
			[]string{passwordErrUpper, passwordErrNumber, passwordErrSymbol, passwordErrInvalidChar([]string{"=", "="})},
		},
		{
			"no lower or number or symbol and invalid chars",
			"AAAAAAAA==",
			[]string{passwordErrLower, passwordErrNumber, passwordErrSymbol, passwordErrInvalidChar([]string{"=", "="})},
		},
		// five error messages - only invalid chars
		{
			"only invalid chars",
			"========",
			[]string{passwordErrUpper, passwordErrLower, passwordErrNumber, passwordErrSymbol, passwordErrInvalidChar([]string{"=", "=", "=", "=", "=", "=", "=", "="})},
		},
		// invalid chars - only invalid error
		{
			"only invalid error",
			"iCLU%Bk#FTWpu92E3y#Y=",
			[]string{passwordErrInvalidChar([]string{"="})},
		},
		// normal use case
		{
			"bitwarden0",
			"kW5n&iCLU%Bk#FTWpu92E3y#Y@rNjSL$",
			[]string{},
		},
		{
			"bitwarden1",
			"r%EZi4%JRgEw!2YtW5tTi2nUdb3%CET#",
			[]string{},
		},
		{
			"bitwarden2",
			"UX@@mB!S$oB6SMJzqcivtt8Z4MdKxEh2",
			[]string{},
		},
		{
			"bitwarden3",
			"jcDpgvn@5VVQ5#&raxoHmaxRrn$%b!vL",
			[]string{},
		},
		{
			"bitwarden4",
			"qnZ#7DA5&@!yT85DT8ZDw%3RQM#SS3@C",
			[]string{},
		},
		{
			"bitwarden5",
			"kW5n&iCLU%Bk#FTWpu92E3y#Y@rNjSL$",
			[]string{},
		},
		{
			"bitwarden6",
			"PvGJms*KJ6yzY6Wraoj$^h$oSAiwZGrd",
			[]string{},
		},
		{
			"bitwarden7",
			"@hkQ*toCT%f5yTAFW9uUsh%o$LgqBx2j",
			[]string{},
		},
		{
			"bitwarden8",
			"A6cFSUqo#%swz^C^Dc9XmEw5jpAwBszC",
			[]string{},
		},
		{
			"bitwarden9",
			"VeiEb32f#VmxQqPf29gP5QptBPC*22rn",
			[]string{},
		},
	}

	for _, test := range table {
		t.Run(test.desc, func(t *testing.T) {
			got := CheckPassword(test.given)
			if !testSliceEquality(t, got, test.want) {
				t.Errorf(errorFormat, got, test.want, test.given)
			}
		})
	}
}

func TestCheckEmail(t *testing.T) {
	table := []struct {
		desc  string
		given string
		want  []string
	}{
		//@ issues - formatting
		{
			"no @",
			"example",
			[]string{emailErrAt},
		},
		{
			"more than one @",
			"example@example@example.com",
			[]string{emailErrAt},
		},
		// domain issues
		{
			"invalid domain",
			"example@fjkalhfjkldsa.com",
			[]string{emailErrDomain},
		},
		{
			"valid domain trailing space",
			"example@example.com ",
			[]string{emailErrDomain},
		},
		// punctuation
		{
			"double period",
			"example..test@example.com",
			[]string{emailErrPunctuation},
		},
		{
			"double hyphen",
			"example--test@example.com",
			[]string{emailErrPunctuation},
		},
		{
			"double underscore",
			"example__test@example.com",
			[]string{emailErrPunctuation},
		},
		// false positives in punctuation
		{
			"period letter period",
			"example.h.test@example.com",
			[]string{},
		},
		{
			"hyphen letter hyphen",
			"example-h-test@example.com",
			[]string{},
		},
		{
			"underscore letter underscore",
			"example_h_test@example.com",
			[]string{},
		},
		{
			"alternating letter and punctuation",
			"a.b-c_d-e_f.@example.com",
			[]string{},
		}, // found index check bug
		// invalid characters
		{
			"invalid character #",
			"hello#world@example.com",
			[]string{emailErrInvalidChar([]string{"#"})},
		},
		{
			"invalid character # near punctuation",
			"hello.#.world@example.com",
			[]string{emailErrInvalidChar([]string{"#"})},
		},
		// normal use case
		{
			"use case 0",
			"prettyandsimple@example.com",
			[]string{},
		},
		{
			"use case 1",
			"very.common@example.com",
			[]string{},
		},
		{
			"use case 2",
			"other.email-with-dash@example.com",
			[]string{},
		},
		{
			"use case 3",
			"x@example.com",
			[]string{},
		},
		{
			"use case 4",
			"user@yahoo.com",
			[]string{},
		},
		{
			"use case 5",
			"example@hotmail.com",
			[]string{},
		},
	}

	for _, test := range table {
		t.Run(test.desc, func(t *testing.T) {
			got := CheckEmail(test.given)
			if !testSliceEquality(t, got, test.want) {
				t.Errorf(errorFormat, got, test.want, test.given)
			}
		})
	}
}

func BenchmarkValidEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckEmail("very.common@example.com")
	}
}

func BenchmarkValidPasswordLen32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckPassword("wp*Vr5#rSwDQ4XKafa8a*Xt3Vgd8X3^s")
	}
}
