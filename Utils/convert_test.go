package Utils

import "testing"

func Test_Convert(t *testing.T) {
	val := MustInt(float64(1000), -1)
	if val == -1 {
		t.Fail()
		return
	}

	val64 := MustInt64(float32(1000), -1)
	if val64 == -1 {
		t.Fail()
		return
	}

	valU64 := MustUInt64(float32(1000), 0xffffff)
	if valU64 == 0xffffff {
		t.Fail()
		return
	}

	valF64 := MustFloat64(int(3000), 0.01)
	if valF64 == 0.01 {
		t.Fail()
		return
	}

	valstr := MustString(float32(1000))
	if len(valstr) <= 0 {
		t.Fail()
		return
	}
}

func Test_ConvertOther(t *testing.T) {
	val64 := MustInt64("1000", -1)
	if val64 == -1 {
		t.Fail()
		return
	}

	val64 = MustInt64("10", -1)
	if val64 == -1 {
		t.Fail()
		return
	}

	val64 = MustInt64("200000000", -1)
	if val64 == -1 {
		t.Fail()
		return
	}

	valF64 := MustFloat64("1000", 0.01)
	if valF64 == 0.01 {
		t.Fail()
		return
	}

	valF64 = MustFloat64("4.345", 0.01)
	if valF64 == 0.01 {
		t.Fail()
		return
	}

	valBool := MustBool("true", false)
	if valBool == false {
		t.Fail()
		return
	}
}

func Benchmark_ConvertOther(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := MustInt("10000000", -1)
		if val == -1 {
			return
		}
	}
}

func Benchmark_Convert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := MustInt(float64(1000), -1)
		if val == -1 {
			return
		}
	}
}

func Benchmark_ConvertOtherParllal(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			val := MustInt("10000000", -1)
			if val == -1 {
				return
			}
		}
	})
}

func Benchmark_ConvertParllal(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			val := MustInt(float64(1000), -1)
			if val == -1 {
				return
			}
		}
	})
}
