
# 生成api服务
.PHONY: swag swagLocal

swag:
    # swag 生成 json 文档并上传 yapi
	swag init --generalInfo pkg/swagger/swag.go --output data/swagger && yapi import --config `pwd`/pkg/swagger/yapi-import.json

swagLocal:
    # swag 生成 json文档
	swag init --generalInfo pkg/swagger/swag.go --output data/swagger
