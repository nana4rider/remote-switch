# remote-switch

PC Remote Switch

## 概要

PCの電源をON/OFFするためのAPIです。
* 電源のONは、Wake-on-LANで行う。
* 電源のOFFは、SSHでログインしコマンドを実行することで行う。
* 電源状態の確認は、pingコマンドで行う。
* PCを参照/登録/更新/削除するAPIを提供。

## API

### PCの一覧を取得します
```http
GET /computers
```
response
```json
[
  {
    "id": 1,
    "name": "pc1",
    "ssh_user": "user1",
    "ssh_key": null,
    "ssh_port": null,
    "ip_address": "192.168.1.100",
    "mac_address": "11:22:33:44:55:66"
  },
  {
    "id": 2,
    "name": "pc2",
    "ssh_user": "user2",
    "ssh_key": "/path/to/ssh_key",
    "ssh_port": 2222,
    "ip_address": "192.168.1.200",
    "mac_address": "aa:bb:cc:dd:ee:ff"
  }
]
```

### PCを取得します
```http
GET /computers/:id
```
response
```json
{
  "id": 1,
  "name": "pc1",
  "ssh_user": "user1",
  "ssh_key": null,
  "ssh_port": null,
  "ip_address": "192.168.1.100",
  "mac_address": "11:22:33:44:55:66"
}
```

### PCを登録します
```http
POST /computers
```
request
```json
{
  "name": "pc3",
  "ip_address": "192.168.1.150",
  "ssh_user": "user3",
  "ssh_key": null, // default $HOME/.ssh/id_rsa
  "ssh_port": null, // default 22
  "mac_address": "11:22:33:44:55:99" // default "arp -a ip_address"
}
```
response
```json
{
  "id": 3,
  "name": "pc3",
  "ssh_user": "user3",
  "ssh_key": null,
  "ssh_port": null,
  "ip_address": "192.168.1.150",
  "mac_address": "11:22:33:44:55:99"
}
```

### PCを更新します
```http
PUT /computers/:id
```
request
```json
{
  "ssh_user": "user3_2",
  "ip_address": "192.168.1.151"
}
```

### PCを削除します
```http
DELETE /computers/:id
```

### PCの起動状態を取得します
```http
GET /computers/:id/power
```
response
```json
{
  "state": "ON or OFF"
}
```

### PCの電源状態を変更します
```http
PUT /computers/:id/power
```
request
```json
{
  "state": "ON or OFF"
}
```
response
```json
{
  "message": "",
  "result": true
}
```
