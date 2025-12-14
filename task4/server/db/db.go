package db

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"metanode.com/homework/server/config"
)

var (
	dbInstance *gorm.DB
	once       sync.Once // 确保只初始化一次
)

// GetDB 获取全局数据库连接实例
func GetDB() *gorm.DB {
	once.Do(func() {
		dbInstance = initDB()
	})
	return dbInstance
}

// initDB 初始化数据库连接
func initDB() *gorm.DB {
	cfg := config.GetDatabaseConfig()

	// 构建 DSN (Data Source Name)
	dsn := buildDSN(cfg)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 可以在这里配置 GORM 的其他选项
		// SkipDefaultTransaction: true, // 跳过默认事务（提升性能）
		// PrepareStmt:             true,  // 缓存预编译语句（提升性能）
	})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取数据库实例失败: %v", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxOpenConns(cfg.MaxConnections)     // 最大打开连接数
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections) // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(0)                   // 连接最大生命周期（0表示无限制）

	log.Println("数据库连接成功")

	return db
}

// buildDSN 构建数据库连接字符串
func buildDSN(cfg config.DatabaseConfig) string {
	return cfg.User + ":" + cfg.Password +
		"@tcp(" + cfg.Host + ":" + cfg.Port + ")/" +
		cfg.Name + "?charset=" + cfg.Charset +
		"&parseTime=True&loc=Local"
}
