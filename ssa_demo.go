package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func main() {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo | packages.NeedImports,
		Dir:  ".", // 确保这是项目根目录
	}

	// 加载包和它的依赖
	pkgs, err := packages.Load(cfg, "test/im", "fmt")
	if err != nil {
		log.Fatalf("Failed to load packages: %s", err)
	}
	if len(pkgs) == 0 {
		log.Fatalf("No packages found.")
	}

	// 创建 SSA 程序
	ssaProg, ssaPkgs := ssautil.AllPackages(pkgs, ssa.BuilderMode(ssa.SanityCheckFunctions))
	ssaProg.Build()

	// 打印每个包中的函数
	for _, ssaPkg := range ssaPkgs {
		if ssaPkg != nil {
			ssaPkg.Build()
			for _, member := range ssaPkg.Members {
				if fn, ok := member.(*ssa.Function); ok && fn.Name() == "main" {
					fmt.Fprintln(os.Stdout, fn.String())
					fn.WriteTo(os.Stdout)
				}
			}
		}
	}
}
