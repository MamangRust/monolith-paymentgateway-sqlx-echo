package main

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	path := os.Args[1]
	if err := generateDocs(path); err != nil {
		fmt.Printf("Error generating documentation: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: docgen <path-to-go-file-or-folder>")
	fmt.Println("\nGenerates Markdown documentation for Go packages")
	fmt.Println("Features:")
	fmt.Println("  - Package overview")
	fmt.Println("  - Type documentation with fields")
	fmt.Println("  - Method signatures")
	fmt.Println("  - Function documentation")
	fmt.Println("  - Proper Markdown formatting")
	fmt.Println("  - Outputs to <package-name>.md file")
}

func generateDocs(path string) error {
	// Check if path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("path does not exist: %s", path)
	}

	fs := token.NewFileSet()
	pkgs, err := parser.ParseDir(fs, path, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse directory: %w", err)
	}

	if len(pkgs) == 0 {
		return fmt.Errorf("no Go packages found in directory: %s", path)
	}

	for pkgName, pkg := range pkgs {
		// Skip test packages
		if strings.HasSuffix(pkgName, "_test") {
			continue
		}

		docPkg := doc.New(pkg, "./", doc.AllDecls)
		if err := generatePackageDocs(docPkg, pkgName, path); err != nil {
			return fmt.Errorf("failed to generate docs for package %s: %w", pkgName, err)
		}
		fmt.Printf("✅ Generated documentation for package '%s' -> %s.md\n", pkgName, pkgName)
	}
	return nil
}

func generatePackageDocs(docPkg *doc.Package, pkgName string, sourcePath string) error {
	// Create output file
	filename := fmt.Sprintf("%s.md", pkgName)
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", filename, err)
	}
	defer file.Close()

	// Write package header
	writeToFile(file, fmt.Sprintf("# 📦 Package `%s`\n\n", pkgName))
	writeToFile(file, fmt.Sprintf("**Source Path:** `%s`\n\n", sourcePath))

	// Package overview
	if docPkg.Doc != "" {
		writeToFile(file, "## 📝 Overview\n\n")
		writeToFile(file, fmt.Sprintf("%s\n\n", strings.TrimSpace(docPkg.Doc)))
	}

	// Constants
	if len(docPkg.Consts) > 0 {
		writeToFile(file, "## 🔢 Constants\n\n")
		for _, c := range docPkg.Consts {
			printDeclarations(file, "Const", c.Doc, c.Decl)
		}
	}

	// Variables
	if len(docPkg.Vars) > 0 {
		writeToFile(file, "## 🏷️ Variables\n\n")
		for _, v := range docPkg.Vars {
			printDeclarations(file, "Var", v.Doc, v.Decl)
		}
	}

	// Types
	if len(docPkg.Types) > 0 {
		writeToFile(file, "## 🧩 Types\n\n")
		for _, t := range docPkg.Types {
			printType(file, t)
		}
	}

	// Functions
	if len(docPkg.Funcs) > 0 {
		writeToFile(file, "## 🚀 Functions\n\n")
		for _, f := range docPkg.Funcs {
			printFunc(file, f)
		}
	}

	return nil
}

func writeToFile(file *os.File, content string) {
	if _, err := file.WriteString(content); err != nil {
		fmt.Printf("Warning: failed to write to file: %v\n", err)
	}
}

func printType(file *os.File, t *doc.Type) {
	writeToFile(file, fmt.Sprintf("### `%s`\n\n", t.Name))

	if t.Doc != "" {
		writeToFile(file, fmt.Sprintf("%s\n\n", strings.TrimSpace(t.Doc)))
	}

	// Print type definition
	writeToFile(file, "```go\n")
	switch spec := t.Decl.Specs[0].(type) {
	case *ast.TypeSpec:
		switch typ := spec.Type.(type) {
		case *ast.StructType:
			writeToFile(file, fmt.Sprintf("type %s struct {\n", t.Name))
			if typ.Fields != nil {
				for _, f := range typ.Fields.List {
					printField(file, f)
				}
			}
			writeToFile(file, "}\n")
		case *ast.InterfaceType:
			writeToFile(file, fmt.Sprintf("type %s interface {\n", t.Name))
			if typ.Methods != nil {
				for _, f := range typ.Methods.List {
					printField(file, f)
				}
			}
			writeToFile(file, "}\n")
		default:
			writeToFile(file, fmt.Sprintf("type %s %s\n", t.Name, astString(typ)))
		}
	}
	writeToFile(file, "```\n\n")

	// Print methods
	if len(t.Methods) > 0 {
		writeToFile(file, "#### Methods\n\n")
		for _, m := range t.Methods {
			printMethod(file, m)
		}
	}
}

func printField(file *os.File, f *ast.Field) {
	var names []string
	for _, name := range f.Names {
		names = append(names, name.Name)
	}

	fieldStr := "\t"
	if len(names) > 0 {
		fieldStr += fmt.Sprintf("%s ", strings.Join(names, ", "))
	}
	fieldStr += astString(f.Type)

	if f.Tag != nil {
		fieldStr += fmt.Sprintf(" %s", f.Tag.Value)
	}
	fieldStr += "\n"

	writeToFile(file, fieldStr)
}

func printMethod(file *os.File, m *doc.Func) {
	writeToFile(file, fmt.Sprintf("##### `%s`\n\n", m.Name))
	if m.Doc != "" {
		writeToFile(file, fmt.Sprintf("%s\n\n", strings.TrimSpace(m.Doc)))
	}
	writeToFile(file, "```go\n")

	// Convert function declaration to string
	fset := token.NewFileSet()
	var funcStr strings.Builder
	if err := printer.Fprint(&funcStr, fset, m.Decl); err == nil {
		writeToFile(file, funcStr.String())
	}
	writeToFile(file, "\n```\n\n")
}

func printFunc(file *os.File, f *doc.Func) {
	writeToFile(file, fmt.Sprintf("### `%s`\n\n", f.Name))
	if f.Doc != "" {
		writeToFile(file, fmt.Sprintf("%s\n\n", strings.TrimSpace(f.Doc)))
	}
	writeToFile(file, "```go\n")

	fset := token.NewFileSet()
	var funcStr strings.Builder
	if err := printer.Fprint(&funcStr, fset, f.Decl); err == nil {
		writeToFile(file, funcStr.String())
	}
	writeToFile(file, "\n```\n\n")
}

func printDeclarations(file *os.File, kind, doc string, decl *ast.GenDecl) {
	if doc != "" {
		writeToFile(file, fmt.Sprintf("**%s:**\n\n%s\n\n", kind, strings.TrimSpace(doc)))
	}

	writeToFile(file, "```go\n")
	fset := token.NewFileSet()
	var declStr strings.Builder
	if err := printer.Fprint(&declStr, fset, decl); err == nil {
		writeToFile(file, declStr.String())
	}
	writeToFile(file, "\n```\n\n")
}

func astString(n ast.Node) string {
	if n == nil {
		return ""
	}

	switch x := n.(type) {
	case *ast.Ident:
		return x.Name
	case *ast.StarExpr:
		return "*" + astString(x.X)
	case *ast.ArrayType:
		return "[]" + astString(x.Elt)
	case *ast.SelectorExpr:
		return astString(x.X) + "." + x.Sel.Name
	case *ast.MapType:
		return "map[" + astString(x.Key) + "]" + astString(x.Value)
	case *ast.FuncType:
		params := ""
		results := ""
		if x.Params != nil {
			params = astString(x.Params)
		}
		if x.Results != nil {
			results = " " + astString(x.Results)
		}
		return "func" + params + results
	case *ast.FieldList:
		if x == nil || len(x.List) == 0 {
			return "()"
		}
		var fields []string
		for _, f := range x.List {
			typeStr := astString(f.Type)
			if len(f.Names) > 0 {
				var names []string
				for _, name := range f.Names {
					names = append(names, name.Name)
				}
				fields = append(fields, strings.Join(names, ", ")+" "+typeStr)
			} else {
				fields = append(fields, typeStr)
			}
		}
		return "(" + strings.Join(fields, ", ") + ")"
	case *ast.ChanType:
		dir := ""
		switch x.Dir {
		case ast.SEND:
			dir = "chan<- "
		case ast.RECV:
			dir = "<-chan "
		default:
			dir = "chan "
		}
		return dir + astString(x.Value)
	case *ast.Ellipsis:
		return "..." + astString(x.Elt)
	default:
		fset := token.NewFileSet()
		var buf strings.Builder
		if err := printer.Fprint(&buf, fset, n); err == nil {
			return buf.String()
		}
		return fmt.Sprintf("(%T)", n)
	}
}
