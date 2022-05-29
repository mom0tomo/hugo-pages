---
title: >
  【fish shell】anyenvを入れた環境でbrew doctorするとWarning: "config" cripts exist outside your system...が出る件の対策
date: 2022-05-29T10:17:50+09:00
draft: false
tags: ["fish"]
images: ["images/articles/avatar.png"]
description: >
  anyenvを入れた環境でbrew doctorするとWarning: "config" cripts exist outside your system...というwarnが出るので簡単にエイリアスを設定して解消したいと思いました。しかし、fishでコマンドと同じ名前のエイリアスを設定する際は無限ループのエラーを回避するために、ちょっと工夫がいりました。
---

## 結論
`~/.config/fish/config.fish` の設定は下記のようにします。

```:shell
alias brew="PATH=/usr/local/bin:/usr/bin:/bin:/usr/local/sbin:/usr/sbin:/sbin command brew $argv"
```

## やりたかったこと
下記の記事と同じ状況で同じwarnが出たので、簡単にエイリアスを設定して解消したいと思いました（anyenvを入れた環境でbrew doctorするとWarning: "config" cripts exist outside your system...が出る）。

[【M1 Mac】anyenvを入れた環境でbrew doctorしたらWarning: "config" scripts exist outside your system...が出た時の対処法](https://zenn.dev/ryuu/scraps/fddefc2ca60f88)

しかし、fishでコマンドと同じ名前のエイリアスを設定するのはちょっと工夫がいりました。下記のように、コマンドと同じ名前のエイリアスを設定しようとすると、無限ループのエラーになります。

```:shell
# NGな例
alias brew="PATH=/usr/local/bin:/usr/bin:/bin:/usr/local/sbin:/usr/sbin:/sbin command brew $argv"
```
```:shell
fish: The function 'brew' calls itself immediately, which would result in an infinite loop.
PATH=/usr/local/bin:/usr/bin:/bin:/usr/local/sbin:/usr/sbin:/sbin brew
                                                                  ^
```

## 調べたこと
この理由は、fishでは `alias` コマンドは関数のシンプルなラッパーであるため、コマンドと同じ名前のエイリアスを設定しようとすると、関数内で関数（自分自身）を呼び出す動作をするためです。

[alias - create a function — fish-shell 3.4.1 documentation](https://fishshell.com/docs/current/cmds/alias.html)


[github.com issue | aliases don't work, fish complains of recursive function call · Issue #224 · fish-shell/fish-](https://github.com/fish-shell/fish-shell/issues/224)


つまりこうなっています。（このような関数はそのままfish shell上で実行して試すことができます）
```:shell
function hoge_command
  hoge_command &argv
end
```

fishでコマンドと同じ名前のエイリアスを設定したい場合は、 `command` をつけることで無限ループ問題が解決できます。 実行したいコマンド（今回はbrew）の前に `command` をつけることにより、関数ではなくコマンド（External command）を実行します。

[command - run a program — fish-shell 3.4.1 documentation](https://fishshell.com/docs/current/cmds/command.html#cmd-command)

つまりこのように書けばよいということでした。
```:shell
function hoge_command
  command hoge_command &argv
end
```

***

今回の場合PATHの設定も必要なので、PATHを設定した後に `command` をつけてbrewを実行するようにしました。また、短い処理なのでエイリアスとしてワンライナーにしました。

```:shell
alias brew="PATH=/usr/local/bin:/usr/bin:/bin:/usr/local/sbin:/usr/sbin:/sbin command brew $argv"
```
