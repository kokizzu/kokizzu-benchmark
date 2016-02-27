package main

import (
	"bytes"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/kokizzu/gokil/D"
	"gitlab.com/kokizzu/gokil/I"
	"gitlab.com/kokizzu/gokil/L"
	"gitlab.com/kokizzu/gokil/M"
	"gitlab.com/kokizzu/gokil/S"
)

const TABLE_NAME = `test3`

const CREATE_TABLE = `
CREATE TABLE ` + TABLE_NAME + `(
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  created_by BIGINT,
  updated_by BIGINT,
  deleted_by BIGINT,
  unique_id VARCHAR(256) UNIQUE,
  created_at timestamp default current_timestamp,
  data JSON
);`

const DROP_TABLE = `DROP TABLE IF EXISTS test3 CASCADE`

func GenInsertMySql(table string, kvparams M.SX) (string, []interface{}) {
	query := bytes.Buffer{}
	query.WriteString(`INSERT INTO ` + table + `( `)
	len := 0
	params := []interface{}{}
	for key, val := range kvparams {
		if len > 0 {
			query.WriteString(`, `)
		}
		query.WriteString(key)
		params = append(params, val)
		len++
	}
	query.WriteString(` ) VALUES ( `)
	for z := 1; z <= len; z++ {
		if z > 1 {
			query.WriteString(`, `)
		}
		query.WriteString(`?`)
	}
	query.WriteString(` )`)
	return query.String(), params
}

func main() {
	MY := D.NewMyConn(`root`, `test2`, `tcp(:3306)`, ``)
	var t time.Time

	MY.DoTransaction(func(tx *D.Tx) string {
		tx.DoExec(DROP_TABLE)
		return ``
	})

	query := `SELECT COALESCE((SELECT ROUND(((data_length + index_length) / 1024 ), 2) AS "KB"
	FROM information_schema.TABLES
	WHERE table_schema = ` + D.Z(TABLE_NAME) + `),0)`
	L.Print(MY.QFloat(query))

	MY.DoTransaction(func(tx *D.Tx) string {
		tx.DoExec(CREATE_TABLE)
		return ``
	})

	// JSON columns cannot be indexed. https://dev.mysql.com/doc/refman/5.7/en/json.html

	benchmarks := M.SI{
		`micro`: 100,
		`tiny`:  1000,
		`small`: 10000,
	}
	for key, count := range benchmarks {
		t = time.Now()
		for z := int64(0); z < count; z++ {
			data := M.SX{
				`name`: S.RandomPassword(16),
				`age`:  rand.Int63(),
			}
			MY.DoTransaction(func(tx *D.Tx) string {
				query, params := GenInsertMySql(TABLE_NAME, M.SX{
					`data`:       M.ToJson(data),
					`created_by`: 1,
					`unique_id`:  key + `-` + I.ToS(z),
				})
				tx.DoExec(query, params...)
				return ``
			})
		}
		L.TimeTrack(t, `INSERT `+key+` `+I.ToS(count))
	}

	query = `SELECT ROUND(((data_length + index_length) / 1024 / 1024), 2) AS "MB"
FROM information_schema.TABLES
WHERE table_schema = "test2"`
	L.Print(MY.QFloat(query))

	for key, count := range benchmarks {
		t = time.Now()

		found := int64(0)
		not_found := int64(0)
		for z := int64(0); z < count*40; z++ {
			unique_id := key + `-` + I.ToS(rand.Int63()%count)
			m := MY.QFirstMap(`SELECT id, unique_id, data->'$.name', (data->"$.age") age FROM ` + TABLE_NAME + ` WHERE unique_id = ` + D.Z(unique_id))
			if len(m) > 0 {
				found++
			} else {
				not_found++
			}
		}
		L.Print(`found/not: ` + I.ToS(found) + `/` + I.ToS(not_found))
		L.TimeTrack(t, `SEARCH `+key+` `+I.ToS(count))
	}

}
