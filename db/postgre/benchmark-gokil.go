package main

import (
	"math/rand"
	"time"

	_ "github.com/lib/pq"
	"gitlab.com/kokizzu/gokil/D"
	"gitlab.com/kokizzu/gokil/I"
	"gitlab.com/kokizzu/gokil/K"
	"gitlab.com/kokizzu/gokil/L"
	"gitlab.com/kokizzu/gokil/M"
	"gitlab.com/kokizzu/gokil/S"
)

const TABLE_NAME = `test3`

const CREATE_TABLE = `
CREATE TABLE ` + TABLE_NAME + `(
  id BIGSERIAL PRIMARY KEY,
  created_by BIGINT,
  updated_by BIGINT,
  deleted_by BIGINT,
  unique_id VARCHAR(256) UNIQUE,
  created_at timestamp default current_timestamp,
  data JSONB
);`

const CREATE_INDEX1 = `CREATE INDEX ` + TABLE_NAME + `__data__name ON ` + TABLE_NAME + ` ( (data->>'name') )`
const CREATE_INDEX2 = `CREATE INDEX ` + TABLE_NAME + `__data__age ON ` + TABLE_NAME + ` ( ((data->>'age')::BIGINT) )`

const DROP_TABLE = `DROP TABLE IF EXISTS test3 CASCADE`

func main() {
	PG := D.NewPgConn(`postgres`, `test2`)
	var t time.Time

		PG.DoTransaction(func(tx *D.Tx) string {
			tx.DoExec(DROP_TABLE)
			return ``
		})

		init_size := K.RunCmd(`psql`, `-U`, `test2`, `-c`, `\l+ test2`)
		L.Print(string(init_size))

		PG.DoTransaction(func(tx *D.Tx) string {
			tx.DoExec(CREATE_TABLE)
			return ``
		})

		PG.DoTransaction(func(tx *D.Tx) string {
			tx.DoExec(CREATE_INDEX1)
			tx.DoExec(CREATE_INDEX2)
			return ``
		})

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
				PG.DoInsert(1, TABLE_NAME, M.SX{
					`data`:      M.ToJson(data),
					`unique_id`: key + `-` + I.ToS(z),
				})
			}
			L.TimeTrack(t, `INSERT `+key+` `+I.ToS(count))
		}

		final_size := K.RunCmd(`psql`, `-U`, `test2`, `-c`, `\l+ test2`)
		L.Print(string(final_size))

	for key, count := range benchmarks {
		t = time.Now()

		found := int64(0)
		not_found := int64(0)
		for z := int64(0); z < count*40; z++ {
			unique_id := key + `-` + I.ToS(rand.Int63()%count)
			m := PG.QFirstMap(`SELECT id, unique_id, data->>'name', (data->>'age')::BIGINT age FROM ` + TABLE_NAME + ` WHERE unique_id = ` + D.Z(unique_id))
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
