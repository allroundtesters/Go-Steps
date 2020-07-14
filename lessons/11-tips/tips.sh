#！/bin/sh

git rev-parse --short head
# build with commit id
COMMIT_ID=`git rev-parse --short head`
go build -ldflags "-X main.Build ${COMMIT_ID}" main.go