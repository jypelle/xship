.PHONY: install asset

install:
	tinygo flash -target=pygamer cmd/xship/main.go

asset:
	go run ./cmd/assetconvert
