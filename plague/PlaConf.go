package plague

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"gopkg.in/mgo.v2"
)

const (
	MongoMainOne    = "127.0.0.1:27017" //this one is for alpha server
	MongoMainDB     = "plague_of_corona"
	MongoCollMain   = "everyone_info"
	MongoCollAdmin  = "admins_info"
	MongoCollComp   = "companies_info"
	MongoCollPlague = "plague_info"

	RedisHost = "127.0.0.1" //this one is for alpha version
	RedisPort = 6379
	RedisDB   = 0
)

var DataStoreCrazyHandle DataStore
var ClientOfRedis *redis.Client
var EveryoneInfoColl *mgo.Collection
var AdminInfoColl *mgo.Collection
var CompInfoColl *mgo.Collection
var PlagueColl *mgo.Collection

/**************
redis client
***************/
func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     RedisHost + ":" + strconv.Itoa(RedisPort),
		Password: "",
		DB:       RedisDB,
	})
	pong, err := client.Ping().Result()
	CheckWithWarn(err)
	fmt.Printf("* redis connection ping:%s\n", pong)
	return client
}

/**************
mgo connection
***************/
func MongoClient() *mgo.Session {
	session, err := mgo.Dial(MongoMainOne)
	CheckWithWarn(err)
	fmt.Printf("* mongo connection:%s\n", MongoMainOne)
	session.SetMode(mgo.Monotonic, true)
	return session
}

//InitDataStore as named (MongoDB)
func InitDataStore() {
	DataStoreCrazyHandle.session = MongoClient()
	EveryoneInfoColl = DataStoreCrazyHandle.session.DB(MongoMainDB).C(MongoCollMain)
	AdminInfoColl = DataStoreCrazyHandle.session.DB(MongoMainDB).C(MongoCollAdmin)
	CompInfoColl = DataStoreCrazyHandle.session.DB(MongoMainDB).C(MongoCollComp)
	PlagueColl = DataStoreCrazyHandle.session.DB(MongoMainDB).C(MongoCollPlague)
}

func CurrentDateString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func CurrentMilliseconds() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func CheckAndPanic(err error) {
	if err != nil {
		fmt.Println("err&panic::" + err.Error())
		panic(err)
	}
}

func CheckWithWarn(err error) {
	if err != nil {
		fmt.Println("warning::" + err.Error())
	}
}

type DataStore struct {
	session *mgo.Session
}

func Hashing(rtn string) string {
	hasher := md5.New()
	io.WriteString(hasher, rtn)
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}
