package main

import (
	"fmt"
	"time"

	plague "./plague"
)

func main() {
	startMainFuncPlagues() //online prod for host pooling
}

func startMainFuncPlagues() {
	MonstersInitPlagues()
	InitStarter()
	go plague.GetPlagueInfo()
	plague.InitMonsterServerPlagues() //the server
}

func InitStarter() {
	plague.ClientOfRedis = plague.RedisClient()
	plague.InitDataStore()
}

//MonstersInit : all monsters
func MonstersInitPlagues() {
	//BANNERS + TIME
	fmt.Println("============================")
	fmt.Println("# MONSTERS ARE REAL. BEWARE.")
	fmt.Println("============================")
	fmt.Println(time.Now().Format("2006-01-02")) //show current time
	fmt.Println("============================")
	fmt.Println("============================")
	fmt.Println("# \"FOR ALL THE SATANIC POWERS I WILL GAIN FOR ETERNITY.\"")
	fmt.Println("============================")
	/*BANNER for start*/
}
