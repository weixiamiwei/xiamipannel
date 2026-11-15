package global

import "os"

// FREE_LOGIN 免登录模式（自定义分支特性）
// 通过环境变量 ANGE_FREE_LOGIN=true 开启：
//  1. 登录接口忽略账号密码，自动以第一个启用用户（优先管理员）身份签发 token
//  2. LoginInterceptor 对所有受保护接口自动放行
//  3. 前端登录页检测到 freeLogin 标志后自动登录跳转
//
// 关闭环境变量即可完全恢复原始登录行为。
var FREE_LOGIN = false

func init() {
	switch os.Getenv("ANGE_FREE_LOGIN") {
	case "true", "1", "TRUE", "True":
		FREE_LOGIN = true
	}
}
