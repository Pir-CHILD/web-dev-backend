# "好社区"各类信息标识
## 开发业务模型及对应的后台数据库表简要说明
### 用户类
**USERINFO** - Table
- 用户标识$*$
`CHAR(12)`
- 用户名
`VARCHAR`
- 登录密码
    - 不小于$6$位
    - 不少于$2$个数字
    - 不能全部大\小写
`VARCHAR`
- 用户类型（0：系统管理员、1：普通用户）
`INT`
- 用户姓名
`VARCHAR`
- 证件类型（0：学生证、1：身份证）
`INT`
- 证件号码
`VARCHAR`
- 手机号码（11位数字）
`CHAR(11)`
- 用户级别（0：一般、1：VIP）
`INT`
- 用户简介
`TEXT`
- 注册城市（北京）
`VARCHAR`
- 注册社区（0：海淀区、1：朝阳区、2：昌平区）
`INT`
- 注册时间
`DATETIME`
- 修改时间
`DATETIME`
> 在后台建立一个管理员，用户名：admin, 密码: admin
### "劳您驾"请求信息类
**HELPME** - Table
- 请求标识$*$
`CHAR(12)`
- 发布用户标识
`CHAR(12)`
- 请求类型（0：小时工、1：搬重物、2：上下班搭车、3：社区服务自愿者等）
`INT`
- 请求主题名称
`VARCHAR`
- 请求描述
`TEXT`
- 请求人数
`INT`
- 请求结束日期
`DATETIME`
- 请求介绍照片（可空）
`VARCHAR`
- 创建时间
`DATETIME`
- 修改时间
`DATETIME`
- 状态（0：已完成、1：待响应、2：已取消、3：到期未达成）
`INT`
### 响应类
**HELPYOU** - Table
- 响应标识$*$
`CHAR(12)`
- 请求标识
`CHAR(12)`
- 响应用户标识
`CHAR(12)`
- 响应描述
`TEXT`
- 创建时间
`DATETIME`
- 修改时间
`DATETIME`
- 状态（0：待接受、1：同意、2：拒绝、3：取消）
`INT`
### "劳您驾"帮忙成功明细表
**HELPMEDONE** - Table
- 请求标识$*$
`CHAR(12)`
- 发布用户标识
`CHAR(12)`
- 响应用户标识
`CHAR(12)`
- 达成日期
`DATETIME`
- 发布者支付中介费（人数*3元）
`INT`
- 响应者支付中介费（1元）
`INT`
### 中介收益汇总表
**INTERMEDIARYINCOME** - Table
- 月份（YYYY-MM）
`DATETIME`
- 地域（省-市）
`VARCHAR`
- 社区
`INT`
- 请求帮忙类型
`INT`
- 达成笔数
`INT`
- 中介费收入金额
`INT`