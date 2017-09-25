package main

import "log"

var (
	fooKey   = "foo"
	fooValue = "bar"

	mooKey   = "moo"
	mooValue = int32(7)

	gooKey   = "goo"
	gooValue = map[string]interface{}{"boo": "baz"}
)

func main() {
	to := 20000

	for i := 0; i < to; i++ {

		m := map[string]interface{}{
			fooKey: fooValue,
			mooKey: mooValue,
			gooKey: gooValue,
		}

		if m["foo"] != "bar" || m["moo"] != int32(7) || m["goo"] == nil {
			log.Fatalf("your test is broken: %v", m)
		}
	}

	for i := 0; i < to; i++ {
		m := make(map[string]interface{})
		m[fooKey] = fooValue
		m[mooKey] = mooValue
		m[gooKey] = gooValue

		if m["foo"] != "bar" || m["moo"] != int32(7) || m["goo"] == nil {
			log.Fatalf("your test is broken: %v", m)
		}
	}

	for i := 0; i < to; i++ {
		m := make(map[string]interface{})
		SetString(m, fooKey, fooValue)
		SetInt32(m, mooKey, mooValue)
		SetMap(m, gooKey, gooValue)

		if m["foo"] != "bar" || m["moo"] != int32(7) || m["goo"] == nil {
			log.Fatalf("your test is broken: %v", m)
		}
	}
}

func SetString(m map[string]interface{}, field string, value string) {
	m[field] = value
}

func SetInt32(m map[string]interface{}, field string, value int32) {
	m[field] = value
}

func SetMap(m map[string]interface{}, field string, value map[string]interface{}) {
	m[field] = value
}
