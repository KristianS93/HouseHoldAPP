package web

import (
	"testing"
)

func BenchmarkPasswordNoRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validPassword("wp*Vr5#rSwDQ4XKafa8a*Xt3Vgd8X3^s")
	}
}

func TestValidation(t *testing.T) {
	if errs := validPassword("wp*Vr5#rSwDQ4XKafa8a*Xt3Vgd8X3^s"); errs != nil {
		t.Error("noregex: failed")
	}
}

func TestValidEmail(t *testing.T) {
	type test struct {
		email    string
		expected bool
	}
	tests := []test{
		{"prettyandsimple@example.com", true},
		{"very.common@example.com", true},
		{"disposable.style.email.with+symbol@example.com", false},
		{"other.email-with-dash@example.com", true},
		{"fully-qualified-domain@example.com", true},
		{"x@example.com", true},
		{"example-indeed@example.com", true},
		{"admin@hotmail.dk", true},
		{"example@hotmail.com", true},
		{"user@student.aau.dk", true},
		{"user@yahoo.com", true},
		{"Abc.example.com", false},
		{"A@b@c@example.com", false},
		{`a"b(c)d,e:f;gi[j\k]l@example.com`, false},
		{"1234567890123456789012345678901234567890123456789012345678901234+x@example.com", false},
		{"john..doe@example.com", false},
		{"john.doe@example..com", false},
		{" krathermand@gmail.com", false},
		{"krathermand@gmail.com ", false},
	}

	for _, v := range tests {
		b := validEmail(v.email)
		if b != v.expected {
			t.Error("validEmail: failed test with -", v.email, v.expected, b)
		}
	}
}

func BenchmarkValidEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validEmail("very.common@example.com")
	}
}
