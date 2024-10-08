package main

import (
	"go/ast"
	"go/types"

	"github.com/omni-network/omni/lib/errors"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var arrayCastAnalyzer = &analysis.Analyzer{
	Name: "arraycast",
	Doc:  "Ensures panic-prone array casting isn't used",
	Run: func(p *analysis.Pass) (interface{}, error) {
		if isTestPass(p) { // Skip tests for this analyser
			return nil, nil //nolint:nilnil // API requires nil-nil return
		}
		if p.Pkg.Name() == "cast" { // Skip cast package
			return nil, nil //nolint:nilnil // API requires nil-nil return
		}

		i, ok := p.ResultOf[inspect.Analyzer].(*inspector.Inspector)
		if !ok {
			return nil, errors.New("analyzer is not of type *inspector.Inspector")
		}
		diags, err := arrayCast(p.TypesInfo, i)
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

func arrayCast(info *types.Info, i *inspector.Inspector) ([]analysis.Diagnostic, error) {
	var diags []analysis.Diagnostic
	var err error
	filter := []ast.Node{new(ast.CallExpr)}
	i.Preorder(filter, func(n ast.Node) {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			err = errors.New("expected *ast.CallExpr")
			return
		} else if len(call.Args) != 1 {
			return // Casts should only have one argument
		}

		_, ok = isArrayCast(info, call)
		if !ok {
			return
		}

		if !isSliceArg(info, call.Args[0]) {
			return
		}

		diags = append(diags, analysis.Diagnostic{
			Pos:     call.Pos(),
			Message: "panicable array casting, use cast.Array* instead",
		})
	})

	return diags, err
}

func isSliceArg(info *types.Info, arg ast.Expr) bool {
	argType := info.TypeOf(arg)
	if _, ok := argType.(*types.Slice); ok {
		return true
	}

	if _, ok := argType.Underlying().(*types.Slice); ok {
		return true
	}

	return false
}

func isArrayCast(info *types.Info, call *ast.CallExpr) (int64, bool) {
	funType := info.TypeOf(call.Fun)
	if a, ok := funType.(*types.Array); ok {
		return a.Len(), true
	}

	if a, ok := funType.Underlying().(*types.Array); ok {
		return a.Len(), true
	}

	return 0, false
}
