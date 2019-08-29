package hamming

// Source: exercism/problem-specifications
// Commit: 4119671 Hamming: Add a tests to avoid wrong recursion solution (#1450)
// Problem Specifications Version: 2.3.0

var testCases = []struct {
	s1          string
	s2          string
	want        int
	expectError bool
}{
	{ // empty strands
		"",
		"",
		0,
		false,
	},
	{ // single letter identical strands
		"A",
		"A",
		0,
		false,
	},
	{ // single letter different strands
		"G",
		"T",
		1,
		false,
	},
	{ // long identical strands
		"GGACTGAAATCTG",
		"GGACTGAAATCTG",
		0,
		false,
	},
	{ // long different strands
		"GGACGGATTCTG",
		"AGGACGGATTCT",
		9,
		false,
	},
	{ // disallow first strand longer
		"AATG",
		"AAA",
		0,
		true,
	},
	{ // disallow second strand longer
		"ATA",
		"AGTG",
		0,
		true,
	},
	{ // disallow left empty strand
		"",
		"G",
		0,
		true,
	},
	{ // disallow right empty strand
		"G",
		"",
		0,
		true,
	},
	{ // non-ascii input
		"aüa",
		"aüa",
		0,
		false,
	},
	{ // more non-ascii input
		"aüa",
		"aÃa",
		1,
		false,
	},
	{ // more non-ascii input
		"😬",
		"😬",
		0,
		false,
	},
	{ // more non-ascii input
		"😬",
		"😭",
		1,
		false,
	},
	{ // more non-ascii input
		"⌘",
		"⌘",
		0,
		false,
	},
	{ // more non-ascii input
		"⌘",
		"日",
		1,
		false,
	},
	{ // more non-ascii input
		"日本語",
		"日本語",
		0,
		false,
	},
	{ // more non-ascii input
		"日本語",
		"日語本",
		2,
		false,
	},
}
