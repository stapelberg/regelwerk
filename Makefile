.PHONY: midna dr

all: midna

midna:
	go install -tags rw_midna && systemctl --user restart regelwerk

# to deploy permanently: cd ~/gokrazy/dr && make update
dr:
	gok run --instance=dr

test:
	go test -v
