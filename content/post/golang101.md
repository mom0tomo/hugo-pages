---
title: "Golang 101 #1でTeaching Assistantをした"
date: 2019-07-20T21:45:57+09:00
draft: false
tags: ["Go", "event"]
images: ["images/articles/avatar.png"]
description: "@nasa9084氏が講師として開催したGolang101というイベントにTAとして参加してきました。ふだんWomenWhoGoTokyoのスタッフとして初心者向けの勉強会のスタッフをしていますが、それ以外のコミュニティでTAをするのは初めてだったので、勉強になりました。"
---
## はじめに
[@nasa9084](https://twitter.com/nasa9084)氏が講師として開催した[Golang101](https://golang101.connpass.com/event/135889/)というイベントの第1回に、TAとして参加してきました。

ふだんわたしはWomenWhoGoTokyoのスタッフとして初心者向けの勉強会で教えたり教わったりする立場ですが、それ以外のコミュニティでTAをするのは初めてでした。

<blockquote class="twitter-tweet" data-lang="ja"><p lang="ja" dir="ltr">スタバの看板があります <a href="https://twitter.com/hashtag/golang101?src=hash&amp;ref_src=twsrc%5Etfw">#golang101</a> <a href="https://t.co/nqdrV2H7DG">pic.twitter.com/nqdrV2H7DG</a></p>&mdash; nasa9084@某某某某(0x1a) (@nasa9084) <a href="https://twitter.com/nasa9084/status/1152416589146288128?ref_src=twsrc%5Etfw">2019年7月20日</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>


<blockquote class="twitter-tweet" data-lang="ja"><p lang="ja" dir="ltr">今日は久しぶりに自分以外男の人しかいない勉強会だ</p>&mdash; Momo Watanabe (@mom0tomo) <a href="https://twitter.com/mom0tomo/status/1152437873838137349?ref_src=twsrc%5Etfw">2019年7月20日</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

## 会の様子
nasa9084氏による力作！の資料をもとに、Go言語の構文等が基礎から丁寧に説明されました。

最初に書いたHello Worldを出力するプログラムのgo buildはみんな問題なく実行できて、開発環境構築で詰まった人もおらず順調な滑り出しでした。みんなちゃんと事前準備してきてえらい。


型推論ができる `:=`演算子がよく利用されることについて「型を明示的に書いた方がよいのではないか？」という質問がありましたが、この演算子を用いた宣言と代入を行えるのは関数内だけなので、何もしなくてもスコープは限定的になるので問題ないと説明してもらいました。確かにそうだ。

また講義前半の質問タイムでは、sliceについての質問が複数挙がりました。slice、言葉の説明だけで理解するのは難しいですよね...可変長の配列として使うことが多く、使ってみると意外と理解しやすいので、[the go playground](https://play.golang.org/)でいろいろ試してみるのがおすすめです。

ちなみにわたしも「sliceは参照渡しなので、sliceに新しく値を代入すると背後にある（Sliceの元になっている）配列の値も変わる」というのを理解するために、playgroundで実験しました。

https://play.golang.org/p/80k2ohNW7LT
```
package main

import "fmt"

func main() {
  var ns2 = []int{10, 20, 30, 40, 50}
  fmt.Println(ns2) // [10 20 30 40 50]

  ns3 := ns2[1:2]
  fmt.Println(ns3) // [20]
  
  ns3[0] = 100
  fmt.Println(ns3) // [100]
  
  fmt.Println(ns2) // [10 100 30 40 50]
}
```

資料に添付されているplaygroundリンク内のプログラムを書き換えて実行しても元の内容は破棄されません。そしてシェアするときには新たにリンクが発行されるので、気楽に使ってみるのがオススメです。

***

またポインタ型の話では、最初に理解しておくとつまずかないポイントとして「ポインタは定義するときも*を使用し、ポインタから値をとるときも*を使用する」ので混乱しがちだから気をつけようという解説がありました。

そしてGoのゼロ値のお話。

<blockquote class="twitter-tweet" data-lang="ja"><p lang="ja" dir="ltr">また出てきたゼロ値！コンポジット型（構造体とか）のゼロ値についてもここに書いてあるよ<a href="https://t.co/n2iOP1QM1G">https://t.co/n2iOP1QM1G</a><a href="https://twitter.com/hashtag/golang101?src=hash&amp;ref_src=twsrc%5Etfw">#golang101</a></p>&mdash; Momo Watanabe (@mom0tomo) <a href="https://twitter.com/mom0tomo/status/1152456089746591745?ref_src=twsrc%5Etfw">2019年7月20日</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

さらにオブジェクト指向の言語などではおなじみのclassがGoにはなくて、代わりに構造体を使うというお話。

このあたりから実際に自分で書いて振り返りしてみるとよさそうな内容になってきました。ポインタやslice、構造体などはthe go playgroundでいろいろ書き換えて実行してみると、一気に身近になると思います。

***

続くハンズオンでは、コマンドライン引数から値を取得してFizzBuzzしてみようという課題をみんなでやってみました。

コマンドライン引数の取得方法や、割り算して余りを出すための演算子、strconv.Atoiは戻り値が2つあること、switchはデフォルトでbreakすること、forの書き方など、初心者がハマるポイントがたくさん詰まっていてTAとして張り切ってヘルプできるとても良い課題でした。

早く終わった人は追加実装もしていました。

<blockquote class="twitter-tweet" data-lang="ja"><p lang="ja" dir="ltr">課題1のポイント：<br>- 割り算してあまりを返すには%を使う<br>- flag.Arg()を取得する前にはflag.Parse()するのを忘れずに<br>- strconv.Atoiは戻り値が二つある（文字列からコンバートした数値と、err）<br>- 戻り値が複数ある場合はnum, err := strconv.Atoiみたいに書ける<br>  <a href="https://twitter.com/hashtag/golang101?src=hash&amp;ref_src=twsrc%5Etfw">#golang101</a></p>&mdash; Momo Watanabe (@mom0tomo) <a href="https://twitter.com/mom0tomo/status/1152486283605762048?ref_src=twsrc%5Etfw">2019年7月20日</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

ほんとうはこの後もアプリケーションをつくってみるハンズオン課題があったのですが、やはり４時間では尺が足りない、ということで各自持ち帰ってやってみることになりました。

<blockquote class="twitter-tweet" data-lang="ja"><p lang="ja" dir="ltr">課題を解いてみて質問があったら<a href="https://twitter.com/nasa9084?ref_src=twsrc%5Etfw">@nasa9084</a> 氏にDM・メンションしてみよう！<a href="https://twitter.com/hashtag/golang101?src=hash&amp;ref_src=twsrc%5Etfw">#golang101</a></p>&mdash; Momo Watanabe (@mom0tomo) <a href="https://twitter.com/mom0tomo/status/1152494188442734592?ref_src=twsrc%5Etfw">2019年7月20日</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

## おわりに
すごく楽しかったです！

公務員からプログラマになるまで寄り道しながらよちよち歩いてきた自分が、イベントでTAをやることになるとは感慨深いです。

そして今回のイベントはnasa9084氏のていねいな説明のおかげで改めてGoの理解が進みました。すばらしい資料と課題をありがとうございました。

フィードバック次第で次回も開催してくれるかもしれないので、参加者のみなさんはぜひイベント後アンケートにもお答えくださいね！

