package models

// BriefMessage 摘要信息
type BriefMessage struct {
	Code    int
	Message string
}

// --------------------------------------------------------------
// ---------------- 500xxx 服务端错误 ----------------------------
// --------------------------------------------------------------

const (
	// InternalGetBDInstanceErr InternalGetBDInstanceErr
	InternalGetBDInstanceErr = "获取数据库实例失败"
)

// ErrDataBase DB Error
var ErrDataBase = &BriefMessage{500000, "数据库出现错误，请稍后重试或者连接管理员处理"}

// ErrGenToken ErrGenToken
var ErrGenToken = &BriefMessage{500000, "token生成失败"}

// ErrGenSalt ErrGenSalt
var ErrGenSalt = &BriefMessage{500000, "salt生成失败"}

// ErrGenRsa ErrGenRsa
var ErrGenRsa = &BriefMessage{500000, "rsa生成失败"}

// ErrCreateDBData ErrCreateDBData
var ErrCreateDBData = &BriefMessage{500000, "创建数据失败"}

// ErrSearchDBData ErrSearchDBData
var ErrSearchDBData = &BriefMessage{500000, "数据查询失败"}

// ErrDelData ErrDelData
var ErrDelData = &BriefMessage{500000, "删除数据失败"}
var ErrCheckDBData = &BriefMessage{500000, "核对数据时出错"}

// ErrUpdateData ErrUpdateData
var ErrUpdateData = &BriefMessage{500000, "更新数据失败"}

// ErrTmplParse ErrTmplParse
var ErrTmplParse = &BriefMessage{500000, "解析模块错误"}

// ErrSyncNgxTmpl ErrSyncNgxTmpl
var ErrSyncNgxTmpl = &BriefMessage{500000, "同步NGINX配置出错"}

// ErrWriteNgxTmpl ErrWriteNgxTmpl
var ErrWriteNgxTmpl = &BriefMessage{500000, "写入NGINX配置文件失败"}

// ErrDisableDomain ErrDisableDomain
var ErrDisableDomain = &BriefMessage{500000, "设置域名状态为不可用失败"}

// ErrEnableDomain ErrEnableDomain
var ErrEnableDomain = &BriefMessage{500000, "设置域名状态为可用失败"}

// ErrCnameExist ErrCnameExist
var ErrCnameExist = &BriefMessage{500000, "cname已经存在"}

//
var ErrJobDataFormat = &BriefMessage{500000, "解析任务列表数据出错"}

// 数据是空的
var ErrDataIsNil = &BriefMessage{500000, "数据是nil，不能使用"}

// 数据是空的
var ErrDataMarshal = &BriefMessage{500000, "序列化数据失败"}

// 写入数据失败
var ErrDataWriteDisk = &BriefMessage{500000, "写入数据到磁盘失败"}

// 写入数据失败
var ErrReloadPrometheus = &BriefMessage{500000, "重新加载监控服务失败"}

//
var ErrReadFile = &BriefMessage{500000, "读取文件失败"}

// ErrGenUUID ErrGenUUID
var ErrGenUUID = &BriefMessage{500000, "uuid生成失败"}

// ErrMonitorApi ErrMonitorApi
var ErrMonitorApi = &BriefMessage{500000, "调用监控服务的API出现错误"}

// ErrMonitorApi ErrMonitorApi
var ErrFileList = &BriefMessage{500000, "读取文件列表失败"}

// ErrMonitorApi ErrMonitorApi
var ErrRFileContent = &BriefMessage{500000, "读取文件内容失败"}

// ErrGroupNotEmpty ErrGroupNotEmpty
var ErrGroupNotEmpty = &BriefMessage{500000, "还有IP属于该组，不允许删除"}

// ErrDataParse ErrDataParse
var ErrDataParse = &BriefMessage{500000, "解析数据出错"}

// --------------------------------------------------------------
// ---------------- 400xxx 客户端错误 ----------------------------
// --------------------------------------------------------------

// ErrPwdNotMatch ErrPwdNotMatch
var ErrPwdNotMatch = &BriefMessage{400000, "输入的用户密码不正确"}

// ErrUserNotExist ErrUserNotExist
var ErrUserNotExist = &BriefMessage{400000, "用户名不存在"}

// ErrRoleNotExist ErrRoleNotExist
var ErrRoleNotExist = &BriefMessage{400000, "角色不存在"}

// ErrTokenIsNull  ErrTokenIsNull
var ErrTokenIsNull = &BriefMessage{401000, "token验证失败"}

// ErrTokenNoFound  ErrTokenNoFound
var ErrTokenNoFound = &BriefMessage{401000, "不存在的token"}

// ErrPostData  ErrPostData
var ErrPostData = &BriefMessage{400000, "提交的信息不正确"}

// ErrDataNoExist  ErrDataNoExist
var ErrDataNoExist = &BriefMessage{400000, "指向数据不存在"}

// ErrQueryData  ErrQueryData
var ErrQueryData = &BriefMessage{400000, "查询参数不正确"}

// ErrUserExist  ErrUserExist
var ErrUserExist = &BriefMessage{400000, "用户名已经注册"}

// ErrDataExist  ErrDataExist
var ErrDataExist = &BriefMessage{400000, "数据已经存在"}

// ErrNameOrPwdEmpty ErrNameOrPwdEmpty
var ErrNameOrPwdEmpty = &BriefMessage{400000, "用户名或者密码不允许为空"}

// ErrDomainNotExist ErrDomainNotExist
var ErrDomainNotExist = &BriefMessage{400000, "指定域名不存在"}

// ErrSplitParma ErrSplitParma
var ErrSplitParma = &BriefMessage{400000, "分页参数不正确"}

// ErrCount ErrCount
var ErrCount = &BriefMessage{400000, "统计总数时出错"}

// ErrUploadFileFormName ErrUploadFileFormName
var ErrUploadFileFormName = &BriefMessage{400000, "获取上传文件表单(file)错误"}

// ErrUploadFileType ErrUploadFileType
var ErrUploadFileType = &BriefMessage{400000, "必须提供上传文件类型"}

// ErrUploadFileTypeErr ErrUploadFileTypeErr
var ErrUploadFileTypeErr = &BriefMessage{400000, "提供的上传文件类型错误"}

// ErrSaveUploadFile ErrSaveUploadFile
var ErrSaveUploadFile = &BriefMessage{400000, "保存文件错误"}

// ErrMkdir ErrMkdir
var ErrMkdir = &BriefMessage{400000, "创建文件夹失败"}

// ErrChgPassword ErrChgPassword
var ErrChgPassword = &BriefMessage{400000, "提交的密码信息有误"}

//
var ErrIPAddr = &BriefMessage{400000, "不是正确的IP地址"}
var ErrJobTypeEmpty = &BriefMessage{400000, "IP的分组类型不允许为空"}
var ErrOrderIDExist = &BriefMessage{400000, "排序号已经存在"}
var ErrUnSupport = &BriefMessage{400000, "不支持的操作类型"}
var ErrHaveDataNoAllowToDel = &BriefMessage{400000, "因存在相关联数据，不允许删除"}

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
