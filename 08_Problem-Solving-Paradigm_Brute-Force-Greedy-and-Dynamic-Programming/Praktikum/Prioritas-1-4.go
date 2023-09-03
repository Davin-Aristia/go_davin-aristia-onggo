package main

import(
	"fmt"
	"math"
	"strconv"
)

func Kubik(a, b, c, d float64) int {
	tolerance := 1e-6
    maxIterations := 100
    x := 1.0
    for i := 0; i < maxIterations; i++ {
        f := a*x*x*x + b*x*x + c*x + d
        df := 3*a*x*x + 2*b*x + c
        x -= f / df

        if math.Abs(f) < tolerance {
            return int(math.Round(x))
        }
    }
    return int(math.NaN())
}

func Kuadrat(a, b, c float64) int {
    discriminant := b*b - 4*a*c

    if discriminant < 0 {
        return int(math.NaN())
    } else if discriminant == 0 {
        root := -b / (2 * a)
        return int(math.Max(root, root))
    } else {
        sqrtDiscriminant := math.Sqrt(discriminant)
        root1 := (-b + sqrtDiscriminant) / (2 * a)
        root2 := (-b - sqrtDiscriminant) / (2 * a)
        return int(math.Max(root1, root2))
    }
}

func SimpleEquations(a, b, c int) interface{}{
	root := Kubik(-2, float64(4*a), float64(-3*a*a+c), float64(-2*b+a*a*a-a*c))

    if math.IsNaN(float64(root)){
      return "no solution"
    } else {
      z := a-root
      y := Kuadrat(-1, float64(root), float64(-b/z))
      x := a-z-y
      if x+y+z == a && x*y*z == b && x*x+y*y+z*z==c{
        return strconv.Itoa(x) + " " + strconv.Itoa(y) + " " + strconv.Itoa(z)
      }else{
        return "no solution"
      }
    }
}


func main() {
	fmt.Println(SimpleEquations(1, 2, 3)) // no solution
	fmt.Println(SimpleEquations(6, 6, 14)) // 1 2 3
	fmt.Println(SimpleEquations(18, 210, 110)) // 5 6 7
	fmt.Println(SimpleEquations(24, 512, 192)) // 8 8 8
}