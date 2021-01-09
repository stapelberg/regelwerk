.PHONY: midna dr

all: midna

midna:
	go install -tags rw_midna && systemctl --user restart regelwerk

dr:
	CGO_ENABLED=0 GOARCH=arm64 go install
	# TODO: cpu -host=dr home/michael/go/bin/linux_arm64/regelwerk
	# to deploy: cd ~/gokrazy/dr && make update
