package constants

const (
	DEFAULT_CONFIG_PATH = "/app/conf/config-dev.yaml"

	SUCCESS_CODE                = "1"                       //接口请求成功码
	SUCCESS_CODE_ZY             = "200"                     //接口请求成功码
	ERR_REQUEST_PARAM           = "-1"                      //请求参数错误
	ERR_RESPONSE                = "0"                       //响应错误
	ERR_REQUEST_NULL            = "301"                     //参数为空
	ERR_DATA_NULL               = "302"                     //无结果
	ERR_REQUEST_INVALID         = "303"                     //无效的 JSON 数据
	ENV_KUBERNETES_SERVICE_HOST = "KUBERNETES_SERVICE_HOST" //k8s当前服务host
	ERR_SERVICE                 = 500                       //无结果
	ERR_NOT_FOUND               = 331                       //内容未找到
	ERR_CONTENT_UNIQUE          = 332                       //内容重复
)
