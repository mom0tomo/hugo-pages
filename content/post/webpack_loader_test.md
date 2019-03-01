---
title: "webpack.config.jsでloader対象を指定するプロパティ名が test なのはなぜ？"
date: 2019-03-01T00:08:28+09:00
draft: false
tags: ["Webpack"]
images: ["images/articles/avatar.png"]
description: "webpack.config.jsで、loaderの対象とするファイル等を指定するプロパティの名前がtestなのがとても気になるのですが、これはJavaScriptの正規表現オブジェクトのtestメソッドに由来するという説があるようです。真相はわかりません。"
---
webpack.config.jsで、loaderの対象とするファイル等を指定するプロパティの名前が"test"なのが気になって気になってしょうがないです。

```
module: {
 loaders: [{
  test: /\.js$/, // 指定した正規表現にマッチする場合のみ、ローダーを適用する。
  loader: 'babel'
 }]
```

なんでtestなんだ...と思って色々調べたところ、どうやらJavaScriptでは正規表現オブジェクトに `test()`というメソッドがあるらしいことを知りました。

[Github issue | What is a "test" -- why is named "test"?](https://github.com/webpack/webpack/issues/866)

[MDN web Docs | RegExp.prototype.test()](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/RegExp/test)

`test()` メソッドを使うことで、指定した文字列が正規表現に一致するか調べることができるそうです。

***

ローダーの対象を指定するために、正規表現で対象となるファイルを表現するので、 `test()`メソッドを連想する...というのが理由なのでしょうか？

それにしてもこの命名はわかりにくい気がします...includeとかだったらわかりやすくてとてもいいのですが...

_それって納得いくかァ～～～おい？オレはぜーんぜん納得いかねえ……_

もし理由を詳しくご存知の方がいたら、教えていただきたいです。