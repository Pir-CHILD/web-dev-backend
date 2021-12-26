# "好社区"各类信息标识
## 开发业务模型及对应的后台数据库表简要说明
### 用户类
**USERINFO** - Table
- 用户标识$*$
*userId*
`CHAR(12)`
- 用户名
*userName*
`VARCHAR`
- 登录密码
    - 不小于$6$位
    - 不少于$2$个数字
    - 不能全部大\小写
*userPwd*
`VARCHAR`
- 用户类型（0：系统管理员、1：普通用户）
*userType*
`INT`
- 用户姓名
*userXingming*
`VARCHAR`
- 证件类型（0：学生证、1：身份证）
*certiType*
`INT`
- 证件号码
*certiNumber*
`VARCHAR`
- 手机号码（11位数字）
*userPhone*
`CHAR(11)`
- 用户级别（0：一般、1：VIP）
*userLevel*
`INT`
- 用户简介
*userIntroduction*
`TEXT`
- 注册城市（北京）
*userCity*
`VARCHAR`
- 注册社区（0：海淀区、1：朝阳区、2：昌平区）
*userCommunity*
`INT`
- 注册时间
*registerTime*
`DATETIME`
- 修改时间
*changeTime*
`DATETIME`
> 在后台建立一个管理员，用户名：admin, 密码: admin
### "劳您驾"请求信息类
**HELPME** - Table
- 请求标识$*$
*helpMeId*
`CHAR(12)`
- 发布用户标识
*helpMeUserId*
`CHAR(12)`
- 请求类型（0：小时工、1：搬重物、2：上下班搭车、3：社区服务自愿者等）
*helpMeType*
`INT`
- 请求主题名称
*helpMeTheme*
`VARCHAR`
- 请求描述
*helpMeDescription*
`TEXT`
- 请求人数
*helpMePeople*
`INT`
- 请求结束日期
*helpMeFinishTime*
`DATETIME`
- 请求介绍照片（可空）
*helpMeIntroPhoto*
`VARCHAR`
- 创建时间
*helpMeCreateTime*
`DATETIME`
- 修改时间
*helpMeChangeTime*
`DATETIME`
- 状态（0：已完成、1：待响应、2：已取消、3：到期未达成）
*helpMeState*
`INT`
### 响应类
**HELPYOU** - Table
- 响应标识$*$
*helpYouId*
`CHAR(12)`
- 请求标识
*helpMeId*
`CHAR(12)`
- 响应用户标识
*helpYouUserId*
`CHAR(12)`
- 响应描述
*helpYouDescription*
`TEXT`
- 创建时间
*helpYouCreateTime*
`DATETIME`
- 修改时间
*helpYouChangeTime*
`DATETIME`
- 状态（0：待接受、1：同意、2：拒绝、3：取消）
*helpYouState*
`INT`
### "劳您驾"帮忙成功明细表
**HELPMEDONE** - Table
- 请求标识$*$
*helpMeId*
`CHAR(12)`
- 发布用户标识
*helpMeUserId*
`CHAR(12)`
- 响应用户标识
*helpYouUserId*
`CHAR(12)`
- 达成日期
*helpYouFinishTime*
`DATETIME`
- 发布者支付中介费（人数*3元）
*helpMeUserSpend*
`INT`
- 响应者支付中介费（1元）
*helpYouUserSpend*
`INT`
### 中介收益汇总表
**INTERMEDIARYINCOME** - Table
- 月份（YYYY-MM）$*$
*month*
`CHAR(7)`
- 地域（省-市）
*provinceCity*
`VARCHAR`
- 社区
*community*
`INT`
- 请求帮忙类型
*helpMeType*
`INT`
- 达成笔数
*finishNumber*
`INT`
- 中介费收入金额
*interIncome*
`INT`