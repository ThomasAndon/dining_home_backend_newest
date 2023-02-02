package croner

import "strings"

func SendMsg(toUserId []string, title, content, url string) {
	userIds := strings.Join(toUserId, "|")

}
