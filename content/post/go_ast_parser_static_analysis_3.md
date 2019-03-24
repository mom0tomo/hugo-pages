---
title: "Goの標準パッケージではじめる静的解析入門③実践編"
date: 2019-03-24T19:32:28+09:00
draft: false
tags: ["Go", "静的解析"]
images: ["images/articles/avatar.png"]
description: "Go言語では、標準パッケージであるgoパッケージが字句解析や構文解析を行う機能を提供しています。go/astやgo/parserを使って構文解析のはじめの一歩を踏み出してみます。今回は静的解析を使って実践的な処理を行ってみます。"
---

## はじめに
[前回](https://mom0tomo.github.io/post/go_ast_parser_static_analysis_2)は字句解析と構文解析を実際にやってみました。

今回は取得した抽象構文木(AST)を使って実践的な処理を行ってみます。

***

## ファイルを準備する
解析するためのサンプルファイル(example.go)と、解析処理を書くファイル(main.go)を準備します。

なお、処理を見やすくするためにエラーを潰しているところがあります。

```
$ mkdir example
$ touch example/example.go 

$ touch main.go
```

今回はexample.goの内容を以下のようにします。
```
package example

import (
  "fmt"
  "time"
)

func example() {
  fmt.Println("Now :", time.Now())
}
```

またmain.goには、いったん抽象構文木(AST)を取得する処理まで書いておきます。

```go
package main

import (
  "fmt"
  "go/parser"
  "go/token"
)

func main() {
  fset := token.NewFileSet()
  // ファイルをパースしてを抽象構文木(AST)に変換する
  f, _ := parser.ParseFile(fset, "./example/example.go", nil, parser.Mode(0)))
}
```
[parser.Mode](https://godoc.org/go/parser#example-ParseFile)というのはソースコードのパースしたい部分を指定したり、エラーを報告するかどうか選択するするためのものです。

必要な処理だけあらかじめ指定しておけば、パース処理が無駄なくなりますね。

下記の値が指定できます。
```
const (
    PackageClauseOnly Mode             = 1 << iota // stop parsing after package clause
    ImportsOnly                                    // stop parsing after import declarations
    ParseComments                                  // parse comments and add them to AST
    Trace                                          // print a trace of parsed productions
    DeclarationErrors                              // report declaration errors
    SpuriousErrors                                 // same as AllErrors, for backward-compatibility
    AllErrors         = SpuriousErrors             // report all errors (not just the first 10 on different lines)
)
```

今回のように、とりあえず何も指定せず全てパースしたいときは`parser.Mode(0))`またはただの 0としておきます。

## 抽象構文木(AST)を探索して結果を出力する
前回の構文解析で取得した抽象構文木(AST)を使い、ノードを探索して結果を出力してみます。

ファイル内のコメントだけ抜き出したり、関数名だけ調べたりできるのは便利でうれしいですよね！

```go
func main() {
  fset := token.NewFileSet()
  // package名とimportされているパッケージが取得できれば良いのでImportsOnlyモードにする
  f, _ := parser.ParseFile(fset, "./example/example.go", nil, parser.ImportsOnly)

  // package名を出力する
  fmt.Println(f.Name)

  // importされているパッケージを出力する
  for _, s := range f.Imports {
    fmt.Println(s.Path.Value)
  }
}
```

出力結果は以下のようになります。
```
"fmt"
"time"
```

ほかにもCommentなどいろいろ取得できます。

取得できるものの一覧は、[ast/File](https://godoc.org/go/ast#File)に書いてあります。
```
type File struct {
    Doc        *CommentGroup   // associated documentation; or nil
    Package    token.Pos       // position of "package" keyword
    Name       *Ident          // package name
    Decls      []Decl          // top-level declarations; or nil
    Scope      *Scope          // package scope (this file only)
    Imports    []*ImportSpec   // imports in this file
    Unresolved []*Ident        // unresolved identifiers in this file
    Comments   []*CommentGroup // list of all comments in the source file
}
```


***

## 抽象構文木(AST)のソースコード内の位置情報を取得する

ファイル名や行番号など、ソースコード内の位置情報を取得して使いたいことがあります。

先ほどの解析結果に行番号の情報を追加して出力してみます。

ソースコード内のノードの位置情報を取得するには、[token/Pos](https://godoc.org/go/token#Pos)を使います。

また今回はノードの位置情報だけでなくファイルとその行数なども出力したいので、より詳細な情報を出力できる[token/Position](https://godoc.org/go/token#Position)を使います。

`token.Position`構造体はファイル名、行番号、カラム位置の情報を持っています。

```
type Position struct {
    Filename string // filename, if any
    Offset   int    // offset, starting at 0
    Line     int    // line number, starting at 1
    Column   int    // column number, starting at 1 (byte count)
}
```

なお前回も補足しましたが、ノードの位置情報は`token.FileSet`を元に相対的に決まります。

```go
func main() {
  fset := token.NewFileSet()
  f, _ := parser.ParseFile(fset, "./example/example.go", nil, parser.ImportsOnly)

  // importされているパッケージのファイル名/行数/カラム数と、パッケージ名を出力する
  // token.Pos()はノードがソースコード上に占める位置を指す
  for _, s := range f.Imports {
    fmt.Println(fset.Position(s.Pos()), s.Path.Value)
  }
}
```

出力結果です。
```
./example/example.go:4:2 "fmt"
./example/example.go:5:2 "time"
```

ファイル名と行数などの情報を付与したことで、出力結果がよりわかりやすくなりました。

***

## 抽象構文木(AST)を再帰的にトラバースする

今度は、取得した抽象構文木(AST)のすべてのノードを探索（トラバース）して、再帰的に処理してみます。

トラバースには[ast/Inspect](https://godoc.org/go/ast#example-Inspect)を使います。`ast/Inspect`は抽象構文木（AST)のノードに対する（深さ優先）探索を行います。

```go
func main() {
  fset := token.NewFileSet()
  f, _ := parser.ParseFile(fset, "./example/example.go", nil, 0)

  ast.Inspect(f, func(n ast.Node) bool {
    var s string
    // 型によって処理を分岐する
    switch x := n.(type) {
    case *ast.BasicLit: // リテラル
      s = x.Value
    case *ast.Ident: // 識別子名
      s = x.Name 
    }
    if s != "" {
      fmt.Println(fset.Position(n.Pos()), s)
    }
    return true
  })
}
```

次のような結果が得られます。

```
./example/example.go:1:9 example
./example/example.go:4:2 "fmt"
./example/example.go:5:2 "time"
./example/example.go:8:6 example
./example/example.go:9:2 fmt
./example/example.go:9:6 Println
./example/example.go:9:14 "Now :"
./example/example.go:9:23 time
./example/example.go:9:28 Now
```

example.goのソースコードにおける抽象構文木(AST)のすべてのノードを探索し、リテラルと識別子（パッケージ、関数、フィールドや変数など）の名前のみを取得して、一つずつ位置情報とともに出力しています。

***

## おわりに
抽象構文木(AST)を使って、静的解析の実践的な処理を行ってみました。

<br>

なお、この記事は[golang.tokyo #22+Okayama.go/Sendai.go](https://golangtokyo.connpass.com/event/122263/)のイベントに参加して書きました。

イベント当日の講義資料はこちらです。

- [A Tour of Static Analysis | by tenntennさん](https://docs.google.com/presentation/d/13FcaQiFnUBk6xb1cNyNUmEIhqTXYMJAs3oUpxXP5_dM/edit#slide=id.g5475518c10_0_0)

ハンズオンの時間では、golang.tokyoのコードラボ上の教材を進めている方もたくさんいらっしゃいました。

これから静的解析を始めてみたい方は、こちらのコードラボから始めてみるのもおすすめです。

- [golang.tokyo コードラボ | 静的解析をはじめよう - Gopherをさがせ！](https://golangtokyo.github.io/codelab/find-gophers/)
- [型チェックでインターフェースを実装しているか調べよう](https://golangtokyo.github.io/codelab/first-step-type-check/)
