package models

import (
	"pro_cfg_manager/config"
)

func GetTest() (interface{}, *BriefMessage) {
	id, bf := getJobGroupID("默认子组", 100)
	config.Log.Printf("%v %v", id, bf)
	return nil, Success
}
