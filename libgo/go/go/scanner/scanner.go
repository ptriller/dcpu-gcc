// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// A scanner for Go source text. Takes a []byte as source which can
// then be tokenized through repeated calls to the Scan function.
// For a sample use of a scanner, see the implementation of Tokenize.
//
package scanner

import (
	"bytes"
	"go/token"
	"strconv"
	"unicode"
	"utf8"
)


// A Scanner holds the scanner's internal state while processing
// a given text.  It can be allocated as part of another data
// structure but must be initialized via Init before use. For
// a sample use, see the implementation of Tokenize.
//
type Scanner struct {
	// immutable state
	src  []byte       // source
	err  ErrorHandler // error reporting; or nil
	mode uint         // scanning mode

	// scanning state
	filename string // current filename; may change via //line filename:line comment
	line     int    // current line
	column   int    // current column

	ch         int  // current character
	offset     int  // character offset
	rdOffset   int  // reading offset (position after current character)
	insertSemi bool // insert a semicolon before next newline

	// public state - ok to modify
	ErrorCount int // number of errors encountered
}


// Read the next Unicode char into S.ch.
// S.ch < 0 means end-of-file.
//
func (S *Scanner) next() {
	S.column++
	if S.rdOffset < len(S.src) {
		S.offset = S.rdOffset
		if S.ch == '\n' {
			S.line++
			S.column = 1
		}
		r, w := int(S.src[S.rdOffset]), 1
		switch {
		case r == 0:
			S.error("illegal character NUL")
		case r >= 0x80:
			// not ASCII
			r, w = utf8.DecodeRune(S.src[S.rdOffset:])
			if r == utf8.RuneError && w == 1 {
				S.error("illegal UTF-8 encoding")
			}
		}
		S.rdOffset += w
		S.ch = r
	} else {
		S.offset = len(S.src)
		if S.ch == '\n' {
			S.column = 1
		}
		S.ch = -1 // eof
	}
}


// The mode parameter to the Init function is a set of flags (or 0).
// They control scanner behavior.
//
const (
	ScanComments      = 1 << iota // return comments as COMMENT tokens
	AllowIllegalChars             // do not report an error for illegal chars
	InsertSemis                   // automatically insert semicolons
)


// Init prepares the scanner S to tokenize the text src. Calls to Scan
// will use the error handler err if they encounter a syntax error and
// err is not nil. Also, for each error encountered, the Scanner field
// ErrorCount is incremented by one. The filename parameter is used as
// filename in the token.Position returned by Scan for each token. The
// mode parameter determines how comments and illegal characters are
// handled.
//
func (S *Scanner) Init(filename string, src []byte, err ErrorHandler, mode uint) {
	// Explicitly initialize all fields since a scanner may be reused.
	S.src = src
	S.err = err
	S.mode = mode

	S.filename = filename
	S.line = 1
	S.column = 0

	S.ch = ' '
	S.offset = 0
	S.rdOffset = 0
	S.insertSemi = false
	S.ErrorCount = 0

	S.next()
}


func charString(ch int) string {
	var s string
	switch ch {
	case -1:
		return `EOF`
	case '\a':
		s = `\a`
	case '\b':
		s = `\b`
	case '\f':
		s = `\f`
	case '\n':
		s = `\n`
	case '\r':
		s = `\r`
	case '\t':
		s = `\t`
	case '\v':
		s = `\v`
	case '\\':
		s = `\\`
	case '\'':
		s = `\'`
	default:
		s = string(ch)
	}
	return "'" + s + "' (U+" + strconv.Itob(ch, 16) + ")"
}


func (S *Scanner) error(msg string) {
	S.errorAt(token.Position{S.filename, S.offset, S.line, S.column}, msg)
}


func (S *Scanner) errorAt(pos token.Position, msg string) {
	if S.err != nil {
		S.err.Error(pos, msg)
	}
	S.ErrorCount++
}


var prefix = []byte("//line ")

func (S *Scanner) interpretLineComment(text []byte) {
	if bytes.HasPrefix(text, prefix) {
		// get filename and line number, if any
		if i := bytes.Index(text, []byte{':'}); i > 0 {
			if line, err := strconv.Atoi(string(text[i+1:])); err == nil && line > 0 {
				// valid //line filename:line comment;
				// update scanner position
				S.filename = string(text[len(prefix):i])
				S.line = line - 1 // -1 since the '\n' has not been consumed yet
			}
		}
	}
}


func (S *Scanner) scanComment() {
	// initial '/' already consumed; S.ch == '/' || S.ch == '*'
	offs := S.offset - 1 // position of initial '/'
	col := S.column - 1
	pos := token.Position{S.filename, S.offset - 1, S.line, S.column - 1}

	if S.ch == '/' {
		//-style comment
		S.next()
		for S.ch != '\n' && S.ch >= 0 {
			S.next()
		}
		if col == 1 {
			// comment starts at the beginning of the current line
			S.interpretLineComment(S.src[offs:S.offset])
		}
		return
	}

	/*-style comment */
	S.next()
	for S.ch >= 0 {
		ch := S.ch
		S.next()
		if ch == '*' && S.ch == '/' {
			S.next()
			return
		}
	}

	S.errorAt(pos, "comment not terminated")
}


func (S *Scanner) findLineEnd() bool {
	// initial '/' already consumed

	defer func(line, col, offs int) {
		// reset scanner state to where it was upon calling findLineEnd
		// (we don't scan //line comments and ignore errors thus
		// S.filename and S.ErrorCount don't change)
		S.line = line
		S.column = col
		S.ch = '/'
		S.offset = offs
		S.rdOffset = offs + 1
		S.next() // consume initial '/' again
	}(S.line, S.column-1, S.offset-1)

	// read ahead until a newline, EOF, or non-comment token is found
	for S.ch == '/' || S.ch == '*' {
		if S.ch == '/' {
			//-style comment always contains a newline
			return true
		}
		/*-style comment: look for newline */
		S.next()
		for S.ch >= 0 {
			ch := S.ch
			if ch == '\n' {
				return true
			}
			S.next()
			if ch == '*' && S.ch == '/' {
				S.next()
				break
			}
		}
		S.skipWhitespace() // S.insertSemi is set
		if S.ch < 0 || S.ch == '\n' {
			return true
		}
		if S.ch != '/' {
			// non-comment token
			return false
		}
		S.next() // consume '/'
	}

	return false
}


func isLetter(ch int) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= 0x80 && unicode.IsLetter(ch)
}


func isDigit(ch int) bool {
	return '0' <= ch && ch <= '9' || ch >= 0x80 && unicode.IsDigit(ch)
}


func (S *Scanner) scanIdentifier() token.Token {
	offs := S.offset
	for isLetter(S.ch) || isDigit(S.ch) {
		S.next()
	}
	return token.Lookup(S.src[offs:S.offset])
}


func digitVal(ch int) int {
	switch {
	case '0' <= ch && ch <= '9':
		return ch - '0'
	case 'a' <= ch && ch <= 'f':
		return ch - 'a' + 10
	case 'A' <= ch && ch <= 'F':
		return ch - 'A' + 10
	}
	return 16 // larger than any legal digit val
}


func (S *Scanner) scanMantissa(base int) {
	for digitVal(S.ch) < base {
		S.next()
	}
}


func (S *Scanner) scanNumber(seenDecimalPoint bool) token.Token {
	// digitVal(S.ch) < 10
	tok := token.INT

	if seenDecimalPoint {
		tok = token.FLOAT
		S.scanMantissa(10)
		goto exponent
	}

	if S.ch == '0' {
		// int or float
		pos := token.Position{S.filename, S.offset, S.line, S.column}
		S.next()
		if S.ch == 'x' || S.ch == 'X' {
			// hexadecimal int
			S.next()
			S.scanMantissa(16)
		} else {
			// octal int or float
			seenDecimalDigit := false
			S.scanMantissa(8)
			if S.ch == '8' || S.ch == '9' {
				// illegal octal int or float
				seenDecimalDigit = true
				S.scanMantissa(10)
			}
			if S.ch == '.' || S.ch == 'e' || S.ch == 'E' || S.ch == 'i' {
				goto fraction
			}
			// octal int
			if seenDecimalDigit {
				S.errorAt(pos, "illegal octal number")
			}
		}
		goto exit
	}

	// decimal int or float
	S.scanMantissa(10)

fraction:
	if S.ch == '.' {
		tok = token.FLOAT
		S.next()
		S.scanMantissa(10)
	}

exponent:
	if S.ch == 'e' || S.ch == 'E' {
		tok = token.FLOAT
		S.next()
		if S.ch == '-' || S.ch == '+' {
			S.next()
		}
		S.scanMantissa(10)
	}

	if S.ch == 'i' {
		tok = token.IMAG
		S.next()
	}

exit:
	return tok
}


func (S *Scanner) scanEscape(quote int) {
	pos := token.Position{S.filename, S.offset, S.line, S.column}

	var i, base, max uint32
	switch S.ch {
	case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', quote:
		S.next()
		return
	case '0', '1', '2', '3', '4', '5', '6', '7':
		i, base, max = 3, 8, 255
	case 'x':
		S.next()
		i, base, max = 2, 16, 255
	case 'u':
		S.next()
		i, base, max = 4, 16, unicode.MaxRune
	case 'U':
		S.next()
		i, base, max = 8, 16, unicode.MaxRune
	default:
		S.next() // always make progress
		S.errorAt(pos, "unknown escape sequence")
		return
	}

	var x uint32
	for ; i > 0 && S.ch != quote && S.ch >= 0; i-- {
		d := uint32(digitVal(S.ch))
		if d >= base {
			S.error("illegal character in escape sequence")
			break
		}
		x = x*base + d
		S.next()
	}
	// in case of an error, consume remaining chars
	for ; i > 0 && S.ch != quote && S.ch >= 0; i-- {
		S.next()
	}
	if x > max || 0xd800 <= x && x < 0xe000 {
		S.errorAt(pos, "escape sequence is invalid Unicode code point")
	}
}


func (S *Scanner) scanChar() {
	// '\'' opening already consumed
	pos := token.Position{S.filename, S.offset - 1, S.line, S.column - 1}

	n := 0
	for S.ch != '\'' {
		ch := S.ch
		n++
		S.next()
		if ch == '\n' || ch < 0 {
			S.errorAt(pos, "character literal not terminated")
			n = 1
			break
		}
		if ch == '\\' {
			S.scanEscape('\'')
		}
	}

	S.next()

	if n != 1 {
		S.errorAt(pos, "illegal character literal")
	}
}


func (S *Scanner) scanString() {
	// '"' opening already consumed
	pos := token.Position{S.filename, S.offset - 1, S.line, S.column - 1}

	for S.ch != '"' {
		ch := S.ch
		S.next()
		if ch == '\n' || ch < 0 {
			S.errorAt(pos, "string not terminated")
			break
		}
		if ch == '\\' {
			S.scanEscape('"')
		}
	}

	S.next()
}


func (S *Scanner) scanRawString() {
	// '`' opening already consumed
	pos := token.Position{S.filename, S.offset - 1, S.line, S.column - 1}

	for S.ch != '`' {
		ch := S.ch
		S.next()
		if ch < 0 {
			S.errorAt(pos, "string not terminated")
			break
		}
	}

	S.next()
}


func (S *Scanner) skipWhitespace() {
	for S.ch == ' ' || S.ch == '\t' || S.ch == '\n' && !S.insertSemi || S.ch == '\r' {
		S.next()
	}
}


// Helper functions for scanning multi-byte tokens such as >> += >>= .
// Different routines recognize different length tok_i based on matches
// of ch_i. If a token ends in '=', the result is tok1 or tok3
// respectively. Otherwise, the result is tok0 if there was no other
// matching character, or tok2 if the matching character was ch2.

func (S *Scanner) switch2(tok0, tok1 token.Token) token.Token {
	if S.ch == '=' {
		S.next()
		return tok1
	}
	return tok0
}


func (S *Scanner) switch3(tok0, tok1 token.Token, ch2 int, tok2 token.Token) token.Token {
	if S.ch == '=' {
		S.next()
		return tok1
	}
	if S.ch == ch2 {
		S.next()
		return tok2
	}
	return tok0
}


func (S *Scanner) switch4(tok0, tok1 token.Token, ch2 int, tok2, tok3 token.Token) token.Token {
	if S.ch == '=' {
		S.next()
		return tok1
	}
	if S.ch == ch2 {
		S.next()
		if S.ch == '=' {
			S.next()
			return tok3
		}
		return tok2
	}
	return tok0
}


var newline = []byte{'\n'}

// Scan scans the next token and returns the token position pos,
// the token tok, and the literal text lit corresponding to the
// token. The source end is indicated by token.EOF.
//
// If the returned token is token.SEMICOLON, the corresponding
// literal value is ";" if the semicolon was present in the source,
// and "\n" if the semicolon was inserted because of a newline or
// at EOF.
//
// For more tolerant parsing, Scan will return a valid token if
// possible even if a syntax error was encountered. Thus, even
// if the resulting token sequence contains no illegal tokens,
// a client may not assume that no error occurred. Instead it
// must check the scanner's ErrorCount or the number of calls
// of the error handler, if there was one installed.
//
func (S *Scanner) Scan() (pos token.Position, tok token.Token, lit []byte) {
scanAgain:
	S.skipWhitespace()

	// current token start
	insertSemi := false
	pos, tok = token.Position{S.filename, S.offset, S.line, S.column}, token.ILLEGAL
	offs := S.offset

	// determine token value
	switch ch := S.ch; {
	case isLetter(ch):
		tok = S.scanIdentifier()
		switch tok {
		case token.IDENT, token.BREAK, token.CONTINUE, token.FALLTHROUGH, token.RETURN:
			insertSemi = true
		}
	case digitVal(ch) < 10:
		insertSemi = true
		tok = S.scanNumber(false)
	default:
		S.next() // always make progress
		switch ch {
		case -1:
			if S.insertSemi {
				S.insertSemi = false // EOF consumed
				return pos, token.SEMICOLON, newline
			}
			tok = token.EOF
		case '\n':
			// we only reach here if S.insertSemi was
			// set in the first place and exited early
			// from S.skipWhitespace()
			S.insertSemi = false // newline consumed
			return pos, token.SEMICOLON, newline
		case '"':
			insertSemi = true
			tok = token.STRING
			S.scanString()
		case '\'':
			insertSemi = true
			tok = token.CHAR
			S.scanChar()
		case '`':
			insertSemi = true
			tok = token.STRING
			S.scanRawString()
		case ':':
			tok = S.switch2(token.COLON, token.DEFINE)
		case '.':
			if digitVal(S.ch) < 10 {
				insertSemi = true
				tok = S.scanNumber(true)
			} else if S.ch == '.' {
				S.next()
				if S.ch == '.' {
					S.next()
					tok = token.ELLIPSIS
				}
			} else {
				tok = token.PERIOD
			}
		case ',':
			tok = token.COMMA
		case ';':
			tok = token.SEMICOLON
		case '(':
			tok = token.LPAREN
		case ')':
			insertSemi = true
			tok = token.RPAREN
		case '[':
			tok = token.LBRACK
		case ']':
			insertSemi = true
			tok = token.RBRACK
		case '{':
			tok = token.LBRACE
		case '}':
			insertSemi = true
			tok = token.RBRACE
		case '+':
			tok = S.switch3(token.ADD, token.ADD_ASSIGN, '+', token.INC)
			if tok == token.INC {
				insertSemi = true
			}
		case '-':
			tok = S.switch3(token.SUB, token.SUB_ASSIGN, '-', token.DEC)
			if tok == token.DEC {
				insertSemi = true
			}
		case '*':
			tok = S.switch2(token.MUL, token.MUL_ASSIGN)
		case '/':
			if S.ch == '/' || S.ch == '*' {
				// comment
				line := S.line
				col := S.column - 1 // beginning of comment
				if S.insertSemi && S.findLineEnd() {
					// reset position to the beginning of the comment
					S.line = line
					S.column = col
					S.ch = '/'
					S.offset = offs
					S.rdOffset = offs + 1
					S.insertSemi = false // newline consumed
					return pos, token.SEMICOLON, newline
				}
				S.scanComment()
				if S.mode&ScanComments == 0 {
					// skip comment
					S.insertSemi = false // newline consumed
					goto scanAgain
				}
				tok = token.COMMENT
			} else {
				tok = S.switch2(token.QUO, token.QUO_ASSIGN)
			}
		case '%':
			tok = S.switch2(token.REM, token.REM_ASSIGN)
		case '^':
			tok = S.switch2(token.XOR, token.XOR_ASSIGN)
		case '<':
			if S.ch == '-' {
				S.next()
				tok = token.ARROW
			} else {
				tok = S.switch4(token.LSS, token.LEQ, '<', token.SHL, token.SHL_ASSIGN)
			}
		case '>':
			tok = S.switch4(token.GTR, token.GEQ, '>', token.SHR, token.SHR_ASSIGN)
		case '=':
			tok = S.switch2(token.ASSIGN, token.EQL)
		case '!':
			tok = S.switch2(token.NOT, token.NEQ)
		case '&':
			if S.ch == '^' {
				S.next()
				tok = S.switch2(token.AND_NOT, token.AND_NOT_ASSIGN)
			} else {
				tok = S.switch3(token.AND, token.AND_ASSIGN, '&', token.LAND)
			}
		case '|':
			tok = S.switch3(token.OR, token.OR_ASSIGN, '|', token.LOR)
		default:
			if S.mode&AllowIllegalChars == 0 {
				S.errorAt(pos, "illegal character "+charString(ch))
			}
			insertSemi = S.insertSemi // preserve insertSemi info
		}
	}

	if S.mode&InsertSemis != 0 {
		S.insertSemi = insertSemi
	}
	return pos, tok, S.src[offs:S.offset]
}


// Tokenize calls a function f with the token position, token value, and token
// text for each token in the source src. The other parameters have the same
// meaning as for the Init function. Tokenize keeps scanning until f returns
// false (usually when the token value is token.EOF). The result is the number
// of errors encountered.
//
func Tokenize(filename string, src []byte, err ErrorHandler, mode uint, f func(pos token.Position, tok token.Token, lit []byte) bool) int {
	var s Scanner
	s.Init(filename, src, err, mode)
	for f(s.Scan()) {
		// action happens in f
	}
	return s.ErrorCount
}