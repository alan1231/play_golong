package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "go_sql/db"      // 導入 db 包，負責數據庫連接
    "go_sql/queries" // 導入 queries 包，負責查詢邏輯
)

// 設置資料表名稱
var tableName = "flavors" // 替換成實際的資料表名稱

// 查詢資料並返回 JSON 格式
func queryDataHandler(w http.ResponseWriter, r *http.Request) {
    // 使用 ConnectDB 函數連接資料庫
    dbConn, err := db.ConnectDB()  // 使用 db 包中的 ConnectDB 函數
    if err != nil {
        http.Error(w, "資料庫連接失敗", http.StatusInternalServerError)
        return
    }
    defer dbConn.Close()

    // 構建查詢語句
    query := fmt.Sprintf("SELECT * FROM `%s`", tableName)

    // 調用 queries 包中的 QueryData 函數來執行查詢
    data, err := queries.QueryData(dbConn, query)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // 設置 Content-Type 並返回 JSON 格式的數據
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}

func main() {
    // 設置路由並啟動服務器
    http.HandleFunc("/data", queryDataHandler)
    log.Println("伺服器正在運行：http://localhost:8080/data")
    log.Fatal(http.ListenAndServe(":8080", nil))
}