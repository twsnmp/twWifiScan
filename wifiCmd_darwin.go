//go:build darwin

package main

// Dummy
func wifiScanCmd() ([]*wifiAPInfo, error) {
	r := []*wifiAPInfo{}
	return r, nil
}
