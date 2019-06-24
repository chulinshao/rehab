package datasource

import (
	"fmt"
	config2 "github.com/chulinshao/rehab/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"log"
	"sync"
)

var (
	engine *xorm.Engine
	lock   sync.Mutex
)

func GetEngine() *xorm.Engine {
	if engine != nil {
		return engine
	}
	lock.Lock()
	defer lock.Unlock()

	if engine != nil {
		return engine
	}
	config := config2.GetConfig()
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", config.Db.Username, config.Db.Password,
		config.Db.Ip, config.Db.Port, config.Db.DbName)
	dbEngine, err := xorm.NewEngine(config.Db.DriverName, driveSource)
	if err != nil {
		log.Fatal("init db engine error")
	} else {
		dbEngine.SetMaxIdleConns(config.Db.MaxIdleConns)
		dbEngine.SetMaxOpenConns(config.Db.MaxOpenConns)
		engine = dbEngine
		engine.ShowSQL(config.Db.ShowLog)
		if config.Db.Debuglog {
			engine.Logger().SetLevel(core.LOG_DEBUG)
		}
	}

	// 设置前缀
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "dzy_")
	engine.SetTableMapper(tbMapper)
	return engine
}
