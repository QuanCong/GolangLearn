package unit

import "testing"

func TestAdd(t *testing.T) {
	value := Add(1, 2)
	if value != 3 {
		t.Error("expect3 but", value)
	}

}

func BenchmarkAdd(b *testing.B) {
	b.StopTimer();
	//do something ignore
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}