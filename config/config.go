package config

import (
	"github.com/CyrivlClth/snowflake"
	"github.com/CyrivlClth/snowflake/idgen"
)

var idGen idgen.IDGenerator

// Init 初始化config
func Init(workerID, dataCenterID int64) (err error) {
	idGen, err = snowflake.New(workerID, dataCenterID)
	return
}

// IDGenerator 获取全局ID生成器
func IDGenerator() idgen.IDGenerator {
	return idGen
}
