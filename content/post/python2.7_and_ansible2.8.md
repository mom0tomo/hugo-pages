---
title: "Python2.7_and_ansible2"
date: 2020-04-05T22:18:27+09:00
draft: false
tags: ["Python","Ansible"]
images: ["images/articles/avatar.png"]
description: ""
---
新しいPC上で古いバージョンのAnsibleとPythonを操作する必要があったので備忘録にまとめました。

## 環境

- MacOS Catalina 10.5.3
- Python2.7.11
- Ansible2.8


## Python2.7系のインストール
pyenvでインストールしたいところですが、OSのバージョンが新しすぎてBUILD FAILEDになります。

ちなみに環境変数を指定している理由は、openssl関連でエラーになるポイントがあるので回避するためです。

<div class="iframely-embed"><div class="iframely-responsive" style="height: 140px; padding-bottom: 0;"><a href="https://github.com/pyenv/pyenv/issues/1184" data-iframely-url="//cdn.iframe.ly/XO0kylq"></a></div></div><script async src="//cdn.iframe.ly/embed.js" charset="utf-8"></script>

```
$ CONFIGURE_OPTS="--with-openssl='(brew --prefix openssl)'" pyenv install 2.7.11
ERROR: The Python ssl extension was not compiled. Missing the OpenSSL lib?

Please consult to the Wiki page to fix the problem.
https://github.com/pyenv/pyenv/wiki/Common-build-problems

...

BUILD FAILED (OS X 10.15.3 using python-build 20180424)

...

```


仕方がないので公式のダウンロードリンクからinstallerをダウンロードします。
https://www.python.org/downloads/

installerを実行して、終了したらpythonのバージョンを確認します。

```
$ python --version
Python 2.7.11
```

## Ansible2.8系のインストール
2020年4月時点ではbrewからインストールできます。

```
$ brew search ansible
==> Formulae
ansible                            ansible-cmdb                       ansible-lint                       ansible@2.8                        terraform-provisioner-ansible
==> Casks
homebrew/cask/ansible-dk

$ brew install ansible@2.8

If you need to have ansible@2.8 first in your PATH run:

echo 'set -g fish_user_paths "/usr/local/opt/ansible@2.8/bin" $fish_user_paths' >> ~/.config/fish/config.fish
```

shellでfishを使っている場合PATHの設定が必要だそうなので設定します。
```
$ echo 'set -g fish_user_paths "/usr/local/opt/ansible@2.8/bin" $fish_user_paths' >> ~/.config/fish/config.fish
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