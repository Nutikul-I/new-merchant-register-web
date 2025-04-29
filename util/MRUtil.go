package util

import (
	"math/rand"
	"register-service/model"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

func GenMerchantUID(count int) string {

	var passwordData string

	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano())
		digit := rand.Intn(10) // Generates a random number between 0 and 9
		passwordData += string('0' + digit)
	}

	return passwordData
}

func GenPassword(count int) string {

	rand.Seed(time.Now().UnixNano())

	var alphabet []rune = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	password := randomString(8, alphabet)

	return password
}

func randomString(n int, alphabet []rune) string {

	alphabetSize := len(alphabet)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		ch := alphabet[rand.Intn(alphabetSize)]
		sb.WriteRune(ch)
	}

	s := sb.String()
	return s
}

func GetChannel() ([]model.ChannelMerchant, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	var channel = []model.ChannelMerchant{}

	txt_channel := viper.GetString("CHANNEL_MERCHANT")

	res1 := strings.Split(txt_channel, "#")

	for _, s := range res1 {
		log.Debugf("s ==> %#v", s)
		res2 := strings.Split(s, "$$$")
		res3 := strings.Split(res2[0], "$.$")
		for _, e := range res3 {
			var data = model.ChannelMerchant{}
			data.ChannelCode = e
			data.ChannelType = res2[1]

			channel = append(channel, data)
		}
	}

	return channel, nil
}
