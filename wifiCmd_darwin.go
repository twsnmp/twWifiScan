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
		i := 1
		for i < len(fs)-1 {
			if strings.Contains(fs[i], ":") {
				_, errParse := strconv.Atoi(fs[i+1])
				if errParse == nil {
					break
				}
			}
			i++
		}
		if i >= len(fs)-1 {
			continue
		}
		ssid := strings.Join(fs[0:i], " ")
		rssi, errParse := strconv.Atoi(fs[i+1])
		if errParse != nil {
			continue
		}
		if rssi > 0 {
			continue
		}
		r = append(r, &wifiAPInfo{
			SSID:    ssid,
			BSSID:   fs[i],
			RSSI:    fs[i+1],
			Channel: fs[i+2],
			Info:    strings.Join(fs[i+3:], ";"),
		})
	}
	return r, nil
}
