package connections

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

// Paginator allows to paginate a collection
type Paginator struct {
	CursorPrefix string
}

// OffsetToCursor creates the cursor string from an offset
func (p *Paginator) OffsetToCursor(offset int) string {
	str := fmt.Sprintf("%v%v", p.CursorPrefix, offset)
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// CursorToOffset re-derives the offset from the cursor string.
func (p *Paginator) CursorToOffset(cursor string) (int, error) {
	str := ""
	b, err := base64.StdEncoding.DecodeString(cursor)
	if err == nil {
		str = string(b)
	}
	str = strings.Replace(str, p.CursorPrefix, "", -1)
	offset, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("invalid cursor %s: %s", cursor, err)
	}
	return offset, nil
}
