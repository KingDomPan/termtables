// Copyright 2012-2013 Apcera Inc. All rights reserved.
package termtables

import (
	"testing"
)

func DisplayFailedOutput(actual, expected string) string {
	return "Output didn't match expected\r\n\r\n" +
		"Actual:\r\n\r\n" +
		actual + "\r\n" +
		"Expected:\r\n\r\n" +
		expected
}

func checkRendersTo(t *testing.T, table *Table, expected string) {
	output := table.Render()
	if output != expected {
		t.Fatal(DisplayFailedOutput(output, expected))
	}
}

func TestCreateTable(t *testing.T) {
	expected := "" +
		"+-----------+-------+\r\n" +
		"| Name      | Value |\r\n" +
		"+-----------+-------+\r\n" +
		"| hey       | you   |\r\n" +
		"| ken       | 1234  |\r\n" +
		"| derek     | 3.14  |\r\n" +
		"| derek too | 3.15  |\r\n" +
		"| escaping  | rox%% |\r\n" +
		"+-----------+-------+\r\n"

	table := CreateTable()

	table.AddHeaders("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)
	table.AddRow("escaping", "rox%%")

	checkRendersTo(t, table, expected)
}

func TestStyleResets(t *testing.T) {
	expected := "" +
		"+-----------+-------+\r\n" +
		"| Name      | Value |\r\n" +
		"+-----------+-------+\r\n" +
		"| hey       | you   |\r\n" +
		"| ken       | 1234  |\r\n" +
		"| derek     | 3.14  |\r\n" +
		"| derek too | 3.15  |\r\n" +
		"+-----------+-------+\r\n"

	table := CreateTable()
	table.UTF8Box()
	table.Style.setAsciiBoxStyle()

	table.AddHeaders("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)

	checkRendersTo(t, table, expected)
}

func TestTableWithHeader(t *testing.T) {
	expected := "" +
		"+-------------------+\r\n" +
		"|      Example      |\r\n" +
		"+-----------+-------+\r\n" +
		"| Name      | Value |\r\n" +
		"+-----------+-------+\r\n" +
		"| hey       | you   |\r\n" +
		"| ken       | 1234  |\r\n" +
		"| derek     | 3.14  |\r\n" +
		"| derek too | 3.15  |\r\n" +
		"+-----------+-------+\r\n"

	table := CreateTable()

	table.AddTitle("Example")
	table.AddHeaders("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)

	checkRendersTo(t, table, expected)
}

func TestTableTitleWidthAdjusts(t *testing.T) {
	expected := "" +
		"+---------------------------+\r\n" +
		"| Example My Foo Bar'd Test |\r\n" +
		"+-----------+---------------+\r\n" +
		"| Name      | Value         |\r\n" +
		"+-----------+---------------+\r\n" +
		"| hey       | you           |\r\n" +
		"| ken       | 1234          |\r\n" +
		"| derek     | 3.14          |\r\n" +
		"| derek too | 3.15          |\r\n" +
		"+-----------+---------------+\r\n"

	table := CreateTable()

	table.AddTitle("Example My Foo Bar'd Test")
	table.AddHeaders("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)

	checkRendersTo(t, table, expected)
}

func TestTableHeaderWidthAdjusts(t *testing.T) {
	expected := "" +
		"+---------------+---------------------+\r\n" +
		"| Slightly Long | More than 2 columns |\r\n" +
		"+---------------+---------------------+\r\n" +
		"| a             | b                   |\r\n" +
		"+---------------+---------------------+\r\n"

	table := CreateTable()

	table.AddHeaders("Slightly Long", "More than 2 columns")
	table.AddRow("a", "b")

	checkRendersTo(t, table, expected)
}

func TestTableWithNoHeaders(t *testing.T) {
	expected := "" +
		"+-----------+------+\r\n" +
		"| hey       | you  |\r\n" +
		"| ken       | 1234 |\r\n" +
		"| derek     | 3.14 |\r\n" +
		"| derek too | 3.15 |\r\n" +
		"+-----------+------+\r\n"

	table := CreateTable()

	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)

	checkRendersTo(t, table, expected)
}

func TestTableUnicodeWidths(t *testing.T) {
	expected := "" +
		"+-----------+------+\r\n" +
		"| Name      | Cost |\r\n" +
		"+-----------+------+\r\n" +
		"| Currency  | ¤10  |\r\n" +
		"| US Dollar | $30  |\r\n" +
		"| Euro      | €27  |\r\n" +
		"| Thai      | ฿70  |\r\n" +
		"+-----------+------+\r\n"

	table := CreateTable()
	table.AddHeaders("Name", "Cost")
	table.AddRow("Currency", "¤10")
	table.AddRow("US Dollar", "$30")
	table.AddRow("Euro", "€27")
	table.AddRow("Thai", "฿70")

	checkRendersTo(t, table, expected)
}

func TestTableInUTF8(t *testing.T) {
	expected := "" +
		"╭───────────────────╮\r\n" +
		"│      Example      │\r\n" +
		"├───────────┬───────┤\r\n" +
		"│ Name      │ Value │\r\n" +
		"├───────────┼───────┤\r\n" +
		"│ hey       │ you   │\r\n" +
		"│ ken       │ 1234  │\r\n" +
		"│ derek     │ 3.14  │\r\n" +
		"│ derek too │ 3.15  │\r\n" +
		"│ escaping  │ rox%% │\r\n" +
		"╰───────────┴───────╯\r\n"

	table := CreateTable()
	table.UTF8Box()

	table.AddTitle("Example")
	table.AddHeaders("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)
	table.AddRow("escaping", "rox%%")

	checkRendersTo(t, table, expected)
}

func TestTableUnicodeUTF8AndSGR(t *testing.T) {
	// at present, this mostly just tests that alignment still works
	expected := "" +
		"╭───────────────────────╮\r\n" +
		"│       \033[1mFanciness\033[0m       │\r\n" +
		"├──────────┬────────────┤\r\n" +
		"│ \033[31mred\033[0m      │ \033[32mgreen\033[0m      │\r\n" +
		"├──────────┼────────────┤\r\n" +
		"│ plain    │ text       │\r\n" +
		"│ Καλημέρα │ κόσμε      │\r\n" +
		"│ \033[1mvery\033[0m     │ \033[4munderlined\033[0m │\r\n" +
		"│ a\033[1mb\033[0mc      │ \033[45mmagenta\033[0m    │\r\n" +
		"│ \033[31m→\033[0m        │ \033[32m←\033[0m          │\r\n" +
		"╰──────────┴────────────╯\r\n"

	sgred := func(in string, sgrPm string) string {
		return "\033[" + sgrPm + "m" + in + "\033[0m"
	}
	bold := func(in string) string { return sgred(in, "1") }

	table := CreateTable()
	table.UTF8Box()

	table.AddTitle(bold("Fanciness"))
	table.AddHeaders(sgred("red", "31"), sgred("green", "32"))
	table.AddRow("plain", "text")
	table.AddRow("Καλημέρα", "κόσμε") // from http://plan9.bell-labs.com/sys/doc/utf.html
	table.AddRow(bold("very"), sgred("underlined", "4"))
	table.AddRow("a"+bold("b")+"c", sgred("magenta", "45"))
	table.AddRow(sgred("→", "31"), sgred("←", "32"))
	// TODO: in future, if we start detecting presence of SGR sequences, we
	// should ensure that the SGR reset is done at the end of the cell content,
	// so that SGR doesn't "bleed across" (cells or rows).  We would then add
	// tests for that here.
	//
	// Of course, at that point, we'd also want to support automatic HTML
	// styling conversion too, so would need a test for that also.

	checkRendersTo(t, table, expected)
}

func TestTableInMarkdown(t *testing.T) {
	expected := "" +
		"Table: Example\r\n\r\n" +
		"| Name  | Value |\r\n" +
		"| ----- | ----- |\r\n" +
		"| hey   | you   |\r\n" +
		"| a &#x7c; b | esc   |\r\n" +
		"| esc   | rox%% |\r\n"

	table := CreateTable()
	table.SetModeMarkdown()

	table.AddTitle("Example")
	table.AddHeaders("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("a | b", "esc")
	table.AddRow("esc", "rox%%")

	checkRendersTo(t, table, expected)
}

func TestTitleUnicodeWidths(t *testing.T) {
	expected := "" +
		"+-------+\r\n" +
		"| ← 5 → |\r\n" +
		"+---+---+\r\n" +
		"| a | b |\r\n" +
		"| c | d |\r\n" +
		"| e | 3 |\r\n" +
		"+---+---+\r\n"

	// minimum width for a table of two columns is 9 characters, given
	// one space of padding, and non-empty tables.

	table := CreateTable()

	// We have 4 characters down for left and right columns and padding, so
	// a width of 5 for us should match the minimum per the columns

	// 5 characters; each arrow is three octets in UTF-8, giving 9 bytes
	// so, same in character-count-width, longer in bytes
	table.AddTitle("← 5 →")

	// a single character per cell, here; use ASCII characters
	table.AddRow("a", "b")
	table.AddRow("c", "d")
	table.AddRow("e", 3)

	checkRendersTo(t, table, expected)
}

// We identified two error conditions wherein length wrapping would not correctly
// wrap width when, for instance, in a two-column table, the longest row in the
// right-hand column was not the same as the longest row in the left-hand column.
// This tests that we correctly accumulate the maximum width across all rows of
// the termtable and adjust width accordingly.
func TestTableWidthHandling(t *testing.T) {
	expected := "" +
		"+-----------------------------------------+\r\n" +
		"|        Example... to Fix My Test        |\r\n" +
		"+-----------------+-----------------------+\r\n" +
		"| hey foo bar baz | you                   |\r\n" +
		"| ken             | you should write code |\r\n" +
		"| derek           | 3.14                  |\r\n" +
		"| derek too       | 3.15                  |\r\n" +
		"+-----------------+-----------------------+\r\n"

	table := CreateTable()

	table.AddTitle("Example... to Fix My Test")
	table.AddRow("hey foo bar baz", "you")
	table.AddRow("ken", "you should write code")
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)

	output := table.Render()
	if output != expected {
		t.Fatal(DisplayFailedOutput(output, expected))
	}

}

func TestTableWidthHandling_SecondErrorCondition(t *testing.T) {
	expected := "" +
		"+----------------------------------------+\r\n" +
		"|       Example... to Fix My Test        |\r\n" +
		"+-----------------+----------------------+\r\n" +
		"| hey foo bar baz | you                  |\r\n" +
		"| ken             | you should sell cod! |\r\n" +
		"| derek           | 3.14                 |\r\n" +
		"| derek too       | 3.15                 |\r\n" +
		"+-----------------+----------------------+\r\n"

	table := CreateTable()

	table.AddTitle("Example... to Fix My Test")
	table.AddRow("hey foo bar baz", "you")
	table.AddRow("ken", "you should sell cod!")
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)

	output := table.Render()
	if output != expected {
		t.Fatal(DisplayFailedOutput(output, expected))
	}
}

func TestTableAlignPostsetting(t *testing.T) {
	expected := "" +
		"+-----------+-------+\r\n" +
		"| Name      | Value |\r\n" +
		"+-----------+-------+\r\n" +
		"|       hey | you   |\r\n" +
		"|       ken | 1234  |\r\n" +
		"|     derek | 3.14  |\r\n" +
		"| derek too | 3.15  |\r\n" +
		"|  escaping | rox%% |\r\n" +
		"+-----------+-------+\r\n"

	table := CreateTable()

	table.AddHeaders("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)
	table.AddRow("escaping", "rox%%")

	table.SetAlign(AlignRight, 1)

	checkRendersTo(t, table, expected)
}

func TestTableMissingCells(t *testing.T) {
	expected := "" +
		"+----------+---------+---------+\r\n" +
		"| Name     | Value 1 | Value 2 |\r\n" +
		"+----------+---------+---------+\r\n" +
		"| hey      | you     | person  |\r\n" +
		"| ken      | 1234    |\r\n" +
		"| escaping | rox%s%% |\r\n" +
		"+----------+---------+---------+\r\n"
		// FIXME: missing extra cells there

	table := CreateTable()

	table.AddHeaders("Name", "Value 1", "Value 2")
	table.AddRow("hey", "you", "person")
	table.AddRow("ken", 1234)
	table.AddRow("escaping", "rox%s%%")

	checkRendersTo(t, table, expected)
}

// We don't yet support combining characters, double-width characters or
// anything to do with estimating a tty-style "character width" for what in
// Unicode is a grapheme cluster.  This disabled test shows what we want
// to support, but don't yet.
func TestTableWithCombiningChars(t *testing.T) {
	t.Skip("FIXME: not implemented: grapheme cluster support & combining characters")
	expected := "" +
		"+------+---+\r\n" +
		"| noel | 1 |\r\n" +
		"| noël | 2 |\r\n" +
		"| noël | 3 |\r\n" +
		"+------+---+\r\n"

	table := CreateTable()

	table.AddRow("noel", "1")
	table.AddRow("noe\u0308l", "2") // LATIN SMALL LETTER E  +  COMBINING DIAERESIS
	table.AddRow("noël", "3")       // Hex EB; LATIN SMALL LETTER E WITH DIAERESIS

	checkRendersTo(t, table, expected)
}

// another unicode length issue
func TestTableWithFullwidthChars(t *testing.T) {
	t.Skip("FIXME: not implemented: grapheme cluster support & widechars")
	expected := "" +
		"+----------+------------+\r\n" +
		"| wide     | not really |\r\n" +
		"| ｗｉｄｅ | fullwidth  |\r\n" +
		"+----------+------------+\r\n"

	table := CreateTable()
	table.AddRow("wide", "not really")
	table.AddRow("ｗｉｄｅ", "fullwidth") // FULLWIDTH LATIN SMALL LETTER <X>

	checkRendersTo(t, table, expected)
}

func TestTableMultipleAddHeader(t *testing.T) {
	expected := "" +
		"+--------------+--------+-------+\r\n" +
		"| First column | Second | Third |\r\n" +
		"+--------------+--------+-------+\r\n" +
		"| 2            | 3      | 5     |\r\n" +
		"+--------------+--------+-------+\r\n"

	table := CreateTable()
	table.AddHeaders("First column", "Second")
	table.AddHeaders("Third")
	table.AddRow(2, 3, 5)

	checkRendersTo(t, table, expected)
}
