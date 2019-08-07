---
title: "Macとfish shellで mecab-golang を使う"
date: 2019-08-07T15:28:48+09:00
draft: false
tags: ["Go", "MeCab"]
images: ["images/articles/avatar.png"]
description: "オープンソース形態素解析エンジンMeCabのGo言語用ラッパー mecab-golangをMacにインストールしたときの手順をまとめました。fish shellを使っていたので環境変数の設定でちょっと詰まりました。"
---

## はじめに
ふと人工無能を作りたいと思い立って、オープンソース形態素解析エンジン [MeCab](https://taku910.github.io/mecab/)をインストールしてみました。

コードはGoで書きたかったので、他の方が使っているところを何度か見かけたGoのラッパー [mecab-golang](https://github.com/bluele/mecab-golang)を使うことにしました。

<div class="iframely-embed"><div class="iframely-responsive" style="height: 140px; padding-bottom: 0;"><a href="https://github.com/bluele/mecab-golang" data-iframely-url="//cdn.iframe.ly/5GNgBw4"></a></div></div><script async src="//cdn.iframe.ly/embed.js" charset="utf-8"></script>

Mac & fish shellで初期設定をするときにちょっと詰まったので手順をまとめます。

## MeCabのインストール
brewを使いました。

```
$ brew install mecab
```

MeCabだけインストールした状態で `mecab-golang`のExampleを実行しようとするとエラーになります。

```go
$ go mod init .
$ go run main.go

# github.com/bluele/mecab-golang
Undefined symbols for architecture x86_64:
```
これは後述のように環境変数を指定することで解消できます。

***

環境変数を設定して再度実行しようとすると、下記のエラーになりました。

```
panic: mecab_model is not created: param.cpp(69) [ifs] no such file or directory: /usr/local/lib/mecab/dic/ipadic/dicrc

goroutine 1 [running]:
main.main()
  /Users/200491/dev/bruno-bot/main.go:37 +0xd8
exit status 2
```

MeCabのインストールと併せて辞書をインストールする必要があります。

```
$ brew install mecab-ipadic
```

辞書をインストールすることでエラーが解消できました。

```go
$ go run main.go

すもも 名詞,一般,*,*,*,*,すもも,スモモ,スモモ
もも 名詞,一般,*,*,*,*,もも,モモ,モモ
もも 名詞,一般,*,*,*,*,もも,モモ,モモ
うち 名詞,非自立,副詞可能,*,*,*,うち,ウチ,ウチ
```

## mecab-golangの設定
環境変数を設定します。bashやzshならREADMEに書いてある通りに `.bashrc`などに設定するだけですが、fishを使っていると少し書き方が変わります。

```
$ set -x CGO_LDFLAGS '-L/usr/local/Cellar/mecab/0.996/lib -lmecab -lstdc++'
$ set -x CGO_CFLAGS '-I/usr/local/Cellar/mecab/0.996/include'
```

`config.fish`に設定しました。


またpathは下記のコマンドで確認しました。
```
$ mecab-config --libs
-L/usr/local/Cellar/mecab/0.996/lib -lmecab -lstdc++

$ mecab-config --inc-dir
/usr/local/Cellar/mecab/0.996/include
```

mecab-golangのREADMEには `"`mecab-config --libs`"`のようにコマンドを使って指定する方法も書いてありますが、私の場合fishでうまく動かなかったので、パスを直接指定しています。


## おわりに
MeCabと併せて辞書をインストールすること、環境変数の設定をすることを忘れなければ特に詰まることなく設定できます。


