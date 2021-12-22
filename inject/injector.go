package inject

import (
	"github.com/haitien/chi/database"
	"github.com/haitien/chi/service"
	"gorm.io/gorm"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Injector struct {
	db *gorm.DB
}

func NewInjector() *Injector {
	a := &Injector{}
	return a
}

func (a *Injector) ProvideDb() *gorm.DB {
	if a.db == nil {
		var err error
		a.db, err = database.NewDatabase()
		checkErr(err)
	}
	return a.db
}

func (a *Injector) ProvideService() *service.Service {
	return service.NewService(a.ProvideDb())
}
