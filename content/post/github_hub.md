---
title: "GitHubのhubを使ってコマンドラインからリポジトリを作る"
date: 2019-03-31T20:31:10+09:00
draft: false
tags: ["Github"]
images: ["images/articles/avatar.png"]
description: "Githubのコマンドラインツールhubを使うと、簡単にコマンドラインからリポジトリを作ることができます。"
---

Githubのコマンドラインツール、[hub](https://github.com/github/hub)がとても便利です。

勉強用のコードなど、ちょっと書いてからGitHubに保存することがよくあります。
いままではブラウザでGitHubを開いてリポジトリを作っていたのですが、hubを使うと簡単にコマンドラインからリポジトリを作ることができます。

hubのインストール方法や使い方はシンプルで、[README](https://github.com/github/hub)を読めばすぐに使うことができます。

***

## 使い方

brewでインストールします。

```shell
$ brew install hub

$ hub version
git version 1.7.6
hub version 2.2.3
```

最初に自分のGitHubアカウントとひもづける設定をします。

hub コマンドを叩くとuser nameとパスワードを聞かれるので、設定します。

```shell
$ hub

github.com username: mom0tomo
github.com password for mom0tomo (never stored):
```


## よく使うコマンド

`hub create`コマンドで作業中のディレクトリと同じ名前のリポジトリをGitHub上につくり、originとして指定します。

```shell
$ hub cretate

Updating origin
Warning: Permanently added the RSA host key for IP address '192.30.255.113' to the list of known hosts.
https://github.com/mom0tomo/sample-repo
```

`hub browse`コマンドで作業中のディレクトリと同じ名前のリポジトリをブラウザで開くことができます。

```shell
$  hub browse
```
