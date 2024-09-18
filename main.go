package main

import (
    "log"
    "net/http"
    "go_sql/api"  // 導入 api 包
)

func main() {
    // 設置路徑 /data 並處理請求
    http.HandleFunc("/data/api/list", api.QueryDataHandler)  // 使用 api 包中的 QueryDataHandler
    http.HandleFunc("/data/api/delete", api.DeleteRecordHandler) // 使用 DeleteRecordHandler


    // 啟動服務器，監聽端口 8081
    log.Println("伺服器正在運行：http://localhost:8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}