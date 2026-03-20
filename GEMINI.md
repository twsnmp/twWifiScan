# twWifiScan Project Overview

`twWifiScan` is a Wi-Fi Access Point (AP) sensor designed for TWSNMP FC. It scans for surrounding wireless LAN access points and transmits the collected data via Syslog or MQTT. It also monitors and reports the host system's resource usage.

## Main Technologies
- **Language:** Go (1.24.0+)
- **Communication Protocols:** Syslog (UDP), MQTT (JSON)
- **External Libraries:**
  - `github.com/eclipse/paho.mqtt.golang`: MQTT client.
  - `github.com/shirou/gopsutil`: System resource monitoring.
  - `golang.org/x/text`: Japanese character encoding support (for Windows `netsh` output).
- **Build System:** Makefile, GoReleaser (configured via `.goreleaser.yaml`).

## Architecture
The application is structured into several concurrent components managed by goroutines:
1. **Wi-Fi Scanner (`wifiScan.go`):** Periodically triggers platform-specific scan commands.
2. **Platform Adapters:**
   - `wifiCmd_linux.go`: Uses `iwlist <iface> scan`.
   - `wifiCmd_windows.go`: Uses `netsh.exe wlan show networks mode=Bssid`.
   - `wifiCmd_darwin.go`: Dummy implementation (currently returns empty).
3. **Data Exporters:**
   - `syslog.go`: Sends formatted text logs via UDP.
   - `mqtt.go`: Publishes JSON-encoded data to a broker.
4. **Monitor (`monitor.go`):** Collects CPU, memory, load, and network stats.
5. **Vendor Lookup (`vendor.go`):** Contains a large OUI map for MAC-to-vendor resolution.

## Building and Running

### Build
To build the executables for multiple platforms (Windows, Linux AMD64/ARM):
```bash
make
```
The binaries are generated in the `dist/` directory.

### Run
The sensor requires at least a monitoring interface and a destination (Syslog or MQTT).

```bash
# Example for Linux
sudo ./twWifiScan -iface wlan0 -syslog 192.168.1.1 -mqtt tcp://broker.example.com:1883
```

#### Key Flags:
- `-iface`: Network interface to use for scanning (default: `wlan0`).
- `-interval`: Reporting interval in seconds (default: `600`).
- `-syslog`: Comma-separated list of syslog servers.
- `-mqtt`: MQTT broker URL.
- `-debug`: Enable verbose debug logging.

## Development Conventions
- **Concurrent Programming:** Uses channels (`syslogCh`, `mqttCh`) to decouple data collection from transmission.
- **Error Handling:** Errors in background goroutines are logged but typically do not halt the entire process.
- **Cross-Platform:** Uses build tags (`//go:build`) to separate OS-specific Wi-Fi scanning logic.
- **Vendor Mapping:** The `ouiMap` in `vendor.go` is statically compiled; updates require modifying this file.

## Key Files
- `main.go`: Orchestrates the application lifecycle and parses flags.
- `wifiScan.go`: Logic for processing AP data and tracking changes.
- `syslog.go` / `mqtt.go`: Protocol-specific transmission logic.
- `monitor.go`: System resource monitoring integration.
- `vendor.go`: Large OUI database for vendor identification.
