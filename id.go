package bmm

import (
	"fmt"
	"strconv"
)

type ID int64

func (id ID) String() string {
	return fmt.Sprintf("%d", id)
}

func Parse(s string) (ID, error) {
	id, err := strconv.ParseInt(s, 10, 64)
	return ID(id), err
}

func MustParse(s string) ID {
	id, err := Parse(s)
	if err != nil {
		panic(err)
	}

	return id
}
