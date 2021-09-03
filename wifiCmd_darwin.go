package main

import (
	"os/exec"
	"strconv"
	"strings"
)

func wifiScanCmd() ([]*wifiAPInfo, error) {
	r := []*wifiAPInfo{}
	cmd := exec.Command("/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport", "-s")
	o, err := cmd.Output()
	if err != nil {
		return r, err
	}
	for _, l := range strings.Split(string(o), "\n") {
		fs := strings.Fields(l)
		if len(fs) < 6 {
			continue
		}
		rssi, errParse := strconv.Atoi(fs[2])
		if errParse != nil {
			continue
		}
		if rssi > 0 {
			continue
		}
		r = append(r, &wifiAPInfo{
			SSID:    fs[0],
			BSSID:   fs[1],
			RSSI:    fs[2],
			Channel: fs[3],
			Info:    strings.Join(fs[4:], ";"),
		})
	}
	return r, nil
}
