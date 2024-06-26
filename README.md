# twWifiScan
Wifi AP sensor for TWSNMP FC  
TWSNMP FCのためのWifiアクセスポイントセンサー

[![Godoc Reference](https://godoc.org/github.com/twsnmp/twWifiScan?status.svg)](http://godoc.org/github.com/twsnmp/twWifiScan)
[![Go Report Card](https://goreportcard.com/badge/twsnmp/twWifiScan)](https://goreportcard.com/report/twsnmp/twWifiScan)

## Overview

周辺にある無線LANのアクセスポイントの情報をTWSNMP FCなどへ  
syslogで送信するためのセンサープログラムです。  
現在のバージョンでは以下の情報を取得できます。

- モニタしたパケット数の統計情報
- センサーのリソース
- 無線LANアクセスポイントのBSSID,SSID,RSSI,Channel,暗号化の有無などの情報

## Status

v1.0.0をリリースしました。(2021/9/8)  
（基本的な機能の動作する状態）  
v1.0.1をリリースしました。(2021/10/31) 
（Mac OS版のバグフィックス）  

## Build

ビルドはmakeで行います。
```
$make
```
以下のターゲットが指定できます。
```
  all        全実行ファイルのビルド（省略可能）
  mac        Mac用の実行ファイルのビルド
  clean      ビルドした実行ファイルの削除
  zip        リリース用のZIPファイルを作成
```

```
$make
```
を実行すれば、MacOS,Windows,Linux(amd64),Linux(arm)用の実行ファイルが、  
`dist`のディレクトリに作成されます。


配布用のZIPファイルを作成するためには、
```
$make zip
```
を実行します。ZIPファイルが`dist/`ディレクトリに作成されます。

## Run

### 使用方法

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

syslogの送信先はカンマ区切りで複数指定できます。  
:に続けてポート番号を指定することもできます。

```
-syslog 192.168.1.1,192.168.1.2:5514
```


### 起動方法

起動するためには、モニタするLAN I/F(-iface)とsyslogの送信先(-syslog)が必要です。

Mac OS,Windows,Linuxの環境では以下のコマンドで起動できます。  
（例はLinux場合）

```
#./twWifiScan -iface wlan0 -syslog 192.168.1.1
```

## TWSNMP FCのパッケージ

TWSNMP FCのパッケージにtwWifiScanが含まれています。  
Windows/Mac OS/Linux(amd64,arm)があります。  
詳しくは、  
https://note.com/twsnmp/n/nc6e49c284afb  
を見てください。

## Copyright

see ./LICENSE

```
Copyright 2021 Masayuki Yamai
```
