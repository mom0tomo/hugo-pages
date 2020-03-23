---
title: "Sequel Pro Nightly buildをbrewでインストールできない（IPv6関連のエラー）"
date: 2019-12-31T23:01:02+09:00
draft: false
tags: ["MySQL"]
images: ["images/articles/avatar.png"]
description: ""
---
## 公式の最新版ではMySQL8.0 を使えない
Sequel Pro 公式の最新版 v1.1.2 で MySQL8.0 が使えない。データベースに接続できなかったり、接続できるようになってもテーブルの情報が表示されなかったりする問題がある。

[Qiita | MySQL8.0.x に Sequel Pro で接続する](https://qiita.com/ucan-lab/items/b68db1db931c954da921)

この問題を解決するためにはナイトリービルド版をインストールする必要がある。Macでのナイトリービルド版のインストール方法はふたつある。

- 公式サイトからインストールする
https://sequelpro.com/test-builds

- Homebrewからインストールする
```
$ brew cask install homebrew/cask-versions/sequel-pro-nightly
```

brewでインストールする方法を試したところエラーになった。
```
$ brew cask install homebrew/cask-versions/sequel-pro-nightly

Updating Homebrew...
Error: Download failed on Cask 'sequel-pro-nightly' with message: Failed to open TCP connection to sequelpro.com:443 (Hostname not known: sequelpro.com)
```
関連するissueが上がっている。

[Error when installing `sequel-pro-nightly` #52773](https://github.com/Homebrew/homebrew-cask/issues/52773)

IPv6を介したダウンロードができないためこのエラーが発生するそうだ。RubyのURIライブラリでIPv6がサポートされていないことが原因とのこと。

FixしたというPRが2018/12/4にマージされているが（`require 'resolv-replace'`している）、2019年末現在もエラーが発生している。

[fixed ipv6 download issue sequel-pro-nightly #6626](https://github.com/Homebrew/homebrew-cask-versions/pull/6626)

IPv6のアドレスを指定する場合はブラケット `[]` でアドレスの文字列を囲む必要があり、こちらに対応してもらわないといけないようだ。

[ruby-doc.org | URI::Generic](https://ruby-doc.org/stdlib-2.6.1/libdoc/uri/rdoc/URI/Generic.html)

## MacでIPv6 を無効化する

とりあえずの対応として、Githubのissueに書いてあった方法を使ってMacでIPv6 を無効化した。

```
$ networksetup -setv6off Wi-Fi
```

有効化する場合のオプションはちょっとトリッキー。`on` ではなく `automatic` 。
```
$ networksetup -setv6automatic Wi-Fi
```

```
$ brew cask audit sequel-pro-nightly

audit for sequel-pro-nightly: passed
```

これでとりあえずダウンロードできるようになる。