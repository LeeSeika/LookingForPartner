set -ex

VERSION=1.0.0

echo $VERSION

LDFLAGS="-X main.Version=${VERSION}"

UNAME_S=$(uname -s)
GOOS=linux
if [ $UNAME_S = "Darwin" ]; then
	GOOS=darwin
fi

GOOS=${GOOS} GOARCH=amd64 go build -ldflags "${LDFLAGS}" -o ../target/user-api ./service/user/api/user.go
GOOS=${GOOS} GOARCH=amd64 go build -ldflags "${LDFLAGS}" -o ../target/user-rpc ./service/user/rpc/user.go

GOOS=${GOOS} GOARCH=amd64 go build -ldflags "${LDFLAGS}" -o ../target/post-api ./service/post/api/post.go
GOOS=${GOOS} GOARCH=amd64 go build -ldflags "${LDFLAGS}" -o ../target/post-rpc ./service/post/rpc/post.go