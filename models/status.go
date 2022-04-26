package models

import "fmt"

// BriefMessage 摘要信息
type BriefMessage struct {
	Code    int
	Message string
}

func (bf *BriefMessage) Append(s string) *BriefMessage {
	return &BriefMessage{
		Code:    bf.Code,
		Message: fmt.Sprintf("%s: %s", bf.Message, s),
	}
}

func (bf *BriefMessage) String() string {
	return bf.Message
}

// --------------------------------------------------------------
// ---------------- 500xxx 服务端错误 ----------------------------
// --------------------------------------------------------------

const (
	InternalGetBDInstanceErr = "获取数据库实例失败"
)

var ErrDataBase = &BriefMessage{500000, "数据库出现错误，请稍后重试或者连接管理员处理"}
var ErrGenToken = &BriefMessage{500000, "token生成失败"}
var ErrLoginExpire = &BriefMessage{500000, "登录Session已经过期"}
var ErrGenSalt = &BriefMessage{500000, "salt生成失败"}
var ErrGenRsa = &BriefMessage{500000, "rsa生成失败"}
var ErrCreateDBData = &BriefMessage{500000, "创建数据失败"}
var ErrTransaction = &BriefMessage{500000, "执行事务出错"}
var ErrSearchDBData = &BriefMessage{500000, "数据查询失败"}
var ErrDelData = &BriefMessage{500000, "删除数据失败"}
var ErrCheckDBData = &BriefMessage{500000, "核对数据时出错"}
var ErrUpdateData = &BriefMessage{500000, "更新数据失败"}
var ErrTmplParse = &BriefMessage{500000, "解析模块错误"}
var ErrSyncNgxTmpl = &BriefMessage{500000, "同步NGINX配置出错"}
var ErrWriteNgxTmpl = &BriefMessage{500000, "写入NGINX配置文件失败"}
var ErrDisableDomain = &BriefMessage{500000, "设置域名状态为不可用失败"}
var ErrEnableDomain = &BriefMessage{500000, "设置域名状态为可用失败"}
var ErrCnameExist = &BriefMessage{500000, "cname已经存在"}
var ErrJobDataFormat = &BriefMessage{500000, "解析任务列表数据出错"}
var ErrDataIsNil = &BriefMessage{500000, "数据是nil，不能使用"}
var ErrDataMarshal = &BriefMessage{500000, "序列化数据失败"}
var ErrDataWriteDisk = &BriefMessage{500000, "写入数据到磁盘失败"}
var ErrReloadPrometheus = &BriefMessage{500000, "重新加载监控服务失败"}
var ErrReadFile = &BriefMessage{500000, "读取文件失败"}
var ErrGenUUID = &BriefMessage{500000, "uuid生成失败"}
var ErrMonitorApi = &BriefMessage{500000, "调用监控服务的API出现错误"}
var ErrFileList = &BriefMessage{500000, "读取文件列表失败"}
var ErrRFileContent = &BriefMessage{500000, "读取文件内容失败"}
var ErrGroupNotEmpty = &BriefMessage{500000, "还有IP属于该组，不允许删除"}
var ErrDataParse = &BriefMessage{500000, "解析数据出错"}
var ErrUpGrader = &BriefMessage{500000, "提升GET请求为webSocket协议时出错"}
var ErrGetControlField = &BriefMessage{500000, "获取控制参数时出错"}
var ErrHaveTaskRunning = &BriefMessage{500000, "有命令正在运行，请稍后重试"}
var ErrConvertDataType = &BriefMessage{500000, "转换数据类型出错"}
var ErrHaveInstanceRunning = &BriefMessage{500000, "有一个实例正在运行，请稍后重试"}
var ErrHaveLine = &BriefMessage{500000, "此机房下还有线路存在，不允许直接删除"}
var ErrReset = &BriefMessage{500000, "重置出现错误"}

// --------------------------------------------------------------
// ---------------- 400xxx 客户端错误 ----------------------------
// --------------------------------------------------------------

var ErrPwdNotMatch = &BriefMessage{400000, "输入的用户密码不正确"}
var ErrUserNotExist = &BriefMessage{400000, "用户名不存在"}
var ErrUserDisabled = &BriefMessage{400000, "账号已经被禁用"}
var ErrGroupDisabled = &BriefMessage{400000, "账号所属用户组已经被禁用"}
var ErrPassword = &BriefMessage{400000, "密码错误"}
var ErrPasswordSame = &BriefMessage{400000, "两次输入的密码不匹配"}
var ErrRoleNotExist = &BriefMessage{400000, "角色不存在"}
var ErrTokenIsNull = &BriefMessage{401000, "token验证失败"}
var ErrTokenNoFound = &BriefMessage{401000, "不存在的token"}
var ErrPostData = &BriefMessage{400000, "提交的信息不正确"}
var ErrDataNoExist = &BriefMessage{400000, "指向数据不存在"}
var ErrQueryData = &BriefMessage{400000, "查询参数不正确"}
var ErrUserExist = &BriefMessage{400000, "用户名已经注册"}
var ErrDataExist = &BriefMessage{400000, "数据已经存在"}
var ErrNameOrPwdEmpty = &BriefMessage{400000, "用户名或者密码不允许为空"}
var ErrDomainNotExist = &BriefMessage{400000, "指定域名不存在"}
var ErrSplitParma = &BriefMessage{400000, "分页参数不正确"}
var ErrCount = &BriefMessage{400000, "统计总数时出错"}
var ErrUploadFileFormName = &BriefMessage{400000, "获取上传文件表单(file)错误"}
var ErrUploadFileType = &BriefMessage{400000, "必须提供上传文件类型"}
var ErrUploadFileTypeErr = &BriefMessage{400000, "提供的上传文件类型错误"}
var ErrSaveUploadFile = &BriefMessage{400000, "保存文件错误"}
var ErrMkdir = &BriefMessage{400000, "创建文件夹失败"}
var ErrChgPassword = &BriefMessage{400000, "提交的密码信息有误"}
var ErrIPAddr = &BriefMessage{400000, "不是正确的IP地址"}
var ErrJobTypeEmpty = &BriefMessage{400000, "IP的分组类型不允许为空"}
var ErrOrderIDExist = &BriefMessage{400000, "排序号已经存在"}
var ErrUnSupport = &BriefMessage{400000, "不支持的操作类型"}
var ErrHaveDataNoAllowToDel = &BriefMessage{400000, "因存在相关联数据，不允许删除"}
var ErrHaveDataNoAllowToDisabled = &BriefMessage{400000, "因存在相关联数据，不允许禁用"}
var ErrUploadFile = &BriefMessage{400000, "未找到上传文件的实例"}
var ErrTooLarge = &BriefMessage{400000, "文件超过限定大小"}
var ErrFileFormat = &BriefMessage{400000, "文件格式错误"}
var ErrParseFileToYaml = &BriefMessage{400000, "解析文件成Yaml格式数据时，出现错误"}
var ErrNoResetKey = &BriefMessage{400000, "重置KEY还没有生成"}
var ErrResetKeyDiff = &BriefMessage{400000, "重置KEY不匹配"}
var ErrAlreadyRunning = &BriefMessage{400000, "服务正在运行，请稍后重试"}
var ErrNoPrivRequest = &BriefMessage{400000, "未授权的访问"}
var ErrNoVaildData = &BriefMessage{400000, "未找到有效的数据"}
var ErrNoDefined = &BriefMessage{400000, "没有定义URL"}
var ErrIPAddrPost = &BriefMessage{400000, "提交的IP地址不正确"}

// --------------------------------------------------------------
// ---------------- 300xxx 链接错误 ------------------------------
// --------------------------------------------------------------

// --------------------------------------------------------------
// ---------------- 200xxx 成功 ---------------------------------
// --------------------------------------------------------------

// Success Success
var Success = &BriefMessage{200000, "成功"}

// --------------------------------------------------------------
// ---------------- 100xxx 不知道 -------------------------------
// --------------------------------------------------------------
