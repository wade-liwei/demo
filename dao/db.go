package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	singleDB *sql.DB
)

//  UPDATE user SET password=PASSWORD("root") WHERE user='root';
//  update user set authentication_string=PASSWORD("root") where User='root'; #更改密码
//  update user set plugin="mysql_native_password"; #如果没这一行可能也会报一个错误，因此需要运行这一行
//  flush privileges; #更新所有操作权限
//  quit

func GetSqlInstance() (*sql.DB, error) {

	if singleDB != nil {
		return singleDB, nil
	}

	db, err := ConnectMysqlDatabase()
	if err != nil {
		return nil, err
	}
	singleDB = db

	return singleDB, nil
}

func ConnectMysqlDatabase() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/atlasDemo")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
