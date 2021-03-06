package model

import (
	"github.com/ZRothschild/goIris/app/model/base"
	"time"
)

// 用户
type User struct {
	base.Base
	AuthId        string    `gorm:"not null;default:'';type:varchar(60);column:auth_id;comment:认证ID" json:"authId"`
	Nickname      string    `gorm:"not null;default:'';type:varchar(60);column:nickname;comment:昵称" json:"nickname"`
	Email         string    `gorm:"not null;default:'';type:varchar(60);column:email;comment:用户email;index" json:"email"`
	Password      string    `gorm:"not null;default:'';type:varchar(90);column:password;comment:密码" json:"password"`
	HeadImg       string    `gorm:"not null;default:'';type:varchar(80);column:head_tmg;comment:头像" json:"headImg"`
	Phone         string    `gorm:"not null;default:'';type:varchar(20);column:phone;comment:手机号码" json:"phone"`
	Autograph     string    `gorm:"not null;default:'';type:varchar(255);column:autograph;comment:个新签名" json:"autograph"`
	Github        string    `gorm:"not null;default:'';type:varchar(60);column:github;comment:GitHub 账号" json:"github"`
	Wechat        string    `gorm:"not null;default:'';type:varchar(60);column:wechat;comment:微信号" json:"wechat"`
	Alipay        string    `gorm:"not null;default:'';type:varchar(60);column:alipay;comment:支付宝号" json:"alipay"`
	School        string    `gorm:"not null;default:'';type:varchar(60);column:school;comment:毕业学校" json:"school"`
	MyWeb         string    `gorm:"not null;default:'';type:varchar(60);column:my_web;comment:自己的网站" json:"myWeb"`
	City          string    `gorm:"not null;default:'';type:varchar(60);column:city;comment:所在城市" json:"city"`
	Company       string    `gorm:"not null;default:'';type:varchar(60);column:company;comment:所在公司" json:"company"`
	Sign          string    `gorm:"not null;default:'';type:varchar(60);column:sign;comment:签名信息" json:"sign"`
	Sex           uint8     `gorm:"not null;default:1;column:sex;comment:1 未定 2 男 3 女" json:"sex"`
	VerifyType    uint8     `gorm:"not null;default:1;column:verify_type;comment:暂定 1 默认值" json:"verifyType"`
	AuthType      uint8     `gorm:"not null;default:1;column:auth_type;comment:1 未认证  2 github认证 3 qq认证 4 微信认证" json:"authType"`
	Status        uint8     `gorm:"not null;default:1;column:status;comment:1 正常 2 非正常用户" json:"status"`
	GithubId      uint64    `gorm:"not null;default:0;column:github_id;comment:github id" json:"githubId"`
	Experience    uint64    `gorm:"not null;default:0;column:experience;comment:用户经验值" json:"experience"`
	Currency      uint64    `gorm:"not null;default:0;column:currency;comment:用户当前艾币数量" json:"currency"`
	TotalCurrency uint64    `gorm:"not null;default:0;column:total_currency;comment:用户总共获得的艾币数量" json:"totalCurrency"`
	Level         uint64    `gorm:"not null;default:0;column:level;comment:用户等级" json:"level"`
	AttentionNums uint64    `gorm:"not null;default:0;column:attention_nums;comment:被关注数" json:"attentionNums"`
	LoginTime     time.Time `gorm:"type:date;column:login_time;comment:登录时间" json:"loginTime"`
	Birthday      time.Time `gorm:"type:date;column:birthday;comment:生日" json:"birthday"`
}

// 用户示例
func NewUser() *User {
	return &User{}
}

// 表名
func (m *User) TableName() string {
	return "users"
}
