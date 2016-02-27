package main

import (
	"math/rand"
	"time"

	"github.com/garyburd/redigo/redis"
	"gitlab.com/kokizzu/gokil/I"
	"gitlab.com/kokizzu/gokil/K"
	"gitlab.com/kokizzu/gokil/L"
	"gitlab.com/kokizzu/gokil/M"
	"gitlab.com/kokizzu/gokil/S"
	"gitlab.com/kokizzu/gokil/T"
)

const TABLE_NAME = `test3`

func main() {
	redisConnectors := redis.Pool{
		MaxIdle:   8,
		MaxActive: 4096,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(`tcp`, `:6379`)
			L.PanicIf(err, `failed to connect in-memory database`)
			_, err = c.Do(`SELECT`, 3)
			L.PanicIf(err, `failed to select default in-memory database`)
			return c, err
		},
	}

	RE := redisConnectors.Get()
	defer RE.Close()

	var t time.Time

	//	query := K.RunCmd(`bash`, `-c`, `echo`, `db.test2.stats()`, `|`, `mongo`)
	//	L.Print(string(query))

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

			name := S.RandomPassword(16)
			age := rand.Int63()
			err := RE.Send(`MULTI`)
			L.PanicIf(err, `MULTI`)
			coded := K.ToGOB64(M.SX{
				`created_at`: T.NowStr(),
				`created_by`: 1,
				`data`: M.SX{
					`name`: name,
					`age`:  age,
				},
			})
			err = RE.Send(`SET`, unique_id, coded)
			L.PanicIf(err, `SET`)
			err = RE.Send(`HSET`, `names_`+name, unique_id, 1)
			L.PanicIf(err, `HSET names`)
			err = RE.Send(`HSET`, `ages_`+I.ToS(age), unique_id, 1)
			L.PanicIf(err, `HSET age`)
			err = RE.Send(`EXEC`)
			L.PanicIf(err, `EXEC`)
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
			rec, err := RE.Do(`GET`, unique_id)
			if err == nil {
				found++
				if val, ok := rec.([]uint8); ok {
					_ = K.FromGOB64(string(val))
				}
			} else {
				not_found++
			}
		}
		L.Print(`found/not: ` + I.ToS(found) + `/` + I.ToS(not_found))
		L.TimeTrack(t, `SEARCH `+key+` `+I.ToS(count))
	}

}
