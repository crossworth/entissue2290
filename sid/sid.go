package sid

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type ID string

func (i ID) Value() (driver.Value, error) {
	r, err := strconv.ParseInt(string(i), 10, 64)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (i *ID) Scan(src interface{}) error {
	switch v := src.(type) {
	case nil:
		return nil
	case int64:
		*i = ID(fmt.Sprint(v))
		return nil
	}
	return errors.New("not a valid base62")
}

func New() ID {
	return NewLength(10)
}
func NewLength(len int) ID {
	return ID(RandStringRunes(len))
}

// from: https://stackoverflow.com/a/31832326/1241602
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
