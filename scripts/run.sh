eval "$(cat .env | sed 's/^/export /')"

echo "program pid list" > ../services.pid

../target/user-rpc -f ${USER_RPC_CONFIG_PATH} | tee ../log/user-rpc.log & echo "user-rpc:$!" >> ../services.pid
sleep 2

../target/post-rpc -f ${POST_RPC_CONFIG_PATH} | tee ../log/post-rpc.log & echo "post-rpc:$!" >> ../services.pid
sleep 2

../target/user-api -f ${USER_API_CONFIG_PATH} | tee ../log/user-api.log & echo "user-api:$!" >> ../services.pid
sleep 2

../target/post-api -f ${POST_API_CONFIG_PATH} | tee ../log/post-api.log & echo "post-api:$!" >> ../services.pid