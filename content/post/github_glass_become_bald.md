---
title: "リポジトリを削除するとGitHubの草は禿げる"
date: 2019-03-22T16:42:20+09:00
draft: false
tags: ["Github"]
images: ["images/articles/avatar.png"]
description: "フォルダ整理の感覚でいらなくなったGitHubからリポジトリを削除していたのですが、リポジトリを削除するとせっかく生やした草もまとめて禿げるということを先日知りました。
また、masterにマージしないとcommitしていても草は生えないということも最近知りました。"
---

フォルダ整理の感覚でいらなくなったリポジトリをGitHubから削除していたのですが、リポジトリを削除するとせっかく生やした草もまとめて禿げるということを先日知りました。

試してみた結果です。

before↓{{< figure src="/images/articles/github_glasses.png" width="50px" class="glass">}}

after↓{{< figure src="/images/articles/github_glasses_bald.png" width="50px" class="glass">}}

これを知ったときわりとショックだったので、いまはGitHubもSNSみたいなものだし...と思って古いリポジトリも削除しないようにしています。(古いツイートを消さないのと同じ感覚)

***

また、masterにマージしないとcommitは積んでいても草は生えないということも最近知りました。

プログラマなので一応毎日仕事でcommitは積んでいるのですが、ここ2週間ほど重ためのタスクをやっていて、未だmasterにマージできていないブランチがあります。
そのため、commitは積んでいても草が禿げている日があります。

この辺↓
{{< figure src="/images/articles/github_no_glasses.png" width="50px" class="glass">}}

こちらは、masterにマージするとそれぞれのcommitをした日まで遡って草が生えるそうです。


草の生える条件について、詳しくはこちらに書いてあります。

[GitHub | Viewing contributions on your profile - User Documentation](https://help.github.com/en/articles/viewing-contributions-on-your-profile)

***

GitHubは確かにSNSの一種です。

気になるユーザーのコードを見てフォローしたり、リポジトリにいいねをできる機能などはTwitterやInstagramと変わりません。

ただプログラマの自分としては、毎日使うコードレビューの機能はもう少し改善の余地があるのではないかと思います。

GitHubではコードレビュー時に気軽にいいねをつける機能がないので、コードレビューではコードの改善点や問題点を指摘することが中心になります。

「そのコードいいね！」とおもったところで、軽ーいインタラクションでいいねできるようなUXがあればいいなと思っています。

いいねされたことを知る必要がない人は、ミュートもできるとよさそう。
