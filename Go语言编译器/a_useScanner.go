package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	src := []byte("cos(x)+2i*sin(x) // Euler")
	// 初始化 scanner
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil /*there is no handler*/, scanner.ScanComments)
	// 扫描
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}
