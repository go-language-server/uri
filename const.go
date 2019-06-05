// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uri

const (
	// Null is the null.
	Null = 0

	// Tab is the `\t` character.
	Tab = 9

	// LineFeed is the `\n` character.
	LineFeed = 10

	// CarriageReturn is the `\r` character.
	CarriageReturn = 13

	// Space is the whitespace.
	Space = 32

	// ExclamationMark is the `!` character.
	ExclamationMark = 33

	// DoubleQuote is the `"` character.
	DoubleQuote = 34

	// Hash is the `#` character.
	Hash = 35

	// DollarSign is the `$` character.
	DollarSign = 36

	// PercentSign is the `%` character.
	PercentSign = 37

	// Ampersand is the `&` character.
	Ampersand = 38

	// SingleQuote is the `'` character.
	SingleQuote = 39

	// OpenParen is the `(` character.
	OpenParen = 40

	// CloseParen is the `)` character.
	CloseParen = 41

	// Asterisk is the `*` character.
	Asterisk = 42

	// Plus is the `+` character.
	Plus = 43

	// Comma is the `,` character.
	Comma = 44

	// Dash is the `-` character.
	Dash = 45

	// Period is the `.` character.
	Period = 46

	// Slash is the `/` character.
	Slash = 47

	// Digit0 is the 0.
	Digit0 = 48
	// Digit1 is the 1.
	Digit1 = 49
	// Digit2 is the 2.
	Digit2 = 50
	// Digit3 is the 3.
	Digit3 = 51
	// Digit4 is the 4.
	Digit4 = 52
	// Digit5 is the 5.
	Digit5 = 53
	// Digit6 is the 6.
	Digit6 = 54
	// Digit7 is the 7.
	Digit7 = 55
	// Digit8 is the 8.
	Digit8 = 56
	// Digit9 is the 9.
	Digit9 = 57

	// Colon is the `:` character.
	Colon = 58

	// Semicolon is the `;` character.
	Semicolon = 59

	// LessThan is the `<` character.
	LessThan = 60

	// Equals is the `=` character.
	Equals = 61

	// GreaterThan is the `>` character.
	GreaterThan = 62

	// QuestionMark is the `?` character.
	QuestionMark = 63

	// AtSign is the `@` character.
	AtSign = 64

	// UpperA is the A.
	UpperA = 65
	// UpperB is the B.
	UpperB = 66
	// UpperC is the C.
	UpperC = 67
	// UpperD is the D.
	UpperD = 68
	// UpperE is the E.
	UpperE = 69
	// UpperF is the F.
	UpperF = 70
	// UpperG is the G.
	UpperG = 71
	// UpperH is the H.
	UpperH = 72
	// UpperI is the I.
	UpperI = 73
	// UpperJ is the J.
	UpperJ = 74
	// UpperK is the K.
	UpperK = 75
	// UpperL is the L.
	UpperL = 76
	// UpperM is the M.
	UpperM = 77
	// UpperN is the N.
	UpperN = 78
	// UpperO is the O.
	UpperO = 79
	// UpperP is the P.
	UpperP = 80
	// UpperQ is the Q.
	UpperQ = 81
	// UpperR is the R.
	UpperR = 82
	// UpperS is the S.
	UpperS = 83
	// UpperT is the T.
	UpperT = 84
	// UpperU is the U.
	UpperU = 85
	// UpperV is the V.
	UpperV = 86
	// UpperW is the W.
	UpperW = 87
	// UpperX is the X.
	UpperX = 88
	// UpperY is the Y.
	UpperY = 89
	// UpperZ is the Z.
	UpperZ = 90

	// OpenSquareBracket is the `[` character.
	OpenSquareBracket = 91

	// Backslash is the `\` character.
	Backslash = 92

	// CloseSquareBracket is the `]` character.
	CloseSquareBracket = 93

	// Caret is the `^` character.
	Caret = 94

	// Underline is the `_` character.
	Underline = 95

	// BackTick is the ` character.
	BackTick = 96

	// LowerA is the a.
	LowerA = 97
	// LowerB is the b.
	LowerB = 98
	// LowerC is the c.
	LowerC = 99
	// LowerD is the d.
	LowerD = 100
	// LowerA is the e.
	LowerE = 101
	// LowerF is the f.
	LowerF = 102
	// LowerG is the g.
	LowerG = 103
	// LowerH is the h.
	LowerH = 104
	// LowerI is the i.
	LowerI = 105
	// LowerJ is the j.
	LowerJ = 106
	// LowerK is the k.
	LowerK = 107
	// LowerL is the l.
	LowerL = 108
	// LowerM is the m.
	LowerM = 109
	// LowerN is the n.
	LowerN = 110
	// LowerO is the o.
	LowerO = 111
	// LowerP is the p.
	LowerP = 112
	// LowerQ is the q.
	LowerQ = 113
	// LowerR is the r.
	LowerR = 114
	// LowerS is the s.
	LowerS = 115
	// LowerT is the t.
	LowerT = 116
	// LowerU is the u.
	LowerU = 117
	// LowerV is the v.
	LowerV = 118
	// LowerW is the w.
	LowerW = 119
	// LowerX is the x.
	LowerX = 120
	// LowerY is the y.
	LowerY = 121
	// LowerZ is the z.
	LowerZ = 122

	// OpenCurlyBrace is the `{` character.
	OpenCurlyBrace = 123

	// Pipe is the `|` character.
	Pipe = 124

	// CloseCurlyBrace is the `}` character.
	CloseCurlyBrace = 125

	// Tilde is the `~` character.
	Tilde = 126

	// U_Combining_Grave_Accent U+0300 Combining Grave Accent.
	U_Combining_Grave_Accent = 0x0300
	// U_Combining_Acute_Accent U+0301 Combining Acute Accent.
	U_Combining_Acute_Accent = 0x0301
	// U_Combining_Circumflex_Accent U+0302 Combining Circumflex Accent.
	U_Combining_Circumflex_Accent = 0x0302
	// U_Combining_Tilde U+0303 Combining Tilde.
	U_Combining_Tilde = 0x0303
	// U_Combining_Macron U+0304 Combining Macron.
	U_Combining_Macron = 0x0304
	// U_Combining_Overline U+0305 Combining Overline.
	U_Combining_Overline = 0x0305
	// U_Combining_Breve U+0306 Combining Breve.
	U_Combining_Breve = 0x0306
	// U_Combining_Dot_Above U+0307 Combining Dot Above.
	U_Combining_Dot_Above = 0x0307
	// U_Combining_Diaeresis U+0308 Combining Diaeresis@.
	U_Combining_Diaeresis = 0x0308
	// U_Combining_Hook_Above U+0309 Combining Hook Above.
	U_Combining_Hook_Above = 0x0309
	// U_Combining_Ring_Above U+030A Combining Ring Above.
	U_Combining_Ring_Above = 0x030A
	// U_Combining_Double_Acute_Accent U+030B Combining Double Acute Accent.
	U_Combining_Double_Acute_Accent = 0x030B
	// U_Combining_Caron U+030C Combining Caron.
	U_Combining_Caron = 0x030C
	// U_Combining_Vertical_Line_Above U+030D Combining Vertical Line Above.
	U_Combining_Vertical_Line_Above = 0x030D
	// U_Combining_Double_Vertical_Line_Above U+030E Combining Double Vertical Line Above.
	U_Combining_Double_Vertical_Line_Above = 0x030E
	// U_Combining_Double_Grave_Accent U+030F Combining Double Grave Accent.
	U_Combining_Double_Grave_Accent = 0x030F
	// U_Combining_Candrabindu U+0310 Combining Candrabindu.
	U_Combining_Candrabindu = 0x0310
	// U_Combining_Inverted_Breve U+0311 Combining Inverted Breve.
	U_Combining_Inverted_Breve = 0x0311
	// U_Combining_Turned_Comma_Above U+0312 Combining Turned Comma Above.
	U_Combining_Turned_Comma_Above = 0x0312
	// U_Combining_Comma_Above U+0313 Combining Comma Above.
	U_Combining_Comma_Above = 0x0313
	// U_Combining_Reversed_Comma_Above U+0314 Combining Reversed Comma Above.
	U_Combining_Reversed_Comma_Above = 0x0314
	// U_Combining_Comma_Above_Right U+0315 Combining Comma Above Right.
	U_Combining_Comma_Above_Right = 0x0315
	// U_Combining_Grave_Accent_Below U+0316 Combining Grave Accent Below.
	U_Combining_Grave_Accent_Below = 0x0316
	// U_Combining_Acute_Accent_Below U+0317 Combining Acute Accent Below.
	U_Combining_Acute_Accent_Below = 0x0317
	// U_Combining_Left_Tack_Below U+0318 Combining Left Tack Below.
	U_Combining_Left_Tack_Below = 0x0318
	// U_Combining_Right_Tack_Below U+0319 Combining Right Tack Below.
	U_Combining_Right_Tack_Below = 0x0319
	// U_Combining_Left_Angle_Above U+031A Combining Left Angle Above.
	U_Combining_Left_Angle_Above = 0x031A
	// U_Combining_Horn U+031B Combining Horn.
	U_Combining_Horn = 0x031B
	// U_Combining_Left_Half_Ring_Below U+031C Combining Left Half Ring Below.
	U_Combining_Left_Half_Ring_Below = 0x031C
	// U_Combining_Up_Tack_Below U+031D Combining Up Tack Below.
	U_Combining_Up_Tack_Below = 0x031D
	// U_Combining_Down_Tack_Below U+031E Combining Down Tack Below.
	U_Combining_Down_Tack_Below = 0x031E
	// U_Combining_Plus_Sign_Below U+031F Combining Plus Sign Below.
	U_Combining_Plus_Sign_Below = 0x031F
	// U_Combining_Minus_Sign_Below U+0320 Combining Minus Sign Below.
	U_Combining_Minus_Sign_Below = 0x0320
	// U_Combining_Palatalized_Hook_Below U+0321 Combining Palatalized Hook Below.
	U_Combining_Palatalized_Hook_Below = 0x0321
	// U_Combining_Retroflex_Hook_Below U+0322 Combining Retroflex Hook Below.
	U_Combining_Retroflex_Hook_Below = 0x0322
	// U_Combining_Dot_Below U+0323 Combining Dot Below.
	U_Combining_Dot_Below = 0x0323
	// U_Combining_Diaeresis_Below U+0324 Combining Diaeresis Below.
	U_Combining_Diaeresis_Below = 0x0324
	// U_Combining_Ring_Below U+0325 Combining Ring Below.
	U_Combining_Ring_Below = 0x0325
	// U_Combining_Comma_Below U+0326 Combining Comma Below.
	U_Combining_Comma_Below = 0x0326
	// U_Combining_Cedilla U+0327 Combining Cedilla.
	U_Combining_Cedilla = 0x0327
	// U_Combining_Ogonek U+0328 Combining Ogonek.
	U_Combining_Ogonek = 0x0328
	// U_Combining_Vertical_Line_Below U+0329 Combining Vertical Line Below.
	U_Combining_Vertical_Line_Below = 0x0329
	// U_Combining_Bridge_Below U+032A Combining Bridge Below.
	U_Combining_Bridge_Below = 0x032A
	// U_Combining_Inverted_Double_Arch_Below U+032B Combining Inverted Double Arch Below.
	U_Combining_Inverted_Double_Arch_Below = 0x032B
	// U_Combining_Caron_Below U+032C Combining Caron Below.
	U_Combining_Caron_Below = 0x032C
	// U_Combining_Circumflex_Accent_Below U+032D Combining Circumflex Accent Below.
	U_Combining_Circumflex_Accent_Below = 0x032D
	// U_Combining_Breve_Below U+032E Combining Breve Below.
	U_Combining_Breve_Below = 0x032E
	// U_Combining_Inverted_Breve_Below U+032F Combining Inverted Breve Below.
	U_Combining_Inverted_Breve_Below = 0x032F
	// U_Combining_Tilde_Below U+0330 Combining Tilde Below.
	U_Combining_Tilde_Below = 0x0330
	// U_Combining_Macron_Below U+0331 Combining Macron Below.
	U_Combining_Macron_Below = 0x0331
	// U_Combining_Low_Line U+0332 Combining Low Line.
	U_Combining_Low_Line = 0x0332
	// U_Combining_Double_Low_Line U+0333 Combining Double Low Line.
	U_Combining_Double_Low_Line = 0x0333
	// U_Combining_Tilde_Overlay U+0334 Combining Tilde Overlay.
	U_Combining_Tilde_Overlay = 0x0334
	// U_Combining_Short_Stroke_Overlay U+0335 Combining Short Stroke Overlay.
	U_Combining_Short_Stroke_Overlay = 0x0335
	// U_Combining_Long_Stroke_Overlay U+0336 Combining Long Stroke Overlay.
	U_Combining_Long_Stroke_Overlay = 0x0336
	// U_Combining_Short_Solidus_Overlay U+0337 Combining Short Solidus Overlay.
	U_Combining_Short_Solidus_Overlay = 0x0337
	// U_Combining_Long_Solidus_Overlay U+0338 Combining Long Solidus Overlay.
	U_Combining_Long_Solidus_Overlay = 0x0338
	// U_Combining_Right_Half_Ring_Below U+0339 Combining Right Half Ring Below.
	U_Combining_Right_Half_Ring_Below = 0x0339
	// U_Combining_Inverted_Bridge_Below U+033A Combining Inverted Bridge Below.
	U_Combining_Inverted_Bridge_Below = 0x033A
	// U_Combining_Square_Below U+033B Combining Square Below.
	U_Combining_Square_Below = 0x033B
	// U_Combining_Seagull_Below U+033C Combining Seagull Below.
	U_Combining_Seagull_Below = 0x033C
	// U_Combining_X_Above U+033D Combining X Above.
	U_Combining_X_Above = 0x033D
	// U_Combining_Vertical_Tilde U+033E Combining Vertical Tilde.
	U_Combining_Vertical_Tilde = 0x033E
	// U_Combining_Double_Overline U+033F Combining Double Overline.
	U_Combining_Double_Overline = 0x033F
	// U_Combining_Grave_Tone_Mark U+0340 Combining Grave Tone Mark.
	U_Combining_Grave_Tone_Mark = 0x0340
	// U_Combining_Acute_Tone_Mark U+0341 Combining Acute Tone Mark.
	U_Combining_Acute_Tone_Mark = 0x0341
	// U_Combining_Greek_Perispomeni U+0342 Combining Greek Perispomeni.
	U_Combining_Greek_Perispomeni = 0x0342
	// U_Combining_Greek_Koronis U+0343 Combining Greek Koronis.
	U_Combining_Greek_Koronis = 0x0343
	// U_Combining_Greek_Dialytika_Tonos U+0344 Combining Greek Dialytika Tonos.
	U_Combining_Greek_Dialytika_Tonos = 0x0344
	// U_Combining_Greek_Ypogegrammeni U+0345 Combining Greek Ypogegrammeni.
	U_Combining_Greek_Ypogegrammeni = 0x0345
	// U_Combining_Bridge_Above U+0346 Combining Bridge Above.
	U_Combining_Bridge_Above = 0x0346
	// U_Combining_Equals_Sign_Below U+0347 Combining Equals Sign Below.
	U_Combining_Equals_Sign_Below = 0x0347
	// U_Combining_Double_Vertical_Line_Below U+0348 Combining Double Vertical Line Below.
	U_Combining_Double_Vertical_Line_Below = 0x0348
	// U_Combining_Left_Angle_Below U+0349 Combining Left Angle Below.
	U_Combining_Left_Angle_Below = 0x0349
	// U_Combining_Not_Tilde_Above U+034A Combining Not Tilde Above.
	U_Combining_Not_Tilde_Above = 0x034A
	// U_Combining_Homothetic_Above U+034B Combining Homothetic Above.
	U_Combining_Homothetic_Above = 0x034B
	// U_Combining_Almost_Equal_To_Above U+034C Combining Almost Equal To Above.
	U_Combining_Almost_Equal_To_Above = 0x034C
	// U_Combining_Left_Right_Arrow_Below U+034D Combining Left Right Arrow Below.
	U_Combining_Left_Right_Arrow_Below = 0x034D
	// U_Combining_Upwards_Arrow_Below U+034E Combining Upwards Arrow Below.
	U_Combining_Upwards_Arrow_Below = 0x034E
	// U_Combining_Grapheme_Joiner U+034F Combining Grapheme Joiner.
	U_Combining_Grapheme_Joiner = 0x034F
	// U_Combining_Right_Arrowhead_Above U+0350 Combining Right Arrowhead Above.
	U_Combining_Right_Arrowhead_Above = 0x0350
	// U_Combining_Left_Half_Ring_Above U+0351 Combining Left Half Ring Above.
	U_Combining_Left_Half_Ring_Above = 0x0351
	// U_Combining_Fermata U+0352 Combining Fermata.
	U_Combining_Fermata = 0x0352
	// U_Combining_X_Below U+0353 Combining X Below.
	U_Combining_X_Below = 0x0353
	// U_Combining_Left_Arrowhead_Below U+0354 Combining Left Arrowhead Below.
	U_Combining_Left_Arrowhead_Below = 0x0354
	// U_Combining_Right_Arrowhead_Below U+0355 Combining Right Arrowhead Below.
	U_Combining_Right_Arrowhead_Below = 0x0355
	// U_Combining_Right_Arrowhead_And_Up_Arrowhead_Below U+0356 Combining Right Arrowhead And Up Arrowhead Below.
	U_Combining_Right_Arrowhead_And_Up_Arrowhead_Below = 0x0356
	// U_Combining_Right_Half_Ring_Above U+0357 Combining Right Half Ring Above.
	U_Combining_Right_Half_Ring_Above = 0x0357
	// U_Combining_Dot_Above_Right U+0358 Combining Dot Above Right.
	U_Combining_Dot_Above_Right = 0x0358
	// U_Combining_Asterisk_Below U+0359 Combining Asterisk Below.
	U_Combining_Asterisk_Below = 0x0359
	// U_Combining_Double_Ring_Below U+035A Combining Double Ring Below.
	U_Combining_Double_Ring_Below = 0x035A
	// U_Combining_Zigzag_Above U+035B Combining Zigzag Above.
	U_Combining_Zigzag_Above = 0x035B
	// U_Combining_Double_Breve_Below U+035C Combining Double Breve Below.
	U_Combining_Double_Breve_Below = 0x035C
	// U_Combining_Double_Breve U+035D Combining Double Breve.
	U_Combining_Double_Breve = 0x035D
	// U_Combining_Double_Macron U+035E Combining Double Macron.
	U_Combining_Double_Macron = 0x035E
	// U_Combining_Double_Macron_Below U+035F Combining Double Macron Below.
	U_Combining_Double_Macron_Below = 0x035F
	// U_Combining_Double_Tilde U+0360 Combining Double Tilde.
	U_Combining_Double_Tilde = 0x0360
	// U_Combining_Double_Inverted_Breve U+0361 Combining Double Inverted Breve.
	U_Combining_Double_Inverted_Breve = 0x0361
	// U_Combining_Double_Rightwards_Arrow_Below U+0362 Combining Double Rightwards Arrow Below.
	U_Combining_Double_Rightwards_Arrow_Below = 0x0362
	// U_Combining_Latin_Small_Letter_A U+0363 Combining Latin Small Letter A.
	U_Combining_Latin_Small_Letter_A = 0x0363
	// U_Combining_Latin_Small_Letter_E U+0364 Combining Latin Small Letter E.
	U_Combining_Latin_Small_Letter_E = 0x0364
	// U_Combining_Latin_Small_Letter_I U+0365 Combining Latin Small Letter I.
	U_Combining_Latin_Small_Letter_I = 0x0365
	// U_Combining_Latin_Small_Letter_O U+0366 Combining Latin Small Letter O.
	U_Combining_Latin_Small_Letter_O = 0x0366
	// U_Combining_Latin_Small_Letter_U U+0367 Combining Latin Small Letter U.
	U_Combining_Latin_Small_Letter_U = 0x0367
	// U_Combining_Latin_Small_Letter_C U+0368 Combining Latin Small Letter C.
	U_Combining_Latin_Small_Letter_C = 0x0368
	// U_Combining_Latin_Small_Letter_D U+0369 Combining Latin Small Letter D.
	U_Combining_Latin_Small_Letter_D = 0x0369
	// U_Combining_Latin_Small_Letter_H U+036A Combining Latin Small Letter H.
	U_Combining_Latin_Small_Letter_H = 0x036A
	// U_Combining_Latin_Small_Letter_M U+036B Combining Latin Small Letter M.
	U_Combining_Latin_Small_Letter_M = 0x036B
	// U_Combining_Latin_Small_Letter_R U+036C Combining Latin Small Letter R.
	U_Combining_Latin_Small_Letter_R = 0x036C
	// U_Combining_Latin_Small_Letter_T U+036D Combining Latin Small Letter T.
	U_Combining_Latin_Small_Letter_T = 0x036D
	// U_Combining_Latin_Small_Letter_V U+036E Combining Latin Small Letter V.
	U_Combining_Latin_Small_Letter_V = 0x036E
	// U_Combining_Latin_Small_Letter_X U+036F Combining Latin Small Letter X.
	U_Combining_Latin_Small_Letter_X = 0x036F

	// LINE_SEPARATOR_202 is the 8Unicode Character 'LINE SEPARATOR' (U+2028).
	//  http://www.fileformat.info/info/unicode/char/2028/index.htm
	LINE_SEPARATOR_2028 = 8232

	//  http://www.fileformat.info/info/unicode/category/Sk/list.htm

	// U_CIRCUMFLEX U+005E CIRCUMFLEX.
	U_CIRCUMFLEX = 0x005E
	// U_GRAVE_ACCENT U+0060 GRAVE ACCENT.
	U_GRAVE_ACCENT = 0x0060
	// U_DIAERESIS U+00A8 DIAERESIS.
	U_DIAERESIS = 0x00A8
	// U_MACRON U+00AF MACRON.
	U_MACRON = 0x00AF
	// U_ACUTE_ACCENT U+00B4 ACUTE ACCENT.
	U_ACUTE_ACCENT = 0x00B4
	// U_CEDILLA U_CEDILLAU+00B8 CEDILLA.
	U_CEDILLA = 0x00B8
	// U_MODIFIER_LETTER_LEFT_ARROWHEAD U+02C2 MODIFIER LETTER LEFT ARROWHEAD.
	U_MODIFIER_LETTER_LEFT_ARROWHEAD = 0x02C2
	// U_MODIFIER_LETTER_RIGHT_ARROWHEAD U+02C3 MODIFIER LETTER RIGHT ARROWHEAD.
	U_MODIFIER_LETTER_RIGHT_ARROWHEAD = 0x02C3
	// U_MODIFIER_LETTER_UP_ARROWHEAD U+02C4 MODIFIER LETTER UP ARROWHEAD.
	U_MODIFIER_LETTER_UP_ARROWHEAD = 0x02C4
	// U_MODIFIER_LETTER_DOWN_ARROWHEAD U+02C5 MODIFIER LETTER DOWN ARROWHEAD.
	U_MODIFIER_LETTER_DOWN_ARROWHEAD = 0x02C5
	// U_MODIFIER_LETTER_CENTRED_RIGHT_HALF_RING U+02D2 MODIFIER LETTER CENTRED RIGHT HALF RING.
	U_MODIFIER_LETTER_CENTRED_RIGHT_HALF_RING = 0x02D2
	// U_MODIFIER_LETTER_CENTRED_LEFT_HALF_RING U+02D3 MODIFIER LETTER CENTRED LEFT HALF RING.
	U_MODIFIER_LETTER_CENTRED_LEFT_HALF_RING = 0x02D3
	// U_MODIFIER_LETTER_UP_TACK U+02D4 MODIFIER LETTER UP TACK.
	U_MODIFIER_LETTER_UP_TACK = 0x02D4
	// U_MODIFIER_LETTER_DOWN_TACK U+02D5 MODIFIER LETTER DOWN TACK.
	U_MODIFIER_LETTER_DOWN_TACK = 0x02D5
	// U_MODIFIER_LETTER_PLUS_SIGN U+02D6 MODIFIER LETTER PLUS SIGN.
	U_MODIFIER_LETTER_PLUS_SIGN = 0x02D6
	// U_MODIFIER_LETTER_MINUS_SIGN U+02D7 MODIFIER LETTER MINUS SIGN.
	U_MODIFIER_LETTER_MINUS_SIGN = 0x02D7
	// U_BREVE U+02D8 BREVE.
	U_BREVE = 0x02D8
	// U_DOT_ABOVE U+02D9 DOT ABOVE.
	U_DOT_ABOVE = 0x02D9
	// U_RING_ABOVE U+02DA RING ABOVE.
	U_RING_ABOVE = 0x02DA
	// U_OGONEK U+02DB OGONEK.
	U_OGONEK = 0x02DB
	// U_SMALL_TILDE U+02DC SMALL TILDE.
	U_SMALL_TILDE = 0x02DC
	// U_DOUBLE_ACUTE_ACCENT U+02DD DOUBLE ACUTE ACCENT.
	U_DOUBLE_ACUTE_ACCENT = 0x02DD
	// U_MODIFIER_LETTER_RHOTIC_HOOK U+02DE MODIFIER LETTER RHOTIC HOOK.
	U_MODIFIER_LETTER_RHOTIC_HOOK = 0x02DE
	// U_MODIFIER_LETTER_CROSS_ACCENT U+02DF MODIFIER LETTER CROSS ACCENT.
	U_MODIFIER_LETTER_CROSS_ACCENT = 0x02DF
	// U_MODIFIER_LETTER_EXTRA_HIGH_TONE_BAR U+02E5 MODIFIER LETTER EXTRA-HIGH TONE BAR.
	U_MODIFIER_LETTER_EXTRA_HIGH_TONE_BAR = 0x02E5
	// U_MODIFIER_LETTER_HIGH_TONE_BAR U+02E6 MODIFIER LETTER HIGH TONE BAR.
	U_MODIFIER_LETTER_HIGH_TONE_BAR = 0x02E6
	// U_MODIFIER_LETTER_MID_TONE_BAR U+02E7 MODIFIER LETTER MID TONE BAR.
	U_MODIFIER_LETTER_MID_TONE_BAR = 0x02E7
	// U_MODIFIER_LETTER_LOW_TONE_BAR U+02E8 MODIFIER LETTER LOW TONE BAR.
	U_MODIFIER_LETTER_LOW_TONE_BAR = 0x02E8
	// U_MODIFIER_LETTER_EXTRA_LOW_TONE_BAR U+02E9 MODIFIER LETTER EXTRA-LOW TONE BAR.
	U_MODIFIER_LETTER_EXTRA_LOW_TONE_BAR = 0x02E9
	// U_MODIFIER_LETTER_YIN_DEPARTING_TONE_MARK U+02EA MODIFIER LETTER YIN DEPARTING TONE MARK.
	U_MODIFIER_LETTER_YIN_DEPARTING_TONE_MARK = 0x02EA
	// U_MODIFIER_LETTER_YANG_DEPARTING_TONE_MARK U+02EB MODIFIER LETTER YANG DEPARTING TONE MARK.
	U_MODIFIER_LETTER_YANG_DEPARTING_TONE_MARK = 0x02EB
	// U_MODIFIER_LETTER_UNASPIRATED U+02ED MODIFIER LETTER UNASPIRATED.
	U_MODIFIER_LETTER_UNASPIRATED = 0x02ED
	// U_MODIFIER_LETTER_LOW_DOWN_ARROWHEAD U+02EF MODIFIER LETTER LOW DOWN ARROWHEAD.
	U_MODIFIER_LETTER_LOW_DOWN_ARROWHEAD = 0x02EF
	// U_MODIFIER_LETTER_LOW_UP_ARROWHEAD U+02F0 MODIFIER LETTER LOW UP ARROWHEAD.
	U_MODIFIER_LETTER_LOW_UP_ARROWHEAD = 0x02F0
	// U_MODIFIER_LETTER_LOW_LEFT_ARROWHEAD U+02F1 MODIFIER LETTER LOW LEFT ARROWHEAD.
	U_MODIFIER_LETTER_LOW_LEFT_ARROWHEAD = 0x02F1
	// U_MODIFIER_LETTER_LOW_RIGHT_ARROWHEAD U+02F2 MODIFIER LETTER LOW RIGHT ARROWHEAD.
	U_MODIFIER_LETTER_LOW_RIGHT_ARROWHEAD = 0x02F2
	// U_MODIFIER_LETTER_LOW_RING U+02F3 MODIFIER LETTER LOW RING.
	U_MODIFIER_LETTER_LOW_RING = 0x02F3
	// U_MODIFIER_LETTER_MIDDLE_GRAVE_ACCENT U+02F4 MODIFIER LETTER MIDDLE GRAVE ACCENT.
	U_MODIFIER_LETTER_MIDDLE_GRAVE_ACCENT = 0x02F4
	// U_MODIFIER_LETTER_MIDDLE_DOUBLE_GRAVE_ACCENT U+02F5 MODIFIER LETTER MIDDLE DOUBLE GRAVE ACCENT.
	U_MODIFIER_LETTER_MIDDLE_DOUBLE_GRAVE_ACCENT = 0x02F5
	// U_MODIFIER_LETTER_MIDDLE_DOUBLE_ACUTE_ACCENT U+02F6 MODIFIER LETTER MIDDLE DOUBLE ACUTE ACCENT.
	U_MODIFIER_LETTER_MIDDLE_DOUBLE_ACUTE_ACCENT = 0x02F6
	// U_MODIFIER_LETTER_LOW_TILDE U+02F7 MODIFIER LETTER LOW TILDE.
	U_MODIFIER_LETTER_LOW_TILDE = 0x02F7
	// U_MODIFIER_LETTER_RAISED_COLON U+02F8 MODIFIER LETTER RAISED COLON.
	U_MODIFIER_LETTER_RAISED_COLON = 0x02F8
	// U_MODIFIER_LETTER_BEGIN_HIGH_TONE U+02F9 MODIFIER LETTER BEGIN HIGH TONE.
	U_MODIFIER_LETTER_BEGIN_HIGH_TONE = 0x02F9
	// U_MODIFIER_LETTER_END_HIGH_TONE U+02FA MODIFIER LETTER END HIGH TONE.
	U_MODIFIER_LETTER_END_HIGH_TONE = 0x02FA
	// U_MODIFIER_LETTER_BEGIN_LOW_TONE U+02FB MODIFIER LETTER BEGIN LOW TONE.
	U_MODIFIER_LETTER_BEGIN_LOW_TONE = 0x02FB
	// U_MODIFIER_LETTER_END_LOW_TONE U+02FC MODIFIER LETTER END LOW TONE.
	U_MODIFIER_LETTER_END_LOW_TONE = 0x02FC
	// U_MODIFIER_LETTER_SHELF U+02FD MODIFIER LETTER SHELF.
	U_MODIFIER_LETTER_SHELF = 0x02FD
	// U_MODIFIER_LETTER_OPEN_SHELF U+02FE MODIFIER LETTER OPEN SHELF.
	U_MODIFIER_LETTER_OPEN_SHELF = 0x02FE
	// U_MODIFIER_LETTER_LOW_LEFT_ARROW U+02FF MODIFIER LETTER LOW LEFT ARROW.
	U_MODIFIER_LETTER_LOW_LEFT_ARROW = 0x02FF
	// U_GREEK_LOWER_NUMERAL_SIGN U+0375 GREEK LOWER NUMERAL SIGN.
	U_GREEK_LOWER_NUMERAL_SIGN = 0x0375
	// U_GREEK_TONOS U+0384 GREEK TONOS.
	U_GREEK_TONOS = 0x0384
	// U_GREEK_DIALYTIKA_TONOS U+0385 GREEK DIALYTIKA TONOS.
	U_GREEK_DIALYTIKA_TONOS = 0x0385
	// U_GREEK_KORONIS U+1FBD GREEK KORONIS.
	U_GREEK_KORONIS = 0x1FBD
	// U_GREEK_PSILI U+1FBF GREEK PSILI.
	U_GREEK_PSILI = 0x1FBF
	// U_GREEK_PERISPOMENI U+1FC0 GREEK PERISPOMENI.
	U_GREEK_PERISPOMENI = 0x1FC0
	// U_GREEK_DIALYTIKA_AND_PERISPOMENI U+1FC1 GREEK DIALYTIKA AND PERISPOMENI.
	U_GREEK_DIALYTIKA_AND_PERISPOMENI = 0x1FC1
	// U_GREEK_PSILI_AND_VARIA U+1FCD GREEK PSILI AND VARIA.
	U_GREEK_PSILI_AND_VARIA = 0x1FCD
	// U_GREEK_PSILI_AND_OXIA U+1FCE GREEK PSILI AND OXIA.
	U_GREEK_PSILI_AND_OXIA = 0x1FCE
	// U_GREEK_PSILI_AND_PERISPOMENI U+1FCF GREEK PSILI AND PERISPOMENI.
	U_GREEK_PSILI_AND_PERISPOMENI = 0x1FCF
	// U_GREEK_DASIA_AND_VARIA U+1FDD GREEK DASIA AND VARIA.
	U_GREEK_DASIA_AND_VARIA = 0x1FDD
	// U_GREEK_DASIA_AND_OXIA U+1FDE GREEK DASIA AND OXIA.
	U_GREEK_DASIA_AND_OXIA = 0x1FDE
	// U_GREEK_DASIA_AND_PERISPOMENI U+1FDF GREEK DASIA AND PERISPOMENI.
	U_GREEK_DASIA_AND_PERISPOMENI = 0x1FDF
	// U_GREEK_DIALYTIKA_AND_VARIA U+1FED GREEK DIALYTIKA AND VARIA.
	U_GREEK_DIALYTIKA_AND_VARIA = 0x1FED
	// U_GREEK_DIALYTIKA_AND_OXIA U+1FEE GREEK DIALYTIKA AND OXIA.
	U_GREEK_DIALYTIKA_AND_OXIA = 0x1FEE
	// U_GREEK_VARIA U+1FEF GREEK VARIA.@
	U_GREEK_VARIA = 0x1FEF
	// U_GREEK_OXIA U+1FFD GREEK OXIA.
	U_GREEK_OXIA = 0x1FFD
	// U_GREEK_DASIA U+1FFE GREEK DASIA.
	U_GREEK_DASIA = 0x1FFE

	// U_OVERLINE Unicode Character 'OVERLINE'.
	U_OVERLINE = 0x203E

	// UTF8_BOM Unicode Character 'ZERO WIDTH NO-BREAK SPACE' (U+FEFF).
	//  http://www.fileformat.info/info/unicode/char/feff/index.htm
	UTF8_BOM = 65279
)
