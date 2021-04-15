---
title: "AWS EC2でSSHするポートを変更する"
date: 2021-04-15T19:23:12+09:00
draft: false
tags: ["AWS", "EC2"]
images: ["images/articles/avatar.png"]
description: SSHのデフォルトポートは22番です。セキュリティ対策のため、踏み台サーバーなどアクセス元IPによる制限を加えないサーバーを立てる場合、SSHポート番号をwell known portの22から変更したいときがあります。EC2の場合、セキュリティグループ上の変更も忘れずに行います。
---
SSHのデフォルトポートは22番です。セキュリティ対策のため、踏み台サーバーなどアクセス元IPによる制限を加えないサーバーを立てる場合、SSHポート番号をwell known portsの22番から変更したいときがあります。

今回はポートを2222に変更してみます。

## 設定ファイルを変更する
まず、EC2インスタンスにデフォルトポートでSSH接続します。接続方法とコマンドはAWSコンソールで見ることができるので省略します。

次に、設定ファイルを更新します。

```bash
# 設定ファイルのバックアップを取る
$ cp /etc/ssh/ssh_config /tmp

# 設定ファイルでポートを変更する
$ sudo vi /etc/ssh/ssh_config

# バックアップしたファイルと編集後のファイルを比較する
$ diff -u /tmp/ssh_config /etc/ssh/ssh_config
- Port 22
+ Port 2222
```

sshdを再起動して設定を反映します。

```bash
$ sudo systemctl reload sshd
```

サービス再起動後、起動状態を確認します。
```
$ sudo systemctl status sshd
● sshd.service - OpenSSH server daemon
   Loaded: loaded (/usr/lib/systemd/system/sshd.service; enabled; vendor preset: enabled)
   Active: active (running) since Sun 2021-04-15 19:48:43 UTC; 1 days ago
     Docs: man:sshd(8)
           man:sshd_config(5)
  Process: 3763 ExecReload=/bin/kill -HUP $MAINPID (code=exited, status=0/SUCCESS)
 Main PID: 3239 (sshd)
   CGroup: /system.slice/sshd.service
           └─3239 /usr/sbin/sshd -D
...
```

念のため、設定が完了し動作確認できるまでこのSSHコネクションは切断せずにおきます。もし設定にミスがあってサーバーに接続できなくなった場合、この接続から修正を行います。

## セキュリティグループを更新する
EC2インスタンスに接続して上記の変更を行った後、AWSコンソールから「セキュリティグループ」にアクセスして、「インバウンドルール」のポートを2222に変更します。

![ポートを22から2222に変更する](/images/articles/aws_ec2_ssh.png)

ここまで終わったら、ターミナルの新規ウィンドウを開き新しいセッションを張って、対象サーバーにポート2222でSSH接続できることを確認します。

