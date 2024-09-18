package api

import (
    "fmt"
    "net/http"
    "strconv"
    "go_sql/db" // 使用資料庫連接的包
)

// DeleteRecordHandler 處理 /delete 路徑，用於刪除指定 id 的記錄
func DeleteRecordHandler(w http.ResponseWriter, r *http.Request) {
    // 確認是否有傳入 id 參數
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "無效的 ID", http.StatusBadRequest)
        return
    }

    // 將 id 參數轉換為整數
    id, err := strconv.Atoi(idStr)
    if err != nil || id <= 0 {
        http.Error(w, "無效的 ID", http.StatusBadRequest)
        return
    }

    // 連接資料庫
    dbConn, err := db.ConnectDB()
    if err != nil {
        http.Error(w, "資料庫連接失敗", http.StatusInternalServerError)
        return
    }
    defer dbConn.Close()

    // 執行 DELETE 查詢
    deleteQuery := fmt.Sprintf("DELETE FROM `%s` WHERE id = ?", "flavors")
    result, err := dbConn.Exec(deleteQuery, id)
    if err != nil {
        http.Error(w, "刪除失敗: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // 確認是否有刪除成功
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        http.Error(w, "獲取刪除結果失敗: "+err.Error(), http.StatusInternalServerError)
        return
    }

    if rowsAffected > 0 {
        fmt.Fprintf(w, "記錄已刪除，id: %d", id)
    } else {
        fmt.Fprintf(w, "沒有找到對應的記錄，id: %d", id)
    }
}