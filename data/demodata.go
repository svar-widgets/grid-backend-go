package data

import (
	"time"
)

func dataDown(d *DAO) {
	d.mustExec("DELETE from films")
}

func formatDate(date string) *time.Time {
	t, _ := time.Parse(time.RFC3339, date)
	return &t
}

func dataUp(d *DAO) {
	db := d.GetDB()

	db.Create(&Film{Name: "Best film ever", Year: 2018, Votes: 950})
	db.Create(&Film{Name: "Not so good one", Year: 2019, Votes: 432})
	db.Create(&Film{Name: "Boring one", Year: 2018, Votes: 138})
}
