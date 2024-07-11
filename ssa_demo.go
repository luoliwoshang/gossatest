package main

import (
	"fmt"
	"log"

	"golang.org/x/tools/go/packages"
)

func main() {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo,
		Dir:  ".", // 确保这是项目根目录
	}

	pkgs, err := packages.Load(cfg, "test/im")
	if err != nil {
		log.Fatalf("Failed to load packages: %s", err)
	}
	if len(pkgs) == 0 {
		log.Fatalf("No packages found.")
	}

	// 打印包中的所有Go文件
	for _, pkg := range pkgs {
		fmt.Printf("Package: %s, Files: %v\n", pkg.PkgPath, pkg.GoFiles)
	}
}
