---
title: "GAE / Cloud Functions + Datastore + Go で “rpc error: code = PermissionDenied desc = Missing or insufficient permissions.” エラー"
date: 2022-04-15T22:23:12+09:00
draft: false
tags: ["GCP"]
images: ["images/articles/avatar.png"]
description: GCPのプロジェクトIDをプログラムの中にハードコードしている場合、異なるプロジェクトIDをデプロイ対象に指定すると、デプロイ後にエラーになります。エラーの内容を見るとDatastoreの権限が不足しているように見えますが、実際はプロジェクトIDが間違っているだけの場合があります。
---
今回は特殊な要件で作業していたので、珍しいエラーを引きました。０から構築した場合は発生しないのでケースは多くないと思いますがメモとして残しておきます。


## 概要

GAEとDatastoreを利用したGoのサンプルプログラムがありました。これをバージョンアップ・修正しデプロイしようとしたところ、表題のエラーになりました。

```
$ gcloud app deploy --project {プロジェクトID} --version {バージョン} .
...
（デプロイは成功する）
```

デプロイしたプログラムの動作確認のためにブラウザからアクセスすると、下記のエラーメッセージが表示されます。

`rpc error: code = PermissionDenied desc = Missing or insufficient permissions.`

権限不足というエラーメッセージと、下記のstackoverflowのポストを見て、なにかしらの権限が不足しているのではないかと思い調べましたが、原因は違うところにありました。

https://stackoverflow.com/questions/61651865/getting-error-message-rpc-error-code-permissiondenied-desc-missing-or-ins


## 原因と対処法

サンプルプログラムの中にGCPのプロジェクトIDがハードコードされており、それとは別のプロジェクトIDを指定してデプロイしたためにエラーになっていました。ハードコードされた部分を修正し、今回のデプロイ対象のプロジェクトIDを指定することでエラーが解消されました。このエラーはGAEだけでなくCloud Functionsでも発生しました。

プログラムの中にプロジェクトIDをハードコードしていると、今回のような場合に気づくのが難しくなります。
このため、環境変数からプロジェクトIDを読みこむようにプログラムを修正するのが望ましいです。

GAEの場合、　`GOOGLE_CLOUD_PROJECT` という環境変数を使うことができます。`GOOGLE_CLOUD_PROJECT` には実行環境のGCP プロジェクトIDが入ります。

https://cloud.google.com/appengine/docs/standard/go/runtime#environment_variables

 `os.Getenv("GOOGLE_CLOUD_PROJECT")` で、動的にプロジェクトIDを取得するようにプログラムを修正しました。
