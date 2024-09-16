package db

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

// ConnectDB 函數負責連接 MySQL 資料庫
func ConnectDB() (*sql.DB, error) {
    dbHost := "localhost"
    dbUsername := "root"
    dbPassword := "00000000"
    database := "flavors"

    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbUsername, dbPassword, dbHost, database)

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("連接失敗: %v", err)
    }

    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("無法連接到數據庫: %v", err)
    }

    return db, nil
}