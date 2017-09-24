package main

import "log"

func main() {
	to := 20000

	for i := 0; i < to; i++ {

		m := map[string]interface{}{
			"foo": "bar",
			"moo": int32(7),
			"goo": map[string]interface{}{"boo": "baz"},
		}

		if m["foo"] != "bar" || m["moo"] != int32(7) || m["goo"] == nil {
			log.Fatalf("your test is broken: %v", m)
		}
	}

	for i := 0; i < to; i++ {
		m := make(map[string]interface{})
		m["foo"] = "bar"
		m["moo"] = int32(7)
		m["goo"] = map[string]interface{}{"boo": "baz"}

		if m["foo"] != "bar" || m["moo"] != int32(7) || m["goo"] == nil {
			log.Fatalf("your test is broken: %v", m)
		}
	}

	for i := 0; i < to; i++ {
		m := make(map[string]interface{})
		SetString(m, "foo", "bar")
		SetInt32(m, "moo", int32(7))
		SetMap(m, "goo", map[string]interface{}{"boo": "baz"})

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
