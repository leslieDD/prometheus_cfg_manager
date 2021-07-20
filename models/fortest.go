package models

import (
	"encoding/json"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/utils"
)

func GetTest() (interface{}, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	l := Label{}
	if err := db.Table("annotations").Where("id=416").First(&l).Error; err != nil {
		return nil, ErrSearchDBData
	}
	j, _ := json.Marshal(&l)
	utils.WIoutilByte("tmpfile.txt", j)
	// for _, c := range l.Value {
	// fmt.Printf(" %v %c \n", c, c)
	// }
	// l.Value = utils.ClearCrLR(l.Value)
	return l, Success
	// id, bf := getJobGroupID("默认子组", 100)
	// config.Log.Printf("%v %v", id, bf)
	// return nil, Success
}
