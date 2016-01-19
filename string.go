/*
* @Author               : ruoli
* @Email                : ruooooooli@gmail.com
* @Date                 : 2016-01-19 18:26:46
* @Last Modified by     : ruoli
* @Last Modified time   : 2016-01-19 18:49:05
 */
package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func MD5(str string) string {
	hashMD5 := md5.New()
	io.WriteString(hashMD5, str)
	return fmt.Sprintf("%x", hashMD5.Sum(nil))
}

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

/* End of file  : string.go */
/* Location     : /string.go */
