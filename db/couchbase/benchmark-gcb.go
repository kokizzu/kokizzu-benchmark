package main

import (
	"math/rand"
	"time"

	"github.com/couchbase/go-couchbase"
	// github.com/couchbase/gocb requires go 1.5+ maybe
	"gitlab.com/kokizzu/gokil/I"
	"gitlab.com/kokizzu/gokil/L"
	"gitlab.com/kokizzu/gokil/M"
	"gitlab.com/kokizzu/gokil/S"
)

const TABLE_NAME = `test3`

func main() {
	CB, err := couchbase.Connect("http://Administrator:Administrator1@127.0.0.1:8091/")
	L.PanicIf(err, `connect`)

	var t time.Time

	pool, err := CB.GetPool(`default`)
	L.PanicIf(err, `pool`)

	//	query := K.RunCmd(`bash`, `-c`, `echo`, `db.test2.stats()`, `|`, `mongo`)
	//	L.Print(string(query))

	buck, err := pool.GetBucket(TABLE_NAME)
	L.PanicIf(err, `bucket`)

	// TODO: create index http://developer.couchbase.com/documentation/server/4.0/n1ql/n1ql-language-reference/createindex.html

	benchmarks := M.SI{
		`micro`: 100,
		`tiny`:  1000,
		`small`: 10000,
	}
	for key, count := range benchmarks {
		t = time.Now()
		for z := int64(0); z < count; z++ {
			unique_id := key + `-` + I.ToS(z)
			err := buck.Set(unique_id, 0, map[string]interface{}{
				`created_by`: 1,
				`data`: map[string]interface{}{
					`name`: S.RandomPassword(16),
					`age`:  rand.Int63(),
				},
				`created_at`: time.Now(),
			})
			_ = err
			//L.PanicIf(err, `set `+unique_id)
			// TODO: add transaction?
		}
		L.TimeTrack(t, `INSERT `+key+` `+I.ToS(count))
	}

	//	query = K.RunCmd(`bash`, `-c`, `echo`, `db.test2.stats()`, `|`, `mongo`)
	//	L.Print(string(query))

	for key, count := range benchmarks {
		t = time.Now()

		found := int64(0)
		not_found := int64(0)
		for z := int64(0); z < count*40; z++ {
			unique_id := key + `-` + I.ToS(rand.Int63()%count)
			res := map[string]interface{}{}
			err := buck.Get(unique_id, res)
			if err == nil && len(res) == 0 {
				found++
			} else {
				not_found++
			}
		}
		L.Print(`found/not: ` + I.ToS(found) + `/` + I.ToS(not_found))
		L.TimeTrack(t, `SEARCH `+key+` `+I.ToS(count))
	}

}
