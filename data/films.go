package data

import (
	"fmt"

	"gorm.io/gorm"
)

type FilmUpdate struct {
	Row struct {
		Name  string   `json:"name"`
		Votes FuzzyInt `json:"votes"`
		Year  FuzzyInt `json:"year"`
	} `json:"row"`
}

func NewFilmsDAO(db *gorm.DB) *FilmsDAO {
	return &FilmsDAO{db}
}

type FilmsDAO struct {
	db *gorm.DB
}

func (m *FilmsDAO) GetAll() ([]Film, error) {
	films := make([]Film, 0)
	err := m.db.
		Order("`id` asc").
		Find(&films).Error

	return films, err
}

func (m *FilmsDAO) Delete(id int) error {
	err := m.db.Delete(&Film{}, id).Error
	return err
}

func (m *FilmsDAO) Update(id int, f FilmUpdate) (*Film, error) {
	c := Film{}
	err := m.db.Find(&c, id).Error
	if err != nil || c.ID == 0 {
		return nil, err
	}

	c.Name = f.Row.Name
	c.Votes = int(f.Row.Votes)
	c.Year = int(f.Row.Year)

	err = m.db.Save(&c).Error

	return &c, err
}

func (m *FilmsDAO) Patch(id int, p Patch) error {
	if p.Key != "name" && p.Key != "year" && p.Key != "votes" {
		return fmt.Errorf("not supported field: %s", p.Key)
	}

	return m.db.Model(Film{ID: id}).Update(p.Key, p.Value).Error
}

func (m *FilmsDAO) Add(f FilmUpdate) (*Film, error) {
	c := Film{}
	c.Name = f.Row.Name
	c.Votes = int(f.Row.Votes)
	c.Year = int(f.Row.Year)

	err := m.db.Create(&c).Error
	return &c, err
}
