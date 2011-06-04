// Copyright (c) 2010 CZ.NIC z.s.p.o. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// blame: jnml, labs.nic.cz


package main


import (
	"github.com/cznic/lex"
	"github.com/cznic/lexer"
	"fmt"
	"strings"
)


type renderGo struct {
	noRender
	scStates map[int]bool
}


func (r *renderGo) prolog(l *lex.L) {
	for _, state := range l.StartConditionsStates {
		r.scStates[int(state.Index)] = true
	}
	for _, state := range l.StartConditionsBolStates {
		r.scStates[int(state.Index)] = true
	}
	for _, line := range l.DefCode {
		r.w.Write([]byte(line))
	}
	r.wprintf("\nyystate0:\n")
	if action0 := l.Rules[0].Action; action0 != "" {
		r.w.Write([]byte(action0))
	}
	scNames := map[int]string{}
	for name, i := range l.StartConditions {
		scNames[i] = name
	}
	if len(l.StartConditionsStates) > 1 || len(l.StartConditionsBolStates) != 0 {
		if len(l.StartConditionsBolStates) == 0 {
			r.wprintf("\n\nswitch yyt := %s; yyt {\n", l.YYT)
		} else {
			r.wprintf("\n\nswitch yyt, yyb := %s, %s; yyt {\n", l.YYT, l.YYB)
		}
		r.wprintf("default:\npanic(fmt.Errorf(`invalid start condition %%d`, yyt))\n")
		for sc, state := range l.StartConditionsStates {
			r.wprintf("case %d: // start condition: %s\n", sc, scNames[sc])
			if state, ok := l.StartConditionsBolStates[sc]; ok {
				r.wprintf("if yyb { goto yystart%d }\n", state.Index)
			}
			r.wprintf("goto yystart%d\n", state.Index)
		}
		r.wprintf("}\n\n")
	} else {
		r.wprintf("\n\ngoto yystart%d\n\n", l.StartConditionsStates[0].Index)
	}
}


func (r *renderGo) rules(l *lex.L) {
	for i := 1; i < len(l.Rules); i++ {
		rule := l.Rules[i]
		r.wprintf("yyrule%d: // %s\n", i, rule.Pattern)
		act := strings.TrimSpace(rule.Action)
		if act != "" && act != "|" {
			r.wprintf("{\n")
			r.w.Write([]byte(rule.Action))
		}
		if act != "|" {
			r.wprintf("\ngoto yystate0\n")
		}
		if act != "" && act != "|" {
			r.wprintf("}\n")
		}
	}
	r.wprintf(`panic("unreachable")` + "\n")
}


func (r *renderGo) scanFail(l *lex.L) {
	r.wprintf("\ngoto yyabort // silence unused label error\n")
	r.wprintf("\nyyabort: // no lexem recognized\n")
}


func (r *renderGo) userCode(l *lex.L) {
	if userCode := l.UserCode; userCode != "" {
		r.w.Write([]byte(userCode))
	}
}


func (r *renderGo) defaultTransition(l *lex.L, state *lexer.NfaState) (defaultEdge *lexer.RangesEdge) {
	r.wprintf("default:\n")
	if rule, ok := l.Accepts[state]; ok {
		r.wprintf("goto yyrule%d\n", rule)
		return
	}

	cases := map[int]bool{}
	for i := 0; i < 256; i++ {
		cases[i] = true
	}
	for _, edge0 := range state.Consuming {
		switch edge := edge0.(type) {
		default:
			panic("unexpected type")
		case *lexer.RuneEdge:
			cases[edge.Rune] = false, false
		case *lexer.RangesEdge:
			if defaultEdge == nil || len(edge.Ranges.R32) > len(defaultEdge.Ranges.R32) {
				defaultEdge = edge
			}
			for _, rng := range edge.Ranges.R32 {
				for c := rng.Lo; c <= rng.Hi; c += rng.Stride {
					cases[int(c)] = false, false
				}
			}
		}
	}
	if len(cases) != 0 {
		r.wprintf("goto yyabort\n")
		return nil
	}

	if defaultEdge != nil {
		r.wprintf("goto yystate%d // %s\n", defaultEdge.Target().Index, r.rangesEdgeString(defaultEdge, l))
		return
	}

	panic("internal error")
}


func (r *renderGo) rangesEdgeString(edge *lexer.RangesEdge, l *lex.L) string {
	a := []string{}
	for _, rng := range edge.Ranges.R32 {
		if rng.Stride != 1 {
			panic("internal error")
		}

		if rng.Hi-rng.Lo == 1 {
			a = append(a, fmt.Sprintf("%s == %s || %s == %s", l.YYC, q(rng.Lo), l.YYC, q(rng.Hi)))
			continue
		}

		if rng.Hi-rng.Lo > 0 {
			a = append(a, fmt.Sprintf("%s >= %s && %s <= %s", l.YYC, q(rng.Lo), l.YYC, q(rng.Hi)))
			continue
		}

		// rng.Hi == rng.Lo
		a = append(a, fmt.Sprintf("%s == %s", l.YYC, q(rng.Lo)))
	}
	return strings.Replace(strings.Join(a, " || "), "%", "%%", -1)
}


func (r *renderGo) transitions(l *lex.L, state *lexer.NfaState) {
	r.wprintf("switch {\n")
	var defaultEdge lexer.Edger = r.defaultTransition(l, state)
	for _, edge0 := range state.Consuming {
		if edge0 == defaultEdge {
			continue
		}

		r.wprintf("case ")
		switch edge := edge0.(type) {
		default:
			panic("unexpected type")
		case *lexer.RuneEdge:
			r.wprintf("%s == %s", l.YYC, q(uint32(edge.Rune)))
		case *lexer.RangesEdge:
			r.wprintf(r.rangesEdgeString(edge, l))
		}
		r.wprintf(":\ngoto yystate%d\n", edge0.Target().Index)
	}
	r.wprintf("}\n\n")
}

func (r *renderGo) states(l *lex.L) {
	for _, state := range l.Dfa {
		iState := int(state.Index)
		if _, ok := r.scStates[iState]; ok {
			r.wprintf("goto yystate%d // silence unused label error\n", iState)
		}
		r.wprintf("yystate%d:\n", iState)
		rule, ok := l.Accepts[state]
		if !ok || !l.Rules[rule].EOL {
			r.wprintf("%s\n", l.YYN)
		}
		if _, ok := r.scStates[iState]; ok {
			r.wprintf("yystart%d:\n", iState)
		}
		if len(state.Consuming) != 0 {
			r.transitions(l, state)
		} else {
			if rule, ok := l.Accepts[state]; ok {
				r.wprintf("goto yyrule%d\n\n", rule)
			} else {
				panic("internal error")
			}
		}
	}
}


func (r renderGo) render(srcname string, l *lex.L) {
	r.prolog(l)
	r.states(l)
	r.rules(l)
	r.scanFail(l)
	r.userCode(l)
}
