---
title: "【2020年3月】新しいMacのセットアップと入れたアプリ"
date: 2020-04-01T13:41:06+09:00
draft: false
tags: ["Mac"]
images: ["images/articles/avatar.png"]
description: ""
---
## dotfiles
gitで管理しており、install.shでシンボリックリンクを貼るだけ。
```
$ git clone https://github.com/mom0tomo/dotfiles.git
```

***

## homebrew
なにはなくともhomebrew。
```
$ /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

***

## キーの変更
邪魔で使わないCaps LockをCtrlキーに変えたい。JISキーボードと同じ場所にCtrlキーをおきたい。

キャプチャのように設定したら変えられました。

![System Preferences](/images/articles/setup_new_mac_2020_03/preference1.png)
![Keyboard](/images/articles/setup_new_mac_2020_03/preference2.png)
![Modifier Keys...](/images/articles/setup_new_mac_2020_03/preference3.png)


## Screen Shotの場所を変更する
デスクトップがごちゃごちゃするのが嫌なので`~/sshot`という場所にスクリーンショットが収まるようにしています。
```
$ mkdir ~/sshot

$ defaults write com.apple.screencapture location ~/SS/;killall SystemUIServer
```

元の場所（デスクトップ）に戻すときは

```
$ defaults delete com.apple.screencapture location;killall SystemUIServer
```

***

## Go言語のダウンロード
いまやGOPATHの指定など難しいことはなくなり、ダウンロードボタンをポチッとして開くだけで完結するようになりました。
https://golang.org/dl/

***

## brewで入れるツール

### fish
2017年から、元同僚に勧められて[fish](https://fishshell.com/)を使っています。
```
$ brew install fish
```

パッケージマネージャー[fisherman](https://github.com/fisherman/fisherman/wiki/%E6%97%A5%E6%9C%AC%E8%AA%9E)も併せてインストール。
```
$ curl -Lo ~/.config/fish/functions/fisher.fish --create-dirs git.io/fisher
```

### git関係
git logの代わりに見やすい[tig](https://jonas.github.io/tig/)を使っています。
```
$ brew insatll tig
```

### ブログ用
このブログはGo製の静的サイトジェネレータ[hugo](https://gohugo.io/)を使っています。
```
$ brew inastall hugo
```

## デスクトップアプリ類
### Slack
[Slack](https://slack.com/intl/ja-jp/downloads/osx)

### Chrome
[Chrome](https://www.google.co.jp/chrome/browser/desktop/index.html?brand=CHBD&gclid=EAIaIQobChMI45X78s-m2AIVVR0rCh33TgUeEAAYASAAEgKUBfD_BwE)

Developerツールや拡張機能が便利でメインブラウザとして使っています。Safariは検証時しか使いません。

### iterm2
慣れ親しんだ画面[iterm2](https://www.iterm2.com/)です。

### Alfred3
[Alfread](https://www.alfredapp.com/)は未だランチャーと計算機くらいしか使いこなせていないです...

### JetBrains
Toolboxでまとめて管理したいので、まずJetBrains ToolboxをダウンロードしてToolBox上で各IDEをインストールします。
[Toolboxをダウンロード | Mac](https://www.jetbrains.com/ja-jp/toolbox-app/download/?_ga=2.43875813.1372113674.1585751570-1832300423.1585132437#section=mac)

どんなふうに便利かはこちらに書いてあります。Updateやインストールなど、各IDEの管理をToolboxアプリでまとめて管理できます。
<div class="iframely-embed"><div class="iframely-responsive" style="padding-bottom: 52.3333%; padding-top: 120px;"><a href="https://blog.jetbrains.com/jp/2018/10/09/1294" data-iframely-url="//cdn.iframe.ly/Rquz79B"></a></div></div><script async src="//cdn.iframe.ly/embed.js" charset="utf-8"></script>

### Sublime Text 3
主にテキスト系の編集（このブログなど）に使っています　[Download Sublime Text 3](https://www.sublimetext.com/3)。はじめて使ったエディタなので今でも好きで文章を書くときはいつも使っています。

### Notion
メモ書きを保存するのに使っています。仕事関連で打ったコマンドのメモ、大学の課題のメモや、プライベートで考えたことのメモなど。いまのところ無料枠で間に合っています。

[Desktop App](https://www.notion.so/)

### Grammarly
英語の文法やスペリングのミス、不自然な表現などを直してくれるアプリです。

とても使い勝手が良く、英語で文章を書くときには欠かせません。12,000円くらいの年間プランを使っています。月々払いより少し安いです。

https://www.grammarly.com/

[Grammarly for macOS | Chrome拡張](https://app.grammarly.com/apps)
