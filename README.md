# Memomail

![version](https://img.shields.io/badge/version-bata-orange.svg)
[![GoDoc](https://godoc.org/github.com/kamiazya/memomail?status.svg)](https://godoc.org/github.com/kamiazya/memomail)
[![Go Report Card](https://goreportcard.com/badge/github.com/kamiazya/memomail)](https://goreportcard.com/report/github.com/kamiazya/memomail)
[![Build Status](https://travis-ci.org/kamiazya/memomail.svg?branch=master)](https://travis-ci.org/kamiazya/memomail)
[![Maintainability](https://api.codeclimate.com/v1/badges/8e65ec7e5a316caafb80/maintainability)](https://codeclimate.com/github/kamiazya/memomail/maintainability)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

Memomail is a tool to write a memo to send mail.

> Memomailはメモを書いて、メールで送るツールです。

## Demo

[![asciicast](https://asciinema.org/a/rqcoRmVrfOQpxp9igwPnGSjwj.png)](https://asciinema.org/a/rqcoRmVrfOQpxp9igwPnGSjwj)

The above is a demo to write a memo using Memomail and send it by email.

For the demonstration, I set up a mail server with [Docker](https://hub.docker.com/r/kamiazya/mailcatcher/) so that I can check mail from `localhost:1080`.

> 上記はMemomailを使ってメモを書き、メールで送信するデモです。
>
> デモ用に[Docker](https://hub.docker.com/r/kamiazya/mailcatcher/)でメールサーバーを立て、`localhost:1080`からメールを確認できるようにしています。

## Overview

Easily take notes with your favorite editor and send them to your mailbox.

Of course, you can attach attachments as well.

> 好きなエディタで簡単にメモを取って、自分のメールボックスに送信します。
>
> もちろん、添付ファイルもつけることができます。

## Features

- Write a note in your editor.
- And send a mail to you.

> 好きなエディタで書いたメモをメールで送信します。

## Usage

```bash
$ memomail
open editor and you write a note.

$ echo "hello, meilmemo" | mailmemo
# send stdin
# message will be "hello, meilmemo".

$ mailmemo -a hoge.txt -a huga.log
# send mail with some attachments
```

### Config File

It is necessary to tell `meilmemo` the location of the mail server in the following installation of the setting file.

> 設定ファイルを下記のパスに設置し、メールサーバーの位置を`meilmemo`に伝える必要があります。

- `/etc/memomail/memomail.yml`
- `$HOME/.config/memomail/memomail.yml`
- `$PWD/memomail.yml`

The default configuration file is generated in `$HOME/.config/memomail/memomail.yml` the first time you start it.

> 最初に起動したときに、 `$HOME/.config/memomail/memomail.yml` にデフォルトの設定ファイルが生成されます。

This setting is described in YAML format.

> この設定はYAML形式で記述します。
> 下記はサンプルです。

```yml memomail.yml
mailer:
  # about server
  host: localhost
  port: 25

  # about account
  email-address: test@example.com
  username: yourname
  password: secrEtpassw0rd
```

You can also specify an editor to use by default.

> また、デフォルトで使うエディタを指定することもできます。

```yml memomail.yml
editor: vim
```

A sample file is in `memomail.sample.yml`.

> サンプルファイルは`memomail.sample.yml`にあります。

## Installation

```bash
$ go install -v github.com/kamiazya/memomail
github.com/kamiazya/memomail
```

## License

[MIT](https://github.com/kamiazya/mailcatcher/blob/master/LICENSE)

## Author

[kamaizya](https://github.com/kamiazya)
