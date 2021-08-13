package main

import (
	"fmt"
	"goService/components"
	DemandPlanSearch "goService/model"

	//"gopkg.in/redis.v4"

	_ "github.com/go-sql-driver/mysql" //只是用这个包力的init
)

func main() {

	components.Init()

	//demandPlanSearch := model.DemandPlanSearch{}

	demandPlanSearch := DemandPlanSearch.Get(13268)

	fmt.Printf("std: %+v\n", demandPlanSearch)

}
