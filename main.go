package main

import (
    "log"
    "net/http"
    "go_sql/api"  // 導入 api 包
)

func main() {
    // 設置路徑 /data 並處理請求
    http.HandleFunc("/data/api", api.QueryDataHandler)  // 使用 api 包中的 QueryDataHandler

    // 啟動服務器，監聽端口 8081
    log.Println("伺服器正在運行：http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}