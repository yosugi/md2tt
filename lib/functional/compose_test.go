package functional

import (
    "testing"

    "github.com/yosugi/md2tt/lib/freeze"
)

func TestComposeWithBind(test *testing.T) {
    appendOne := func (str string) freeze.Freeze {
        return freeze.Wrap(str + "1")
    }

    appendTwo := func (str string) freeze.Freeze {
        return freeze.WrapFrozen(str + "2")
    }

    appendThree := func (str string) freeze.Freeze {
        return freeze.Wrap(str + "3")
    }

    fn := Compose(
        freeze.Then(appendOne),
        freeze.Then(appendTwo),
        freeze.Then(appendThree),
    )

    wrappedText := freeze.Wrap("test")
    actual := fn(wrappedText)
    expect := freeze.WrapFrozen("test32")
    if (actual != expect) {
        test.Errorf("{ actual: %v, expect: %v }", actual, expect)
    }

    fn = Compose(
        freeze.Then(appendThree),
        freeze.Then(appendTwo),
        freeze.Then(appendOne),
    )

    wrappedText = freeze.Wrap("test")
    actual = fn(wrappedText)
    expect = freeze.WrapFrozen("test12")
    if (actual != expect) {
        test.Errorf("{ actual: %v, expect: %v }", actual, expect)
    }
}

func TestComposeWithoutThen(test *testing.T) {
    appendOne := func (freezeStr freeze.Freeze) freeze.Freeze {
        return freezeStr.Bind(func (str string) freeze.Freeze {
            return freeze.Wrap(str + "1")
        })
    }

    appendTwo := func (freezeStr freeze.Freeze) freeze.Freeze {
        return freezeStr.Bind(func (str string) freeze.Freeze {
            return freeze.WrapFrozen(str + "2")
        })
    }

    appendThree := func (freezeStr freeze.Freeze) freeze.Freeze {
        return freezeStr.Bind(func (str string) freeze.Freeze {
            return freeze.Wrap(str + "3")
        })
    }

    fn := Compose(
        appendOne,
        appendTwo,
        appendThree,
    )

    wrappedText := freeze.Wrap("test")
    actual := fn(wrappedText)
    expect := freeze.WrapFrozen("test32")
    if (actual != expect) {
        test.Errorf("{ actual: %v, expect: %v }", actual, expect)
    }

    fn = Compose(
        appendThree,
        appendTwo,
        appendOne,
    )
    wrappedText = freeze.Wrap("test")
    actual = fn(wrappedText)
    expect = freeze.WrapFrozen("test12")
    if (actual != expect) {
        test.Errorf("{ actual: %v, expect: %v }", actual, expect)
    }
}
