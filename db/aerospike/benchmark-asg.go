package main

import (
	"math/rand"
	"time"

	"github.com/aerospike/aerospike-client-go"
	"gitlab.com/kokizzu/gokil/I"
	"gitlab.com/kokizzu/gokil/L"
	"gitlab.com/kokizzu/gokil/M"
	"gitlab.com/kokizzu/gokil/S"
	"gitlab.com/kokizzu/gokil/T"
)

const TABLE_NAME = `test3`

func main() {
	client, err := aerospike.NewClient("127.0.0.1", 3000)
	L.PanicIf(err, `newclient`)

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
			key, err := aerospike.NewKey(`test2`, TABLE_NAME, unique_id)
			L.PanicIf(err, `newkey`)

			bin1 := aerospike.NewBin("name", name)
			bin2 := aerospike.NewBin("age", age)
			bin3 := aerospike.NewBin("created_by", 1)
			bin4 := aerospike.NewBin("created_at", T.NowStr())
			// TODO: create index, but no asc for ArchLinux

			// Write a record
			err = client.PutBins(nil, key, bin1, bin2, bin3, bin4)
			L.PanicIf(err, `putbins`)

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
			key, err := aerospike.NewKey(`test2`, TABLE_NAME, unique_id)
			L.PanicIf(err, `newkey`)
			rec, err := client.Get(nil, key)
			if err == nil && len(rec.Bins) != 0 {
				found++
			} else {
				not_found++
			}
		}
		L.Print(`found/not: ` + I.ToS(found) + `/` + I.ToS(not_found))
		L.TimeTrack(t, `SEARCH `+key+` `+I.ToS(count))
	}

	client.Close()
}
