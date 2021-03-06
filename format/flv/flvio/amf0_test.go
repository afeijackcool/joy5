package flvio

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, a, b interface{}, args ...interface{}) {
	equal := func() bool {
		av := reflect.ValueOf(a)
		if av.Kind() == reflect.Func {
			return fmt.Sprint(a) == fmt.Sprint(b)
		}
		return reflect.DeepEqual(a, b)
	}()
	if !equal {
		t.Fail()
	}
}

func TestDecodeMetaData(t *testing.T) {
	data := []byte{
		0x02, 0x00, 0x0d, 0x40, 0x73, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x46, 0x72, 0x61, 0x6d, 0x65,

		0x02, 0x00, 0x0a, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61,

		0x08, 0x00, 0x00, 0x00, 0x0d,

		0x00, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x00, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68,
		0x00, 0x40, 0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x00, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
		0x00, 0x40, 0x86, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x00, 0x0d, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x72, 0x61, 0x74, 0x65,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x00, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65,
		0x00, 0x40, 0x4d, 0xf8, 0x53, 0xe2, 0x55, 0x6b, 0x28,

		0x00, 0x0c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x69, 0x64,
		0x00, 0x40, 0x1c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x00, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x72, 0x61, 0x74, 0x65,
		0x00, 0x40, 0x57, 0x58, 0x90, 0x00, 0x00, 0x00, 0x00,

		0x00, 0x0f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x65,
		0x00, 0x40, 0xe7, 0x70, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x00, 0x0f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x69, 0x7a, 0x65,
		0x00, 0x40, 0x30, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x00, 0x06, 0x73, 0x74, 0x65, 0x72, 0x65, 0x6f,
		0x01, 0x01,

		0x00, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x69, 0x64,
		0x00, 0x40, 0x24, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x00, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72,
		0x02, 0x00, 0x0d, 0x4c, 0x61, 0x76, 0x66, 0x35, 0x36, 0x2e, 0x33, 0x36, 0x2e, 0x31, 0x30, 0x30,

		0x00, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x69, 0x7a, 0x65,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x00, 0x00,
		0x09,
	}

	var n int = 0
	var start int = 0
	val, err := parseAMF0Val(0, data, &start)
	assertEqual(t, err, nil)
	assertEqual(t, val.(string), "@setDataFrame")
	n = start

	start = 0
	val, err = parseAMF0Val(0, data[n:], &start)
	assertEqual(t, err, nil)
	assertEqual(t, val.(string), "onMetaData")
	n += start

	start = 0
	val, err = parseAMF0Val(0, data[n:], &start)

	assertEqual(t, int(val.(AMFMap).Get("duration").V.(float64)), 0)
	assertEqual(t, int(val.(AMFMap).Get("width").V.(float64)), 1280)
	assertEqual(t, int(val.(AMFMap).Get("height").V.(float64)), 720)
	assertEqual(t, int(val.(AMFMap).Get("videodatarate").V.(float64)), 0)
	assertEqual(t, val.(AMFMap).Get("framerate").V.(float64), 59.94005994005994)
	assertEqual(t, int(val.(AMFMap).Get("videocodecid").V.(float64)), 7)
	assertEqual(t, val.(AMFMap).Get("audiodatarate").V, 93.3837890625)
	assertEqual(t, int(val.(AMFMap).Get("audiosamplerate").V.(float64)), 48000)
	assertEqual(t, int(val.(AMFMap).Get("audiosamplesize").V.(float64)), 16)
	assertEqual(t, val.(AMFMap).Get("stereo").V.(bool), true)
	assertEqual(t, int(val.(AMFMap).Get("audiocodecid").V.(float64)), 10)
	assertEqual(t, val.(AMFMap).Get("encoder").V.(string), "Lavf56.36.100")
	assertEqual(t, int(val.(AMFMap).Get("filesize").V.(float64)), 0)
}

func TestDecoderConnect(t *testing.T) {
	data := []byte{
		// connect ap
		0x02, 0x00, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,

		0x00, 0x3f, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x03,

		0x00, 0x03, 0x61, 0x70, 0x70,
		0x02, 0x00, 0x02, 0x61, 0x70,

		0x00, 0x04, 0x74, 0x79, 0x70, 0x65,
		0x02, 0x00, 0x0a, 0x6e, 0x6f, 0x6e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,

		0x00, 0x08, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x56, 0x65, 0x72,
		0x02, 0x00, 0x24, 0x46, 0x4d, 0x4c, 0x45, 0x2f, 0x33, 0x2e, 0x30, 0x20, 0x28, 0x63, 0x6f, 0x6d,
		0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x3b, 0x20, 0x4c, 0x61, 0x76, 0x66, 0x35, 0x36, 0x2e,
		0x31, 0x35, 0x2e, 0x31, 0x30, 0x32, 0x29,

		0x00, 0x05, 0x74, 0x63, 0x55, 0x72, 0x6c,
		0x02, 0x00, 0x1c, 0x72, 0x74, 0x6d, 0x70, 0x3a, 0x2f, 0x2f, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36,
		0x38, 0x2e, 0x31, 0x2e, 0x32, 0x33, 0x33, 0x3a, 0x31, 0x39, 0x33, 0x35, 0x2f, 0x61, 0x70,

		0x00, 0x00,
		0x09,
	}

	var n int = 0
	var start int = 0
	val, err := parseAMF0Val(0, data, &start)
	assertEqual(t, val.(string), "connect")
	assertEqual(t, err, nil)
	n += start

	start = 0
	val, err = parseAMF0Val(0, data[n:], &start)
	assertEqual(t, int(val.(float64)), 1)
	assertEqual(t, err, nil)
	n += start

	start = 0
	val, err = parseAMF0Val(0, data[n:], &start)
	assertEqual(t, err, nil)
	assertEqual(t, val.(AMFMap).Get("app").V, "ap")
	assertEqual(t, val.(AMFMap).Get("tcUrl").V, "rtmp://192.168.1.233:1935/ap")
}

func TestDecodeFromBytes(t *testing.T) {
	data := []byte{
		0x02, 0x00, 0x07, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,

		0x00, 0x3f, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x03,
		0x00, 0x06, 0x66, 0x6d, 0x73, 0x56, 0x65, 0x72,
		0x02, 0x00, 0x0d, 0x46, 0x4d, 0x53, 0x2f, 0x33, 0x2c, 0x35, 0x2c, 0x33, 0x2c, 0x38, 0x38, 0x38,

		0x00, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
		0x00, 0x40, 0x5f, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x00, 0x04, 0x6d, 0x6f, 0x64, 0x65,
		0x00, 0x3f, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x00, 0x00,
		0x09,

		0x03,
		0x00, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
		0x02, 0x00, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,

		0x00, 0x04, 0x63, 0x6f, 0x64, 0x65,
		0x02, 0x00, 0x1d, 0x4e, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,

		0x00, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
		0x02, 0x00, 0x14, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64,

		0x00, 0x0e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x00, 0x04, 0x64, 0x61, 0x74, 0x61,
		0x08,
		0x00, 0x00, 0x00, 0x00,

		0x00, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
		0x02, 0x00, 0x09, 0x33, 0x2c, 0x35, 0x2c, 0x33, 0x2c, 0x38, 0x38, 0x38,

		0x00, 0x07, 0x73, 0x72, 0x73, 0x5f, 0x73, 0x69, 0x67,
		0x02, 0x00, 0x03, 0x53, 0x52, 0x53,

		0x00, 0x0a, 0x73, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
		0x02, 0x00, 0x34, 0x53, 0x52, 0x53, 0x20, 0x31, 0x2e, 0x30, 0x2e, 0x31, 0x30, 0x20, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x76, 0x69, 0x70, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x72, 0x74, 0x6d, 0x70, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x29,

		0x00, 0x0b, 0x73, 0x72, 0x73, 0x5f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
		0x02, 0x00, 0x15, 0x54, 0x68, 0x65, 0x20, 0x4d, 0x49, 0x54, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x20, 0x28, 0x4d, 0x49, 0x54, 0x29,

		0x00, 0x08, 0x73, 0x72, 0x73, 0x5f, 0x72, 0x6f, 0x6c, 0x65,
		0x02, 0x00, 0x12, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,

		0x00, 0x07, 0x73, 0x72, 0x73, 0x5f, 0x75, 0x72, 0x6c,
		0x02, 0x00, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x76, 0x69, 0x70, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x72, 0x74, 0x6d, 0x70, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,

		0x00, 0x0b, 0x73, 0x72, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
		0x02, 0x00, 0x06, 0x31, 0x2e, 0x30, 0x2e, 0x31, 0x30,

		0x00, 0x08, 0x73, 0x72, 0x73, 0x5f, 0x73, 0x69, 0x74, 0x65,
		0x02, 0x00, 0x1c, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x63, 0x73, 0x64, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x77, 0x69, 0x6e, 0x5f, 0x6c, 0x69, 0x6e,

		0x00, 0x09, 0x73, 0x72, 0x73, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
		0x02, 0x00, 0x12, 0x77, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x40, 0x76, 0x69, 0x70, 0x2e, 0x31, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d,

		0x00, 0x0d, 0x73, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74,
		0x02, 0x00, 0x1e, 0x43, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x20, 0x28, 0x63, 0x29, 0x20, 0x32, 0x30, 0x31, 0x33, 0x2d, 0x32, 0x30, 0x31, 0x34, 0x20, 0x77, 0x69, 0x6e, 0x6c, 0x69, 0x6e,

		0x00, 0x0b, 0x73, 0x72, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
		0x02, 0x00, 0x06, 0x77, 0x69, 0x6e, 0x6c, 0x69, 0x6e,

		0x00, 0x0b, 0x73, 0x72, 0x73, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73,
		0x02, 0x00, 0x0b, 0x77, 0x65, 0x6e, 0x6a, 0x69, 0x65, 0x2e, 0x7a, 0x68, 0x61, 0x6f,

		0x00, 0x0d, 0x73, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x70,
		0x02, 0x00, 0x0b, 0x31, 0x37, 0x32, 0x2e, 0x31, 0x37, 0x2e, 0x30, 0x2e, 0x31, 0x30,

		0x00, 0x07, 0x73, 0x72, 0x73, 0x5f, 0x70, 0x69, 0x64,
		0x00, 0x3f, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x00, 0x06, 0x73, 0x72, 0x73, 0x5f, 0x69, 0x64,
		0x00, 0x40, 0x5a, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x00, 0x00,
		0x09,

		0x00, 0x00,
		0x09,
	}

	var n int = 0
	var start int = 0
	val, err := parseAMF0Val(0, data, &start)
	n += start
	assertEqual(t, err, nil)
	assertEqual(t, val.(string), "_result")

	start = 0
	val, err = parseAMF0Val(0, data[n:], &start)
	n += start
	assertEqual(t, err, nil)
	log.Println(val)
	assertEqual(t, int(val.(float64)), 1)

	start = 0
	val, err = parseAMF0Val(0, data[n:], &start)
	n += start
	assertEqual(t, err, nil)
	log.Println(val)
	assertEqual(t, val.(AMFMap).Get("fmsVer").V, "FMS/3,5,3,888")
	assertEqual(t, int(val.(AMFMap).Get("capabilities").V.(float64)), 127)
	assertEqual(t, int(val.(AMFMap).Get("mode").V.(float64)), 1)

	start = 0
	val, err = parseAMF0Val(0, data[n:], &start)

	// fmt.Println(val)
	print(val)
	n += start
	assertEqual(t, err, nil)
	assertEqual(t, val.(AMFMap).Get("level").V, "status")
	assertEqual(t, val.(AMFMap).Get("code").V, "NetConnection.Connect.Success")
	assertEqual(t, val.(AMFMap).Get("description").V, "Connection succeeded")
	assertEqual(t, int(val.(AMFMap).Get("objectEncoding").V.(float64)), 0)
	assertEqual(t, val.(AMFMap).Get("data").V.(AMFMap).Get("version").V, "3,5,3,888")
	assertEqual(t, int(val.(AMFMap).Get("data").V.(AMFMap).Get("srs_pid").V.(float64)), 1)

	assertEqual(t, val.(AMFMap).Get("notUsed") == nil, true)
}

func TestDecodeNull(t *testing.T) {
	data := []byte{
		2, 0, 4, 112, 108, 97, 121,
		0, 64, 8, 0, 0, 0, 0, 0, 0,
		5,
		2, 0, 1, 49,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	var n int = 0
	var start int = 0
	val, err := parseAMF0Val(0, data, &start)
	n += start
	assertEqual(t, err, nil)
	assertEqual(t, val.(string), "play")

	start = 0
	val, err = parseAMF0Val(0, data[n:], &start)
	n += start
	assertEqual(t, err, nil)
	assertEqual(t, int(val.(float64)), 3)

	start = 0
	val, err = parseAMF0Val(0, data[n:], &start)
	n += start
	assertEqual(t, err, nil)
	assertEqual(t, val, nil)

	start = 0
	val, err = parseAMF0Val(0, data[n:], &start)
	n += start
	assertEqual(t, err, nil)
	assertEqual(t, val.(string), "1")

	start = 0
	val, err = parseAMF0Val(0, data[n:], &start)
	n += start
	assertEqual(t, err, nil)
	assertEqual(t, int(val.(float64)), 0)
}

func TestEncodeToBytes(t *testing.T) {
	v := AMFMap{}
	v = v.Set("fmsVer", "FMS/3,5,3,888")
	v = v.Set("capabilities", float64(127.0))
	v = v.Set("mode", 1)
	v = v.Set("data", AMFMap{{"data", 123}})
	v = v.Set("null", 0)
	v = v.Set("st", AMFMap{{"float64", 1}})

	var n int
	FillAMF0Val(nil, &n, v)
	b := make([]byte, n)
	n = 0
	FillAMF0Val(b, &n, v)

	var nn int = 0
	var start int = 0
	val, err := parseAMF0Val(0, b, &start)
	nn += start
	assertEqual(t, err, nil)
	assertEqual(t, val.(AMFMap).Get("fmsVer").V, "FMS/3,5,3,888")

	start = 0
	val, err = parseAMF0Val(0, b, &start)
	nn += start
	assertEqual(t, err, nil)
	assertEqual(t, val.(AMFMap).Get("capabilities").V.(float64), float64(127.0))

	start = 0
	val, err = parseAMF0Val(0, b, &start)
	nn += start
	assertEqual(t, err, nil)
	assertEqual(t, val.(AMFMap).Get("mode").V.(float64), float64(1))

	start = 0
	val, err = parseAMF0Val(0, b, &start)
	nn += start
	assertEqual(t, err, nil)
	assertEqual(t, val.(AMFMap).Get("data").V.(AMFMap).Get("data").V.(float64), float64(123))

	start = 0
	val, err = parseAMF0Val(0, b, &start)
	nn += start
	assertEqual(t, err, nil)
	assertEqual(t, val.(AMFMap).Get("null").V.(float64), float64(0))

	start = 0
	val, err = parseAMF0Val(0, b, &start)
	nn += start
	assertEqual(t, err, nil)
	assertEqual(t, val.(AMFMap).Get("st").V.(AMFMap).Get("float64").V.(float64), float64(1))
}

func TestLenAMF0Val(t *testing.T) {
	v := uint8(1)

	var n int
	FillAMF0Val(nil, &n, v)
	assertEqual(t, n, 9)

	b := make([]byte, n)

	n = 0
	FillAMF0Val(b, &n, v)
	assertEqual(t, b, []byte{0x0, 0x3f, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})
}

func TestFillAMF0Val(t *testing.T) {
	v1 := float64(2.0)

	var n int
	FillAMF0Val(nil, &n, v1)
	assertEqual(t, n, 9)

	b := make([]byte, n)

	n = 0
	FillAMF0Val(b, &n, v1)
	assertEqual(t, b, []byte{0x0, 0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})
}

func TestAMFMap(t *testing.T) {
	m := AMFMap{}
	m = m.Set("a", "b")
	m = m.Set("c", "d")
	m = m.Del("a")
	assertEqual(t, len(m), 1)
}
