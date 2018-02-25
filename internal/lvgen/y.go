//line sunrpc.y:20

// Copyright 2017 The go-libvirt Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//
// Code generated by goyacc. DO NOT EDIT.
//
// To regenerate, run 'go generate' in internal/lvgen.
//

package lvgen

import __yyfmt__ "fmt"

//line sunrpc.y:40
import (
//"fmt"
)

//line sunrpc.y:49
type yySymType struct {
	yys int
	val string
}

const BOOL = 57346
const CASE = 57347
const CONST = 57348
const DEFAULT = 57349
const DOUBLE = 57350
const ENUM = 57351
const FLOAT = 57352
const OPAQUE = 57353
const STRING = 57354
const STRUCT = 57355
const SWITCH = 57356
const TYPEDEF = 57357
const UNION = 57358
const UNSIGNED = 57359
const VOID = 57360
const HYPER = 57361
const INT = 57362
const SHORT = 57363
const CHAR = 57364
const IDENTIFIER = 57365
const CONSTANT = 57366
const ERROR = 57367
const PROGRAM = 57368
const VERSION = 57369
const METADATACOMMENT = 57370

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"BOOL",
	"CASE",
	"CONST",
	"DEFAULT",
	"DOUBLE",
	"ENUM",
	"FLOAT",
	"OPAQUE",
	"STRING",
	"STRUCT",
	"SWITCH",
	"TYPEDEF",
	"UNION",
	"UNSIGNED",
	"VOID",
	"HYPER",
	"INT",
	"SHORT",
	"CHAR",
	"IDENTIFIER",
	"CONSTANT",
	"ERROR",
	"PROGRAM",
	"VERSION",
	"METADATACOMMENT",
	"';'",
	"'{'",
	"'}'",
	"','",
	"'='",
	"'['",
	"']'",
	"'<'",
	"'>'",
	"'*'",
	"'('",
	"')'",
	"':'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line sunrpc.y:267

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 153

var yyAct = [...]int{

	86, 79, 36, 115, 107, 78, 62, 32, 68, 55,
	57, 133, 130, 132, 104, 121, 87, 88, 80, 64,
	37, 102, 75, 31, 76, 101, 114, 135, 119, 72,
	92, 41, 97, 89, 63, 40, 10, 39, 43, 42,
	13, 73, 30, 14, 38, 122, 48, 49, 50, 51,
	47, 111, 93, 81, 71, 11, 110, 138, 10, 65,
	99, 54, 13, 52, 12, 14, 29, 131, 123, 74,
	77, 112, 94, 59, 82, 15, 90, 91, 58, 70,
	16, 64, 85, 96, 87, 88, 60, 61, 95, 84,
	100, 98, 48, 49, 50, 51, 113, 59, 106, 27,
	103, 25, 109, 105, 23, 20, 18, 67, 46, 8,
	108, 45, 7, 44, 4, 109, 2, 120, 124, 117,
	126, 118, 83, 69, 127, 8, 26, 128, 7, 125,
	4, 129, 134, 28, 116, 136, 137, 53, 24, 66,
	22, 35, 34, 33, 21, 19, 56, 17, 9, 6,
	5, 3, 1,
}
var yyPact = [...]int{

	49, -1000, -1000, 51, -1000, -1000, -1000, -1000, -1000, -1000,
	83, 82, -1000, 81, 78, 76, 49, 36, -1000, 9,
	-1000, 27, 33, -1000, -1000, -1000, 31, -1000, -1000, 50,
	63, -1000, -1000, -1000, -1000, -1000, -4, -1000, 73, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 93, 52, 23, -3, 8, 74, -1000,
	-1000, -1000, -12, 58, -1000, -1000, 27, -21, 22, 45,
	66, -1000, 50, 61, 0, 61, -7, -1000, 21, 43,
	27, -1, 52, 30, -1000, -1000, -1000, -1000, -1000, 61,
	-10, -16, -1000, -1000, 27, -26, 58, 61, -1000, 27,
	-1000, -1000, -1000, -1000, 26, -1000, -1000, 20, 42, 3,
	114, -5, 27, -24, -1000, 14, 39, 61, -1000, 61,
	-1000, 27, -1000, 114, -1000, -29, 38, -27, -1000, -30,
	27, -1000, -6, 27, -1000, 61, -1000, 28, -1000,
}
var yyPgo = [...]int{

	0, 152, 116, 0, 151, 113, 150, 149, 111, 108,
	148, 147, 9, 146, 10, 145, 144, 1, 7, 143,
	142, 141, 2, 6, 20, 140, 139, 5, 138, 137,
	3, 134, 131, 129, 126, 8, 123, 122, 4, 110,
	96,
}
var yyR1 = [...]int{

	0, 1, 3, 3, 2, 2, 4, 4, 4, 4,
	4, 4, 5, 12, 12, 13, 13, 13, 11, 14,
	6, 6, 15, 16, 7, 17, 17, 17, 17, 18,
	22, 22, 22, 22, 22, 22, 22, 22, 22, 22,
	22, 24, 24, 24, 24, 23, 19, 20, 20, 21,
	26, 8, 25, 27, 27, 29, 9, 28, 30, 30,
	32, 31, 33, 31, 10, 34, 35, 35, 36, 37,
	38, 38, 39, 40,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 2, 3, 1, 1, 1, 1,
	1, 1, 5, 1, 3, 1, 4, 3, 1, 1,
	4, 4, 1, 0, 3, 1, 1, 1, 1, 2,
	1, 2, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 5, 5, 4, 3,
	0, 6, 1, 2, 3, 0, 10, 1, 2, 3,
	0, 5, 0, 4, 7, 1, 2, 3, 8, 1,
	2, 3, 8, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -4, -5, -6, -7, -8, -9, -10,
	9, 6, 15, 13, 16, 26, 29, -11, 23, -15,
	23, -16, -25, 23, -28, 23, -34, 23, -2, 30,
	33, -17, -18, -19, -20, -21, -22, -24, 17, 10,
	8, 4, 12, 11, -5, -8, -9, 23, 19, 20,
	21, 22, 30, -29, 30, -12, -13, -14, 28, 23,
	23, 24, -23, 38, 23, -24, -26, 14, -35, -36,
	27, 31, 32, 33, -14, 34, 36, -23, -27, -17,
	39, 31, 29, -37, 23, -12, -3, 23, 24, 33,
	-3, -3, 37, 31, 29, -18, -22, 33, -35, 30,
	-3, 35, 37, -27, 40, -23, -3, -38, -39, -22,
	30, 31, 29, -40, 23, -30, -31, 5, 7, 33,
	-38, 39, 31, 29, -3, -33, -3, -22, -30, -32,
	41, 29, 40, 41, -17, 33, -17, -3, 29,
}
var yyDef = [...]int{

	0, -2, 1, 0, 6, 7, 8, 9, 10, 11,
	0, 0, 23, 0, 0, 0, 4, 0, 18, 0,
	22, 0, 0, 52, 55, 57, 0, 65, 5, 0,
	0, 24, 25, 26, 27, 28, 0, 30, 0, 32,
	33, 34, 35, 36, 37, 38, 39, 40, 41, 42,
	43, 44, 50, 0, 0, 0, 13, 15, 0, 19,
	20, 21, 29, 0, 45, 31, 0, 0, 0, 0,
	0, 12, 0, 0, 0, 0, 0, 49, 0, 0,
	0, 0, 66, 0, 69, 14, 17, 2, 3, 0,
	0, 0, 48, 51, 53, 0, 0, 0, 67, 0,
	16, 46, 47, 54, 0, 29, 64, 0, 0, 0,
	0, 0, 70, 0, 73, 0, 0, 0, 62, 0,
	71, 0, 56, 58, 60, 0, 0, 0, 59, 0,
	0, 68, 0, 0, 63, 0, 61, 0, 72,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	39, 40, 38, 3, 32, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 41, 29,
	36, 33, 37, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 34, 3, 35, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 30, 3, 31,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 12:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sunrpc.y:87
		{
			StartEnum(yyDollar[2].val)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sunrpc.y:96
		{
			err := AddEnumAutoVal(yyDollar[1].val)
			if err != nil {
				yylex.Error(err.Error())
				return 1
			}
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sunrpc.y:103
		{
			err := AddEnumValMeta(yyDollar[2].val, yyDollar[4].val, yyDollar[1].val)
			if err != nil {
				yylex.Error(err.Error())
				return 1
			}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sunrpc.y:110
		{
			err := AddEnumVal(yyDollar[1].val, yyDollar[3].val)
			if err != nil {
				yylex.Error(err.Error())
				return 1
			}
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sunrpc.y:132
		{
			err := AddConst(yyDollar[2].val, yyDollar[4].val)
			if err != nil {
				yylex.Error(err.Error())
				return 1
			}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sunrpc.y:146
		{
			StartTypedef()
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sunrpc.y:157
		{
			AddDeclaration(yyDollar[2].val, yyDollar[1].val)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sunrpc.y:162
		{
			yyVAL.val = "u" + yyDollar[2].val
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sunrpc.y:163
		{
			yyVAL.val = "float32"
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sunrpc.y:164
		{
			yyVAL.val = "float64"
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sunrpc.y:165
		{
			yyVAL.val = "bool"
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sunrpc.y:166
		{
			yyVAL.val = "string"
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sunrpc.y:167
		{
			yyVAL.val = "byte"
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sunrpc.y:175
		{
			yyVAL.val = "int64"
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sunrpc.y:176
		{
			yyVAL.val = "int32"
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sunrpc.y:177
		{
			yyVAL.val = "int16"
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sunrpc.y:178
		{
			yyVAL.val = "int8"
		}
	case 46:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sunrpc.y:186
		{
			AddFixedArray(yyDollar[2].val, yyDollar[1].val, yyDollar[4].val)
		}
	case 47:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sunrpc.y:190
		{
			AddVariableArray(yyDollar[2].val, yyDollar[1].val, yyDollar[4].val)
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sunrpc.y:191
		{
			AddVariableArray(yyDollar[2].val, yyDollar[1].val, "")
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sunrpc.y:199
		{
			AddOptValue(yyDollar[3].val, yyDollar[1].val)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sunrpc.y:203
		{
			StartStruct(yyDollar[2].val)
		}
	case 51:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sunrpc.y:203
		{
			AddStruct()
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sunrpc.y:216
		{
			StartUnion(yyDollar[2].val)
		}
	case 56:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line sunrpc.y:216
		{
			AddUnion()
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sunrpc.y:229
		{
			StartCase(yyDollar[2].val)
		}
	case 61:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sunrpc.y:229
		{
			AddCase()
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sunrpc.y:230
		{
			StartCase("default")
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sunrpc.y:230
		{
			AddCase()
		}
	}
	goto yystack /* stack new state and value */
}
