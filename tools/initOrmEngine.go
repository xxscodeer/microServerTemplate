package tools

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"microServerTemplate/domain/model"
)
var DbEngine *gorm.DB
func InitOrmEngine(dbConfig MysqlConfig){
	conn := dbConfig.DbUser+":"+dbConfig.DbPwd+"@/"+dbConfig.DbName+"?charset="+dbConfig.Charset+"&parseTime="+dbConfig.ParseTime+"&loc="+dbConfig.Loc

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{
		//SkipDefaultTransaction: true,
		PrepareStmt: true,
	})
	if err !=nil{
		log.Fatalln("conn mysql fail err",err)
	}
	//初始化表
	err = db.AutoMigrate(&model.Xxxx{})
	if err !=nil{
		log.Fatalln("init table fail,",err)
	}
	//赋值
	DbEngine = db
}
