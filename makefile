NAME="uca-edt"

build_release: build_linux build_windows

build_linux:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/$(NAME) -v && upx bin/$(NAME)

build_windows:
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o bin/$(NAME).exe -v && upx bin/$(NAME).exe