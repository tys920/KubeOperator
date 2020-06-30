package model

import (
	"github.com/KubeOperator/KubeOperator/pkg/model/common"
	uuid "github.com/satori/go.uuid"
)

type Zone struct {
	common.BaseModel
	ID       string `json:"id" gorm:"type:varchar(64)"`
	Name     string `json:"name" gorm:"type:varchar(256);not null;unique"`
	Vars     string `json:"vars" gorm:"type longtext(0)"`
	Status   string `json:"status" gorm:"type:varchar(64)"`
	IpUsed   string `json:"ipUsed" gorm:"type longtext(0)"`
	RegionID string `json:"regionId" gorm:"type:int(64)"`
}

func (z *Zone) BeforeCreate() (err error) {
	z.ID = uuid.NewV4().String()
	return err
}

func (z Zone) TableName() string {
	return "ko_zone"
}
