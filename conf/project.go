package conf

import (
	"sync"
	"time"
)

const SysTimeForm = "2006-01-02 15:04:05"
const SysTimeFormShort = "2006-01-02"

var SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")

var SignSecret = []byte("1234567890abcedfg")

var CookieSecret = "hellolottery"

var LoginUser = new(sync.Map)
