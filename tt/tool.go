package tt

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"math/rand"
	"reflect"
	"runtime/debug"
	"time"
)

func T1() {
	glog.Debug("T1", gtime.Datetime())
	gcache.Set("T1.Start", time.Now(), 0) // 改成先进后出, 剥洋葱的方式, 嵌套多个
}

func T2() {
	v := gcache.Get("T1.Start")
	elapsed := time.Now().Sub(gconv.Time(v))
	glog.Debug("T2 elapsed = ", elapsed)
}
func Sleep(ms time.Duration) {
	glog.Debug("Sleep", time.Millisecond*time.Duration(ms))
	time.Sleep(time.Millisecond * time.Duration(ms))
}
func TryExpect() {
	if e := recover(); e != nil {
		glog.Errorf("%s: %s", e, debug.Stack()) // line 20
	}
}
func Reverse(s []interface{}) []interface{} {
	//https://github.com/golang/go/wiki/SliceTricks#reversing
	a := make([]interface{}, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func ReverseString(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}
func IsPointer(i interface{}) {
	if err := gvalid.Check(reflect.TypeOf(i), `regex:^\*`, "is not Pointer"); err != nil {
		glog.Fatal(err.Map())
	}
}

func Show(i interface{}) {
	fmt.Printf("Type:(%T) +Value=(%+v)\n", i, i)
	if i == nil {
	} else if GetType(i) == GetType("") {
	} else if GetType(i) == GetType(errors.New("")) {
	} else if GetType(i) == GetType(1) {
	} else if GetType(i) == GetType(1.0) {
	} else {
		fmt.Printf("Value=%v\n", i)
		fmt.Printf("+Value=%s\n\n", gconv.String(i))
	}
}
func ShowType(i interface{}) {
	fmt.Printf("Type=%T\n", i)
	//fmt.Println(reflect.TypeOf(i))
}
func GetType(i interface{}) reflect.Type {
	//fmt.Printf("%T\n", i)
	return reflect.TypeOf(i)
}

func RandomString(length int) string {
	// https://learnku.com/go/t/36699
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
