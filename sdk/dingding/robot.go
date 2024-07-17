package dingding

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/hopeio/utils/net/http/client"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	//VERSION is SDK version
	VERSION = "0.1"

	//ROOT is the root url
	ROOT = "https://oapi.dingtalk.com/"
)

func SendRobotMessage(accessToken, secret, title, content string, msgType MsgType) error {
	signUrl, err := RobotUrl(accessToken, secret)
	if err != nil {
		return err
	}
	body := strings.NewReader(msgType.Body(title, content))

	return client.Post(ROOT+signUrl, body, nil)
}

func RobotUrl(accessToken, secret string) (string, error) {
	if accessToken == "" {
		return "", errors.New("token不能为为空")
	}
	if secret != "" {
		// 密钥加签处理
		now := time.Now().UnixNano() / int64(time.Millisecond)
		timestampStr := strconv.FormatInt(now, 10)
		h := hmac.New(sha256.New, []byte(secret))
		h.Write([]byte(timestampStr + "\n" + secret))
		sum := h.Sum(nil)
		return fmt.Sprintf("robot/send?access_token=%s&timestamp=%s&sign=%s", accessToken, timestampStr, url.QueryEscape(base64.StdEncoding.EncodeToString(sum))), nil
	}
	return fmt.Sprintf("robot/send?access_token=%s", accessToken), nil
}

// SendRobotTextMessage can send a text message to a group chat
func SendRobotTextMessage(accessToken string, content string) error {
	return SendRobotMessage(accessToken, "", "", content, MsgTypeText)
}

func SendRobotTextMessageWithSecret(accessToken, secret, content string) error {
	if secret == "" {
		return errors.New("secret不能为空")
	}
	return SendRobotMessage(accessToken, secret, "", content, MsgTypeText)
}

func SendRobotMarkDownMessage(token, title, content string) error {
	return SendRobotMessage(token, "", title, content, MsgTypeMarkdown)
}

func SendRobotMarkDownMessageWithSecret(token, secret, title, content string) error {
	if secret == "" {
		return errors.New("secret不能为空")
	}
	return SendRobotMessage(token, secret, title, content, MsgTypeMarkdown)
}
