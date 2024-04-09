package pwd

import (
	"crypto/md5"
	"fmt"
	"github.com/scott-x/pwd/util"
	"strings"
)

var options *util.Options

func init() {
	//custom options for password
	options = &util.Options{10, 10000, 50, md5.New}
}

// password
func GenPWD(pwdFromUser string) string {
	// Using custom options
	salt, encodedPwd := util.Encode(pwdFromUser, options)
	return fmt.Sprintf("%s_%s", salt, encodedPwd)
}

func VerifyPWD(pwdFromUser, pwdFromDB string) bool {
	arr := strings.Split(pwdFromDB, "_")
	return util.Verify(pwdFromUser, arr[0], arr[1], options)
}
