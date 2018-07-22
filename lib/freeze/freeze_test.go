package freeze

import "testing"

func TestWrap(test *testing.T) {
    testText := "test"

    expect := modifiable{text: testText}
    actual := Wrap(testText)
    if (actual != expect) {
        test.Errorf("{ actual: %v, expect: %v }", actual, expect)
    }
}

func TestWrapFrozen(test *testing.T) {
    testText := "test"

    expect := frozen{text: testText}
    actual := WrapFrozen(testText)
    if (actual != expect) {
        test.Errorf("{ actual: %v, expect: %v }", actual, expect)
    }
}

func TestUnwrap(test *testing.T) {
    testText := "test"

    expect := testText
    freeze := Wrap(testText)
    actual := freeze.Unwrap()
    if (actual != expect) {
        test.Errorf("{ actual: %v, expect: %v }", actual, expect)
    }

    expect = testText
    freeze = WrapFrozen(testText)
    actual = freeze.Unwrap()
    if (actual != expect) {
        test.Errorf("{ actual: %v, expect: %v }", actual, expect)
    }
}

func TestIsFrozen(test *testing.T) {
    freeze := Wrap("test")
    expect := false
    actual := freeze.IsFrozen()
    if (actual != expect) {
        test.Errorf("{ actual: %v, expect: %v }", actual, expect)
    }

    freeze = WrapFrozen("test")
    expect = true
    actual = freeze.IsFrozen()
    if (actual != expect) {
        test.Errorf("{ actual: %v, expect: %v }", actual, expect)
    }
}

func TestBind(test *testing.T) {
    beforeText := "test"
    afterText := "done"

    expect1 := modifiable{text: beforeText}
    actual1 := Wrap(beforeText)

    expect2 := modifiable{text: afterText}
    actual2 := actual1.Bind(func (text string) Freeze {
        return Wrap(afterText)
    })

    // immutability test
    if (actual1 != expect1) {
        test.Errorf("{ actual: %v, expect: %v }", actual1, expect1)
    }

    if (actual2 != expect2) {
        test.Errorf("{ actual: %v, expect: %v }", actual2, expect2)
    }

    // frozen になった後は実行されない
    freeze := Wrap(beforeText)
    expect3 := frozen{text: afterText}
    actual3 := freeze.Bind(func (text string) Freeze {
        return WrapFrozen(afterText)
    }).Bind(func (text string) Freeze {
        return WrapFrozen(afterText + "2")
    })

    if (actual3 != expect3) {
        test.Errorf("{ actual: %v, expect: %v }", actual3, expect3)
    }
}
