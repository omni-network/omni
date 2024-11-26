package main

import (
	"bytes"
	"go/ast"
	"go/constant"
	"go/printer"
	"go/token"
	"go/types"
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
		diags, err := logAttrs(i, p.TypesInfo, p.Fset)
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

func logAttrs(i *inspector.Inspector, info *types.Info, fset *token.FileSet) ([]analysis.Diagnostic, error) {
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

		nonNil, ok := selectIndex(selct, noNilIndex)
		if ok && isNil(call.Args[nonNil]) {
			diags = append(diags, analysis.Diagnostic{
				Pos:     call.Args[nonNil].Pos(),
				Message: "info/debug-err called with nil error",
			})
		}

		firstAttr, ok := selectIndex(selct, attrFuncs)
		if !ok {
			return
		}

		if call.Ellipsis != token.NoPos {
			// Ignore attribute ellipses, since it is []any.
			return
		}

		for i := firstAttr; i < len(call.Args); i++ {
			arg := call.Args[i]

			// Attribute keys must be either strings or slog.Attrs
			stringVal, ok := getStringValue(arg, info)
			if !ok {
				if !isStructType(arg, info, "log/slog", "Attr") {
					diags = append(diags, analysis.Diagnostic{
						Pos:     arg.Pos(),
						Message: "bad log/error attribute key: " + format(fset, arg),
					})
				}

				continue
			}

			// Ensure next value isn't an error
			if isError(call.Args, i+1, info) {
				diags = append(diags, analysis.Diagnostic{
					Pos:     arg.Pos(),
					Message: "error attributes not allowed",
				})
			}

			// Skip next value
			i++

			if !isSnakeCase(stringVal) {
				diags = append(diags, analysis.Diagnostic{
					Pos:     arg.Pos(),
					Message: "log/error attribute key must be snake_case; not " + stringVal,
				})
			}
		}
	})

	return diags, err
}

// format returns the string representation of the expression.
func format(fset *token.FileSet, arg ast.Expr) string {
	var out bytes.Buffer
	_ = printer.Fprint(&out, fset, arg)

	return out.String()
}

// getStringValue returns the quoted string value of the expression or false if not a string.
// It supports string literals and values declared in constant or variable specifications.
func getStringValue(expr ast.Expr, info *types.Info) (string, bool) {
	const ignore = `"ignore"` // Dummy response to ignore some expressions
	const typString = "string"

	switch v := expr.(type) {
	case *ast.BasicLit: // Handle string literals
		if v.Kind != token.STRING {
			return "", false
		}

		return v.Value, true
	case *ast.Ident: // Handle identifiers (constants and variables)
		typ, ok := info.Types[v]
		if !ok {
			return "", false
		}
		if typ.Value != nil && typ.Value.Kind() == constant.String {
			return typ.Value.String(), true
		} else if typ.Addressable() && typ.Type.String() == typString {
			return ignore, true // Ignore string variables
		}
	case *ast.SelectorExpr: // Handle other package identifiers (constants and variables)
		obj := info.ObjectOf(v.Sel)
		if obj == nil {
			return "", false
		}

		typConst, ok := obj.(*types.Const)
		if ok && typConst.Val().Kind() == constant.String {
			return typConst.Val().String(), true
		}

		typVar, ok := obj.(*types.Var)
		if ok && typVar.Type().String() == typString {
			return ignore, true // Ignore string variables
		}

		return "", false
	case *ast.CallExpr:
		if info.TypeOf(expr).String() == typString {
			return ignore, true // Ignore any string function call
		}
	}

	return "", false
}

// isStructType returns true if the expression is a pkg.name struct type.
func isStructType(arg ast.Expr, info *types.Info, pkg string, name string) bool {
	// Get the type of the expression from the type-checking information
	typ := info.TypeOf(arg)
	if typ == nil {
		return false
	}

	// Ensure the type is a named type (e.g., pkg.name)
	named, ok := typ.(*types.Named)
	if !ok {
		return false
	}

	// Check if the underlying type is a struct
	_, ok = named.Underlying().(*types.Struct)
	if !ok {
		return false
	}

	// Check the package path and type name
	obj := named.Obj()
	if obj == nil {
		return false
	}

	return obj.Pkg().Path() == pkg && obj.Name() == name
}

// isError checks if the ith index expression in args implements the error interface.
func isError(args []ast.Expr, i int, info *types.Info) bool {
	// Ensure the index is within bounds
	if i >= len(args) {
		return false
	}

	// Get the type of the expression
	expr := args[i]
	typ := info.TypeOf(expr)
	if typ == nil {
		return false
	}

	// Check if the type implements the error interface
	errorType := types.Universe.Lookup("error").Type().Underlying().(*types.Interface)

	return types.Implements(typ, errorType)
}

// isNil returns true if the expression is nil.
func isNil(expr ast.Expr) bool {
	ident, ok := expr.(*ast.Ident)
	if !ok {
		return false
	}

	return ident.Name == "nil" && ident.Obj == nil
}

var snakeRegex = regexp.MustCompile(`"[a-z0-9_]+"`)

func isSnakeCase(value string) bool {
	return snakeRegex.MatchString(value)
}

// selectIndex returns the selector index matching the provided select expr.
func selectIndex(selct *ast.SelectorExpr, selectors map[string]map[string]int) (int, bool) {
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

// noNilIndex is a map of package name to function name to the index that must not be nil.
var noNilIndex = map[string]map[string]int{
	"log": {
		"DebugErr": 2,
		"InfoErr":  2,
	},
}

// attrFuncs is a map of package name to function name to the index of the first slog any attribute.
var attrFuncs = map[string]map[string]int{
	"log": {
		"Debug":    2,
		"DebugErr": 3,
		"Info":     2,
		"InfoErr":  3,
		"Warn":     3,
		"Error":    3,
	},
	"errors": {
		"Wrap": 2,
		"New":  1,
	},
}
