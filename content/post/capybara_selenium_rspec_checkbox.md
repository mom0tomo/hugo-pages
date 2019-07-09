---
title: "Capybaraとchromedriverを使ってテストをしたらcheckboxの要素が探せなかった"
date: 2019-07-10T00:13:50+09:00
draft: false
tags: ["Rails"]
images: ["images/articles/avatar.png"]
description: "Capybaraを使って`js: true`にしたテストを書いているときに、checkboxがチェックできずにハマりました。Selenium上でchromdrverを使ったテストでは、Capybaraの「check」を使ってラベルのテキストで要素を探すことはできないようで、ElementNotFoundになってしまいます。これを避けるための書き方を紹介します。"
---

Capybaraを使ってテストを書いているときに、system specでjs: trueにしたところ、checkboxをcheck/uncheckできなくて困りました。

```ruby
RSpec.describe 'hogehoge', type: :system, js: true do
  ...

  check 'ほげほげ' # ElementNotFoundになってしまう

  ...
end
```

Selenium上でchromedriverを使ったテストでは、Capybaraの「check」を使ってラベルのテキストで要素を探すことはできないようです。

これを避けるためにCapybaraのfindとclickを使います。

```ruby
find('label', text: 'ほげほげ', match: :first).click
```
- [#find (*args, **options, &optional_filter_block) ⇒ Capybara::Node::Element](https://www.rubydoc.info/github/jnicklas/capybara/Capybara/Node/Finders#find-instance_method)
- [#click (keys = [], **options) ⇒ Object](https://www.rubydoc.info/github/jnicklas/capybara/Capybara/Driver/Node#click-instance_method)

<div class="iframely-embed"><div class="iframely-responsive" style="height: 140px; padding-bottom: 0;"><a href="https://stackoverflow.com/questions/10613354/how-do-i-click-this-button-in-capybara" data-iframely-url="//cdn.iframe.ly/RVIA3O0"></a></div></div><script async src="//cdn.iframe.ly/embed.js" charset="utf-8"></script>
