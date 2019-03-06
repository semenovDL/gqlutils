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
func (p *Paginator) OffsetToCursor(offset int) (string, error) {
	if offset < 0 {
		return "", fmt.Errorf("negative offset %d", offset)
	}
	str := fmt.Sprintf("%v%v", p.CursorPrefix, offset)
	return base64.StdEncoding.EncodeToString([]byte(str)), nil
}

// CursorToOffset re-derives the offset from the cursor string.
func (p *Paginator) CursorToOffset(cursor string) (int, error) {
	b, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return -1, fmt.Errorf("can't decode cursor %s from base64", cursor)
	}
	str := string(b)
	str = strings.Replace(str, p.CursorPrefix, "", -1)
	offset, err := strconv.Atoi(str)
	if err != nil {
		return -1, fmt.Errorf("cursor %s is not a number", cursor)
	}
	if offset < 0 {
		return -1, fmt.Errorf("negative offset %d in cursor %s", offset, cursor)
	}
	return offset, nil
}
