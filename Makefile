# 默认执行 all 目标
.DEFAULT_GOAL := all
PHONY: all
all: go.tidy

# Makefile

GO=go

# 定义函数以生成路径
define generate_paths
$(1)RPC_PATH=./service/$(1)/rpc
$(1)API_PATH=./service/$(1)/api
$(1)SRC_FILE=$(1)
endef

userrpc:
	$(eval $(call generate_paths,user))
	$(GO) run $(userRPC_PATH)/$(userSRC_FILE).go -f $(userRPC_PATH)/etc/$(userSRC_FILE).yaml


userapi:
	$(eval $(call generate_paths,user))
	$(GO) run ${userAPI_PATH}/${userSRC_FILE}.go -f ${userAPI_PATH}/etc/${userSRC_FILE}.yaml

run_user: userrpc userapi

productrpc:
	$(eval $(call generate_paths,product))
	$(GO) run ${productRPC_PATH}/${productSRC_FILE}.go -f ${productRPC_PATH}/etc/${productSRC_FILE}.yaml

productapi:
	$(eval $(call generate_paths,product))
	$(GO) run ${productAPI_PATH}/${productSRC_FILE}.go -f ${productAPI_PATH}/etc/${productSRC_FILE}.yaml

run_product: productrpc productapi

payrpc:
	$(eval $(call generate_paths,pay))
	$(GO) run ${payRPC_PATH}/${paySRC_FILE}.go -f ${payRPC_PATH}/etc/${paySRC_FILE}.yaml

payapi:
	$(eval $(call generate_paths,pay))
	$(GO) run ${payAPI_PATH}/${paySRC_FILE}.go -f ${payAPI_PATH}/etc/${paySRC_FILE}.yaml

run_pay: payrpc payapi

orderrpc:
	$(eval $(call generate_paths,order))
	$(GO) run ${orderRPC_PATH}/${orderSRC_FILE}.go -f ${orderRPC_PATH}/etc/${orderSRC_FILE}.yaml

orderapi:
	$(eval $(call generate_paths,order))
	$(GO) run ${orderAPI_PATH}/${orderSRC_FILE}.go -f ${orderAPI_PATH}/etc/${orderSRC_FILE}.yaml

run_order: orderrpc orderapi


