package stat

import (
	"go-back/pkg/dbPostgresComposable"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	*dbPostgresComposable.Db
}

func NewStatRepository(db *dbPostgresComposable.Db) *StatRepository {
	return &StatRepository{
		Db: db,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	currentDate := datatypes.Date(time.Now())
	var stat Stat
	repo.DB.Find(&stat, "link_id = ? and date = ?", linkId, datatypes.Date(currentDate))
	if stat.ID == 0 {
		repo.DB.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks += 1
		repo.Save(&stat)
	}
}
