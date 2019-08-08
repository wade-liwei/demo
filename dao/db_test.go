package dao_test

import (
	"github.com/wade-liwei/demo/dao"
	"testing"
)

func TestConnectDatabase(t *testing.T) {
	db, err := dao.ConnectMysqlDatabase()

	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	rows, err := db.Query("show databases")

	if err != nil {
		t.Error(err)
	}
	defer rows.Close()

	cols, err := rows.Columns()

	if err != nil {
		t.Error(err)
	}

	t.Log("show databases:")
	for k, v := range cols {
		t.Logf("k: %v  v: %v \n", k, v)
	}

	vals := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))

	for k, _ := range vals {
		scans[k] = &vals[k]
	}

	i := 0
	result := make(map[int]map[string]string)

	for rows.Next() {
		rows.Scan(scans...)

		row := make(map[string]string)
		for k, v := range vals {
			key := cols[k]
			t.Log(string(v))
			row[key] = string(v)
		}
		result[i] = row
		i++
	}

	for idx, row := range result {
		t.Logf("row number: %v \n", idx)
		for k, v := range row {
			t.Logf("rowK: %v  rowV: %v \n", k, v)
		}
	}
}
