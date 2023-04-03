// Utilities for generating TypeScript code.

package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var (
	methodDefinition = regexp.MustCompile(`^((static|get|set) )?[a-zA-Z0-9_]+\(.*\)(: .*?)? \{$`)
)

type tsImport struct {
	pkg, defaultImport string
	symbols            map[string]string
}

func (t *tsImport) String() string {
	out := "import"
	from := false
	if t.defaultImport != "" {
		from = true
		out += " "
		out += t.defaultImport
		if len(t.symbols) > 0 {
			out += ","
		}
	}
	if len(t.symbols) > 0 {
		from = true
		out += " {"
		sortedSymbols := make([]string, 0, len(t.symbols))
		for s := range t.symbols {
			sortedSymbols = append(sortedSymbols, s)
		}
		sort.StringSlice(sortedSymbols).Sort()
		for i, s := range sortedSymbols {
			out += " "
			out += s
			if alias := t.symbols[s]; s != alias {
				out += " as "
				out += alias
			}
			if i < len(t.symbols)-1 {
				out += ","
			}
		}
		out += " }"
	}
	if from {
		out += " from"
	}
	out += ` "` + t.pkg + *importModuleSpecifierEnding + `";`
	return out
}

// TS represents TypeScript source code.
type TS struct {
	JS bool

	indentation      string
	buf              string
	lastLine         string
	imports          map[string]*tsImport
	scope            []string
	methodScopeIndex int
}

func (t *TS) Lf(format string, args ...any) {
	t.L(fmt.Sprintf(format, args...))
}

func (t *TS) L(args ...any) {
	content := ""
	for i, a := range args {
		content += fmt.Sprint(a)
		if i < len(args)-1 {
			content += " "
		}
	}

	if content == "" {
		// This code should not be reached since we insert blank lines using
		// `shouldInsertBlanklineBefore`
		panic("manual blank lines are not allowed")
	}

	// Note: not using isBlockEnd here to account for blocks ending and
	// starting in the same line, e.g. "} else {"
	if strings.HasPrefix(content, "}") {
		t.indent(-1, content)
		if t.methodScopeIndex >= len(t.scope) {
			t.methodScopeIndex = -1
		}
	}
	if t.shouldInsertBlanklineBefore(content) {
		t.buf += "\n"
	}
	t.buf += t.indentation + content + "\n"
	t.lastLine = content
	if t.isBlockStart(content) {
		t.indent(+1, content)
		if isMethodOrConstructorDefinitionStart(content) {
			t.methodScopeIndex = len(t.scope) - 1
		}
	}
}

func (t *TS) isBlockStart(content string) bool {
	return !strings.HasPrefix(content, " *") && strings.HasSuffix(content, "{")
}
func (t *TS) isBlockEnd(content string) bool {
	return !strings.HasPrefix(content, " *") && (strings.HasSuffix(content, "}") || strings.HasPrefix(content, "},") || content == "};" || strings.HasPrefix(content, "})()"))
}

func (t *TS) shouldInsertBlanklineBefore(content string) bool {
	// Never start the file with a blank line
	if t.buf == "" {
		return false
	}
	// Don't insert blank lines in generated method implementations
	if t.methodScopeIndex > 0 {
		return false
	}

	// Require blank lines before comment
	if strings.HasPrefix(content, "//") {
		return true
	}
	// Require blank lines before block comments, except after block opening
	if strings.HasPrefix(content, "/**") && !t.isBlockStart(t.lastLine) {
		return true
	}
	// Require blank lines before "export " or start of method definition, except
	// after block comment ending or block opening
	if (strings.HasPrefix(content, "export ") ||
		isMethodOrConstructorDefinitionStart(content)) &&
		!strings.HasSuffix(t.buf, "*/\n") && !t.isBlockStart(t.lastLine) {
		return true
	}
	// Require blank lines before anything that follows a block closing, unless
	// it's another block closing
	if content != "}" && content != "}," && t.isBlockEnd(t.lastLine) {
		return true
	}
	// Require blank line after assigning prototype properties and before
	// returning
	if strings.HasPrefix(content, "return ") && strings.Contains(t.lastLine, ".prototype.") && !strings.HasPrefix(t.lastLine, "Object.defineProperty(") {
		return true
	}

	return false
}

func (t *TS) indent(indent int, content string) {
	nSpaces := indent * 2

	if nSpaces < 0 {
		if len(t.indentation) < -nSpaces {
			logf("WARNING: tried to outdent too many times. Check implementations of isBlockStart/isBlockEnd")
			return
		}
		// Outdent
		t.indentation = t.indentation[:len(t.indentation)+nSpaces]
		t.scope = t.scope[:len(t.scope)-1]
	} else {
		for i := 0; i < nSpaces; i++ {
			t.indentation += " "
		}
		t.scope = append(t.scope, content)
	}
}

func (t *TS) addImport(imp *tsImport) {
	if t.imports == nil {
		t.imports = map[string]*tsImport{}
	}
	existing, ok := t.imports[imp.pkg]
	if !ok {
		t.imports[imp.pkg] = imp
		return
	}
	for s, a := range imp.symbols {
		existing.symbols[s] = a
	}
	if imp.defaultImport != "" {
		existing.defaultImport = imp.defaultImport
	}
}

func (t *TS) BlockComment(comment string) {
	if comment == "" {
		return
	}
	t.Lf("/**")
	for _, line := range strings.Split(strings.TrimRight(comment, "\n"), "\n") {
		// If the line itself contains block comment endings, need to somehow
		// prevent those from closing the comment. For now, we replace them with
		// "* /"
		line = strings.ReplaceAll(line, "*/", "* /")
		t.Lf(" *%s", line)
	}
	t.Lf(" */")
}

func (t *TS) DefaultImport(pkg, nameAndAlias string) {
	t.addImport(&tsImport{pkg: pkg, defaultImport: nameAndAlias})
}

func (g *TS) Import(pkg, symbol, alias string) {
	g.addImport(&tsImport{pkg: pkg, symbols: map[string]string{symbol: alias}})
}

func (g *TS) String() string {
	pkgs := make([]string, 0, len(g.imports))
	for pkg := range g.imports {
		pkgs = append(pkgs, pkg)
	}
	sort.Slice(pkgs, func(i, j int) bool {
		doti := strings.HasPrefix(pkgs[i], ".")
		dotj := strings.HasPrefix(pkgs[j], ".")
		if doti != dotj {
			return dotj
		}
		return pkgs[i] < pkgs[j]
	})
	importSection := ""
	for _, pkg := range pkgs {
		importSection += g.imports[pkg].String() + "\n"
	}
	importSection += "\n"

	jsHeader := ""
	if g.JS {
		jsHeader += "/*eslint-disable*/\n"
		jsHeader = "\"use strict\";\n"
		jsHeader += "\n"
	}

	return jsHeader + importSection + g.buf
}

func isMethodOrConstructorDefinitionStart(line string) bool {
	return methodDefinition.MatchString(line)
}
