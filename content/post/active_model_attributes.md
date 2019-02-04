---
title: "ActiveModelにおけるattributeの扱い"
date: 2019-02-04T22:13:32+09:00
draft: false
tags: ["Ruby", "Rails"]
images: ["images/articles/rails_guides.png"]
description: "ActiveRecordの代わりにActiveModelを使うモデル設計のパターンがあります。ActiveModelはActiveRecordと違い、attributeのような機能を持っていないので、attr_readerの外でインスタンス変数を定義して対処する必要があります。"
---

※ この記事はRails5.2から導入された`ActiveModel::Attributes`については記述していません。

## はじめに

よくあるフォームの実装です。

```
class UserForm < ActiveModel::Model
  WITH_KIDS_OPTIONS = [['こどもの同席あり', true], ['こどもの同席なし', false]].freeze

  attr_reader :name, email

  validates :name, presence: true
  validates :email, presence: true, email: true
  validates :with_kids, presence: true, inclusion: { in: [true, false] }

  def save
    user = User.new(
      name: name,
      email: email,
      with_kids: with_kids
    )
    user.save
  end

  def with_kids
    ActiveRecord::Type::Boolean.new.cast(@with_kids)
  end
end
```

ActiveModel::ModelをincludeしてFormオブジェクトを定義しています。

このコードを見たとき、`with_kids`メソッドを使って`@with_kids`というインスタンス変数をキャストしている部分が何をやっているのか引っかかったので、丁寧に読んでみました。

***

## attr_readerのはたらき
`attr_reader`は、Rubyのメソッドです。

内部的にインスタンス変数を定義し、クラスやモジュールにインスタンス変数を読み出すことを可能にします。

下記のコードは `attr_reader :name, email`と同じ挙動になります。

```
Class UserFrom
  def name
    @name
  end
  def email
    @email
  end
end
```

はじめに出てきたコードでは、`attr_reader`でwith_kids属性だけ指定していません。

その代わり、自分で`with_kids`メソッドを定義しています。

つまり、with_kids属性のみ、`attr_reader`を使わずに独自のメソッドを使ってインスタンス変数を定義しているのです。

### キャストを行う理由

では、なぜwith_kidsメソッドでキャストを行なっているのでしょうか？

キャストしている理由は、UserFormオブジェクトのwith_kids属性がString型に変換されてしまうのを防ぐことが目的です。

定数  `WITH_KIDS_OPTIONS`ではvalueをBoolean型で指定しており、それを守りたいのです。

この処理の背景には、ActiveModelの存在があります。

***

## ActiveModelについて
今回のコードでは、UserFormというクラスを定義しています。

これはRailsの[ActiveModel::Model](https://railsguides.jp/active_model_basics.html#model%E3%83%A2%E3%82%B8%E3%83%A5%E3%83%BC%E3%83%AB)をincludeしています。

このように、FormなどのオブジェクトをActiveModelを使って切り出し、[ActiveRecord](https://railsguides.jp/active_record_basics.html)を使ったモデルが大きくなりすぎないようにする、というリファクタリングのパターンがあります。

[7 Patterns to Refactor Fat ActiveRecord Models ](https://codeclimate.com/blog/7-ways-to-decompose-fat-activerecord-models)

翻訳版は[こちら](https://techracho.bpsinc.jp/hachi8833/2013_11_19/14738)

***

## ActiveRecordのattribute機能
ActiveRecordにあってActiveModelにないものが、`attribute`の機能です。

ActiveRecordでは、`attribute`機能を使って属性の型の指定が簡単にできます。

```
class User < ApplicationRecord
  attribute :with_kids, :boolean, default: false
end

user = User.new
user.with_kids == false  
```

上のコードでは、`with_kids`属性に対し型とデフォルトの値を指定しています。

ActiveModelの場合は自分でこれをやってあげなければなりません。

そこで、with_kidsメソッド内でキャストしているのです。

***

## まとめ

- Rubyのattr_readerは内部的にインスタンス変数を定義する
- ActiveRecordの代わりにActiveModelを使うモデル設計のパターンがある
- ActiveModelはActiveRecordと違い、attributeのような機能を持っていない（※）
- ActiveModelでattributeのようなことをしたいときは、attr_readerの外でインスタンス変数を定義して対処する（※）

<br>

※ 補足： Rails5.2から導入されたActiveModel::Attributesについて

Rails5.2からは`ActiveModel::Attributes`という大変便利なものが使えるようになりました。

これをincludeして使うことで、ActiveRecodeのattributeと同じような機能を使えるようになりました。

`ActiveModel::Attributes`の導入についての[わかりやすい記事はこちら](https://qiita.com/alpaca_taichou/items/bebace92f06af3f32898)です。
