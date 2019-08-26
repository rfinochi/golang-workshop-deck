func TestFibonacci(t *testing.T) {
	var in int = 10
	var want int = 55

	got := Fibonacci(in)
	if got != want {
		t.Errorf("Fibonacci(%d) == %d, want %d", in, got, want)
	}
}