---
title: "GAE上で動くアプリケーションをGo1.12にアップデートする上でやったこと"
date: 2019-08-22T13:34:20+09:00
draft: false
tags: ["Go", "GAE"]
images: ["images/articles/avatar.png"]
description: "Go1.11から依存管理にgo modulesが使えるようになりました。Women Who Go TokyoのハンズオンでつくったSlackbotに利用しているライブラリのバージョン管理をgo modulesを使って行いたいと思ったため、アプリケーションのGoのバージョンを1.9から1.12に一気に上げたときの記録です。"

---
## はじめに
Go1.11から依存管理に [go modules](https://blog.golang.org/using-go-modules) が使えるようになりました。

Women Who Go Tokyo のハンズオンでつくった Slackbot に利用しているライブラリのバージョン管理を go modules を使って行いたいと思ったため、アプリケーションの Go のバージョンを1.9から1.12に一気に上げました。

***

## App Engine SDKの利用をやめる

Go1.9の時点では Cloud SDK に加え App Engine SDK をダウンロードする必要がありました。

Go 1.12では App Engine SDK は 廃止され、Cloud SDKとそのコンポーネントを使う形に変わりました。

https://cloud.google.com/appengine/docs/standard/go/download


Cloud SDK はここからダウンロードできます。
https://cloud.google.com/appengine/docs/standard/go112/quickstart

今回の場合 Cloud SDK はすでにダウンロードしているのでアップデートをします。
```
$ gcloud components update
```
さらにGoのApp Engine拡張機能が含まれている gcloud コンポーネントをインストールします。
```
$ gcloud components install app-engine-go
```

また、不要になった App Engine SDK（go_appengineディレクトリ）を併せて削除しました。

App Engine SDK を脱却し今後は Cloud SDK を利用するので、デプロイのコマンドも下記のように変更します。


（旧）App Engine SDKを利用していたときのデプロイのコマンド
```
$ goapp deploy --application <ROJECT_ID> --version <VERSION_ID> .
```

（新）Cloud SDKを利用したデプロイのコマンド
```
$ gcloud app deploy --project <PROJECT_ID> --version <VERSION_ID>
```

## app.yamlを書き換える

app.yamlを変更します。

```diff
- runtime: go
- api_version: go1.9
- module: slackbot
+ runtime: go112
+ service: slackbot

handlers:
- url: /.*
-  script: _go_app
+  script: auto
```

module は名称が service に変更されました。

またscriptは省略可能で、記載する場合は Go1.9 では `_go_app` とする必要がありましたが、 Go1.12 では指定できるのは `auto`のみとなっています。

[Go1.9 のドキュメント(日本語)](https://cloud.google.com/appengine/docs/standard/go/config/appref?hl=ja)<br>
[Go1.12 のドキュメント(英語のみ)](https://cloud.google.com/appengine/docs/standard/go112/config/appref?hl=ja)

## コードを書き換える
コード上は下記のような変更がありました。

- init 関数ではなくmain 関数を使う。
- `appengine.NewContext()` と `slack.SetHTTPClient()`の代わりに、標準パッケージ `net/http` の `http.ListenAndServe()`を使う。
- `google.golang.org/appengine/log` の代わりに標準パッケージの `log` を使う。
  - `log.Errorf()` や `log.Infof()` など `google.golang.org/appengine` の関数は使えなくなる
  - 標準パッケージの `log.Printf()` を使う

```diff
package slackbot

import (
...
+  "log"
...

-  "google.golang.org/appengine"
-  "google.golang.org/appengine/log"
-  "google.golang.org/appengine/urlfetch"
)

- func init() {
+ func main() {
  http.HandleFunc("/events", eventsHandler)
+  http.ListenAndServe(":8080", nil)
}
...

-  ctx := appengine.NewContext(r)
-  slack.SetHTTPClient(urlfetch.Client(ctx))
...
- log.Errorf(ctx, "ParseEvent: %v", err)
+ log.Printf("ParseEvent: %v", err)
```

## おわりに
ここまで書いてきたコード上の変更は以下の Pull Request にまとまっています。

ハンズオンの教材をつくってくださり、 Women Who Go Tokyo で講師をしていただいている @tenntennさん、 @knsh14さん、レビューしていただいた @sinmetalさん、 @sonatardさんありがとうございます。

<div class="iframely-embed"><div class="iframely-responsive" style="height: 140px; padding-bottom: 0;"><a href="https://github.com/tenntenn/gohandson/pull/33" data-iframely-url="//cdn.iframe.ly/thL9KMz"></a></div></div><script async src="//cdn.iframe.ly/embed.js" charset="utf-8"></script>

