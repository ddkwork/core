// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shell

import (
	"fmt"
	"go/token"

	"cogentcore.org/core/base/logx"
)

// TranspileLine is the main function for parsing a single line of shell input,
// returning a new transpiled line of code that converts Exec code into corresponding
// Go function calls.
func (sh *Shell) TranspileLine(ln string) string {
	if len(ln) == 0 {
		return ln
	}
	toks := sh.TranspileLineTokens(ln)
	paren, brace, brack := toks.BracketDepths()
	sh.ParenDepth += paren
	sh.BraceDepth += brace
	sh.BrackDepth += brack
	// logx.PrintlnDebug("depths: ", sh.ParenDepth, sh.BraceDepth, sh.BrackDepth)
	return toks.Code()
}

// TranspileLineTokens returns the tokens for the full line
func (sh *Shell) TranspileLineTokens(ln string) Tokens {
	if ln == "" {
		return nil
	}
	toks := sh.Tokens(ln)
	n := len(toks)
	ewords, err := ExecWords(ln)
	if err != nil {
		sh.AddError(err)
	}
	logx.PrintlnDebug("\n########## line:\n", ln, "\nTokens:\n", toks.String(), "\nWords:\n", ewords)

	t0 := toks[0]
	_, t0pn := toks.Path(true) // true = first position
	en := len(ewords)
	// f0 := ewords[0]
	// f1 := ""
	// if n > 0 {
	// 	f1 = ewords[1]
	// }

	f0exec := (t0.Tok == token.IDENT && ExecWordIsCommand(ewords[0]))

	// t1idx := max(t0pn, )
	// t1pn := 0
	// t1in := 0
	// if n > 1 {
	// 	if t1idx > 0 {
	// 		end0 := int(toks[t1idx-1].Pos) + len(toks[t1idx-1].String())
	// 		if t1idx < n && int(toks[t1idx].Pos) > end0 { // only if spaced
	// 			_, t1pn = toks[t1idx:].Path(false)
	// 		}
	// 	}
	// }

	switch {
	case t0.Tok == token.LBRACE:
		logx.PrintlnDebug("go:   { } line")
		return sh.TranspileGo(toks[1 : n-1])
	case t0.Tok == token.LBRACK:
		logx.PrintlnDebug("exec: [ ] line")
		return sh.TranspileExec(ewords, false) // it processes the [ ]
	case t0.Tok == token.ILLEGAL:
		logx.PrintlnDebug("exec: illegal")
		return sh.TranspileExec(ewords, false)
	case t0.IsBacktickString():
		logx.PrintlnDebug("exec: backquoted string")
		exe := sh.TranspileExecString(t0.Str, false)
		if n > 1 { // todo: is this an error?
			exe.AddTokens(sh.TranspileGo(toks[1:]))
		}
		return exe
	case t0.Tok == token.IDENT && t0.Str == "command":
		sh.lastCommand = toks[1].Str // 1 is the name -- triggers AddCommand
		toks = toks[2:]              // get rid of first
		toks.Insert(0, token.IDENT, "shell.AddCommand")
		toks.Insert(1, token.LPAREN)
		toks.Insert(2, token.STRING, `"`+sh.lastCommand+`"`)
		toks.Insert(3, token.COMMA)
		toks.Insert(4, token.FUNC)
		toks.Insert(5, token.LPAREN)
		toks.Insert(6, token.IDENT, "args")
		toks.Insert(7, token.ELLIPSIS)
		toks.Insert(8, token.IDENT, "string")
		toks.Insert(9, token.RPAREN)
		toks.AddTokens(sh.TranspileGo(toks[11:]))
	case t0.IsGo():
		if t0.Tok == token.GO {
			if !toks.Contains(token.LPAREN) {
				logx.PrintlnDebug("exec: go command")
				return sh.TranspileExec(ewords, false)
			}
		}
		logx.PrintlnDebug("go    keyword")
		return sh.TranspileGo(toks)
	case t0pn > 0: // path expr
		logx.PrintlnDebug("exec: path...")
		return sh.TranspileExec(ewords, false)
	case t0.Tok == token.STRING:
		logx.PrintlnDebug("exec: string...")
		return sh.TranspileExec(ewords, false)
	case f0exec && en == 1:
		logx.PrintlnDebug("exec: 1 word")
		return sh.TranspileExec(ewords, false)
	case !f0exec: // exec must be IDENT
		logx.PrintlnDebug("go:   not ident")
		return sh.TranspileGo(toks)
	case f0exec && en > 1 && (ewords[1][0] == '=' || ewords[1][0] == ':'):
		logx.PrintlnDebug("go:   assignment or defn")
		return sh.TranspileGo(toks)
	case f0exec: // now any ident
		logx.PrintlnDebug("exec: ident..")
		return sh.TranspileExec(ewords, false)
	// case t0in > 0 && n > t0in && (toks[t1idx].Tok != token.ASSIGN) && (t1in > 0 || t1pn > 0 || toks[t1idx].Tok == token.SUB || toks[t1idx].Tok == token.DEC || toks[t1idx].Tok == token.GTR || toks[t1idx].Tok == token.SHR || toks[t1idx].Tok == token.LSS || toks[t1idx].Tok == token.STRING || toks[t1idx].Tok == token.LBRACE || toks[t1idx].Tok == token.ILLEGAL):
	// 	logx.PrintlnDebug("exec: word non-go...")
	// 	return sh.TranspileExec(ewords, false)
	default:
		logx.PrintlnDebug("go:   default")
		return sh.TranspileGo(toks)
	}
	return toks
}

// TranspileGo returns transpiled tokens assuming Go code.
// Unpacks any backtick encapsulated shell commands.
func (sh *Shell) TranspileGo(toks Tokens) Tokens {
	gtoks := make(Tokens, 0, len(toks)) // return tokens
	for _, tok := range toks {
		if tok.IsBacktickString() {
			gtoks = append(gtoks, sh.TranspileExecString(tok.Str, true)...)
		} else {
			gtoks = append(gtoks, tok)
		}
	}
	return gtoks
}

// TranspileExecString returns transpiled tokens assuming Exec code,
// from a backtick-encoded string, with the given bool indicating
// whether [Output] is needed.
func (sh *Shell) TranspileExecString(str string, output bool) Tokens {
	ewords, err := ExecWords(str[1 : len(str)-1]) // enclosed string
	if err != nil {
		sh.AddError(err)
	}
	return sh.TranspileExec(ewords, output)
}

// TranspileExec returns transpiled tokens assuming Exec code,
// with the given bools indicating the type of run to execute.
func (sh *Shell) TranspileExec(ewords []string, output bool) Tokens {
	n := len(ewords)
	if n == 0 {
		return nil
	}
	etoks := make(Tokens, 0, n+5) // return tokens
	var execTok *Token
	bgJob := false
	noStop := false
	if ewords[0] == "[" {
		ewords = ewords[1:]
		n--
		noStop = true
	}
	startExec := func() {
		bgJob = false
		etoks.Add(token.IDENT, "shell")
		etoks.Add(token.PERIOD)
		switch {
		case output && noStop:
			execTok = etoks.Add(token.IDENT, "OutputErrOK")
		case output && !noStop:
			execTok = etoks.Add(token.IDENT, "Output")
		case !output && noStop:
			execTok = etoks.Add(token.IDENT, "RunErrOK")
		case !output && !noStop:
			execTok = etoks.Add(token.IDENT, "Run")
		}
		etoks.Add(token.LPAREN)
	}
	endExec := func() {
		if bgJob {
			execTok.Str = "Start"
		}
		etoks.DeleteLastComma()
		etoks.Add(token.RPAREN)
	}

	startExec()

	for i := 0; i < n; i++ {
		f := ewords[i]
		switch {
		case f == "{": // embedded go
			if n < i+3 {
				sh.AddError(fmt.Errorf("cosh: no matching right brace } found in exec command line"))
			} else {
				gstr := ewords[i+1]
				etoks.AddTokens(sh.TranspileGo(sh.Tokens(gstr)))
				etoks.Add(token.COMMA)
				i += 2
			}
		case f == "]": // solo is def end
			// just skip
		case f == "&":
			bgJob = true
		case f == ";":
			endExec()
			etoks.Add(token.SEMICOLON)
			etoks.AddTokens(sh.TranspileExec(ewords[i+1:], output))
			return etoks
		default:
			if f[0] == '"' || f[0] == '`' {
				etoks.Add(token.STRING, f)
			} else {
				etoks.Add(token.IDENT, AddQuotes(f)) // mark as an IDENT but add quotes!
			}
			etoks.Add(token.COMMA)
		}
	}
	endExec()
	return etoks
}
