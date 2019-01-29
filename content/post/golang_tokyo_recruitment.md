---
title: "golang.tokyo#21「Goエンジニアの採用」レポート"
date: 2019-01-28T19:35:51+09:00
draft: false
tags: ["Go", "event"]
images: ["images/articles/golanggopher.png"]
description: "golang.tokyoも21回目！今回のgolang.tokyoは「Goエンジニアの採用」がテーマでした。"
---
[golang.tokyo](https://twitter.com/hashtag/golangtokyo?src=hash)も21回目！

今回のgolang.tokyoは「Goエンジニアの採用」がテーマでした。

![golang.tokyo](/images/articles/golangtokyo.jpg)

公開が確認できたスライドについてはタイトルにリンクを貼っています。

***

## Talk 1 社内ベンチャーで採用からチームビルディングで注力していること
最初の発表は、パーソルプロセス&テクノロジー株式会社 seedscompany というシード事業部門でエンジニアの採用を手がけている@ntk1000さん。

求職者側に気をつけてほしいことや、採用時に気をつけていること、採用活動の見直し過程やチームビルディングの手法についてお話しいただきました。

Gopherに限らず、一般的なエンジニアの採用に役に立つ情報でした。

内容とは外れますが、全編シンプルな英語で書かれた大変美しくおしゃれなスライドでした！


## Talk 2 [34人のGoの課題をレビューして感じたこと](https://docs.google.com/presentation/d/1TAwxT9mRmiEjQOZurz-TbXc8SQf0mlEJDS8tL49Q-M4/)
本来はgolang.tokyo募集時に恒例となったDevQuizについて@tenntennさんにお話いただくはずでしたが、体調不良のためご欠席。

準備していた資料は後ほどQiita等に公開予定だそうです。

急遽、golang.tokyo運営メンバーでfreee株式会社所属の@budougumi0617さんにお話いただきました。


今回のDevQuiz枠についてのざっくりとした講評は以下のようなものでした。

総計34人のうち、

- `go fmt`がかかっていなかった 8人
- `golint`がかかっていなかった 8人
- `go.mod`が入っていた（依存関係を考慮していた） 7人
- `_test.go`が入っていた 10人
- 何らかの `flag`を実装していた 24人
- `Makefile`が入っていた 3人
- 実行結果のGIFが入っていた 1人
- `filepath.Walk`使っていた 4人

提出フォーマットが指定した形式と異なっていたり、Macのバイナリやエディタの設定ファイルなど不要なものがはいっているなど、コード以前にレビュアーに苦労をかけるところは事前のチェックで回避できるので気をつけてほしいとのこと（身につまされます）。

提出前のセルフレビューで使える補助ツール(go fmt, go vet, golint, gometalinter, golangci-lint…)や、[コードレビューコメントの日本語訳](https://qiita.com/knsh14/items/8b73b31822c109d4c497)（Goの模範的なコードの書き方集）、[Effective Go](https://golang.org/doc/effective_go.html)(Goコードを明快に書くためのTips)などについて紹介いただきました。

tenntennさんいわく、 コードレビューでは **「lsしたときに受けた印象が変わることはほとんどない」** (ファイル構成で悪い印象を受けると、それがコード読んで良い印象に変わる事はほぼ無い) そうです。

至言ですね。

![tenntenn says](/images/articles/gopherslide.png)


多くの方がgolang.tokyoのタグとともにコードを公開していたので、ご興味のある方は#golangtokyo treeでTwitter検索してみると面白いと思います。

***

## LT1 [200人以上のGoの課題の1人目として受けた際に気をつけたこと](https://speakerdeck.com/sonatard/200ren-yi-shang-falsegofalseke-ti-false1ren-mu-tositeshou-ketaji-niqi-wotuketakoto)

D Technologyの@sonatardさん。

tenntennさんの元同僚で、なんとtenntennさんのつくった採用課題を解いた第一号だったそうです。

採用テストを受ける側として気をつけるべきポイントについてお話しいただきました。

一つ目の技術選定の観点では、フレームワークを採用するかどうかに際しての判断を書くこと、設計や仕様の決定にあたって自分の考えをきちんと表現することが大切とのこと。

二つ目のチーム開発の観点では、読み手を意識したコードを書くこと（テストは大事）、技術的バックグラウンドやGoの基本をどこまでわかっているかなどレビュアーの評価基準になるような情報を出すことが大切とのお話をしていただきました。

***

## LT2 [PHPエンジニアが転職してGoを書き始めたら楽しくなってきた話](https://speakerdeck.com/spre55/phpensiniakazhuan-zhi-sitegowoshu-kishi-metarale-sikunatutekitahua)

Cloud Aceの@spre55さん。

PHPerがGCPのパートナー企業に転職して、Goのコードに慣れながらその良さに気づいていったというお話でした。

GAEはお手軽でいいですよね！

動的型付け言語とは異なるところの多いGoですが、良いレビュアーについてもらうことで学べることが多いと思っています。

***

## LT3 [DockerのTUIツールを作った話](https://docs.google.com/presentation/d/1ty9MbOGT8HThYVC1DVZtmhKS5MyCNGkDdkCTs6vpqzk/)

@gorilla0513さん。

DockerのTUIツール[docui](https://github.com/skanehira/docui)についてご紹介いただきました。

はじめてデモを見せていただきましたが、視覚的に見やすく便利そうでした。

Goを学びながら動くものをつくってみようというのがモチベーションだったそうですが、gocuiにもDockerにも詳しくない中ひとりで開発を行った大変さが伝わってきました。

ひとりきりでの開発でモチベーションを維持しているのがシンブルにすごいと思いました。静かなアピール力もあり、熱い方ですね。


