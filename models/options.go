package models

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"time"
)

type Option struct {
	OptKey   string `json:"opt_key" gorm:"column:opt_key"`
	OptValue string `json:"opt_value" gorm:"column:opt_value"`
}

func GetOptions() (map[string]string, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	opts := []Option{}
	tx := db.Table("options").Find(&opts)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrDelData
	}
	result := map[string]string{}
	for _, o := range opts {
		result[o.OptKey] = o.OptValue
	}
	return result, Success
}

func PutOptions(opts map[string]string) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	for key, value := range opts {
		tx := db.Table("options").
			Where("opt_key", key).
			Update("opt_value", value)
		if tx.Error != nil {
			config.Log.Error(tx.Error)
			return ErrUpdateData
		}
	}
	return Success
}

func getJobGroupID(groupName string, jobID int) (*OnlyID, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	id := OnlyID{ID: -1}
	tx := db.Table("job_group").
		Select("id").
		Where("name=? and jobs_id=?", groupName, jobID).
		Order("update_at desc").
		First(&id)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return &id, Success
}

func doOptions_2() *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	// 事务
	sql := `SELECT 
	jobs.id as jobs_id, 
  machines.id AS machines_id 
FROM machines 
LEFT JOIN jobs 
ON JSON_CONTAINS(machines.job_id, JSON_ARRAY(jobs.id)) 
LEFT JOIN job_group 
ON job_group.jobs_id=jobs.id 
WHERE jobs.is_common=0 AND job_group.name IS NULL `
	jgs := []*JobIDAndMachinesID{}
	tx := db.Table("machines").Raw(sql).Find(&jgs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrSearchDBData
	}
	shiwu := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			shiwu.Rollback()
		}
	}()
	jobsIDMap := map[int]int{}
	for _, jg := range jgs {
		// 检查是否有默认子组，如果有则加入第一个，如果没有则新建一个
		groupID, ok := jobsIDMap[jg.JobsID]
		if !ok {
			// 检查是不是有这么一个子组
			id, bf := getJobGroupID("默认子组", jg.JobsID)
			if bf == Success {
				jobsIDMap[jg.JobsID] = id.ID
			} else {
				newJG := JobGroup{
					ID:       0,
					Name:     "默认子组",
					JobsID:   jg.JobsID,
					Enabled:  true,
					UpdateAt: time.Now(),
				}
				txCreate := shiwu.Table("job_group").Create(&newJG)
				if txCreate.Error != nil {
					config.Log.Error(txCreate.Error)
					db.Rollback()
					return ErrTransaction
				}
				jobsIDMap[jg.JobsID] = newJG.ID
			}
			groupID = id.ID
		}
		jgm := JobGroupIP{
			ID:         0,
			MachinesID: jg.MachinesID,
			JobGroupID: groupID,
			UpdateAt:   time.Now(),
		}
		txCreate2 := shiwu.Table("group_machines").Create(&jgm)
		if txCreate2.Error != nil {
			config.Log.Error(txCreate2.Error)
			db.Rollback()
			return ErrTransaction
		}
	}
	r := shiwu.Commit()
	if r.Error != nil {
		config.Log.Error(r.Error)
		db.Rollback()
		return ErrTransaction
	}
	return Success
}

func doOptions_3() *BriefMessage {
	return Success
}

func doOptions_4() *BriefMessage {
	return Success
}

// publish_at_null_subgroup
func doOptions_1() *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	// 事务
	sql := `SELECT machines.id AS machines_id, jobs.id AS jobs_id FROM machines 
	LEFT JOIN jobs 
	ON JSON_CONTAINS(machines.job_id, JSON_ARRAY(jobs.id))
	WHERE jobs.id NOT IN (SELECT jobs_id FROM job_group)
	`
	jgs := []*JobIDAndMachinesID{}
	tx := db.Table("machines").Raw(sql).Find(&jgs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrSearchDBData
	}
	shiwu := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			shiwu.Rollback()
		}
	}()
	jobsIDMap := map[int]int{}
	for _, jg := range jgs {
		groupID, ok := jobsIDMap[jg.JobsID] // 是唯一
		if !ok {
			newJG := JobGroup{
				ID:       0,
				Name:     "默认子组",
				JobsID:   jg.JobsID,
				Enabled:  true,
				UpdateAt: time.Now(),
			}
			txCreate := shiwu.Table("job_group").Create(&newJG)
			if txCreate.Error != nil {
				config.Log.Error(txCreate.Error)
				db.Rollback()
				return ErrTransaction
			}
			jobsIDMap[jg.JobsID] = newJG.ID
			groupID = newJG.ID
		}
		jgm := JobGroupIP{
			ID:         0,
			MachinesID: jg.MachinesID,
			JobGroupID: groupID,
			UpdateAt:   time.Now(),
		}
		txCreate2 := shiwu.Table("group_machines").Create(&jgm)
		if txCreate2.Error != nil {
			config.Log.Error(txCreate2.Error)
			db.Rollback()
			return ErrTransaction
		}
	}
	r := shiwu.Commit()
	if r.Error != nil {
		config.Log.Error(r.Error)
		db.Rollback()
		return ErrTransaction
	}
	return Success
}

func DoOptionsFunc() *BriefMessage {
	// 在发布JOB组时，针对没有配置任何子组的JOB组，是否为此JOB组的所有IP生成无标签子组
	r, bf := CheckByFiled("publish_at_null_subgroup", "true")
	if bf != Success {
		return bf
	}
	if r {
		if bf := doOptions_1(); bf != Success {
			return bf
		}
	}
	// 在发布JOB组时，针对有配置子组的JOB组，是否为此JOB组的未分组IP生成无标签子组
	r, bf = CheckByFiled("publish_at_remain_subgroup", "true")
	if bf != Success {
		return bf
	}
	if r {
		if bf := doOptions_2(); bf != Success {
			return bf
		}
	}
	r, bf = CheckByFiled("publish_at_empty_nocreate_file", "true")
	if bf != Success {
		return bf
	}
	if r {
		// if bf := GetJobsForOptions(); bf != Success {
		// 	return bf
		// }
	}
	r, bf = CheckByFiled("publish_jobs_also_ips", "true")
	if bf != Success {
		return bf
	}
	if r {
		// if bf := GetJobsForOptions(); bf != Success {
		// 	return bf
		// }
	}
	r, bf = CheckByFiled("publish_jobs_also_reload_srv", "true")
	if bf != Success {
		return bf
	}
	if r {
		// if bf := GetJobsForOptions(); bf != Success {
		// 	return bf
		// }
	}
	return Success

}
