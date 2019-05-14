---
title: "FaradayのURLのパスの書き方でハマった"
date: 2019-05-10T19:36:05+09:00
draft: false
tags: ["Ruby"]
images: ["images/articles/avatar.png"]
description: "Faradayでは、パスを書くときの決まりがありました。先頭にスラッシュを入れていると、FaradayによってURLは絶対パスとして扱われ、ドメイン部分のみがURLとされます。ドメイン以下のパスは破棄されてしまいます。"
---

## 起こったこと

[Faraday](https://github.com/lostisland/faraday)を使って立てたAPIサーバへアクセスするテストを、[webmock](https://github.com/bblimke/webmock)を使って行いました。

下記のような実装にしたところ、テストを実行してもwebmockが期待通り動作しません。

client.rb
```ruby
class Client
  def user
    conn = Faraday.new(url: 'http://example.com/api/v1')
    res = conn.get '/register'
    end
    res.status
  end
end
```

テスト内ではAPIエンドポイントとなるURL(`http://example.com/api/v1/register`)をスタブ化しているのですが、指定したステータスコードの200ではなく404が返ってきます。

client_spec.rb
```ruby
require 'rails_helper'
require 'webmock/rspec'

describe do
  before do
    # 外部API呼び出し処理に対し、スタブを登録する
    stub_request(:get, 'http://example.com/api/v1/register').to_return(status: 200)
  end

    it 'Status Code 200を返すこと' do
      response = Client.new.user
      expect(response).to eq 200
    end
end
```

検証したところ、パスの指定がおかしいという結論に至りました。

***

## 原因と対処法

Faradayでは、パスを書くときの決まりがありました。

client.rb
```ruby
class Client
  def user
    conn = Faraday.new(url: 'http://example.com/api/v1')
    # 先頭にスラッシュを入れてはいけない!
    res = conn.get 'register' #=> GET http://example.com/api/v1/register

    ...
```
Faradayを使ってGETする際に、先頭にスラッシュを入れない状態だと、URLは相対パスとして扱われます。設定したとおりです。

これが期待通りの挙動です。

<br>

一方、先頭にスラッシュを入れていると、FaradayによってURLは絶対パスとして扱われ、ドメイン部分のみ`https://example.com` がURLとされてしまいます。ドメイン以下のパス `/api/v1`は破棄されます。

client.rb
```ruby
class Client
  def user
    conn = Faraday.new(url: 'http://example.com/api/v1')
    # 先頭にスラッシュを入れると/api/v1が破棄されてしまう
    res = conn.get '/register' #=> GET http://example.com/register

```

`http://example.com/register`に対してGETしてしまっています。

この動作に関して、issueが上がっていました。

<div class="iframely-embed"><div class="iframely-responsive" style="height: 140px; padding-bottom: 0;"><a href="https://github.com/lostisland/faraday/issues/293" data-iframely-url="//cdn.iframe.ly/api/iframe?url=https%3A%2F%2Fgithub.com%2Flostisland%2Ffaraday%2Fissues%2F293&key=2621d5600f6d423389afec325dbfc63d"></a></div></div><script async src="//cdn.iframe.ly/embed.js" charset="utf-8"></script>

***

## なぜこの仕様なのだろう
一般にUNIXではルートを起点として絶対パスを指します。

先頭にスラッシュを入れることで、Faradayではルートからの絶対パスを指定したとみなします。

このためWebサイトのルートディレクトリであるドメイン部分を残し、それ以下のパスは破棄されるそうです。

なるほど...。

READMEのサンプルが先頭スラッシュで始まっているので、このような間違いを誘うのかもしれません。

```ruby
conn = Faraday.new(:url => 'http://www.example.com')
response = conn.get '/users' # GET http://www.example.com/users'
```

#### 追記
Faradayの面白い仕様（先頭スラッシュにするとバスを削除しちゃうやつ）、issueあげたらREADMEに挙動を追記してくれました！

<div class="iframely-embed"><div class="iframely-responsive" style="height: 140px; padding-bottom: 0;"><a href="https://github.com/lostisland/faraday/issues/976" data-iframely-url="//cdn.iframe.ly/TkSSFpA"></a></div></div><script async src="//cdn.iframe.ly/embed.js" charset="utf-8"></script>


README:
```ruby
conn = Faraday.new(:url => 'http://www.example.com/api')
response = conn.get 'users'                 # GET http://www.example.com/api/users'

# You can override the path from the connection initializer by using an absolute path
response = conn.get '/users'                # GET http://www.example.com/users'
```

これでもう間違えない。