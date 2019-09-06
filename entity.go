package main

type Entity struct {
	tableName struct{} `sql:"entity" json:"-"` // nolint
	Id        int      `sql:",pk" json:"id"`
	Name      string   `json:"name"`
	Timestamps
}

func (entity *Entity) Get() (err error) {
	err = db.Model(entity).WherePK().
		Select()
	return
}

func (entity *Entity) Create() (err error) {
	err = db.Insert(entity)
	return
}

func (entity *Entity) Update() (err error) {
	err = db.Update(entity)
	return
}

func (entity *Entity) Delete() (err error) {
	err = db.Delete(entity)
	return
}
