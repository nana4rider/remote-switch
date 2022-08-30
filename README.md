# remote-switch

PC Remote Switch

## 概要

PCの電源をON/OFFするためのAPIです。

* 電源のONは、Wake-on-LANで行う。
* 電源のOFFは、SSHでログインしコマンドを実行することで行う。
* 電源状態の確認は、pingコマンドで行う。

## APIドキュメント
https://nana4rider.github.io/remote-switch/

## APIドキュメント更新
`swag i -ot yaml -o .`
