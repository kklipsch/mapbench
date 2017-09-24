package main

import "testing"

func BenchmarkMapInit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := map[string]interface{}{
			"foo": "bar",
			"moo": int32(7),
			"goo": map[string]interface{}{"boo": "baz"},
		}

		if m["foo"] != "bar" || m["moo"] != int32(7) || m["goo"] == nil {
			b.Errorf("your test is broken: %v", m)
		}
	}
}

func BenchmarkMapSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := make(map[string]interface{})
		m["foo"] = "bar"
		m["moo"] = int32(7)
		m["goo"] = map[string]interface{}{"boo": "baz"}

		if m["foo"] != "bar" || m["moo"] != int32(7) || m["goo"] == nil {
			b.Errorf("your test is broken: %v", m)
		}
	}
}

func BenchmarkMapFunctions(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := make(map[string]interface{})
		SetString(m, "foo", "bar")
		SetInt32(m, "moo", int32(7))
		SetMap(m, "goo", map[string]interface{}{"boo": "baz"})

		if m["foo"] != "bar" || m["moo"] != int32(7) || m["goo"] == nil {
			b.Errorf("your test is broken: %v", m)
		}
	}
}

func BenchmarkMapSetAlloc(b *testing.B) {
	m := make(map[string]interface{})
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		m["foo"] = "bar"
		m["moo"] = int32(7)
		m["goo"] = map[string]interface{}{"boo": "baz"}

		if m["foo"] != "bar" || m["moo"] != int32(7) || m["goo"] == nil {
			b.Errorf("your test is broken: %v", m)
		}
	}
}

func BenchmarkMapFunctionsAlloc(b *testing.B) {
	m := make(map[string]interface{})
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		SetString(m, "foo", "bar")
		SetInt32(m, "moo", int32(7))
		SetMap(m, "goo", map[string]interface{}{"boo": "baz"})

		if m["foo"] != "bar" || m["moo"] != int32(7) || m["goo"] == nil {
			b.Errorf("your test is broken: %v", m)
		}
	}
}
