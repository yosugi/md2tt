package main

import (
	"testing"

	"./lib/freeze"
)

func TestHeading(test *testing.T) {
	actual := heading("# hoge")
	expect := freeze.WrapFrozen("h1. hoge")
	if actual != expect {
		test.Errorf("{ actual: %v, expect: %v }", actual, expect)
	}

	actual = heading("## hoge")
	expect = freeze.WrapFrozen("h2. hoge")
	if actual != expect {
		test.Errorf("{ actual: %v, expect: %v }", actual, expect)
	}

	actual = heading("### hoge")
	expect = freeze.WrapFrozen("h3. hoge")
	if actual != expect {
		test.Errorf("{ actual: %v, expect: %v }", actual, expect)
	}

	actual = heading("#### hoge")
	expect = freeze.WrapFrozen("h4. hoge")
	if actual != expect {
		test.Errorf("{ actual: %v, expect: %v }", actual, expect)
	}

	actual = heading("##### hoge")
	expect = freeze.WrapFrozen("h5. hoge")
	if actual != expect {
		test.Errorf("{ actual: %v, expect: %v }", actual, expect)
	}
}

func TestBulletList4(test *testing.T) {
	bulletList4 := bulletList(4)

	actual := bulletList4("* hoge")
	expect := freeze.WrapFrozen("* hoge")
	if actual != expect {
		test.Errorf("{ actual: %v, expect: %v }", actual, expect)
	}

	actual = bulletList4("    - hoge")
	expect = freeze.WrapFrozen("** hoge")
	if actual != expect {
		test.Errorf("{ actual: %v, expect: %v }", actual, expect)
	}

	actual = bulletList4("        + hoge")
	expect = freeze.WrapFrozen("*** hoge")
	if actual != expect {
		test.Errorf("{ actual: %v, expect: %v }", actual, expect)
	}
}

func TestBulletList2(test *testing.T) {
	bulletList2 := bulletList(2)

	actual := bulletList2("* hoge")
	expect := freeze.WrapFrozen("* hoge")
	if actual != expect {
		test.Errorf("{ actual: %v, expect: %v }", actual, expect)
	}

	actual = bulletList2("  - hoge")
	expect = freeze.WrapFrozen("** hoge")
	if actual != expect {
		test.Errorf("{ actual: %v, expect: %v }", actual, expect)
	}

	actual = bulletList2("    + hoge")
	expect = freeze.WrapFrozen("*** hoge")
	if actual != expect {
		test.Errorf("{ actual: %v, expect: %v }", actual, expect)
	}
}

func TestNumberedList4(test *testing.T) {
	numberedList4 := numberedList(4)

	actual := numberedList4("1. hoge")
	expect := freeze.WrapFrozen("# hoge")
	if actual != expect {
		test.Errorf("{ actual: %v, expect: %v }", actual, expect)
	}

	actual = numberedList4("    10. hoge")
	expect = freeze.WrapFrozen("## hoge")
	if actual != expect {
		test.Errorf("{ actual: %v, expect: %v }", actual, expect)
	}

	actual = numberedList4("        100. hoge")
	expect = freeze.WrapFrozen("### hoge")
	if actual != expect {
		test.Errorf("{ actual: %v, expect: %v }", actual, expect)
	}
}
