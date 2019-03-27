package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

const sender = "quota@example.org"
const host = "smtp.example.org"
const pass = "supersecretpassword"

const tpl = "You're using %d bytes, %d%% of your quota!"

const quotaLimit = 1000000000 // 1 GB

var usage = make(map[string]int64)

var notify = func(msg string, username string) {
	auth := smtp.PlainAuth("", sender, pass, host)
	err := smtp.SendMail(host + ":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("sending mail to %s failed, %v", username, err)
	}
}

func bytesUsedBy(username string) int64 {
	return usage[username]
}

func CheckQuota(username string) {
	used := bytesUsedBy(username)
	percent := 100 * used / quotaLimit
	if percent >= 90 {
		msg := fmt.Sprintf(tpl, used, percent)
		notify(msg, username)
	}
}
