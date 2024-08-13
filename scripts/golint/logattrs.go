package main

import (
	"go/ast"
	"go/token"
	"regexp"

	"github.com/omni-network/omni/lib/errors"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var logAttrsAnalyzer = &analysis.Analyzer{
	Name: "logattrs",
	Doc:  "Ensures snake_case log and error attribute keys",
	Run: func(p *analysis.Pass) (interface{}, error) {
		i, ok := p.ResultOf[inspect.Analyzer].(*inspector.Inspector)
		if !ok {
			return nil, errors.New("analyzer is not of type *inspector.Inspector")
		}
		diags, err := logAttrs(i)
		if err != nil {
			return nil, errors.Wrap(err, "detect uint subtract")
		}

		for _, diag := range diags {
			p.Report(diag)
		}

		return nil, nil //nolint:nilnil // API requires nil-nil return
	},
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func logAttrs(i *inspector.Inspector) ([]analysis.Diagnostic, error) {
	var diags []analysis.Diagnostic
	var err error
	filter := []ast.Node{new(ast.CallExpr)}
	i.Preorder(filter, func(n ast.Node) {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			err = errors.New("expected *ast.CallExpr")
			return
		}

		selct, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}

		index, ok := attrIndex(selct, attrFuncs)
		if !ok {
			return
		}

		for i := index; i < len(call.Args); i++ {
			arg := call.Args[i]

			strLiteral, ok := arg.(*ast.BasicLit)
			if !ok {
				continue
			}
			if strLiteral.Kind != token.STRING {
				return // Unexpected
			}

			// Skip next value
			i++

			if !isSnakeCase(strLiteral.Value) {
				diags = append(diags, analysis.Diagnostic{
					Pos:     strLiteral.Pos(),
					Message: "log/error attribute key must be snake_case; not " + strLiteral.Value,
				})
			}
		}
	})

	return diags, err
}

var snakeRegex = regexp.MustCompile(`"[a-z0-9_]+"`)

func isSnakeCase(value string) bool {
	return snakeRegex.MatchString(value)
}

func attrIndex(selct *ast.SelectorExpr, selectors map[string]map[string]int) (int, bool) {
	funIdent, ok := selct.X.(*ast.Ident)
	if !ok {
		return 0, false
	}

	funcs, ok := selectors[funIdent.Name]
	if !ok {
		return 0, false
	}

	index, ok := funcs[selct.Sel.Name]

	return index, ok
}

var attrFuncs = map[string]map[string]int{
	"log": {
		"Debug": 2,
		"Info":  2,
		"Warn":  3,
		"Error": 3,
	},
	"errors": {
		"Wrap": 2,
		"New":  1,
	},
}
