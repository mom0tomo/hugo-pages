---
title: "CapybaraでJSのテストをしようとしたらonClickのhref属性でハマった"
date: 2019-02-10T23:16:50+09:00
draft: false
tags: ["Rails", "React"]
images: ["images/articles/avatar.png"]
description: "Capybaraを使ってReact JSのテストを書いているときに、onClickで生成されるHTMLのaタグの挙動でハマりました。aタグにhrefがないとCapybaraはリンクとみなさないそうです。かといって空のhrefをつけてしまうとクリックイベントではなく画面のリフレッシュをしてしまいます。これを避けるためにテストの書き方を変えます"
---
## はじめに

わたしの所属先のサービスはRails製で、フロントエンドの一部でReactを使っています。

Capybaraを使ってReact JSのテストを書いているときに、onClickで生成されるHTMLのaタグの挙動でハマりました。

Capybaraの`Unable to find visible link "hoge"` のエラーを避けるために、 `click_link`を使わず `find('a', text: 'hoge').click`を使うようにテストを修正しました。

***

## 何が起こったか

Reactとreact-modal（モーダル用ライブラリ。ここでは詳しく紹介しません）をつかってこんなJavaScriptのコードを書きました。

モーダルで選択肢を出し、クリックしたらモーダルを消します。

```
this.state = {
  isUnderagen: true,
}

...

 <Button
  text="20歳未満"
  onClick={() => {
    this.state.isUnderage('true')
  }}
/>

<Button
  text="20歳以上"
  onClick={() => {
    this.state.isUnderage('false')
  }}
/>

const Button = (props: Props) => (
  <div>
    <a onClick={props.onClick}>
      {props.text}
    </a>
  </div>
)
```

そして、Rspecでテストをこのように書きました。
```
click_link '20歳以上'
expect(page).to have_content '20歳以上です'
```

一見問題がなさそうですが、これはエラーになります。
```
Unable to find visible link "20歳以上"
```

「20歳以上」というリンクは生成されているはずなのに「見つからない」と言われています。

***

## aタグに対するCapybaraの見解

この挙動はCapybara側に問題があるのではないか？と思いましたが、仕様でした。

hrefのないaタグはリンクではないので、click_link/click_onの対象としないというのがCapybaraの見解です。

[click_link 'foo' not working on links without href ](https://github.com/teamcapybara/capybara/issues/379)

```
`a` tags without an `href` are not links, they are placeholders for links.
That's how the HTML spec defines it, that's how every modern browser treats them.
Capybara does indeed only click on links which have the `href` attribute,
and imho, that's sensible behaviour.
```

***

## onClickにhrefをつけるとどうなるか
では、React側の `onClick()`にhrefをつければ解決できるのでしょうか？

空のhrefをつけると、画面がリロードされてしまい、意図しない挙動になります。

```
const Button = (props: Props) => (
  <div>
    <a onClick={props.onClick}>
      href=''
      {props.text}
    </a>
  </div>
)

// => 画面がリロードされてしまい、クリックしてもモーダルを閉じることができない
```

そもそもhrefはaタグを使う際に必須とされる属性ではありません。
[<a>: アンカー要素](https://developer.mozilla.org/ja/docs/Web/HTML/Element/a])

そのため、このような飛び先のないクリックイベントの処理ではhrefの定義は不要です。

React公式のサンプルコードでも、必要がなければhref属性を定義していません。

[Handling Events](https://reactjs.org/docs/handling-events.html)

テストにあわせて実装を変えるのはよくないので、テストを修正して対処するようにします。

***

## 回避方法
エラーの回避方法として、 `click_link`をやめて`find('a', text: 'hoge').click` を使います。

テストを直します。

```
find('a', text: '20歳未満').click
expect(page).to have_content '20歳以上です'
```

これで、 `Unable to find visible link`のエラーは回避できました。

***

## まとめ
* aタグにhrefがないとCapybaraはリンクとみなさない。
* かといって空のhrefをつけてしまうと画面のリフレッシュをしてしまう。
* `click_link`を使わず `find('a', text: 'hoge').click`を使う。

