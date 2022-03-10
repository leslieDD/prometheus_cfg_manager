package models

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"time"

	"gorm.io/gorm"
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
	id := OnlyID{}
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

// publish_at_null_subgroup
func doOptions_1() *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	// 事务
	sql := `SELECT machines.id AS machine_id, job_machines.job_id FROM job_machines 
	LEFT JOIN machines 
	ON machines.id=job_machines.machine_id
	WHERE job_machines.job_id NOT IN (SELECT jobs_id FROM job_group)`
	jgs := []*JobIDAndMachinesID{}
	tx := db.Table("job_machines").Raw(sql).Find(&jgs)
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
		groupID, ok := jobsIDMap[jg.JobID] // 是唯一
		if !ok {
			newJG := JobGroup{
				ID:       0,
				Name:     "默认子组",
				JobsID:   jg.JobID,
				Enabled:  true,
				UpdateAt: time.Now(),
			}
			txCreate := shiwu.Table("job_group").Create(&newJG)
			if txCreate.Error != nil {
				config.Log.Error(txCreate.Error)
				db.Rollback()
				return ErrTransaction
			}
			jobsIDMap[jg.JobID] = newJG.ID
			groupID = newJG.ID
		}
		jgm := JobGroupIP{
			MachinesID: jg.MachineID,
			JobGroupID: groupID,
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

func doOptions_2() *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		jobs := []*Jobs{}
		if err := tx.Table("jobs").Find(&jobs).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		for _, j := range jobs {
			sql := `SELECT * FROM job_machines
			LEFT JOIN group_machines
			ON job_machines.machine_id = group_machines.machines_id
			WHERE group_machines.job_group_id IS NULL AND job_machines.job_id=?`
			jgs := []*JobIDAndMachinesID{}
			if err := db.Raw(sql, j.ID).Find(&jgs).Error; err != nil {
				config.Log.Error(err)
				return err
			}

		}
	})
	// 事务
	sql := `SELECT job_machines.machine_id, job_machines.job_id FROM job_machines
	LEFT JOIN group_machines
	ON group_machines.machines_id=job_machines.machine_id
	WHERE group_machines.job_group_id IS NULL`
	jgs := []*JobIDAndMachinesID{}
	tx := db.Table("job_machines").Raw(sql).Find(&jgs)
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
		groupID, ok := jobsIDMap[jg.JobID]
		if !ok {
			// 检查是不是有这么一个子组
			id, bf := getJobGroupID("默认子组", jg.JobID)
			if bf == Success {
				jobsIDMap[jg.JobID] = id.ID
			} else {
				newJG := JobGroup{
					ID:       0,
					Name:     "默认子组",
					JobsID:   jg.JobID,
					Enabled:  true,
					UpdateAt: time.Now(),
				}
				txCreate := shiwu.Table("job_group").Create(&newJG)
				if txCreate.Error != nil {
					config.Log.Error(txCreate.Error)
					db.Rollback()
					return ErrTransaction
				}
				jobsIDMap[jg.JobID] = newJG.ID
			}
			groupID = jobsIDMap[jg.JobID]
		}
		jgm := JobGroupIP{
			// ID:         0,
			MachinesID: jg.MachineID,
			JobGroupID: groupID,
		}
		txCreate2 := shiwu.Exec("INSERT ignore INTO `group_machines` (`job_group_id`,`machines_id`) VALUES (?,?);",
			jgm.JobGroupID,
			jgm.MachinesID,
		)
		// txCreate2 := shiwu.Table("group_machines").Create(&jgm)
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

func doOptions_3() ([]*OnlyIDAndCount, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	idC := []*OnlyIDAndCount{}
	tx := db.Table("job_machines").Raw(`SELECT jobs.id, COUNT(job_machines.machine_id) AS count FROM job_machines 
	LEFT JOIN jobs 
	ON job_machines.job_id=jobs.id
	GROUP BY jobs.id 
	ORDER BY jobs.id `).Find(&idC)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	idDef := []*OnlyIDAndCount{}
	tx = db.Table("jobs").
		Select("jobs.id, (SELECT COUNT(*) FROM machines WHERE enabled=1) AS `count`").
		Where("jobs.is_common=1 AND jobs.enabled=1").Find(&idDef)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	idC = append(idC, idDef...)
	return idC, Success
}

func DoTmplBefore() *BriefMessage {
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
	// 在发布JOB组时，针对有配置子组的JOB组，是否为此JOB组的未分组的IP添加进默认子组
	r, bf = CheckByFiled("publish_at_remain_subgroup", "true")
	if bf != Success {
		return bf
	}
	if r {
		if bf := doOptions_2(); bf != Success {
			return bf
		}
	}
	return Success
}

func DoTmplAfter(needReload bool) *BriefMessage {
	// var needReload bool
	// r, bf := CheckByFiled("publish_jobs_also_reload_srv", "true")
	// if bf != Success {
	// 	return bf
	// }
	// if r {
	// needReload = true
	// }
	// r, bf = CheckByFiled("publish_jobs_also_ips", "true")
	// if bf != Success {
	// 	return bf
	// }
	// if r {
	// 	if bf := publish.Do(needReload); bf != Success {
	// 		return bf
	// 	}
	// }
	if bf := publish.Do(needReload); bf != Success {
		return bf
	}
	if needReload {
		if bf := Reload(); bf != Success {
			return bf
		}
	}
	return Success
}
