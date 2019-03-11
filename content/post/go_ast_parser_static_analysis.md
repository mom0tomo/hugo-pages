---
title: "Goの標準パッケージではじめる静的解析入門 ①準備編"
date: 2019-03-10T23:52:21+09:00
draft: false
tags: ["Go", "標準パッケージ"]
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
- 抽象構文木(AST)のトラバース/ノードの変更

## やらないこと
- 静的解析のモジュール化
- golang.org/x/tools/go/analysisパッケージの解説

実用の視点に立ったモジュール化の手法については、tenntennさんの記事をご覧下さい。

[Goにおける静的解析のモジュール化について](https://tech.mercari.com/entry/2018/12/16/150000)

***

## 利用する標準パッケージ

### go/parser
[go/parser](https://godoc.org/go/parser)

構文解析を行うパッケージです。<br>
ソースコードを字句解析し、構文解析を行い、抽象構文木(AST)まで作ってくれます。

文字列をパースするために、`parser.ParseExpr`関数を使います。

```go
func ParseExpr(x string) (ast.Expr, error)
```
- [parser/ParseExpr](https://godoc.org/go/parser#ParseExpr)

### go/ast
[go/ast](https://godoc.org/go/ast)

抽象構文木(AST)を定義したパッケージです。

抽象構文木（AST）をトラバースするために、`ast.Inspect`関数を使います。

```go
func Inspect(node Node, f func(Node) bool)
```

- [ast/Inspect](https://godoc.org/go/ast#Inspect)

***

次回は実際に字句解析と構文解析を行い、抽象構文木(AST)をつくります。



