


package arrary_test

import "testing"


<<<<<<< HEAD
func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	t.Log(arr[1], arr[2])
=======
func TestArrayInit(t *testing.T){
  var arr[3]int
  arr1:=[4]int{1,2,3,4}
  arr3:=[...]int{1,3,4,5}
  t.Log(arr[1],arr[2])
>>>>>>> e3385d93f9e10317b74fd4eadf6d2da2361fe6e4

  t.Log(arr1,arr3)
}

func TestArrTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for idx, v := range arr3 {
		t.Log(idx, v)
	}

}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3_sec := arr3[3:]
	t.Log(arr3_sec)
}

//slice
