package community

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	communityRes "github.com/flipped-aurora/gin-vue-admin/server/model/community/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type CommunityBaseApi struct{}

// Captcha
// @Tags      communityBase
// @Summary   生成验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=communityRes.CaptchaResponse,msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router    /communityBase/captcha [post]
func (b *CommunityBaseApi) Captcha(c *gin.Context) {
	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	key := c.ClientIP()
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
		oc = true
	}
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		global.GVA_LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(communityRes.CaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.GVA_CONFIG.Captcha.KeyLong,
		OpenCaptcha:   oc,
	}, "验证码获取成功", c)
}

// SendCode
// @Tags      communityBase
// @Summary   发送短信验证码
// @Security  ApiKeyAuth
// @accept    application/x-www-form-urlencoded
// @Produce   application/json
// @Param     phone  formData   string   true  "手机号"
// @Param     type   formData   string   true  "类型：Register,PhoneCodeLogin"
// @Success   200  {object}  response.Response{data=communityRes.CodeResponse,msg=string}  "发送短信验证码"
// @Router    /communityBase/sendCode [post]
func (b *CommunityBaseApi) SendCode(c *gin.Context) {
	phone := c.PostForm("phone")
	codeType := c.PostForm("type")

	if codeType == "" {
		codeType = "PhoneCodeLogin"
	}

	if !CheckMobile(phone) {
		global.GVA_LOG.Error("手机号码格式错误!")
		response.FailWithMessage("手机号码格式错误", c)
		return
	}

	code, err := CreateCaptcha(6)
	if err != nil {
		global.GVA_LOG.Error("生成验证码失败!", zap.Error(err))
		response.FailWithMessage("生成验证码失败", c)
		return
	}

	debug := global.GVA_CONFIG.System.Debug
	key := c.ClientIP() + codeType + phone
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, code, time.Second*time.Duration(60*5))
	} else {
		code = interfaceToString(v)
	}
	if debug == 1 {
		response.OkWithDetailed(communityRes.CodeResponse{
			Phone: phone,
			Code:  code,
		}, "发送成功", c)
	}

	// TODO 发送验证码
}

func verifyCode(c *gin.Context, phone string, code string, codeType string) response.Response {
	key := c.ClientIP() + codeType + phone
	v, ok := global.BlackCache.Get(key)
	if !ok {
		return response.Response{
			Code: -1,
			Data: "",
			Msg:  "验证码无效",
		}
	}

	if code != interfaceToString(v) {
		return response.Response{
			Code: -2,
			Data: "",
			Msg:  "验证码错误",
		}
	}

	global.BlackCache.Set(key, 0, time.Second*time.Duration(1))

	return response.Response{
		Code: 0,
	}
}

func CreateCaptcha(num int) (string, error) {
	str := "1"
	for i := 0; i < num; i++ {
		str += strconv.Itoa(0)
	}
	str10 := str
	int10, err := strconv.ParseInt(str10, 10, 32)
	if err != nil {
		return "", err
	} else {
		j := int32(int10)
		return fmt.Sprintf("%0"+strconv.Itoa(num)+"v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(j)), nil
	}
}

// CheckMobile 检验手机号
func CheckMobile(phone string) bool {
	// 匹配规则
	// ^1第一位为一
	// [345789]{1} 后接一位345789 的数字
	// \\d \d的转义 表示数字 {9} 接9位
	// $ 结束符
	regRuler := "^1[345789]{1}\\d{9}$"

	// 正则调用规则
	reg := regexp.MustCompile(regRuler)

	// 返回 MatchString 是否匹配
	return reg.MatchString(phone)

}

// CheckIdCard 检验身份证
func CheckIdCard(card string) bool {
	//18位身份证 ^(\d{17})([0-9]|X)$
	// 匹配规则
	// (^\d{15}$) 15位身份证
	// (^\d{18}$) 18位身份证
	// (^\d{17}(\d|X|x)$) 18位身份证 最后一位为X的用户
	regRuler := "(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)"

	// 正则调用规则
	reg := regexp.MustCompile(regRuler)

	// 返回 MatchString 是否匹配
	return reg.MatchString(card)
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}

// 类型转换
func interfaceToString(v interface{}) (s string) {
	switch v := v.(type) {
	case string:
		s = v
	default:
		s = ""
	}
	return
}
