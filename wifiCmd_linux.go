package main

import (
	"os/exec"
	"regexp"
	"strings"
)

var bssidReg = regexp.MustCompile(`Address:\s*([0-9:A-F]+)`)
var ssidReg = regexp.MustCompile(`ESSID:"([^"]+)"`)
var rssiReg = regexp.MustCompile(`Signal\s+level=([-0-9]+)\s*dBm`)
var channelReg = regexp.MustCompile(`Channel:\s*(\d+)`)
var ieeeReg = regexp.MustCompile(`IE: IEEE\s*(.+)`)

func wifiScanCmd() ([]*wifiAPInfo, error) {
	r := []*wifiAPInfo{}
	cmd := exec.Command("iwlist", iface, "scan")
	o, err := cmd.Output()
	if err != nil {
		return r, err
	}
	w := &wifiAPInfo{}
	for _, l := range strings.Split(string(o), "\n") {
		if a := bssidReg.FindStringSubmatch(l); len(a) > 1 {
			if w.BSSID != "" {
				r = append(r, w)
				w = &wifiAPInfo{}
			}
			w.BSSID = a[1]
			continue
		}
		if a := ssidReg.FindStringSubmatch(l); len(a) > 1 {
			w.SSID = a[1]
			continue
		}
		if a := rssiReg.FindStringSubmatch(l); len(a) > 1 {
			w.RSSI = a[1]
			continue
		}
		if a := channelReg.FindStringSubmatch(l); len(a) > 1 {
			w.Channel = a[1]
			continue
		}
		if a := ieeeReg.FindStringSubmatch(l); len(a) > 1 {
			w.Info += a[1]
			continue
		}
		if strings.Contains(l, "Encryption key:on") {
			w.Info += "Encrypt;"
		}
	}
	if w.BSSID != "" && w.SSID != "" {
		r = append(r, w)
	}
	return r, nil
}
