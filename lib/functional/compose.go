package functional

import (
    "sort"

    "github.com/yosugi/md2tt/lib/freeze"
)

// (Freeze -> Freeze) -> ... -> (Freeze -> Freeze)
func Compose(funcs ...func(freeze.Freeze) freeze.Freeze) func (freeze.Freeze) freeze.Freeze {
    // cf. https://stackoverflow.com/questions/19239449/how-do-i-reverse-an-array-in-go
	sort.Slice(funcs, func(i, j int) bool {
		return true
	})

    return func (init freeze.Freeze) freeze.Freeze {
        acc := init
        for _, fn := range funcs {
            acc = fn(acc)
        }
        return acc
    }
}
