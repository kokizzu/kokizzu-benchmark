package main

import (
	"math/rand"
	"time"

	"gitlab.com/kokizzu/gokil/I"
	"gitlab.com/kokizzu/gokil/K"
	"gitlab.com/kokizzu/gokil/L"
	"gitlab.com/kokizzu/gokil/M"
	"gitlab.com/kokizzu/gokil/S"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const TABLE_NAME = `test3`

type Person struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	CreatedBy int64
	UpdatedBy int64
	DeletedBy int64
	UniqueId  string
	// created_at ObjectId(Id).getTimeStamp()
	Data map[string]interface{}
}

func main() {
	MO, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer MO.Close()
	MO.SetMode(mgo.Monotonic, true)

	var t time.Time

	coll := MO.DB(`test2`).C(TABLE_NAME)

	coll.DropCollection()

	//	query := K.RunCmd(`bash`, `-c`, `echo`, `db.test2.stats()`, `|`, `mongo`)
	//	L.Print(string(query))

	coll.Create(&mgo.CollectionInfo{})

	coll.EnsureIndex(mgo.Index{
		Key:      []string{"uniqueid"},
		Unique:   true,
		DropDups: true,
	})

	coll.EnsureIndex(mgo.Index{
		Key: []string{"data.name"},
	})
	coll.EnsureIndex(mgo.Index{
		Key: []string{"data.age"},
	})

	benchmarks := M.SI{
		`micro`: 100,
		`tiny`:  1000,
		`small`: 10000,
	}
	for key, count := range benchmarks {
		t = time.Now()
		for z := int64(0); z < count; z++ {
			err := coll.Insert(
				&Person{
					CreatedBy: 1,
					UniqueId:  key + `-` + I.ToS(z),
					Data: map[string]interface{}{
						`name`: S.RandomPassword(16),
						`age`:  rand.Int63(),
					},
				})
			if err != nil {
				panic(err)
			}
			// TODO: add transaction
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
			result := Person{}
			err := coll.Find(bson.M{`uniqueid`: unique_id}).One(&result)

			if err == nil {
				found++
			} else {
				not_found++
			}
		}
		L.Print(`found/not: ` + I.ToS(found) + `/` + I.ToS(not_found))
		L.TimeTrack(t, `SEARCH `+key+` `+I.ToS(count))
	}

}
