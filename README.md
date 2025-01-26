# twWifiScan
Wifi AP sensor for TWSNMP FC  
TWSNMP FCのためのWifiアクセスポイントセンサー

[![Godoc Reference](https://godoc.org/github.com/twsnmp/twWifiScan?status.svg)](http://godoc.org/github.com/twsnmp/twWifiScan)
[![Go Report Card](https://goreportcard.com/badge/twsnmp/twWifiScan)](https://goreportcard.com/report/twsnmp/twWifiScan)

## Overview

Information on the wireless LAN access points around the surrounding area to TWSNMP FC, etc.
It is a sensor program for sending by syslog.
You can get the following information in the current version.

- The statistical information of monitored packets
- Sensor resources
- Prot LAN Access Points BSSID, SSID, RSSI, Channel, information on encryption


周辺にある無線LANのアクセスポイントの情報をTWSNMP FCなどへ  
syslogで送信するためのセンサープログラムです。  
現在のバージョンでは以下の情報を取得できます。

- モニタしたパケット数の統計情報
- センサーのリソース
- 無線LANアクセスポイントのBSSID,SSID,RSSI,Channel,暗号化の有無などの情報

## Status

v1.0.0 has been released.(2021/9/8)
(State in which basic functions operate)
v1.0.1 has been released.(2021/10/31)
(Mac OS version of Bug Fix)
v1.1.0 has beedn released.(2025/1/26)
(Automatic release)

v1.0.0をリリースしました。(2021/9/8)  
（基本的な機能の動作する状態）  
v1.0.1をリリースしました。(2021/10/31) 
（Mac OS版のバグフィックス）  
v1.1.0をリリース (2025/1/26)
(自動リリース)

## Build

build by make command.
ビルドはmakeで行います。

```
$make
```

You can specify the following targets.
以下のターゲットが指定できます。

```
  all        Build a full -executable file (omitted)
  mac        Build of executable file for Mac
  clean      Delete the builded executable file
  zip        Create Zip files for release
```

```
  all        全実行ファイルのビルド（省略可能）
  mac        Mac用の実行ファイルのビルド
  clean      ビルドした実行ファイルの削除
  zip        リリース用のZIPファイルを作成
```

```
$make
```

Execute, execution files for macOS, Windows, Linux (AMD64), Linux (ARM),
It is created in the `dist` directory.

を実行すれば、MacOS,Windows,Linux(amd64),Linux(arm)用の実行ファイルが、  
`dist`のディレクトリに作成されます。


To create a zip file for distribution,
配布用のZIPファイルを作成するためには、

```
$make zip
```

Is executed.The zip file is created in the `dist/` directory.

を実行します。ZIPファイルが`dist/`ディレクトリに作成されます。

## Run

### Usage

```
Usage of ./twWifiScan.app:
Usage of dist/twWifiScan.app:
  -cpuprofile file
    	write cpu profile to file
  -iface string
    	monitor interface (default "wlan0")
  -interval int
    	syslog send interval(sec) (default 600)
  -memprofile file
    	write memory profile to file
  -syslog string
    	syslog destnation list
```

Syslog destinations can be specified multiple by separation of comma.
: You can also specify the port number.

syslogの送信先はカンマ区切りで複数指定できます。  
:に続けてポート番号を指定することもできます。

```
-syslog 192.168.1.1,192.168.1.2:5514
```


To start up, you need a monitoring LAN I/F (-iface) and syslog destination (-syslog).

起動するためには、モニタするLAN I/F(-iface)とsyslogの送信先(-syslog)が必要です。

In the Mac OS, Windows, Linux environment, you can start with the following command.
(In the case of Linux)

Mac OS,Windows,Linuxの環境では以下のコマンドで起動できます。  
（例はLinux場合）

```
#./twWifiScan -iface wlan0 -syslog 192.168.1.1
```

## syslog message example

The sentence of the transmitted syslog message is `local5`.TAG is `twWifiScan`.

送信されるsyslogのメッセージのファシリティーは`local5`です。tagは`twWifiScan`です。

This is an example of a log of APInfo.

アクセスポイントのログの例です。

```
type=APInfo,ssid=F660T-VFyM-X,bssid=FC:C8:97:B0:xx:D5,rssi=-73,Channel=1,info=Encrypt;802.11i/WPA2 Version 1,count=9593,change=0,vendor=zte,ft=2024-12-06T15:48:57+09:00,lt=2025-01-26T16:38:57+09:00
```

You can get information on SSI, signal level, and cryptocation intensity.

SSI、信号のレベル、暗号強度の情報を取得できます。


## TWSNMP FCのパッケージ

TWWIFISCAN is included in the TWSNMP FC package.
There is Windows/Mac OS/Linux (AMD64, Arm).

For more information
https://note.com/twsnmp/n/nc6e49c284afb
Please see

TWSNMP FCのパッケージにtwWifiScanが含まれています。
Windows/Mac OS/Linux(amd64,arm)があります。

詳しくは、  
https://note.com/twsnmp/n/nc6e49c284afb  
を見てください。

## Copyright

see ./LICENSE

```
Copyright 2021-2025 Masayuki Yamai
```
