package tt

import (
	"github.com/gogf/gf/encoding/gcompress"
	"github.com/gogf/gf/encoding/gjson"
	//"github.com/vmihailenco/msgpack/v4"
)

func Pack(src interface{}) (data []byte, err error) {
	//data, err = msgpack.Marshal(src)
	data, err = gjson.Encode(src)
	if err != nil {
		return
	}
	data, err = gcompress.Zlib(data)
	if err != nil {
		return
	}
	//glog.Debugf("%.2f%%", 100*float64(len(data))/float64(len(gconv.String(src))))
	return
}

func Unpack(src []byte) (value interface{}, err error) {
	err = UnpackTo(src, &value)
	return
}

func UnpackTo(src []byte, item interface{}) (err error) {
	IsPointer(item)
	var data []byte
	data, err = gcompress.UnZlib(src)
	if err != nil {
		return err
	}
	err = gjson.DecodeTo(data, item) //err := msgpack.Unmarshal(b3, item)
	if err != nil {
		return err
	}
	return err
}
