package actions

import "strconv"

type Int int

func (i Int) Action() string { return strconv.FormatInt(int64(i), 10) }
