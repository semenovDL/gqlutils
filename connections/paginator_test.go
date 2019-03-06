package connections_test

import (
	"errors"
	"fmt"
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

func TestConnections_Paginator_OffsetToCursor(t *testing.T) {
	testCases := []testCase{
		{"", 0, "MA==", nil},
		{"", 1, "MQ==", nil},
		{"", 2, "Mg==", nil},
		{"", 100000, "MTAwMDAw", nil},
		{"cursor:", 0, "Y3Vyc29yOjA=", nil},
		{"cursor:", 1, "Y3Vyc29yOjE=", nil},
		{"", -1, "", errors.New("negative offset -1")},
		{"", -2, "", errors.New("negative offset -2")},
	}
	for _, tc := range testCases {
		func(tc testCase) {
			t.Run(fmt.Sprintf("%s %d", tc.prefix, tc.offset), func(t *testing.T) {
				t.Parallel()
				p := connections.Paginator{CursorPrefix: tc.prefix}
				value, err := p.OffsetToCursor(tc.offset)
				assert.Equal(t, tc.cursor, value)
				assert.Equal(t, tc.err, err)
			})
		}(tc)
	}
}

func TestConnections_Paginator_CursorToOffset(t *testing.T) {
	testCases := []testCase{
		{"", 0, "MA==", nil},
		{"", 1, "MQ==", nil},
		{"", 2, "Mg==", nil},
		{"", 100000, "MTAwMDAw", nil},
		{"cursor:", 0, "Y3Vyc29yOjA=", nil},
		{"cursor:", 1, "Y3Vyc29yOjE=", nil},
		{"", -1, "M{==", errors.New("can't decode cursor M{== from base64")},
		{"", -1, "LTE=", errors.New("negative offset -1 in cursor LTE=")},
		{"", -1, "LTI=", errors.New("negative offset -2 in cursor LTI=")},
		{"", -1, "YWJj", errors.New("cursor YWJj is not a number")},
	}
	for _, tc := range testCases {
		func(tc testCase) {
			t.Run(fmt.Sprintf("%s %s", tc.prefix, tc.cursor), func(t *testing.T) {
				t.Parallel()
				p := connections.Paginator{CursorPrefix: tc.prefix}
				value, err := p.CursorToOffset(tc.cursor)
				assert.Equal(t, tc.offset, value)
				assert.Equal(t, tc.err, err)
			})
		}(tc)
	}
}
