package repos

import (
	"github.com/jinzhu/gorm"
	. "github.com/rajeshpg/pair-monitor-go/models"
)

type DevPairRepo interface {
	SaveSession(devPair *DevPair) (uint, error)
	AllSessions() ([]DevPair, error)
}

type DevPairDao struct {
	Db *gorm.DB
}

func (dao *DevPairDao) AllSessions() ([]DevPair, error) {
	var pairs []DevPair
	res := dao.Db.Find(&pairs)
	return pairs, res.Error
}

func (dao *DevPairDao) SaveSession(pair *DevPair) (uint, error) {
	res := dao.Db.Create(pair)
	return pair.ID, res.Error
}

