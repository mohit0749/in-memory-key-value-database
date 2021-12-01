package keyvaluestore

import (
	"fmt"
	"testing"
)

func TestNewKeyValueStore(t *testing.T) {
	kvStore := NewKeyValueStore()
	t.Run("test add", func(t *testing.T) {
		kvStore.Add("delhi", "pollution_level", "high")
		kvStore.Add("delhi", "population", "10million")

		kvStore.Add("jakarata", "pollution_level", "high")
		kvStore.Add("jakarata", "latitude", -6.0)
		kvStore.Add("jakarata", "longitude", 106)

		kvStore.Add("banglore", "latitude", 12.94)
		kvStore.Add("banglore", "longitude", 77.64)
		kvStore.Add("banglore", "pollution_level", "high")
		kvStore.Add("banglore", "free_food", true)

		kvStore.Add("india", "capital", "delhi")
		kvStore.Add("india", "population", "1.2B")

		kvStore.Add("crocin", "category", "Cold & Flu")
		kvStore.Add("crocin", "manufacturer", "GSK")

		v, _ := kvStore.Get("delhi")
		fmt.Println(v)
		v, _ = kvStore.Get("jakarata")
		fmt.Println(v)
		v, _ = kvStore.Get("banglore")
		fmt.Println(v)
		v, _ = kvStore.Get("india")
		fmt.Println(v)
		v, _ = kvStore.Get("crocin")
		fmt.Println(v)

		val, err := kvStore.Scan("pollution_level", "high")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(val)

		kvStore.Remove("delhi")
		val, err = kvStore.Scan("pollution_level", "high")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(val)
		v, err = kvStore.Get("delhi")
		fmt.Println(v)
	})
}
