package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"strings"
	"syscall"
	"time"
)

var version = "v1.0.0"
var commit = ""
var syslogDst = ""
var iface = "wlan0"
var syslogInterval = 600
var cpuprofile string
var memprofile string

func init() {
	flag.StringVar(&syslogDst, "syslog", "", "syslog destnation list")
	flag.StringVar(&iface, "iface", "wlan0", "monitor interface")
	flag.IntVar(&syslogInterval, "interval", 600, "syslog send interval(sec)")
	flag.StringVar(&cpuprofile, "cpuprofile", "", "write cpu profile to `file`")
	flag.StringVar(&memprofile, "memprofile", "", "write memory profile to `file`")
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
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatalf("could not create CPU profile: %v", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatalf("could not start CPU profile: %v", err)
		}
		defer pprof.StopCPUProfile()
	}
	if memprofile != "" {
		f, err := os.Create(memprofile)
		if err != nil {
			log.Fatalf("could not create memory profile: %v", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatalf("could not write memory profile:%v", err)
		}
	}
	log.Printf("version=%s", fmt.Sprintf("%s(%s)", version, commit))
	if iface == "" {
		log.Fatalln("no monitor interface")
	}
	if syslogDst == "" {
		log.Fatalln("no syslog distenation")
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go startSyslog(ctx)
	go startWifiScan(ctx)
	<-quit
	syslogCh <- "quit by signal"
	time.Sleep(time.Second * 1)
	log.Println("quit by signal")
	cancel()
	time.Sleep(time.Second * 2)
}
