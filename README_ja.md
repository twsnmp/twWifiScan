# twWifiScan
TWSNMP FCのためのWifiアクセスポイントセンサーです。

[![Godoc Reference](https://godoc.org/github.com/twsnmp/twWifiScan?status.svg)](http://godoc.org/github.com/twsnmp/twWifiScan)
[![Go Report Card](https://goreportcard.com/badge/twsnmp/twWifiScan)](https://goreportcard.com/report/twsnmp/twWifiScan)

[English](./README.md)

![](images/twWifiScan.png)

## 概要

周辺にある無線LANのアクセスポイントの情報をTWSNMP FCなどへ、syslogまたはMQTTで送信するためのセンサープログラムです。
現在のバージョンでは以下の情報を取得できます。

- モニタしたパケット数の統計情報
- センサーのリソース情報
- 無線LANアクセスポイントのBSSID, SSID, RSSI, Channel, 暗号化の有無などの情報

## ステータス

- v1.0.0 (2021/9/8): 基本的な機能をリリース
- v1.0.1 (2021/10/31): macOS版バグ修正
- v1.1.0 (2025/1/26): 自動リリース
- v2.0.0 (現在): MQTT対応、macOS対応削除

## ビルド

`make` コマンドでビルドします。

```bash
$ make
```

以下のターゲットが指定できます。

```
  all        全実行ファイルのビルド
  clean      ビルドした実行ファイルの削除
  zip        リリース用のZIPファイルを作成
```

実行すれば、Windows, Linux (AMD64, ARM, ARM64) 用の実行ファイルが `dist` ディレクトリに作成されます。

配布用のZIPファイルを作成するには以下を実行します。

```bash
$ make zip
```

## 実行

### 使用方法

```bash
Usage of dist/twWifiScan:
  -debug
    	デバッグモードの有効化
  -iface string
    	モニタするインターフェース (デフォルト "wlan0")
  -interval int
    	送信間隔(秒) (デフォルト 600)
  -mqtt string
    	MQTTブローカーURL
  -mqttClientID string
    	MQTTクライアントID (デフォルト "twWifiScan")
  -mqttPassword string
    	MQTTパスワード
  -mqttTopic string
    	MQTTトピックプレフィックス (デフォルト "twWifiScan")
  -mqttUser string
    	MQTTユーザー
  -syslog string
    	syslog送信先リスト
```

syslogの送信先はカンマ区切りで複数指定できます。 `:` に続けてポート番号を指定することも可能です。

```bash
-syslog 192.168.1.1,192.168.1.2:5514
```

起動するためには、モニタするLAN I/F (`-iface`) と、syslog または MQTT の送信先が必要です。

起動例 (Linux):

```bash
# ./twWifiScan -iface wlan0 -syslog 192.168.1.1 -mqtt tcp://broker.example.com:1883
```

## メッセージ例

### syslog
Facility は `local5`、Tag は `twWifiScan` です。

ログ例 (APInfo):
```
type=APInfo,ssid=F660T-VFyM-X,bssid=FC:C8:97:B0:xx:D5,rssi=-73,Channel=1,info=Encrypt;802.11i/WPA2 Version 1,count=9593,change=0,vendor=zte,ft=2024-12-06T15:48:57+09:00,lt=2025-01-26T16:38:57+09:00
```

### MQTT
JSON 形式で送信されます。トピックは指定したプレフィックスに基づいて以下のようになります。
- `<mqttTopic>/APInfo`
- `<mqttTopic>/WifiScanStats`
- `<mqttTopic>/Monitor`

## ライセンス

[LICENSE](./LICENSE) を参照してください。

```
Copyright 2021-2026 Masayuki Yamai
```
