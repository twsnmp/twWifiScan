package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var version = "v2.0.0"
var commit = ""
var syslogDst = ""
var iface = "wlan0"
var interval = 600

var mqttDst = ""
var mqttUser = ""
var mqttPassword = ""
var mqttClientID = ""
var mqttTopic = ""
var debug = false

func init() {
	flag.StringVar(&syslogDst, "syslog", "", "syslog destnation list")
	flag.StringVar(&iface, "iface", "wlan0", "monitor interface")
	flag.IntVar(&interval, "interval", 600, "syslog send interval(sec)")
	flag.StringVar(&mqttDst, "mqtt", "", "mqtt broker url")
	flag.StringVar(&mqttUser, "mqttUser", "", "mqtt user")
	flag.StringVar(&mqttPassword, "mqttPassword", "", "mqtt password")
	flag.StringVar(&mqttClientID, "mqttClientID", "twWifiScan", "mqtt client id")
	flag.StringVar(&mqttTopic, "mqttTopic", "twWifiScan", "mqtt topic prefix")
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.VisitAll(func(f *flag.Flag) {
		if s := os.Getenv("TWWIFISCAN_" + strings.ToUpper(f.Name)); s != "" {
			f.Value.Set(s)
		}
	})
	flag.Parse()
}

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("2006-01-02T15:04:05.999 ") + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
	log.Printf("version=%s", fmt.Sprintf("%s(%s)", version, commit))
	if iface == "" {
		log.Fatalln("no monitor interface")
	}
	if syslogDst == "" && mqttDst == "" {
		log.Fatalln("no syslog and mqtt distenation")
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go startSyslog(ctx)
	go startMQTT(ctx)
	go startWifiScan(ctx)
	<-quit
	syslogCh <- "quit by signal"
	time.Sleep(time.Second * 1)
	log.Println("quit by signal")
	cancel()
	time.Sleep(time.Second * 2)
}
