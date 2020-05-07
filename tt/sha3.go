package tt

import (
	"encoding/hex"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"golang.org/x/crypto/sha3"
	"io"
	"os"
)

func Sha3(v interface{}) string {
	r := sha3.Sum256(gconv.Bytes(v))
	return hex.EncodeToString(r[:])
}

func Sha3File(path string) string {
	f, err := os.Open(path)
	if err != nil {
		glog.Fatal(err)
	}
	defer f.Close()
	h := sha3.New256()
	_, err = io.Copy(h, f)
	if err != nil {
		glog.Fatal(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}
