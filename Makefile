.PHONY: install

install:
	tinygo flash -target=pygamer cmd/xship/main.go
