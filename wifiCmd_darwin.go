//go:build darwin

package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jaisonerick/macwifi"
)

func wifiScanCmd() ([]*wifiAPInfo, error) {
	r := []*wifiAPInfo{}
	networks, err := macwifi.Scan(context.Background())
	if err != nil {
		return r, err
	}
	for _, net := range networks {
		bssid := strings.ToUpper(net.BSSID)
		if bssid == "" {
			continue
		}
		r = append(r, &wifiAPInfo{
			SSID:    net.SSID,
			BSSID:   bssid,
			RSSI:    strconv.Itoa(net.RSSI),
			Channel: strconv.Itoa(net.Channel),
			Info:    fmt.Sprintf("%s;%s;%s;%dMHz", net.ChannelBand.String(), net.Security.String(), net.PHYMode, net.ChannelWidth),
		})
	}
	return r, nil
}
