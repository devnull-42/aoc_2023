package day20

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartA(t *testing.T) {

	var testData = []struct {
		name     string
		raw      string
		expected int
	}{
		{
			name: "1 press cycle",
			raw: `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`,
			expected: 32000000,
		},
		{
			name: "4 press cycle",
			raw: `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`,
			expected: 11687500,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			lines := strings.Split(tt.raw, "\n")
			result := partA(lines)
			assert.Equal(t, tt.expected, result, "PartA(%v) should equal %d", lines, tt.expected)
		})
	}
}
