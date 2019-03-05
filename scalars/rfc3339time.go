package scalars

import (
	"fmt"
	"io"
	"log"
	"time"
)

// RFC3339Time implement scalar type for time in rfc3339 format.
type RFC3339Time time.Time

// String implements Stringer interface
func (dt RFC3339Time) String() string {
	return time.Time(dt).UTC().Format(time.RFC3339)
}

// UnmarshalGQL implements the graphql.Marshaler interface
func (dt *RFC3339Time) UnmarshalGQL(v interface{}) error {
	ts, ok := v.(string)
	if !ok {
		return fmt.Errorf("value must be a strings")
	}

	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		return fmt.Errorf("%s must be in RFC3339 format: %s", ts, err)
	}
	*dt = RFC3339Time(t)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (dt RFC3339Time) MarshalGQL(w io.Writer) {
	if _, err := w.Write([]byte(fmt.Sprintf(`"%s"`, dt))); err != nil {
		log.Printf("[ERROR] when marshaling %s: %s", dt.String(), err)
	}
}
