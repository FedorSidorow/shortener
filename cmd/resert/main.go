package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// StructInfo holds information about a struct that needs a Reset method.
type StructInfo struct {
	Name   string
	Fields []FieldInfo
}

// FieldInfo describes a field of a struct.
type FieldInfo struct {
	Name     string
	Type     string
	IsPtr    bool
	IsSlice  bool
	IsMap    bool
	IsStruct bool
}

// packageMap maps package directory to its structs.
var packageMap = make(map[string][]StructInfo)

func main() {
	root := "." // start from current directory
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Skip directories that are not Go packages (e.g., vendor, .git)
		if d.IsDir() {
			if shouldSkipDir(path) {
				return filepath.SkipDir
			}
			return nil
		}
		// Only process .go files (excluding generated files)
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, ".gen.go") {
			return nil
		}
		// Parse the file
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			log.Printf("Failed to parse %s: %v", path, err)
			return nil
		}
		// Find package directory
		pkgDir := filepath.Dir(path)
		// Process AST for structs with // generate:reset comment
		processFile(fset, f, pkgDir)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	// Generate reset.gen.go for each package that has structs
	for pkgDir, structs := range packageMap {
		if len(structs) > 0 {
			if err := generateResetFile(pkgDir, structs); err != nil {
				log.Printf("Failed to generate reset.gen.go for %s: %v", pkgDir, err)
			}
		}
	}
}

func shouldSkipDir(path string) bool {
	base := filepath.Base(path)
	// Skip hidden directories (except "." and ".."), vendor, testdata, etc.
	if base == "." || base == ".." {
		return false
	}
	if strings.HasPrefix(base, ".") || base == "vendor" || base == "testdata" {
		return true
	}
	return false
}

func processFile(fset *token.FileSet, f *ast.File, pkgDir string) {
	// Iterate over all declarations in the file
	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		// Look for type specs inside this declaration
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}
			// Check for generate:reset comment on the type spec or its parent GenDecl
			hasResetComment := false
			if typeSpec.Doc != nil {
				for _, comment := range typeSpec.Doc.List {
					if strings.Contains(comment.Text, "generate:reset") {
						hasResetComment = true
						break
					}
				}
			}
			if !hasResetComment && genDecl.Doc != nil {
				for _, comment := range genDecl.Doc.List {
					if strings.Contains(comment.Text, "generate:reset") {
						hasResetComment = true
						break
					}
				}
			}
			if !hasResetComment {
				continue
			}
			// Collect fields
			var fields []FieldInfo
			for _, field := range structType.Fields.List {
				if len(field.Names) == 0 {
					// Embedded field (anonymous) - we need to handle it
					// For now, skip because we don't know how to reset it.
					continue
				}
				for _, name := range field.Names {
					fieldType := exprToString(field.Type)
					fields = append(fields, FieldInfo{
						Name:     name.Name,
						Type:     fieldType,
						IsPtr:    isPointer(field.Type),
						IsSlice:  isSlice(field.Type),
						IsMap:    isMap(field.Type),
						IsStruct: isStructType(field.Type),
					})
				}
			}
			packageMap[pkgDir] = append(packageMap[pkgDir], StructInfo{
				Name:   typeSpec.Name.Name,
				Fields: fields,
			})
		}
	}
}

// exprToString converts an ast.Expr to a string representation.
func exprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + exprToString(t.X)
	case *ast.ArrayType:
		if t.Len == nil {
			return "[]" + exprToString(t.Elt)
		}
		// array like [N]T
		return fmt.Sprintf("[%s]%s", exprToString(t.Len), exprToString(t.Elt))
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", exprToString(t.Key), exprToString(t.Value))
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name
	default:
		return fmt.Sprintf("%v", expr)
	}
}

func isPointer(expr ast.Expr) bool {
	_, ok := expr.(*ast.StarExpr)
	return ok
}

func isSlice(expr ast.Expr) bool {
	arr, ok := expr.(*ast.ArrayType)
	return ok && arr.Len == nil
}

func isMap(expr ast.Expr) bool {
	_, ok := expr.(*ast.MapType)
	return ok
}

func isStructType(expr ast.Expr) bool {
	// Determine if the expression is a built-in type.
	// If it's a built-in, it's not a struct.
	switch t := expr.(type) {
	case *ast.Ident:
		name := t.Name
		// Built-in types
		builtins := map[string]bool{
			"bool": true, "string": true,
			"int": true, "int8": true, "int16": true, "int32": true, "int64": true,
			"uint": true, "uint8": true, "uint16": true, "uint32": true, "uint64": true, "uintptr": true,
			"float32": true, "float64": true,
			"complex64": true, "complex128": true,
			"byte": true, "rune": true,
			"error": true,
		}
		if builtins[name] {
			return false
		}
		// Otherwise assume it's a struct (could be a type alias, but we treat as struct)
		return true
	case *ast.SelectorExpr:
		// Qualified identifier (pkg.Type) - assume it's a struct (could be something else)
		return true
	default:
		return false
	}
}

func generateResetFile(pkgDir string, structs []StructInfo) error {
	var buf bytes.Buffer
	buf.WriteString("// Code generated by resert; DO NOT EDIT.\n\n")
	buf.WriteString("package " + filepath.Base(pkgDir) + "\n\n")
	for _, s := range structs {
		buf.WriteString(fmt.Sprintf("func (x *%s) Reset() {\n", s.Name))
		for _, f := range s.Fields {
			buf.WriteString(generateResetStatement("x."+f.Name, f) + "\n")
		}
		buf.WriteString("}\n\n")
	}
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("formatting error: %v", err)
	}
	outPath := filepath.Join(pkgDir, "reset.gen.go")
	return os.WriteFile(outPath, formatted, 0644)
}

func generateResetStatement(fieldPath string, f FieldInfo) string {
	if f.IsPtr {
		underlying := strings.TrimPrefix(f.Type, "*")
		if isBuiltInType(underlying) {
			return fmt.Sprintf("\tif %s != nil {\n\t\t*%s = %s\n\t}", fieldPath, fieldPath, zeroValue(underlying))
		}
		return fmt.Sprintf("\tif %s != nil {\n\t\t*%s = %s{}\n\t}", fieldPath, fieldPath, underlying)
	}
	if f.IsSlice {
		return fmt.Sprintf("\t%s = %s[:0]", fieldPath, fieldPath)
	}
	if f.IsMap {
		return fmt.Sprintf("\tclear(%s)", fieldPath)
	}
	if f.IsStruct {
		return fmt.Sprintf("\t%s = %s{}", fieldPath, f.Type)
	}
	return fmt.Sprintf("\t%s = %s", fieldPath, zeroValue(f.Type))
}

func isBuiltInType(name string) bool {
	builtins := map[string]bool{
		"bool": true, "string": true,
		"int": true, "int8": true, "int16": true, "int32": true, "int64": true,
		"uint": true, "uint8": true, "uint16": true, "uint32": true, "uint64": true, "uintptr": true,
		"float32": true, "float64": true,
		"complex64": true, "complex128": true,
		"byte": true, "rune": true,
		"error": true,
	}
	return builtins[name]
}

func zeroValue(typ string) string {
	switch typ {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
		"float32", "float64", "complex64", "complex128", "byte", "rune":
		return "0"
	case "string":
		return `""`
	case "bool":
		return "false"
	default:
		return typ + "{}"
	}
}
