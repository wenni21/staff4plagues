package plague

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//GoroutineControls : controllable thread routine
type GoroutineControls struct {
	Wait         bool        `json:"wait" bson:"wait"`
	ControlSleep int64       `json:"controlSleep" bson:"controlSleep"`
	Notify       bool        `json:"notify" bson:"notify"`
	Cancel       bool        `json:"cancel" bson:"cancel"`
	Done         bool        `json:"done" bson:"done"`
	On           bool        `json:"on" bson:"on"`
	Get          interface{} `json:"get" bson:"get"`
}

//StrSliceContains to see if it contains
func StrSliceContains(s []string, st string) bool {
	if len(s) > 0 {
		for _, a := range s {
			if strings.Compare(a, st) == 0 {
				return true
			}
		}
	}
	return false
}

//TimeLongStr  gets a timestamp and as string
func TimeLongStr() string {
	return strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
}

//UUidGenAlt is a uuid random gen.
func uuidGenAlt() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func hashing(rtn string) string {
	hasher := md5.New()
	io.WriteString(hasher, rtn)
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}

//HashingMD : hash + MD5
func HashingMD(rtn string) string {
	hasher := md5.New()
	io.WriteString(hasher, rtn)
	hash := hex.EncodeToString(hasher.Sum(nil))
	return strings.ToUpper(hash)
}

//CurrentMillis : get time
func CurrentMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//GenerateID : make an ID
func GenerateID() string {
	uu := uuidGenAlt()
	milli := strconv.FormatInt(CurrentMillis(), 10)
	return hashing(uu + milli)
}

//FormatHourMinute transform hour:minute to milli secs
func FormatHourMinute(hm string) time.Time {
	curDate := time.Now().Format("20060102") + strings.Replace(hm, ":", "", 1)
	stTime, _ := time.Parse("200601021504", curDate)
	return stTime
}

//InitAllMonsters initiate all threads
func InitAllMonsters() {
	time.Sleep(time.Millisecond * 1000 * 3) // wait for the pushing to cache
}

//InitMonsterServer init server only
func InitMonsterServerPlagues() {
	WebForDistributionsPlagues()
}

//ReturnABunchOfStuff ...
func ReturnABunchOfStuff(w http.ResponseWriter, backData map[string]interface{}, statusHTTP int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vb, _ := json.Marshal(backData)
	w.Write(vb)
	w.WriteHeader(statusHTTP)
}

func ReturnABunchOfStuffDirect(w http.ResponseWriter, backData string, statusHTTP int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte(backData))
	w.WriteHeader(statusHTTP)
}

//ReturnWithHeadABunchOfStuff ...
func ReturnWithHeadABunchOfStuff(w http.ResponseWriter, backData map[string]interface{}, statusHTTP int, head map[string]string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if len(head) > 0 {
		for k, v := range head {
			w.Header().Set(k, v)
		}
	}
	vb, _ := json.Marshal(backData)
	w.Write(vb)
	w.WriteHeader(statusHTTP)
}

//FeedBackReturnOk : as named
func FeedBackReturnOk(feedbackContent interface{}) map[string]interface{} {
	res := make(map[string]interface{}, 0)
	res["feed"] = 9000
	res["content"] = feedbackContent
	return res
}

//FeedBackReturnFail : as named
func FeedBackReturnFail(feedbackContent interface{}) map[string]interface{} {
	res := make(map[string]interface{}, 0)
	res["feed"] = 7000
	res["exceptional"] = feedbackContent
	return res
}

func SendBackCommonHead(back string, status int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte(back))
	w.WriteHeader(status)
}

// Inarray is contains.
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panics("InterfaceSlice() given a non-slice type")
	}
	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

func InterfaceMap(mp interface{}) map[string]interface{} {
	rtnMap := make(map[string]interface{}, 0)
	v := mp.(map[string]interface{})
	for k, val := range v {
		rtnMap[k] = val
	}
	return rtnMap
}
