.PHONY: all test clean zip mac

### バージョンの定義
VERSION     := "v1.0.0"
COMMIT      := $(shell git rev-parse --short HEAD)

### コマンドの定義
GO          = go
GO_BUILD    = $(GO) build
GO_TEST     = $(GO) test -v
GO_LDFLAGS  = -ldflags="-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT)"
ZIP          = zip

### ターゲットパラメータ
DIST = dist
SRC = ./main.go ./syslog.go ./monitor.go ./wifiScan.go
TARGETS     = $(DIST)/twWifiScan.exe $(DIST)/twWifiScan.app $(DIST)/twWifiScan $(DIST)/twWifiScan.arm
GO_PKGROOT  = ./...

### PHONY ターゲットのビルドルール
all: $(TARGETS)
test:
	env GOOS=$(GOOS) $(GO_TEST) $(GO_PKGROOT)
clean:
	rm -rf $(TARGETS) $(DIST)/*.zip
mac: $(DIST)/twWifiScan.app
zip: $(TARGETS)
	cd dist && $(ZIP) twWifiScan_win.zip twWifiScan.exe
	cd dist && $(ZIP) twWifiScan_mac.zip twWifiScan.app
	cd dist && $(ZIP) twWifiScan_linux_amd64.zip twWifiScan
	cd dist && $(ZIP) twWifiScan_linux_arm.zip twWifiScan.arm

docker:  dist/twWifiScan Docker/Dockerfile
	cp dist/twWifiScan Docker/
	cd Docker && docker build -t twsnmp/twWifiScan .

### 実行ファイルのビルドルール
$(DIST)/twWifiScan.exe: $(SRC) ./wifiCmd_windows.go
	env GO111MODULE=on GOOS=windows GOARCH=amd64 $(GO_BUILD) $(GO_LDFLAGS) -o $@
$(DIST)/twWifiScan.app: $(SRC) ./wifiCmd_darwin.go
	env GO111MODULE=on GOOS=darwin GOARCH=amd64 $(GO_BUILD) $(GO_LDFLAGS) -o $@
$(DIST)/twWifiScan.arm: $(SRC) ./wifiCmd_linux.go
	env GO111MODULE=on GOOS=linux GOARCH=arm GOARM=7 $(GO_BUILD) $(GO_LDFLAGS) -o $@
$(DIST)/twWifiScan: $(SRC) ./wifiCmd_linux.go
	env GO111MODULE=on GOOS=linux GOARCH=amd64 $(GO_BUILD) $(GO_LDFLAGS) -o $@

