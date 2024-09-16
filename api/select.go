package api

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "go_sql/db"       // 連接資料庫
    "go_sql/queries"  // 查詢資料庫
)

// QueryDataHandler 處理 /data/api 路徑，並根據 URL 查詢參數決定查詢邏輯
func QueryDataHandler(w http.ResponseWriter, r *http.Request) {
    // 獲取請求的詳細信息
    userIP := r.RemoteAddr                  // 用戶 IP 地址
    userPath := r.URL.Path                  // 請求的路徑
    userMethod := r.Method                  // 請求方法 (GET/POST)
    queryParams := r.URL.Query().Encode()   // 查詢參數

    // 記錄用戶的訪問信息
    log.Printf("用戶 IP: %s, 路徑: %s, 方法: %s, 查詢參數: %s", userIP, userPath, userMethod, queryParams)

    // 解析 URL 查詢參數
    tableName := r.URL.Query().Get("select")  // 提取 select 參數
    if tableName == "" {
        http.Error(w, "缺少 select 參數", http.StatusBadRequest)
        return
    }

    // 使用 ConnectDB 函數連接資料庫
    dbConn, err := db.ConnectDB()
    if err != nil {
        http.Error(w, "資料庫連接失敗", http.StatusInternalServerError)
        return
    }
    defer dbConn.Close()

    // 根據 select 參數動態構建查詢語句
    query := fmt.Sprintf("SELECT * FROM `%s`", tableName)

    // 調用 QueryData 函數來執行查詢
    data, err := queries.QueryData(dbConn, query)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // 返回 JSON 格式的查詢結果
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}