// Example iteration package to demonstrate how use godoc
package iteration

// Repeat the input string n times.
func Repeat(input string, n int) (repeated string) {
	for i := 0; i < n; i++ {
		repeated += input
	}
	return
}
