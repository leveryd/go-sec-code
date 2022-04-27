package conf

import (
	"database/sql"
	"log"
	"os"
)

var (
	ConfigFile = ""
	ServerHost = "0.0.0.0"
	ServerPort = 8089

	// DataSourceName [username[:password]@][protocol[(address)]]/dbname[?param1=value1&paramN=valueN]
	DataSourceName = ""

	// Database
	username = "root"
	password = "123456"
	protocol = "tcp"
	address  = "127.0.0.1"
	dbname   = "test"
)

// 环境变量优于配置文件
func set(variable *string, envName string) {
	if value := os.Getenv(envName); value != "" {
		*variable = value
	}
}

func init() {
	set(&username, "username")
	set(&password, "password")
	set(&protocol, "protocol")
	set(&address, "address")
	set(&dbname, "dbname")

	if os.Getenv("DataSourceName") != "" {
		DataSourceName = os.Getenv("DataSourceName")
	} else {
		DataSourceName = username + ":" + password + "@" + protocol + "(" + address + ")/" + dbname + "?charset=utf8&multiStatements=true" // 允许执行多条语句 https://github.com/go-sql-driver/mysql#multistatements
	}

	db, err := sql.Open("mysql", DataSourceName)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	if err != nil {
		log.Println(err)
		return
	}

	// 创建数据库
	{
		_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// 创建表
	{
		createTableSql := `
			use ` + dbname + `;
			CREATE TABLE IF NOT EXISTS user (
				id INT UNSIGNED AUTO_INCREMENT,
				username VARCHAR(255) NOT NULL,
				password VARCHAR(255) NOT NULL,
				PRIMARY KEY (id)
			)ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`
		_, err := db.Exec(createTableSql)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// 初始化表数据
	{
		initTableSql := `
			use ` + dbname + `;
			INSERT INTO user VALUES (1, 'admin', 'admin123');
			INSERT INTO user VALUES (2, 'test', 'test123');
		`
		_, err := db.Exec(initTableSql)
		if err != nil {
			return
		}
	}
}
