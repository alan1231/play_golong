package queries

import (
    "database/sql"
    "fmt"
)

// QueryData 函數，負責執行 SELECT 查詢並返回結果
func QueryData(dbConn *sql.DB, query string) ([]map[string]interface{}, error) {
    rows, err := dbConn.Query(query)
    if err != nil {
        return nil, fmt.Errorf("查詢失敗: %v", err)
    }
    defer rows.Close()

    var data []map[string]interface{}
    cols, err := rows.Columns()
    if err != nil {
        return nil, fmt.Errorf("無法獲取列名: %v", err)
    }

    for rows.Next() {
        row := make(map[string]interface{})
        colValues := make([]interface{}, len(cols))
        colPointers := make([]interface{}, len(cols))

        for i := range colValues {
            colPointers[i] = &colValues[i]
        }

        if err := rows.Scan(colPointers...); err != nil {
            return nil, fmt.Errorf("讀取行數據失敗: %v", err)
        }

        for i, colName := range cols {
            row[colName] = colValues[i]
        }

        data = append(data, row)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("查詢過程中出錯: %v", err)
    }

    return data, nil
}