package main

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/omni-network/omni/lib/errors"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var uintTypes = map[string]bool{
	"uint":   true,
	"uint8":  true,
	"uint16": true,
	"uint32": true,
	"uint64": true,
	"byte":   true,
}

var uintSubtractAnalyzer = &analysis.Analyzer{
	Name: "uintsubtract",
	Doc: "Prevents subtraction of uints, since it may underflow. " +
		"Use a 'uintsub' function with 'a-b' to bypass linter.",
	Run: func(p *analysis.Pass) (interface{}, error) {
		i, ok := p.ResultOf[inspect.Analyzer].(*inspector.Inspector)
		if !ok {
			return nil, errors.New("analyzer is not of type *inspector.Inspector")
		}
		diags, err := uintstract(i, p.TypesInfo)
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

func uintstract(i *inspector.Inspector, typeInfo *types.Info) ([]analysis.Diagnostic, error) {
	var diags []analysis.Diagnostic
	var err error
	filter := []ast.Node{new(ast.BinaryExpr)}
	i.Preorder(filter, func(n ast.Node) {
		bin, ok := n.(*ast.BinaryExpr)
		if !ok {
			err = errors.New("expected *ast.BinaryExpr")
			return
		}

		if bin.Op != token.SUB {
			return // Not a subtraction (-)
		}

		if isVar(bin.X, "a") && isVar(bin.Y, "b") {
			return // Ignore 'a-b'
		}

		leftType := typeInfo.TypeOf(bin.X)
		if leftType == nil {
			err = errors.New("left type not found")
			return
		}

		if uintTypes[leftType.String()] {
			diags = append(diags, analysis.Diagnostic{
				Pos:     bin.OpPos,
				End:     bin.OpPos,
				Message: "uint subtraction",
			})

			return
		}

		rightType := typeInfo.TypeOf(bin.Y)
		if rightType == nil {
			err = errors.New("right type not found")
			return
		}

		if uintTypes[rightType.String()] {
			diags = append(diags, analysis.Diagnostic{
				Pos:     bin.OpPos,
				End:     bin.OpPos,
				Message: "uint subtraction",
			})

			return
		}
	})

	return diags, err
}

func isVar(exp ast.Expr, name string) bool {
	ident, ok := exp.(*ast.Ident)
	if !ok {
		return false
	}

	return ident.Name == name
}
