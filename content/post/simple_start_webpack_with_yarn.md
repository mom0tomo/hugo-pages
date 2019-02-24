---
title: "yarnでシンプルにはじめるwebpack"
date: 2019-02-24T00:23:00+09:00
draft: false
tags: ["Webpack"]
images: ["images/articles/avatar.png"]
description: "モジュールバンドラwebpack入門です。yarnを使ったパッケージのインストール方法、webpackコマンドを使ったビルド方法、watch機能について、できるだけシンプルに書いています。"
---

Reactのモジュールバンドラ（複数のモジュールファイルを一つにまとめて出力するもの）としてwebpackを使っています。

Reactには`create-react-app`という便利な公式のコマンドラインツールがあり、これを使うことでwebpackの設定に触れることなく簡単に練習環境が作れます。

このため、今回はあえてツールに依存することなくシンプルにwebpackをはじめてみました。

***

この記事に書いてあって[公式の Getting Started](https://webpack.js.org/guides/getting-started/)には書いていないこと：

- yarnを使ったパッケージのインストール方法
- webpackコマンドを使ったビルド方法
- watch機能について

***

参考：

- [最新版で学ぶwebpack 4入門](https://ics.media/entry/12140)
- [Qiita: webpack 4 入門](https://qiita.com/soarflat/items/28bf799f7e0335b68186)

***

# 準備

サンプル用のディレクトリを作り、yarnでpackage.jsonを生成します。
```
$ mkdir webpack-example
$ cd webpack-example

$ yarn init -y
```

```
$ cat package.json

{
  "name": "webpack-example",
  "version": "1.0.0",
  "main": "index.js",
  "license": "MIT"
}
```

ソースコードを配置するディレクトリを作ります。

```
$ mkdir src
```

# バンドルしてみる

## webpackをインストールしてパスを通す

yarnでwebpackとwebpack-cliをローカルインストールします。

webpack4から、webpackコマンドを実行するためにwebpack-cliが必要になりました。

```
$ yarn add webpack webpack-cli --dev
```
```diff
$ git diff

-  "license": "MIT"
+  "license": "MIT",
+  "devDependencies": {
+    "webpack": "^4.29.5",
+    "webpack-cli": "^3.2.3"
+  }
```

`yarn.lock`が生成され、 `node_modules/`にソースコードがインストールされていることを確認します。

ローカルインストールしたwebpackのパスを通し、 `webpack`コマンドが使えるようにします。

`.bashrc`にパスの設定を書きます。
```
// ローカルインストールしたwebpackのパスを通す
export PATH=$PATH:./node_modules/.bin
```

sourceコマンドで設定を反映させます。
```
$ source ~/.bashrc
```

webpackコマンドを実行してパスが通ったかどうか試してみます。v4系になっていたらOKです。
```
$ webpack --version

4.29.5
```

***

## エントリーポイントとモジュールを作ってバンドルする
モジュールのファイル `hello.js`をsrc/配下につくります。今回はただ 'Hello'を出すだけのシンプルな関数を書きます。
```js
export function hello() {
  alert("Hello")
}
```

続いて、エントリーポイントとなるファイル `index.js`をsrc/配下につくります。
```js
// helloモジュールを読み込む
import { hello } from "./hello.js";

// hello関数を呼び出す
hello(); 
```

`webpack`コマンドを実行して先ほどのファイル2つをバンドル（ファイルをまとめること）し、一つのファイルに結果を出力します。

警告を防ぐために、いったん `--mode=development`というオプションをつけて実行します。
```
$ webpack --mode=development

Hash: 686abf64329ef2a53cc1
Version: webpack 4.29.5
Time: 76ms
Built at: 2019-02-23 20:58:53
  Asset      Size  Chunks             Chunk Names
main.js  4.53 KiB    main  [emitted]  main
Entrypoint main = main.js
[./src/hello.js] 52 bytes {main} [built]
[./src/index.js] 47 bytes {main} [built]
```

何も設定をしない場合、`dist/main.js`に出力されます。
```
$ ls -la dist/

total 16
drwxr-xr-x  3 mom0tomo  staff    96  2 23 20:57 .
drwxr-xr-x  9 mom0tomo  staff   288  2 23 20:57 ..
-rw-r--r--  1 mom0tomo  staff  4640  2 23 20:58 main.js
```

ここまでで、エントリーポイントのファイルとモジュールのファイルからバンドルファイルを出力するところまでできました。

<br>

#  webpack.config.jsの導入

## modeを設定する
webpack.config.jsファイルを利用することで、細かな設定ができます。

webpack.config.jsがなくてもモジュールをバンドルすることは可能ですが、実用の際は必ず使うので簡単な設定をしてみます。

```
$ touch webpack.config.js
```

まず、先ほどはコマンドのオプションとして指定していたmodeの設定を書きます。

modeにdevelopmentを指定することでソースマップ（バンドルして出力されたコードをデバッグする際に使う）が有効になります。

productionを指定することで、JavaSciptのコードからコメント等が取り除かれ圧縮されます。

ここではdevelopmentを指定します。

```js
module.exports = {
  mode: 'development'
}; 
```
これで、 `webpack`コマンドをオプションなしで実行しても警告が出なくなりました。

***

## watch機能を使う
ファイルの変更を検知して自動でバンドルのコマンドを実行し、ブラウザをリロードするwatchという機能があります。
この機能を使うことで、差分だけをビルドし、バンドルにかかる時間を短くすることができます。

`webpack --watch`というコマンドのオプションで実行できますが、webpack.config.jsに設定を書くことも可能です。

```diff
module.exports = {
-   mode: 'development'
+   mode: 'development',
+   watch: true
};
```

`webpack`コマンドを実行してみると、`webpack is watching the files…`というメッセージが出て変更がリアルタイムで反映されるようになったことがわかります。
```
$ webpack

webpack is watching the files…

Hash: 686abf64329ef2a53cc1
Version: webpack 4.29.5
Time: 94ms
Built at: 2019-02-23 21:44:56
  Asset      Size  Chunks             Chunk Names
main.js  4.53 KiB    main  [emitted]  main
Entrypoint main = main.js
[./src/hello.js] 52 bytes {main} [built]
[./src/index.js] 47 bytes {main} [built]
```

***

この他にもwebpack.config.jsを使うことで、エントリーポイントと出力先を変更したり、コードを遅延読み込みしたり、いろいろな設定が可能になります。

詳しくは[公式のドキュメント](https://webpack.js.org/concepts)を。
