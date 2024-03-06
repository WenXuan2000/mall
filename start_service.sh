#!/bin/bash

# 定义函数以生成路径
generate_paths() {
  RPC_PATH="./service/$1/rpc"
  API_PATH="./service/$1/api"
  SRC_FILE="$1"
}

# 启动服务的函数
start_service() {
  generate_paths "$1"
  go run "${RPC_PATH}/${SRC_FILE}.go" -f "${RPC_PATH}/etc/${SRC_FILE}.yaml" &
  go run "${API_PATH}/${SRC_FILE}.go" -f "${API_PATH}/etc/${SRC_FILE}.yaml" &
}

# 启动所有服务
run_all_services() {
  start_service "user"
  start_service "product"
  start_service "pay"
  start_service "order"
}

# 主程序
run_all_services
wait # 等待所有后台任务完成
