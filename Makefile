.PHONY: start-namenode start-datanode all

start-namenode:
	@echo "Starting NameNode Server..."
	@go run ./hdfs_namenode/cmd/server/main.go &

start-datanode:
	@echo "Starting DataNode Server..."
	@go run ./hdfs_datanode/cmd/datanode/main.go &

all: start-namenode start-datanode
