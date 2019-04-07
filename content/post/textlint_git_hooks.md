---
title: "ブログにtextlintを入れてgit commit時に走るようにした"
date: 2019-04-02T00:07:58+09:00
draft: false
tags: ["Git, textlint"]
images: ["images/articles/avatar.png"]
description: "静的サイトジェネレータhugoを使って書いているブログににtextlintを入れました。git hookを使うことで、git commitやgit push時に自動でリンターをかけて文章のチェックを行うことができます。"
---
## はじめに

静的サイトジェネレータhugoを使っているブログでtypoやスペルミスを繰り返してしまっていたので[textlint](https://github.com/textlint/textlint)を入れることにしました。

textlintは日本語の文章校正ツールで、Markdownなどテキストのリンターとして使います。

詳しい使い方についてはcommitterである@azuさんのブログが詳しいです。

- [textlintで日本語の文章をチェックする](https://efcl.info/2015/09/10/introduce-textlint/)
- [技術文書を書くためのtextlint校正ルールセット](https://efcl.info/2016/07/13/textlint-rule-preset-ja-technical-writing/)

***

## textlintの導入

textlintの初期設定はシンプルです。

```
$ npm init
$ npm install --save-dev textlint

$ ./node_modules/.bin/textlint --init
```

これで初期設定は完了です。

動作確認のためにtextlintコマンドを実行してバージョンを確認してみます。

```
$ textlint -v
v11.2.3
```

インストールされたことが確認できました。

<br>

技術文書向けのtextlintルールプリセットである[preset-ja-technical-writing](https://github.com/textlint-ja/textlint-rule-preset-ja-technical-writing)を導入します。

```shell
$ npm install textlint-rule-preset-ja-technical-writing --save-dev
```

`.textlintrc`にルールを追加します。

```shell
{
  "filters": {},
  "rules": {
    "preset-ja-technical-writing": true
  }
}
```

***

## git hookを利用してcommit時にリンターをかける
[git hooks](https://git-scm.com/book/ja/v1/Git-%E3%81%AE%E3%82%AB%E3%82%B9%E3%82%BF%E3%83%9E%E3%82%A4%E3%82%BA-Git-%E3%83%95%E3%83%83%E3%82%AF)とはgit操作をフックとして他のコマンドやタスクを実行する機能です。

今回はcommit時にtextlintをかけるようにしてみます。

`git init`したときに`.git/hooks/pre-commit.sample`というファイルが自動的に作られているので、それをコピーして`.git/hooks/pre-commit`というファイルをつくります。

```
$ cp .git/hooks/pre-commit.sample .git/hooks/pre-commit
```

pushするときにlintを実行したい場合は同様の方法で`.git/hooks/pre-push`をつくります。

git hooksのアクションはシェルスクリプトで書きます。
また`#!/usr/bin/env ruby`のように宣言することでrubyなど任意の言語で書くこともできます。

私の場合は以下のようなスクリプトを書きました。

```sh
#!/bin/sh
echo run textlint

git diff --diff-filter=ACM --cached --name-only \
| grep -E '\.md' \
| xargs -n 1 -I{} textlint {} --fix

echo finish textlint
```

試してみます。

```sh
$ git commmit -m "published 'ブログにtextlintを入れてgit commit時に走るようにした'"

run textlint
finish textlint
[master d876d00] textlint
 Date: Thu Apr 4 00:31:07 2019 +0900
 1 file changed, 72 insertions(+)
 create mode 100644 content/post/textlint_git_hooks.md
```

git commitのタイミングで対象のファイルを自動的にtextlintにかけることができました。
