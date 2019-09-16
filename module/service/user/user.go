package user

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/os/gtime"
)

const (
	//UserSessionMark 用户session信息标记
	UserSessionMark = "user_info"
)

var (
	// 表对象
	table = g.DB().Table("neu_user").Safe()
)

//SignUp 用户注册
func SignUp(data g.MapStrStr) error {
	// 数据校验
	rules := []string{
		"account @required|length:6,16#账号不能为空|账号长度应当在:min到:max之间",
		"password2@required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间",
		"password @required|length:6,16|same:password2#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等",
	}
	if e := gvalid.CheckMap(data, rules); e != nil {
		return errors.New(e.String())
	}
	if _, ok := data["nickname"]; !ok {
		data["nickname"] = data["account"]
	}
	// 唯一性数据检查
	if !Checkaccount(data["account"]) {
		return fmt.Errorf("账号 %s 已经存在", data["account"])
	}
	if !CheckNickName(data["nickname"]) {
		//return errors.New(fmt.Sprintf("昵称 %s 已经存在", data["nickname"]))
		return fmt.Errorf("昵称 %s 已经存在", data["nickname"])
	}
	// 记录账号创建/注册时间
	if _, ok := data["create_time"]; !ok {
		data["create_time"] = gtime.Now().String()
	}
	if _, err := table.Filter().Data(data).Save(); err != nil {
		return err
	}
	return nil
}

//IsSignedIn 判断用户是否已经登录
func IsSignedIn(session *ghttp.Session) bool {
	return session.Contains(UserSessionMark)
}

//SignIn 用户登录，成功返回用户信息，否则返回nil;
func SignIn(account, password string, session *ghttp.Session) error {

	//判断登录使用的是用户名、手机号、邮箱
	accountField := "username"
	if VerifyMobileFormat(account) {
		accountField = "mobile"
	}
	if VerifyEmailFormat(account) {
		accountField = "email"
	}

	record, err := table.Where(accountField+"=?", account).One()
	if err != nil {
		return err
	}
	if record == nil {
		return errors.New("账号不正确")
	}
	if record["status"].String() != "normal" {
		return errors.New("账号已经被锁定")
	}

	dataPassword := record["password"].String()
	salt := record["salt"].String()
	if dataPassword != getEncryptPassword(password, salt) {
		return errors.New("密码不正确")
	}

	session.Set(UserSessionMark, record)
	return nil
}

//SignOut 用户注销
func SignOut(session *ghttp.Session) {
	session.Remove(UserSessionMark)
}

//Checkaccount 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func Checkaccount(account string) bool {
	if _, err := table.Where("account", account).Count(); err != nil {
		panic(err)
		return false
	}
	return true

}

//CheckNickName 检查昵称是否符合规范(目前仅检查唯一性),存在返回false,否则true
func CheckNickName(nickname string) bool {
	if _, err := table.Where("nickname", nickname).Count(); err != nil {
		return false
	}
	return true
}

//VerifyEmailFormat email verify
func VerifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//VerifyMobileFormat mobile verify
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

/**
 * 获取密码加密后的字符串
 * @param string $password 密码
 * @param string $salt     密码盐
 * @return string
 */
func getEncryptPassword(password, salt string) string {
	tmpcode, _ := gmd5.Encrypt(password)
	retpassword, _ := gmd5.Encrypt(tmpcode + salt)
	return retpassword
}
