package connections_test

import (
	"testing"

	"github.com/semenovDL/gqlutils/connections"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	prefix string
	offset int
	cursor string
	err    error
}

var testCases = []testCase{
	{"", -1, "LTE=", nil},
	{"", 0, "MA==", nil},
	{"", 1, "MQ==", nil},
	{"", 2, "Mg==", nil},
	{"", 100000, "MTAwMDAw", nil},
	{"cursor:", 0, "Y3Vyc29yOjA=", nil},
	{"cursor:", 1, "Y3Vyc29yOjE=", nil},
}

func TestConnections_Paginator_OffsetToCursor(t *testing.T) {
	for _, c := range testCases {
		p := connections.Paginator{CursorPrefix: c.prefix}
		assert.Equal(t, c.cursor, p.OffsetToCursor(c.offset))
	}
}

func TestConnections_Paginator_CursorToOffset(t *testing.T) {
	for _, c := range testCases {
		p := connections.Paginator{CursorPrefix: c.prefix}
		value, err := p.CursorToOffset(c.cursor)
		assert.Equal(t, c.offset, value)
		assert.Equal(t, c.err, err)
	}
}

func TestConnections_Paginator_CursorToOffset_BadCursor(t *testing.T) {
	p := connections.Paginator{}
	value, err := p.CursorToOffset("M{==")
	assert.Equal(t, 0, value)
	assert.Contains(t, err.Error(), "invalid cursor M{==")
}
