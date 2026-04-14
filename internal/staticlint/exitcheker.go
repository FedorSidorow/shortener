package exitcheker

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// ExitAnalyzer Линтер запрещает вызов os.Exit, log.Fatal и panic в функции main
var ExitAnalyzer = &analysis.Analyzer{
	Name: "exitcheck",
	Doc:  "check for direct os.Exit, log.Fatal, and panic calls in main",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {

		for _, decl := range file.Decls {

			fn, ok := decl.(*ast.FuncDecl)

			if !ok {
				continue
			}

			if fn.Name.Name == "main" {
				ast.Inspect(fn.Body, func(node ast.Node) bool {

					switch x := node.(type) {

					case *ast.CallExpr:
						// Check for panic (built-in)
						if ident, ok := x.Fun.(*ast.Ident); ok && ident.Name == "panic" {
							pass.Reportf(x.Fun.Pos(), "direct call to panic in main")
						}
						// Check for selector expressions (package.Func)
						if se, ok := x.Fun.(*ast.SelectorExpr); ok {
							// Use type information to get the actual package
							if pass.TypesInfo != nil {
								obj := pass.TypesInfo.ObjectOf(se.Sel)
								if obj != nil && obj.Pkg() != nil {
									pkgPath := obj.Pkg().Path()
									switch pkgPath {
									case "os":
										if se.Sel.Name == "Exit" {
											pass.Reportf(x.Fun.Pos(), "direct call to os.Exit in main")
										}
									case "log":
										if se.Sel.Name == "Fatal" {
											pass.Reportf(x.Fun.Pos(), "direct call to log.Fatal in main")
										}
									}
								}
							} else {
								// Fallback to identifier-based check
								if pkgIdent, ok := se.X.(*ast.Ident); ok {
									switch pkgIdent.Name {
									case "os":
										if se.Sel.Name == "Exit" {
											pass.Reportf(x.Fun.Pos(), "direct call to os.Exit in main")
										}
									case "log":
										if se.Sel.Name == "Fatal" {
											pass.Reportf(x.Fun.Pos(), "direct call to log.Fatal in main")
										}
									}
								}
							}
						}

					}

					return true
				})
			}
		}
	}

	return nil, nil
}
