set -ex

# get OS
UNAME_S=$(uname -s)
GOOS=linux
if [ $UNAME_S = "Darwin" ]; then
	GOOS=darwin
fi

# get arch
UNAME_M=$(uname -m)
GOARCH=arm64
if [ $UNAME_M == "x86_64" ]; then
   GOARCH=amd64
fi

# build services
services=(user post)

for item in "${services[@]}"; do
  goctl api go -api ./api/"$item".api -dir ./service/"$item"/api/
  goctl rpc protoc ./protobuf/"$item"/"$item".proto --proto_path=./protobuf/ --go_out=plugins=grpc,paths=source_relative:. --go-grpc_out=. --zrpc_out=./service/"$item"/rpc/

  GOOS=${GOOS} GOARCH=${GOARCH} go build  -o ../target/"$item"-rpc ./service/"$item"/rpc/"$item".go
  GOOS=${GOOS} GOARCH=${GOARCH} go build  -o ../target/"$item"-api ./service/"$item"/api/"$item".go
done