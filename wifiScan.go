package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
)

// startWifiScan : start Wifi Scan
func startWifiScan(ctx context.Context) {
	timer := time.NewTicker(time.Second * time.Duration(interval))
	defer timer.Stop()
	log.Println("start wifi scan")
	for {
		select {
		case <-timer.C:
			count := sendReport()
			sendSyslog(fmt.Sprintf("type=Stats,total=%d,count=%d,ps=%.2f,send=%d,param=%s",
				len(apMap), count, float64(count)/float64(interval), syslogCount, iface))
			publishMQTT(&mqttWifiScanStatsDataEnt{
				Time:      time.Now().Format(time.RFC3339),
				Total:     len(apMap),
				Count:     count,
				PS:        float64(count) / float64(interval),
				Interface: iface,
			})
			if debug {
				log.Printf("type=Stats,total=%d,count=%d,ps=%.2f,send=%d,param=%s",
					len(apMap), count, float64(count)/float64(interval), syslogCount, iface)
			}
			syslogCount = 0
			sendMonitor()
		case <-ctx.Done():
			log.Println("stop wifi scan")
			return
		}
	}
}

type wifiAPInfo struct {
	SSID      string
	BSSID     string
	RSSI      string
	Channel   string
	Info      string
	Vendor    string
	Count     int
	Change    int
	FirstTime int64
	LastTime  int64
}

func (e *wifiAPInfo) String() string {
	return fmt.Sprintf("type=APInfo,ssid=%s,bssid=%s,rssi=%s,Channel=%s,info=%s,count=%d,change=%d,vendor=%s,ft=%s,lt=%s",
		e.SSID, e.BSSID, e.RSSI, e.Channel, e.Info, e.Count, e.Change,
		e.Vendor,
		time.Unix(e.FirstTime, 0).Format(time.RFC3339),
		time.Unix(e.LastTime, 0).Format(time.RFC3339),
	)
}

var apMap = make(map[string]*wifiAPInfo)

// syslogでレポートを送信する
func sendReport() int {
	list, err := wifiScanCmd()
	if err != nil {
		log.Printf("wifiScanCmd err=%v", err)
		return 0
	}
	for _, ap := range list {
		ap.BSSID = strings.ToUpper(ap.BSSID)
		ap.Vendor = getVendorFromAddress(ap.BSSID)
		if e, ok := apMap[ap.BSSID]; ok {
			e.Count++
			if e.SSID != ap.SSID || e.Channel != ap.Channel || e.Info != ap.Info {
				e.Change++
			}
			e.Channel = ap.Channel
			e.SSID = ap.SSID
			e.RSSI = ap.RSSI
			e.Info = ap.Info
			e.LastTime = time.Now().Unix()
			sendSyslog(e.String())
			publishMQTT(&mqttApInfoDataEnt{
				Time:      time.Now().Format(time.RFC3339),
				SSID:      e.SSID,
				BSSID:     e.BSSID,
				RSSI:      e.RSSI,
				Channel:   e.Channel,
				Info:      e.Info,
				Vendor:    e.Vendor,
				Count:     e.Count,
				Change:    e.Change,
				FirstTime: time.Unix(e.FirstTime, 0).Format(time.RFC3339),
				LastTime:  time.Unix(e.LastTime, 0).Format(time.RFC3339),
			})
		} else {
			ap.FirstTime = time.Now().Unix()
			ap.LastTime = time.Now().Unix()
			ap.Count = 1
			apMap[ap.BSSID] = ap
			if debug {
				log.Println("new AP", ap.String())
			}
			sendSyslog(ap.String())
			publishMQTT(&mqttApInfoDataEnt{
				Time:      time.Now().Format(time.RFC3339),
				SSID:      ap.SSID,
				BSSID:     ap.BSSID,
				RSSI:      ap.RSSI,
				Channel:   ap.Channel,
				Info:      ap.Info,
				Vendor:    ap.Vendor,
				Count:     ap.Count,
				Change:    ap.Change,
				FirstTime: time.Unix(ap.FirstTime, 0).Format(time.RFC3339),
				LastTime:  time.Unix(ap.LastTime, 0).Format(time.RFC3339),
			})
		}
	}
	return len(list)
}
