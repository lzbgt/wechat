// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package wechat

import (
	"errors"
	"strconv"
	"strings"
)

// 获取微信客户端的版本.
//  @userAgent: 微信内置浏览器的 user-agent;
//  @version:   微信客户端的版本;
//  @err:       错误信息.
func WXVersion(userAgent string) (version float64, err error) {
	// Mozilla/5.0(iphone;CPU iphone OS 5_1_1 like Mac OS X) AppleWebKit/534.46(KHTML,like Geocko)Mobile/9B206 MicroMessenger/5.0
	lastSlashIndex := strings.LastIndex(userAgent, "/")
	versionIndex := lastSlashIndex + 1

	if lastSlashIndex == -1 || versionIndex == len(userAgent) {
		err = errors.New("不是有效的微信浏览器 user-agent")
		return
	}

	version, err = strconv.ParseFloat(userAgent[versionIndex:], 64)
	if err != nil {
		err = errors.New("不是有效的微信浏览器 user-agent")
		return
	}
	return
}