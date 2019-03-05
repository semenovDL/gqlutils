package scalars_test

import (
	"errors"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/semenovDL/gqlutils/scalars"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRFC3339Time_MarshalGQL(t *testing.T) {
	t1, err := time.Parse(time.RFC822Z, "02 Jan 06 15:04 -0700")
	require.Nil(t, err)
	dt1 := scalars.RFC3339Time(t1)
	buf := strings.Builder{}
	dt1.MarshalGQL(&buf)
	assert.Equal(t, `"2006-01-02T22:04:00Z"`, buf.String())
}

type badWriter struct{}

func (*badWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("bad write")
}

func TestRFC3339Time_MarshalGQL_BadWriter(t *testing.T) {
	dt1 := scalars.RFC3339Time(time.Now())
	buf := badWriter{}
	logOut := strings.Builder{}
	log.SetOutput(&logOut)
	dt1.MarshalGQL(&buf)
	assert.Contains(t, logOut.String(), "bad write")
}

func TestRFC3339Time_UnmarshalGQL(t *testing.T) {
	str := "2006-01-02T22:04:00Z"
	dt2 := scalars.RFC3339Time{}
	err := dt2.UnmarshalGQL(str)
	assert.Nil(t, err)
	assert.Equal(t, str, dt2.String())
}

func TestRFC3339Time_UnmarshalGQL_WrongInputType(t *testing.T) {
	dt2 := scalars.RFC3339Time{}
	err := dt2.UnmarshalGQL(1)
	assert.Equal(t, errors.New("value must be a strings"), err)
}

func TestRFC3339Time_UnmarshalGQL_WrongInputFormat(t *testing.T) {
	str := "02 Jan 06 15:04 -0700"
	dt2 := scalars.RFC3339Time{}
	err := dt2.UnmarshalGQL(str)
	assert.Contains(t, err.Error(), "must be in RFC3339 format")
	assert.Equal(t, "0001-01-01T00:00:00Z", dt2.String())
}
