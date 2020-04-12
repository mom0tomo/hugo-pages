---
title: "MacOS CatalinaでPython2.7.13とansible2.8をつかう"
date: 2020-04-05T22:18:27+09:00
draft: false
tags: ["Python","Ansible"]
images: ["images/articles/avatar.png"]
description: "新しいPC上でEOLを迎えたPython2系との古いバージョンのAnsibleを操作する必要があったので備忘録にまとめました。"
---
新しいPC上でEOLを迎えたPython2系との古いバージョンのAnsibleを操作する必要があったので備忘録にまとめました。

## 環境

- MacOS Catalina 10.5.3
- Python2.7.13
- Ansible2.7<=>2.8


## Python2.7.13のインストール
pyenvでインストールします。

pyenv自体はbrewからインストールできます。公式のREADMEに詳しく書いてあるのでこれを読めばうまく設定できるはずです。

<div class="iframely-embed"><div class="iframely-responsive" style="height: 140px; padding-bottom: 0;"><a href="https://github.com/pyenv/pyenv" data-iframely-url="//cdn.iframe.ly/41GCIkX"></a></div></div><script async src="//cdn.iframe.ly/embed.js" charset="utf-8"></script>

PATHを通し、 `pyenv` コマンドを使えるようにします。
```
$ set -Ux PYENV_ROOT $HOME/.pyenv
$ set -Ux fish_user_paths $PYENV_ROOT/bin $fish_user_paths
```

`pyenv init` 設定をshellのconfigに追加します。わたしはfishを使っているので `~/.config/fish/config.fish` です。

```
$ echo -e 'if command -v pyenv 1>/dev/null 2>&1; then\n  pyenv init - | source \nfi' >> ~/.config/fish/config.fish
```

その他のshellでの設定はREADMEの `Basic GitHub Checkout` の章に書いてあります。

***
pyenvでPython2.7.13をインストールします。
```
$ pyenv install 2.7.13
```

インストールしたPythonがきちんと反映されているか `pyenv versions` コマンドで確認します。
```
$ pyenv versions
* system
  2.7.13
```
カレントディレクトリで使うPythonのバージョンを2.7.13に固定します。グローバルにバージョンを固定したい時は `local` を `global` に変えます。
```
$ pyenv local 2.7.13
pyenv versions
  system
* 2.7.13 (set by /Users/mom0tomo/dev/hoge/.python-version)
```

Pythonのバージョンを確認します。
```
$ python --version
Python 2.7.13
```

## Ansible2.7系のインストール
pipでインストールします。

まず、pipのPATHが通っていることを確認します。
```
$ which pip
/Users/mom0tomo/.pyenv/shims/pip
```

```
$ pip install --upgrade pip
$ pip install --upgrade 'ansible>=2.7.0,<2.8.0'
```

ansibleのコマンドが実行できることを確認します。

```
$ ansible --help

Usage: ansible <host-pattern> [options]

Define and run a single task 'playbook' against a set of hosts

Options:

...

`````
これで完了です。