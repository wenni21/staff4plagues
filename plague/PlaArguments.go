package plague

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	PageLimit = 500
)

type EveryoneInfo struct {
	Index                  string                  `json:"index" bson:"index"`
	RecordTime             int64                   `json:"recordTime" bson:"recordTime"`
	RecordName             string                  `json:"recordName" bson:"recordName"`
	StaffName              string                  `json:"staffName" bson:"staffName"`
	StaffNumber            string                  `json:"staffNumber" bson:"staffNumber"`
	CellPhone              string                  `json:"cellphone" bson:"cellphone"`
	Email                  string                  `json:"email" bson:"email"`
	CompanyName            string                  `json:"companyName" bson:"companyName"`
	Department             string                  `json:"department" bson:"department"`
	HealthStatus           string                  `json:"healthStatus" bson:"healthStatus"`
	CurrentLocation        string                  `json:"currentLocation" bson:"currentLocation"`
	CurrentLatLon          string                  `json:"currentLatLon" bson:"currentLatLon"`
	CurrentLocationStatus  int                     `json:"currentLocationStatus" bson:"currentLocationStatus"`
	CurrentLocationDetail  CurrentLocationDetail   `json:"currentLocationDetail" bson:"currentLocationDetail"`
	HasPassedOutbreakZone  int                     `json:"hasPassedOutbreakZone" bson:"hasPassedOutbreakZone"`
	OutbreakZoneTrip       OutbreakZoneTripDetails `json:"outbreakZoneTrip" bson:"outbreakZoneTrip"`
	HasContactWithInfected int                     `json:"hasContactWithInfected" bson:"hasContactWithInfected"`
	InfectedContactDetails InfectedContactDetails  `json:"infectedContactDetails" bson:"infectedContactDetails"`
	LocalWorkPolicy        string                  `json:"localWorkPolicy" bson:"localWorkPolicy"`
}

type OutbreakZoneTripDetails struct {
	When  string `json:"when" bson:"when"`
	Where string `json:"where" bson:"where"`
	How   string `json:"how" bson:"how"`
	What  string `json:"what" bson:"what"`
}

type CurrentLocationDetail struct {
	AppearTime  string `json:"appearTime" bson:"appearTime"`
	Approximity int    `json:"approximity" bson:"approximity"`
	Desc        string `json:"desc" bson:"desc"`
}

type InfectedContactDetails struct {
	Place          string `json:"place" bson:"place"`
	Time           string `json:"time" bson:"time"`
	IsCloseContact int    `json:"isCloseContact" bson:"isCloseContact"`
	Desc           string `json:"desc" bson:"desc"`
}

type FetchFilter struct {
	StaffName   string `json:"staffName" bson:"staffName"`
	StaffNumber string `json:"staffNumber" bson:"staffNumber"`
	Department  string `json:"department" bson:"department"`
	Company     string `json:"company" bson:"company"`
	CellPhone   string `json:"cellphone" bson:"cellphone"`
}

type Plague struct {
	Ticket string `json:"ticket" bson:"ticket"`
	Date   string `json:"date" bson:"date"`
}

type Count struct {
	Count int `json:"count"`
}

func GetCountTotal() string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("GetCountTotal-Error: ", err)
		}
	}()
	c, err := EveryoneInfoColl.Count()
	count := Count{}
	if err != nil {
		fmt.Println(err)
		count.Count = 0
	} else {
		count.Count = c
	}
	cs, _ := json.Marshal(count)
	return string(cs[:])
}

func DeleteOneMan(index string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("DeleteOneMan-Error: ", err)
		}
	}()
	fmt.Println("will delete::" + index)
	EveryoneInfoColl.Remove(bson.M{"index": index})
}

func Recording(info EveryoneInfo) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Recording-Error: ", err)
		}
	}()
	info.Index = HashingMD(info.StaffName + info.StaffNumber)
	when := info.OutbreakZoneTrip.When
	what := info.OutbreakZoneTrip.What
	breaked := Plague{}
	leak := ""
	leakMini := "没有同旅程感染"
	if len(when) > 0 && len(what) > 0 {
		errLeak := PlagueColl.Find(bson.M{"ticket": what, "date": when}).Limit(1).One(&breaked)
		if errLeak != nil {
			fmt.Println(errLeak)
		} else {
			if len(breaked.Ticket) > 0 {
				leak = "在" + breaked.Date + "乘坐" + breaked.Ticket + "的航班/列车 在旅程中有已被感染的乘客同次"
				leakMini = "可能有同旅程感染"
			}
		}
	}
	info.CurrentLatLon = leak
	info.RecordName = info.StaffName + "[" + info.StaffNumber + "] " + info.CurrentLocation + "[" + info.CellPhone + "] (" + leakMini + ")"
	info.RecordTime = CurrentMillis()
	EveryoneInfoColl.Upsert(bson.M{"index": info.Index}, info)
	return 0
}

func GetAllPagedFetches(filter FetchFilter, page int) []EveryoneInfo {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("GetAllPagedFetches-Error: ", err)
		}
	}()
	everyone := make([]EveryoneInfo, 0)
	var err error
	if len(filter.StaffName) > 0 {
		err = EveryoneInfoColl.Find(bson.M{"companyName": filter.Company,
			"$or": []bson.M{bson.M{"staffName": bson.M{"$regex": filter.StaffName}},
				bson.M{"cellphone": bson.M{"$regex": filter.CellPhone}},
				bson.M{"staffNumber": filter.StaffNumber}}}).Sort("-$recordTime").Skip((page - 1) * PageLimit).Limit(PageLimit).All(&everyone)
	} else {
		if strings.Compare(filter.Company, "devops") == 0 {
			err = EveryoneInfoColl.Find(bson.M{}).Sort("-$recordTime").Skip((page - 1) * PageLimit).Limit(PageLimit).All(&everyone)
		} else {
			err = EveryoneInfoColl.Find(bson.M{"companyName": filter.Company}).Sort("-$recordTime").Skip((page - 1) * PageLimit).Limit(PageLimit).All(&everyone)
		}
	}
	if err != nil {
		fmt.Println("fetching everyone error @ GetAllPagedFetches")
		empty := make([]EveryoneInfo, 0)
		return empty
	}
	return everyone
}

func FindInfectedArea(area string, page int) []EveryoneInfo {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("FindInfectedArea-Error: ", err)
		}
	}()
	everyone := make([]EveryoneInfo, 0)
	err := EveryoneInfoColl.Find(bson.M{"currentLocation": bson.M{"$regex": area}}).Sort("-$recordTime").Skip((page - 1) * PageLimit).Limit(PageLimit).All(&everyone)
	if err != nil {
		fmt.Println("fetching everyone error @ FindInfectedArea")
		empty := make([]EveryoneInfo, 0)
		return empty
	}
	return everyone
}

func FindPassedZones() []EveryoneInfo {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("FindPassedZones-Error: ", err)
		}
	}()
	everyone := make([]EveryoneInfo, 0)
	err := EveryoneInfoColl.Find(bson.M{"$or": []bson.M{bson.M{"hasPassedOutbreakZone": 1},
		bson.M{"currentLocationStatus": 1},
		bson.M{"hasContactWithInfected": 1},
		bson.M{"recordName": bson.M{"$regex": "被感染的乘客"}}}}).Sort("-$recordTime").Skip((1 - 1) * 1000).Limit(1000).All(&everyone)
	if err != nil {
		fmt.Println("fetching everyone error @ FindPassedZones")
		empty := make([]EveryoneInfo, 0)
		return empty
	}
	return everyone
}

type AdminInfo struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Company  string `json:"company" bson:"company"`
}

func AdminLogin(adm AdminInfo) string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("AdminLogin-Error: ", err)
		}
	}()
	lgn := AdminInfo{}
	err := AdminInfoColl.Find(bson.M{"username": adm.Username, "password": adm.Password}).Limit(1).One(&lgn)
	if err != nil {
		fmt.Println(err)
		return "NaN"
	}
	if len(lgn.Username) == 0 {
		fmt.Println("cannot find user")
		return "NaN"
	}
	admd, _ := json.Marshal(lgn)
	return string(admd[:])
}

type CompInfo struct {
	Name string `json:"name" bson:"name"`
	Key  string `json:"key" bson:"key"`
}

func GetComps() string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("GetComps-Error: ", err)
		}
	}()
	comps := make([]CompInfo, 0)
	err := CompInfoColl.Find(nil).All(&comps)
	if err != nil {
		fmt.Println(err)
		return "NaN"
	} else {
		fmt.Println(comps)
		rtn, _ := json.Marshal(comps)
		return string(rtn[:])
	}
	return "NaN"
}

func GetPlagueInfo() {
	index := 438995
	for {
		h := make(map[string]string, 0)
		data, head := DoReqsRHB("http://2019ncov.nosugartech.com/data.json?"+strconv.Itoa(index), "GET", "", h)
		fmt.Println(head)
		fmt.Println(string(data[:]))
		if len(data) > 0 {
			parsePlague(data)
			index++
		}
		time.Sleep(time.Millisecond * 2400000)
	}
}

func parsePlague(plague []byte) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("parsePlague-Error: ", err)
		}
	}()
	dataRaw := make(map[string]interface{}, 0)
	err := json.Unmarshal(plague, &dataRaw)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		plagues := make([]Plague, 0)
		dataArr := InterfaceSlice(dataRaw["data"])
		if len(dataArr) > 0 {
			for d := range dataArr {
				p := Plague{}
				m := InterfaceMap(dataArr[d])
				p.Ticket = m["t_no"].(string)
				p.Date = m["t_date"].(string)
				plagues = append(plagues, p)
			}
		}
		for p := range plagues {
			_ = PlagueColl.Remove(bson.M{"ticket": plagues[p].Ticket, "date": plagues[p].Date})
			_ = PlagueColl.Insert(plagues[p])
		}
	}
}
