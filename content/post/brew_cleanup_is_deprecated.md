---
title: "brew 2.0.0からcleanupが自動で走るようになっていた"
date: 2019-02-17T17:00:05+09:00
draft: false
tags: ["homebrew"]
images: ["images/articles/homebrew.png"]
description: "brew 2.0.0から、brew install、brew update、brew reinstallを実行したときにcleanupが自動で走るようになりました。また、何もしなくても30日ごとに定期的に実行されます。これを防ぎたい場合はHOMEBREW_NO_INSTALL_CLEANUPという環境変数を設定します。"
---

2019/2/2のhomebrew 2.0.0リリースノートに書いてありました。

- 30日ごとに自動で `brew cleanup`が走る
- `brew install` / `brew update` / `brew reinstall` を実行すると併せて走る
- これが嫌なときは `HOMEBREW_NO_INSTALL_CLEANUP`という変数を使う

```
brew cleanup is run periodically (every 30 days) and triggers
for individual formula cleanup on reinstall, install or upgrade.

You can opt-out of this behaviour by setting the HOMEBREW_NO_INSTALL_CLEANUP variable.

This addresses a long-standing complaint where users were surprised
by how much disk space Homebrew used if they did not run brew cleanup.
```

[homebrew-2.0.0](https://brew.sh/2019/02/02/homebrew-2.0.0/)

***

brewを使い始めたころから `alias bbb='brew update && brew upgrade && brew cleanup'`というエイリアスの設定をしてきました。

昨秋 `brew upgrade --cleanup`というオプションを見つけて便利だと思って使いはじめたところ、年明けくらいには非推奨になり、この度のアップデートで定期的なcleanupの設定が入ったという次第です。

もうエイリアスもオプションも必要ありません！便利になりましたね。

