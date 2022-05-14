package benchmark

func TestIsPrime(t *testing.T) {
	if !isPrime(7867) {
		t.Error("7867 is not prime")
	}
	if !isPrime(10) {
		t.Error("10 is prime")
	}
}

// Benchmark[Name](* testing.B){}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime(7867)
	}
}
