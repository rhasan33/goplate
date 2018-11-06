# booting up consul, redis, mysql containers
docker-compose up -d consul db redis

# building app
go build -v .

# setting KV, dependecy of app
curl --request PUT --data-binary @config.local.yml http://localhost:8500/v1/kv/goplate

# start app
export GOPLATE_CONSUL_URL="127.0.0.1:8500"
export GOPLATE_CONSUL_PATH="goplate"
./goplate serve
