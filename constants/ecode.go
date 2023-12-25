package constants

import "fmt"

const ParamsValidateErr = 100007

type IECode interface {
	Error() string
	Code() int
	Message() string
}
type ECode struct {
	code int    `json:"code"`
	msg  string `json:"message"`
}

func (e ECode) Error() string {
	return fmt.Sprintf(" %s", e.Message())
}

func (e ECode) Code() int { return int(e.code) }

func (e ECode) Message() string {
	return e.msg
}

func New(code int, msg string) ECode {
	return ECode{code: code, msg: msg}
}

//code 命名规则 code一共六位
//第一位为api版本 v1 为 1
//第二 三位为 app 标识，系统默认提供公共code 为00
//后面三位为错误范围支持999种不同的错误

var (
	// code 正向私有，负向共用
	Success                 = New(200, "success")
	ServerError             = New(100001, "服务器内部错误")
	ParameterError          = New(100002, "参数错误")
	RrcCallError            = New(100003, "RPC调用错误")
	ServerAcquireLockFailed = New(100003, "获取锁资源失败，请稍后再试")
	TokenInfoParseError     = New(100004, "Token 信息解析失败，请稍后再试")
	TokenGenerateError      = New(100004, "Token 生成失败，请稍后再试")
	AppIdNoExist            = New(100005, "appId解析错误")
	UidNoExist              = New(100005, "userId解析错误")
	NoDataNeedError         = New(100006, "没有数据需要操作")
	UsernameNoExist         = New(100007, "username解析错误")
	// user error 第二位 第三位为01
	PathNoExitError     = New(100101, "path参数错误")
	UserNoExit          = New(100102, "用户不存在")
	TableMarkIdNoExist  = New(100103, "table mark 不存在")
	TableViewIdNoExist  = New(100104, "table view 不存在")
	TableViewFiledError = New(100105, "table view 组装字段错误")
	UserWxLoginErr      = New(100106, "微信登录失败")
	BindWxPoneErr       = New(100106, "绑定微信手机号失败")
	WxConfigErr         = New(100107, "微信配置错误")

	// nuser error 第二位 第三位为02
	IdMustError               = New(100201, "id 为比传参数")
	DateFormatError           = New(100202, "日期格式错误")
	InsertDataError           = New(100203, "插入数据库错误")
	HadRefundError            = New(100204, "已经退款，不能再退款")
	UnPayError                = New(100205, "未支付不能退款")
	MustIsPayBillError        = New(100205, "必须是支付单据")
	IncomeAndPayNoExist       = New(100206, "支付单据不存在")
	SmsSendError              = New(100207, "验证码发送错误")
	PayTypeError              = New(100208, "不支持的退款类型")
	SmsCodeMismatch           = New(100209, "验证码错误")
	CreateRefundWxParamsError = New(100210, "创建退款微信参错误")
	RefundCallBackError       = New(100210, "退款回调错误")
	RefundSaveError           = New(100211, "退款记录保存失败")
	PhoneNoExist              = New(100212, "手机号查询不存在")
	UpdateDataError           = New(100203, "更新数据错误")

	// KRY error 第二位 第三位为03
	ConfigTypeError       = New(100301, "配置类型不对")
	KrySignError          = New(100302, "客如云签名错误")
	KryResponseStatusFail = New(100303, "客如云接口返回状态不是200")
	AppIdOrSecretError    = New(100304, "appid或者secret 错误")
	U8ResponseStatusFail  = New(100303, "U8接口返回状态不是200")
	U8OrderHadSync        = New(100304, "u8订单已经同步")
	PointAmountZero       = New(100305, "积分为0")
	MemNoExist            = New(100306, "会员不存在")
	RepeatPushError       = New(100307, "重复推送")
	OperatorPointfail     = New(100308, "操作失败")
	MemLevelNoExist       = New(100309, "会员等级不存在")
	UserHadExist          = New(100310, "用户已经存在")
	NoCondition           = New(100311, "条件不存在")
	NoSupportSource       = New(100312, "暂未支持的来源")
	RegisterMemError      = New(100313, "注册会员失败")

	BillImportNoExist       = New(100401, "薪资关系不存在")
	BillHadDel              = New(100402, "已经被删除")
	MemCareNoExist          = New(100403, "会员护工不存在")
	UserNoExist             = New(100404, "护理员不存在")
	MustOne                 = New(100405, "必须是同一个护理员")
	BillNoExit              = New(100406, "新建没有护理员薪资单")
	DemandNoExist           = New(100407, "需求不存在")
	CreateMemberError       = New(100408, "创建会员错误")
	CreateMemberSourceError = New(100409, "创建会员关系错误")
	DataInitFail            = New(100410, "数据初始化失败")
	StoreCodeMustBeNumber   = New(100411, "store code 必须是数字")
	OrderStatusNoExist      = New(100412, "订单状态不存在")
	CodeHadExpired          = New(100413, "验证码过期")
)
