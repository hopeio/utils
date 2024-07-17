package dbtoentity

import (
	"fmt"
	dbi "github.com/hopeio/utils/dao/database/toentity"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"gorm.io/gorm"
	"strings"
	"testing"
)

var db *gorm.DB

func TestDBToEntity(t *testing.T) {
	MysqlConvert(db, "entity.go")
}

func TestTableToEntity(t *testing.T) {

	MysqlConvertByTable(db, "sku_competition")
}

func TestAst(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "tmpl.go", dbi.Tmpl, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}
	var b strings.Builder
	err = format.Node(&b, fset, f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(b.String())
	b.Reset()
	ty := f.Decls[1].(*ast.GenDecl)
	node := f.Decls[1].(*ast.GenDecl).Specs[0].(*ast.TypeSpec)
	node.Name.Name = "A"
	fileds := node.Type.(*ast.StructType).Fields
	fileds.List = append(fileds.List, &ast.Field{
		Doc: nil,
		Names: []*ast.Ident{
			{
				Name: "D",
				Obj:  &ast.Object{Kind: ast.Var, Name: "D"},
			},
		},
		Type:    &ast.Ident{Name: "time.Time"},
		Tag:     &ast.BasicLit{Kind: token.STRING, Value: `json:"d"`},
		Comment: nil,
	})
	err = format.Node(&b, fset, ty)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(b.String())
}