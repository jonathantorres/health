release:
	rm -fr ./bin && mkdir ./bin
	GOOS=darwin GOARCH=amd64 go build && mv ./health ./bin/health_darwin
	GOOS=freebsd GOARCH=amd64 go build && mv ./health ./bin/health_freebsd
	GOOS=linux GOARCH=amd64 go build && mv ./health ./bin/health_linux
	GOOS=netbsd GOARCH=amd64 go build && mv ./health ./bin/health_netbsd
	GOOS=plan9 GOARCH=amd64 go build && mv ./health ./bin/health_plan9
	GOOS=windows GOARCH=amd64 go build && mv ./health.exe ./bin/health_windows.exe
