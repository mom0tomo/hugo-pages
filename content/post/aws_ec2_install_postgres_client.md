---
title: "Amazon Linux2にPostgreSQLのクライアントモジュールをインストールする"
date: 2021-05-05T22:59:40+09:00
draft: false
tags: ["AWS"]
images: ["images/articles/avatar.png"]
description: EC2(Amazon Linux)上に建てた踏み台サーバーをクライアントとして、RDS(PostgreSQL)に接続したい。EC2をPostgreSQL DBとしてて使うのではなく、RDSに接続するクライアントとして使う場合には、クライアントモジュールをインストールするだけで十分です。
---
# 概要
EC2(Amazon Linux)上に建てた踏み台サーバーをクライアントとして、RDS(PostgreSQL)に接続したいというのが今回の主旨です。
PostgreSQLにはクライアントモジュールとサーバーモジュールの二種類があります。今回のようにEC2をPostgreSQL DBとしてて使うのではなく、RDSに接続するクライアントとして使う場合には、クライアントモジュールをインストールするだけで十分です。この記事では、PostgreSQLクライアントモジュールのみをインストールする手順を記載します。

# 環境
- EC2(接続元踏み台サーバー): Amazon Linux 2 AMI
- RDS(接続先DBサーバー）: Aurora PostgreSQL 11.7

# 手順
EC2に建てたサーバーにSSHで接続し、yumでpostgresのclientモジュールをインストールします。
```
# PostgreSQLのパッケージを確認する
[ec2-user@ip-xxx-xx-xx-xxx ~]$ yum search postgresql | grep client
postgresql.x86_64 : PostgreSQL client programs
                       : clients
postgresql-libs.i686 : The shared libraries required for any PostgreSQL clients
PyGreSQL.x86_64 : A Python client library for PostgreSQL
tcl-pgtcl.x86_64 : A Tcl client library for PostgreSQL
```
postgresql-libs.i686, postgresql.x86_64がPostgreSQLのクライアントプログラムとその依存するパッケージであることがわかります。

しかし、このまま`yum install`しようとすると、EC2ではデフォルトで参照しているPostgreSQLクライアントモジュールのリポジトリが古いため、指定したバージョンのモジュールが参照できずエラーになります。手動でEC2にリポジトリをダウンロードし、任意のバージョンのモジュールをダウンロードできるようにします。

Amazon Linux専用のリポジトリはないので、RPM(Red Hat Package Manager）用のリポジトリをダウンロードしてpgdg-redhat-all.repoを書き換えます。
なお、セキュリティグループでEC2上のサーバーへのアウトバウンドルールを絞っているとwgetする際にRPM接続できずにエラーになるので、HTTPS(Port: 443)を許可するように変更する必要があります。

```
# RPM用のリポジトリを/tmpにダウンロードする
[ec2-user@ip-xxx-xx-xx-xxx ~]$ wget -P /tmp https://download.postgresql.org/pub/repos/yum/reporpms/EL-7-x86_64/pgdg-redhat-repo-latest.noarch.rpm tmp/
[ec2-user@ip-xxx-xx-xx-xxx ~]$ sudo rpm -Uvh --nodeps /tmp/pgdg-redhat-repo-latest.noarch.rpm

# デフォルトのリポジトリを参照しないように設定を変更する
[ec2-user@ip-xxx-xx-xx-xxx ~]$ sudo sed --in-place -e "s/\$releasever/7/g" /etc/yum.repos.d/pgdg-redhat-all.repo
```

ここまでできたらyumでPostgreSQLクライアントモジュールをインストールします。
```
# クライアントモジュールをインストールする
[ec2-user@ip-xxx-xx-xx-xxx ~]$ sudo yum -y install postgresql-libs.i686 postgresql.x86_64
依存性関連をインストールしました:
  cyrus-sasl-lib.i686 0:2.1.26-23.amzn2               glibc.i686 0:2.26-44.amzn2                         keyutils-libs.i686 0:1.5.8-3.amzn2.0.2
  krb5-libs.i686 0:1.15.1-37.amzn2.2.2                libcom_err.i686 0:1.42.9-19.amzn2                  libcrypt.i686 0:2.26-44.amzn2
  libdb.i686 0:5.3.21-24.amzn2.0.3                    libgcc.i686 0:7.3.1-12.amzn2                       libselinux.i686 0:2.5-12.amzn2.0.2
  libsepol.i686 0:2.5-8.1.amzn2.0.2                   libstdc++.i686 0:7.3.1-12.amzn2                    libverto.i686 0:0.2.5-4.amzn2.0.2
  ncurses-libs.i686 0:6.0-8.20170212.amzn2.1.3        nspr.i686 0:4.25.0-2.amzn2                         nss.i686 0:3.53.1-3.amzn2
  nss-pem.i686 0:1.0.3-5.amzn2                        nss-softokn.i686 0:3.53.1-6.amzn2                  nss-softokn-freebl.i686 0:3.53.1-6.amzn2
  nss-util.i686 0:3.53.1-1.amzn2                      openldap.i686 0:2.4.44-22.amzn2                    openssl-libs.i686 1:1.0.2k-19.amzn2.0.6
  pcre.i686 0:8.32-17.amzn2.0.2                       postgresql-libs.x86_64 0:9.2.24-1.amzn2.0.1        readline.i686 0:6.2-10.amzn2.0.2
  sqlite.i686 0:3.7.17-8.amzn2.1.1                    zlib.i686 0:1.2.7-18.amzn2

完了しました!

# インストールしたモジュールのバージョンを確認する
[ec2-user@ip-xxx-xx-xx-xxx ~]$ psql -V
psql (PostgreSQL) 9.2.24
```
ここまでの設定をすることで、psqlコマンドを使ってRDSの接続先DBサーバーRDSと通信できるようになります。
なお、AWSセキュリティグループの設定については記載しませんでしたが、あらかじめ接続元サーバーと接続先サーバーで通信が行えるように、インバウンドルールを設定する必要があります。
```
[ec2-user@ip-xxx-xx-xx-xxx ~]$ psql -h example.ap-northeast-1.rds.amazonaws.com -U example_user -d example_db
```
