package persistence

import (
	"time"

	PO "github.com/Rock-liyi/p2pdb/domain/discovery/repository/po"
	orm "github.com/Rock-liyi/p2pdb/infrastructure/core/orm"
	log "github.com/Rock-liyi/p2pdb/infrastructure/util/log"
	//function "github.com/Rock-liyi/p2pdb/infrastructure/untl/function"
)

var DB *orm.CreateDBFactory = new(orm.CreateDBFactory)

func init() {
	//settings  use  the default  information database
	DB.SetIsDBInformation(true)
}

func CreateInstance(name string) {

}

func GetInstance(name string) {

}

func CreateInstanceInformationTable() bool {
	db := DB.InitDB()
	db.Migrator().CreateTable(&PO.InstanceInformation{})
	return true
}

func CreatePeerNodeInfomationTable() bool {
	db := DB.InitDB()
	db.Migrator().CreateTable(&PO.PeerNodeInfomation{})
	return true
}

func InsertInstanceInformation(InstanceId string, LocalPeerId string, GlobalClockTime int64) bool {
	db := DB.InitDB()

	var instance = &PO.InstanceInformation{}
	instance.InstanceId = InstanceId
	instance.LocalPeerId = LocalPeerId
	instance.GlobalClockTime = GlobalClockTime
	instance.CreatedAt = time.Now()
	instance.UpdatedAt = time.Now()
	err := db.InsertIgnore(instance)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func UpdateInstanceInformation(InstanceId string, LocalPeerId string, GlobalClockTime int64) bool {
	db := DB.InitDB()

	var instance = &PO.InstanceInformation{}
	instance.LocalPeerId = LocalPeerId
	instance.GlobalClockTime = GlobalClockTime + 1
	instance.UpdatedAt = time.Now()
	err := db.Where(instance.GetInstancePrimaryId()+" = ?", InstanceId).Updates(instance)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func GetInstanceInformation() *PO.InstanceInformation {

	db := DB.InitDB()
	var instance = &PO.InstanceInformation{}
	db.First(instance)

	return instance
}
