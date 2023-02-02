package util

import "testing"

func TestGetLatestAT(t *testing.T) {
	redis, err := getLatestAccessTokenToRedis("wwfc9c5ccc385b29d0", "Rl0DapBeVT6GRhC0UYUhwq-vdHJAjSDfPScysdc7lLc")
	if err != nil {
		return
	}
	t.Log(redis)
}

func TestBase64(t *testing.T) {

}
