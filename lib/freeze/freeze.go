package freeze

type Freeze interface {
    IsFrozen() bool
    Unwrap() string
    Bind(fn func (string) Freeze) Freeze
}

// (string -> Freeze) -> (Freeze -> Freeze)
func Then(fn func (string) Freeze) func (Freeze) Freeze {
    return func (freeze Freeze) Freeze {
        return freeze.Bind(fn)
    }
}

type modifiable struct {
    text string
}

// string -> Freeze
func Wrap(text string) Freeze {
    return modifiable{text: text}
}

// modifiable -> bool
func (modifiable modifiable) IsFrozen() bool {
    return false
}

// modifiable -> bool
func (modifiable modifiable) Unwrap() string {
    return modifiable.text
}

// modifiable -> (string -> Freeze) -> Freeze
func (modifiable modifiable) Bind(fn func (string) Freeze) Freeze {
    return fn(modifiable.text)
}

type frozen struct {
    text string
}

// string -> Freeze
func WrapFrozen(text string) Freeze {
    return frozen{text: text}
}

// frozen -> string
func (frozen frozen) Unwrap() string {
    return frozen.text
}

// frozen -> bool
func (frozen frozen) IsFrozen() bool {
    return true
}

// frozen -> (string -> Freeze) -> Freeze
func (frozen frozen) Bind(fn func (string) Freeze) Freeze {
    return frozen;
}
