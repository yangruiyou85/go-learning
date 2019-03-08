package my_map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 0}
	t.Log(m1[2])

	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))

}

func TestAccessNotExistingKey(t *testing.T) {
	//
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 100

	if v, ok := m1[3]; ok {
		t.Logf("m1[3] is %d", v)
	} else {
		t.Log("key 3 is not exist.")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 2, 3: 4, 10: 200}

	for k, v := range m1 {
		t.Log(k, v)
	}
}
