---
title: "Goの標準パッケージではじめる静的解析入門①準備編"
date: 2019-03-10T23:52:21+09:00
draft: false
tags: ["Go", "静的解析"]
images: ["images/articles/gopher.png"]
description: "Go言語では、標準パッケージであるgoパッケージが字句解析や構文解析を行う機能を提供しています。go/astやgo/parserを使って構文解析のはじめの一歩を踏み出してみます。今回は用語や利用するパッケージの説明をして、静的解析をはじめる準備をします。"
---
## はじめに
Go言語では、標準パッケージであるgoパッケージが字句解析や構文解析を行う機能を提供しています。

go/astやgo/parserを使って構文解析のはじめの一歩を踏み出してみます。

### 参考
- [Qiita | goパッケージで簡単に静的解析して世界を広げよう #golang](https://qiita.com/tenntenn/items/868704380455c5090d4b)<br>
静的解析を学ぶ際に参考になる記事のまとめ（tenntennさん筆）

***

## 用語説明
### 静的解析とは
静的コード解析、静的プログラム解析とも言います。<br>
ソースコードを _**実行することなく**_ 解析を行うことを指します。

対義語は動的解析で、ソースコードを実行して解析を行うことを表します。<br>
こちらはプログラマが普段行なっているような検証方法です。
<br>

### 抽象構文木(AST)とは
abstract syntax tree.<br>
構文解析の経過や結果から、言語の意味に関係ない情報を取り除き、意味に関係ある情報のみを取り出したツリー状（木）のデータ構造のことです。<br>
[Wikipedia | 抽象構文木](https://ja.wikipedia.org/wiki/%E6%8A%BD%E8%B1%A1%E6%A7%8B%E6%96%87%E6%9C%A8)


### トラバース(走査)とは
抽象構文木のような木構造の各ノードを探索、調査することです。<br>
ノードとは、下記の図でいうところのA,B,C...が書いてあるそれぞれの◯を指します。<br>
[Wikipedia | 木構造(データ構造)#走査法](https://ja.wikipedia.org/wiki/%E6%9C%A8%E6%A7%8B%E9%80%A0_(%E3%83%87%E3%83%BC%E3%82%BF%E6%A7%8B%E9%80%A0)#%E8%B5%B0%E6%9F%BB%E6%B3%95)

![木構造](/images/articles/tree.png)

<br>

## やること
- 字句解析/構文解析
- 抽象構文木(AST)のトラバース
- 位置情報の取得
- 型情報の取得

静的解析を行うための機能が提供されている、[goパッケージ配下のパッケージ](https://golang.org/pkg/go/)を使います。

## やらないこと
- 静的解析のモジュール化
- golang.org/x/tools/go/analysisパッケージの解説

実用の視点に立ったモジュール化の手法については、tenntennさんの記事をご覧下さい。

[Goにおける静的解析のモジュール化について](https://tech.mercari.com/entry/2018/12/16/150000)

***

## 字句解析で利用するパッケージ

### [go/scanner](https://godoc.org/go/scanner)

字句解析を行う際に使うパッケージです。

`scanner/Scan`関数を使い、ソースコードを読み込み、トークンの位置、トークン、srcで与えた値の文字列リテラルを取得します。

[scanner/Scan](https://godoc.org/go/scanner#Scanner.Scan)
```go
func (s *Scanner) Scan() (pos token.Pos, tok token.Token, lit string)
```

<br>

### [go/token](https://golang.org/pkg/go/token/)

字句解析を行う際に使うパッケージです。

ファイルの位置や長さに関する情報を保持する構造体(token.FileSet型)です。

[token/FileSet](https://godoc.org/go/token#FileSet)
```go
type FileSet struct {
    // contains filtered or unexported fields
}
```

`scanner/Scan`などを使って取得するトークンの位置情報は絶対的なものではなく、`token.FileSet`を基準として相対的に決められます。

<br>

## 構文解析で利用するパッケージ

### [go/parser](https://godoc.org/go/parser)

構文解析を行う際に使うパッケージです。

ソースコードを字句解析し、構文解析を行い、抽象構文木(AST)まで作ってくれます。

文字列をパースするために、`parser.ParseExpr`関数を使います。

[parser/ParseExpr](https://godoc.org/go/parser#ParseExpr)

```go
func ParseExpr(x string) (ast.Expr, error)
```

<br>

### [go/ast](https://godoc.org/go/ast)

抽象構文木(AST)を定義したパッケージです。

抽象構文木（AST）をトラバースするために、`ast.Inspect`関数を使います。

[ast/Inspect](https://godoc.org/go/ast#Inspect)
```go
func Inspect(node Node, f func(Node) bool)
```

***

## おわりに
今回は用語や利用する標準パッケージについて説明しました。

[次回](https://mom0tomo.github.io/post/go_ast_parser_static_analysis_2)は静的解析を行う際の流れを確認します。



