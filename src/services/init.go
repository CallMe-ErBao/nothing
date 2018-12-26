package services

import (
	"meta"
	"pandora/modules/db"
)

func Init() {
	db.Update(meta.TUsers)
	db.Update(meta.EUsers)
	db.Update(meta.Project)

}
