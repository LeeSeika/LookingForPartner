eval "$(cat .env | sed 's/^/export /')"

echo "program pid list" > ../services.pid

../target/user-rpc -f ${USER_RPC_CONFIG_PATH} & echo "user-rpc:$!" >> ../services.pid
sleep 2

../target/comment-rpc -f ${COMMENT_RPC_CONFIG_PATH} & echo "comment-rpc:$!" >> ../services.pid
sleep 2

../target/post-rpc -f ${POST_RPC_CONFIG_PATH} & echo "post-rpc:$!" >> ../services.pid
sleep 2

../target/user-api -f ${USER_API_CONFIG_PATH} & echo "user-api:$!" >> ../services.pid
sleep 2

../target/comment-api -f ${COMMENT_API_CONFIG_PATH} & echo "comment-api:$!" >> ../services.pid
sleep 2

../target/post-api -f ${POST_API_CONFIG_PATH} & echo "post-api:$!" >> ../services.pid