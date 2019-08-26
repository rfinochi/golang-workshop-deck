func BenchmarkFibonacciWrong(b *testing.B) {
        for n := 0; n < b.N; n++ {
                Fibonacci(n)
        }
}

func BenchmarkFibonacciWrong2(b *testing.B) {
        Fibonacci(b.N)
}