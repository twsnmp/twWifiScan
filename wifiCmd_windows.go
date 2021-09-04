package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

/*
SSID 5 : aterm-82aef0-g
    ネットワークの種類            : インフラストラクチャ
    認証          : WPA2-パーソナル
    暗号化              : CCMP
    BSSID 1             : c0:25:a2:c0:31:55
         シグナル           : 48%
         無線タイプ         : 802.11n
         チャネル           : 11
         基本レート (Mbps) : 1 2 5.5 11
         他のレート (Mbps) : 6 9 12 18 24 36 48 54

SSID 9 : Buffalo-G-1358
    Network type            : Infrastructure
    Authentication          : WPA2-Personal
    Encryption              : CCMP
    BSSID 1                 : 58:27:8c:7d:13:58
         Signal             : 25%
         Radio type         : 802.11n
         Channel            : 8
         Basic rates (Mbps) : 1 2 5.5 11
         Other rates (Mbps) : 6 9 12 18 24 36 48 54

*/
var bssidReg = regexp.MustCompile(`BSSID\s+\d+\s+:\s*([0-9:A-Fa-f]+)`)
var ssidReg = regexp.MustCompile(`^SSID\s+\d+\s+:\s*(.+)`)
var rssiReg = regexp.MustCompile(`Signal\s*:\s*(\d+)%`)
var channelReg = regexp.MustCompile(`Channel\s*:\s*(\d+)`)
var radioTypeReg = regexp.MustCompile(`Radio\s+type\*+:\s*(.+)`)
var authReg = regexp.MustCompile(`Authentication\s*:\s*(.+)`)
var encReg = regexp.MustCompile(`Encryption\s*:\s*(.+)`)

var rssiJReg = regexp.MustCompile(`シグナル\s*:\s*(\d+)%`)
var channelJReg = regexp.MustCompile(`チャネル\s*:\s*(\d+)`)
var radioTypeJReg = regexp.MustCompile(`無線タイプ\s*:\s*(.+)`)
var authJReg = regexp.MustCompile(`認証\s*:\s*(.+)`)
var encJReg = regexp.MustCompile(`暗号化\s*:\s*(.+)`)

func wifiScanCmd() ([]*wifiAPInfo, error) {
	r := []*wifiAPInfo{}
	cmd := exec.Command("netsh.exe", "wlan", "show", "networks", "mode=Bssid")
	o, err := cmd.Output()
	if err != nil {
		return r, err
	}
	str, _, err := transform.String(japanese.ShiftJIS.NewDecoder(), string(o))
	if err != nil {
		log.Println(err)
		str = string(o)
	}
	w := &wifiAPInfo{}
	for _, l := range strings.Split(str, "\n") {
		l = strings.TrimSpace(l)
		if a := ssidReg.FindStringSubmatch(l); len(a) > 1 {
			if w.SSID != "" {
				r = append(r, w)
				w = &wifiAPInfo{}
			}
			w.SSID = a[1]
			continue
		}
		if a := bssidReg.FindStringSubmatch(l); len(a) > 1 {
			w.BSSID = strings.ToUpper(a[1])
			continue
		}
		if a := rssiReg.FindStringSubmatch(l); len(a) > 1 {
			if p, err := strconv.ParseInt(a[1], 10, 64); err == nil {
				w.RSSI = fmt.Sprintf("%d", (p/2)-100)
			}
			continue
		}
		if a := channelReg.FindStringSubmatch(l); len(a) > 1 {
			w.Channel = a[1]
			continue
		}
		if a := authReg.FindStringSubmatch(l); len(a) > 1 {
			w.Info += a[1]
			continue
		}
		if a := encReg.FindStringSubmatch(l); len(a) > 1 {
			w.Info += ";" + a[1]
			continue
		}
		if a := radioTypeReg.FindStringSubmatch(l); len(a) > 1 {
			w.Info += ";" + a[1]
			continue
		}
		if a := rssiJReg.FindStringSubmatch(l); len(a) > 1 {
			if p, err := strconv.ParseInt(a[1], 10, 64); err == nil {
				w.RSSI = fmt.Sprintf("%d", (p/2)-100)
			}
			continue
		}
		if a := channelJReg.FindStringSubmatch(l); len(a) > 1 {
			w.Channel = a[1]
			continue
		}
		if a := authJReg.FindStringSubmatch(l); len(a) > 1 {
			w.Info += a[1]
			continue
		}
		if a := encJReg.FindStringSubmatch(l); len(a) > 1 {
			w.Info += ";" + a[1]
			continue
		}
		if a := radioTypeJReg.FindStringSubmatch(l); len(a) > 1 {
			w.Info += ";" + a[1]
			continue
		}
	}
	if w.BSSID != "" && w.SSID != "" && w.RSSI != "" {
		r = append(r, w)
	}
	return r, nil
}
