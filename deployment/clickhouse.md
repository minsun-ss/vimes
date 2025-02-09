For testinng locally.

docker run -d \
  -p 18123:8123 \
  -p 19000:9000 \
  -e CLICKHOUSE_PASSWORD=mypassword \
  --name vimes-clickhouse \
  --ulimit nofile=262144:262144 \
  clickhouse:24.12.3
