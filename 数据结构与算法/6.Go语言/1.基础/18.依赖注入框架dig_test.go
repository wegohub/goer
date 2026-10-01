package base

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"testing"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 扫描目录中的所有Go文件并查找结构体及其注释
func scanStructsInDir(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			err := scanStructsInFile(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// 扫描单个Go文件并查找结构体及其注释
func scanStructsInFile(filename string) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	ast.Inspect(node, func(n ast.Node) bool {
		// 检查是否为结构体
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}
		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		// 获取结构体的注释
		var comments string
		if typeSpec.Doc != nil {
			for _, comment := range typeSpec.Doc.List {
				comments += comment.Text + "\n"
			}
		}

		fmt.Printf("Found struct: %s\n", typeSpec.Name.Name)
		fmt.Printf("Comments: %s\n", comments)

		// 遍历结构体的字段
		for _, field := range structType.Fields.List {
			fieldName := ""
			if len(field.Names) > 0 {
				fieldName = field.Names[0].Name
			}
			fieldType := field.Type
			var fieldComments string
			if field.Doc != nil {
				for _, comment := range field.Doc.List {
					fieldComments += comment.Text + "\n"
				}
			}
			fmt.Printf("  Field: %s %s\n", fieldName, fieldType)
			fmt.Printf("  Field Comments: %s\n", fieldComments)
		}
		return true
	})

	return nil
}

func Test_Inject(t *testing.T) {
	dirs := []string{"../a"}
	for _, dir := range dirs {
		err := scanStructsInDir(dir)
		if err != nil {
			fmt.Printf("Error scanning directory %s: %v\n", dir, err)
		}
	}
}
