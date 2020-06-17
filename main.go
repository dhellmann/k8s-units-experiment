package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/api/resource"
)

func main() {
	for _, s := range []string{
		"1.0",
		"1.23",
		"1.23Ki",
		"1000",
		"1.23Mi",
		"1000000",
		"1.0Gi",
		"1.23Gi",
	} {
		q := resource.MustParse(s)
		i := q.ScaledValue(3)
		fmt.Printf("s = %9q  q = %6s (%v) i = %d\n", s, q.String(), q, i)
	}

	fmt.Printf("%d\n", int(1.23*1000))
}
