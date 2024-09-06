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

if ! test -d "pb"; then
  mkdir pb
fi

# format api
goctl api format --dir ./api/*.api

# gen paginator
protoc ./proto/paginator/paginator.proto --proto_path=./proto/ --go_out=paths=source_relative:./pb

# gen biz service code
biz_services=(user post comment)

for item in "${biz_services[@]}"; do
  goctl api go -api ./api/"$item".api -dir ./service/"$item"/api/
  goctl rpc protoc ./proto/"$item"/"$item".proto --proto_path=./proto/ --go_out=./pb --go-grpc_out=./pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --zrpc_out=./service/"$item"/rpc/
done

# build biz service
for item in "${biz_services[@]}"; do
  GOOS=${GOOS} GOARCH=${GOARCH} go build  -o ../target/"$item"-rpc ./service/"$item"/rpc/"$item".go
  GOOS=${GOOS} GOARCH=${GOARCH} go build  -o ../target/"$item"-api ./service/"$item"/api/"$item".go
done

# gen base service code
base_services=(leaf)

for item in "${base_services[@]}"; do
  goctl rpc protoc ./proto/"$item"/"$item".proto --proto_path=./proto/ --go_out=./pb --go-grpc_out=./pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --zrpc_out=./service/"$item"/rpc/
done

#build base service
for item in "${base_services[@]}"; do
  GOOS=${GOOS} GOARCH=${GOARCH} go build  -o ../target/"$item"-rpc ./service/"$item"/rpc/"$item".go
done