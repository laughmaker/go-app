package main

import (
	"fmt"

	"github.com/gohouse/converter"
)

func main() {
	generate("cat")
}

func generate(table string) {
	t2t := converter.NewTable2Struct()
	dsn := "root:123456@tcp(localhost:3306)/yirimao?charset=utf8mb4"
	path := fmt.Sprintf("../model/%s.go", table)
	err := t2t.Table(table).EnableJsonTag(true).PackageName("model").TagKey("gorm").SavePath(path).Dsn(dsn).Run()
	fmt.Println(err)
}
