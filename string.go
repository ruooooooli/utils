/*
* @Author               : ruoli
* @Email                : ruooooooli@gmail.com
* @Date                 : 2016-01-19 18:26:46
* @Last Modified by     : ruoli
* @Last Modified time   : 2016-01-19 18:34:29
 */
package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5(str string) string {
	hashMD5 := md5.New()
	io.WriteString(hashMD5, str)
	return fmt.Sprintf("%x", hashMD5.Sum(nil))
}

/* End of file  : string.go */
/* Location     : /string.go */
