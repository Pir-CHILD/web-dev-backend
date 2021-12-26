package users

import "time"

type UserInfo struct {
	UserId           string    `json:"userId"` // char(12)
	UserName         string    `json:"userName"`
	UserPwd          string    `json:"userPwd"`
	UserType         int       `json:"userType"` // 0 系统管理员 1 普通用户
	UserXingming     string    `json:"userXingming"`
	CertiType        int       `json:"certiType"` // 0 学生证 1 身份证
	CertiNumber      string    `json:"certiNumber"`
	UserPhone        string    `json:"userPhone"` // char(11)
	UserLevel        int       `json:"userLevel"` // 0 一般 1 VIP
	UserIntroduction string    `json:"userIntroduction"`
	UserCity         string    `json:"userCity"`
	UserCommunity    int       `json:"userCommunity"` // 0 海淀区 1 朝阳区 2 昌平区
	RegisterTime     time.Time `json:"registerTime"`
	ChangeTime       time.Time `json:"changeTime"`
}

type HelpMeInfo struct {
	HelpMeId          string    `json:"helpMeId"`     // char(12)
	HelpMeUserId      string    `json:"helpMeUserId"` // char(12)
	HelpMeType        int       `json:"helpMeType"`   // 0 小时工 1 搬重物 2 上下班搭车 3 社区服务
	HelpMeTheme       string    `json:"helpMeTheme"`
	HelpMeDescription string    `json:"helpMeDescription"`
	HelpMePeople      int       `json:"helpMePeople"`
	HelpMeFinishTime  time.Time `json:"helpMeFinishTime"`
	HelpMeIntroPhoto  string    `json:"helpMeIntroPhoto"`
	HelpMeCreateTime  time.Time `json:"helpMeCreateTime"`
	HelpMeChangeTime  time.Time `json:"helpMeChangeTime"`
	HelpMeState       int       `json:"helpMeState"` // 0 已完成 1 待响应 2 已取消 3 到期未达成
}

type HelpYouInfo struct {
	HelpYouId          string    `json:"helpYouId"`     // char(12)
	HelpMeId           string    `json:"helpMeId"`      // char(12)
	HelpYouUserId      string    `json:"helpYouUserId"` // char(12)
	HelpYouDescription string    `json:"helpYouDescription"`
	HelpYouCreateTime  time.Time `json:"helpYouCreateTime"`
	HelpYouChangeTime  time.Time `json:"helpYouChangeTime"`
	HelpYouState       int       `json:"helpYouState"` // 0 待接受 1 同意 2 拒绝 3 取消
}

type HelpMeDone struct {
	HelpMeId          string    `json:"helpMeId"`      // char(12)
	HelpMeUserId      string    `json:"helpMeUserId"`  // char(12)
	HelpYouUserId     string    `json:"helpYouUserId"` // char(12)
	HelpYouFinishTime time.Time `json:"helpYouFinishTime"`
	HelpMeUserSpend   int       `json:"helpMeUserSpend"` // 人数 * 3
	HelpYouUserSpend  int       `json:"helpYouUserSpend"`
}

type IntermediaryIncomeInfo struct {
	Month        string `json:"month"`        // YYYY-MM
	ProvinceCity string `json:"provinceCity"` // 省-市
	Community    int    `json:"community"`    // 0 海淀区 1 朝阳区 2 昌平区
	HelpMeType   int    `json:"helpMeType"`   // 0 小时工 1 搬重物 2 上下班搭车 3 社区服务
	FinishNumber int    `json:"finishNumber"`
	InterIncome  int    `json:"interIncome"`
}
