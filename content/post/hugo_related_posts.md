---
title: "HugoでRelated Content(関連記事)を表示する"
date: 2019-01-21T20:18:22+09:00
draft: false
tags: ["Hugo"]
images: ["images/avatar.png"]
description: "HugoでRelated Content(関連記事)を表示する方法です。v0.27から公式にサポートされており、コードのサンプルもあるため簡単に設定できます。"
---

## はじめに
このブログはGo製のブログジェネレータHugoを使って運用しています。

ブログによくある「See Also」的なものが表示できないかと思って調べたところ、v0.27から公式にサポートされていました。

<div class="iframely-embed"><div class="iframely-responsive" style="padding-bottom: 50%; padding-top: 120px;"><a href="https://gohugo.io/content-management/related/" data-iframely-url="//cdn.iframe.ly/VUqdS9f"></a></div></div><script async src="//cdn.iframe.ly/embed.js" charset="utf-8"></script>

ここに書いてある通りに設定すれば、簡単に関連記事が表示できます。

<br>
こんな感じ↓
![related_content](/images/articles/related_content.png)
***

## Hugoで関連記事を表示する方法

公式サイトのとおりに、 `layouts/partials/related.html`を作成します。

```
{{ $related := .Site.RegularPages.Related . | first 5 }}
{{ with $related }}
<h3>See Also</h3>
<ul>
  {{ range . }}
  <li><a href="{{ .RelPermalink }}">{{ .Title }}</a></li>
  {{ end }}
</ul>
{{ end }}
```


`layouts/_default/single.html`で読み込みます（articleタグの下がよいと思います）。

```
<article>
...
</article>

{{ partial "related.html" . }}
```

config.tomlを使って細かい設定をすることも可能です（これは必須ではありません）。

以上の変更はこちらのcommitです。

<div class="iframely-embed"><div class="iframely-responsive" style="height: 168px; padding-bottom: 0;"><a href="https://github.com/mom0tomo/hugo-pages/commit/a114c195fb3d41e343b4e21fbd2849f7017bad49" data-iframely-url="//cdn.iframe.ly/aY0q6is"></a></div></div><script async src="//cdn.iframe.ly/embed.js" charset="utf-8"></script>

***

## 詰まったところ

buildに失敗して下記のようなエラーが出ました。

```
Error while rendering "page": template: _default/single.html:35:15:
executing "_default/single.html" at <partial "related.htm...>: 
error calling partial: template: partials/related.html:1:20:
executing "partials/related.html" at <.Site.RegularPages.R...>: 
can't evaluate field Related in type hugolib.Pages
```

原因は、CIで設定しているHugoのバージョンが古かったことです。

わたしがサイトを作ったのは2016年で、そのときからWerker CIの設定ファイルを一度もアップデートしておらず、本番環境のHugoが古いバージョンになっていました。

私が使っていたのはなんとv0.26！「Related Content」のサポートはv0.27からなので、惜しかったです。

単純なミスですが、30分ほどハマってしまいました。

## おわりに
HugoでRelated Content(関連記事)を表示する方法について書きました。

これからはよく使うツールのバージョンアップは定期的に行おうと思います。






