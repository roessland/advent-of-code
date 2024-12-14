package day13

import (
	"embed"
	"fmt"

	"github.com/roessland/advent-of-code/2024/aocutil"
	"github.com/roessland/gopkg/mathutil"
)

//go:embed input*.txt
var Input embed.FS

func Part12(inputName string) (int, int) {
	nums := aocutil.FSGetIntsInStringLines(Input, inputName)

	tokens1a := 0
	tokens1b := 0
	tokens2 := 0
	for i := 0; i < len(nums)-2; i += 4 {
		a1, a2 := nums[i][0], nums[i][1]
		ax, ay := a1, a2
		b1, b2 := nums[i+1][0], nums[i+1][1]
		bx, by := b1, b2
		y1, y2 := nums[i+2][0], nums[i+2][1]
		px, py := y1, y2

		{

			a := 100000*a1 + a2
			b := 100000*b1 + b2
			y := 100000*y1 + y2

			fmt.Println("---")
			fmt.Println("Button A:", ax, ay)
			fmt.Println("Button B:", bx, by)
			fmt.Println("Prize:", px, py)
			fmt.Println(a, b, y)
			// Part 1
			for xb := 100; xb >= 0; xb-- {
				if (y-b*xb)%a == 0 {
					xa := (y - b*xb) / a
					if xa > 100 {
						panic("yolo")
					}
					tokens1a += 3*xa + xb
					fmt.Println("xa=", xa, ", xb=", xb)
					break
				}
			}

			gcd := func(a, b int) int {
				return int(mathutil.GCD(int64(a), int64(b)))
			}

			// Part 1 different strat
			// y1, y2 = y1+offset, y2+offset
			fmt.Println("X: a1 xa + b1 xb = px")
			gx := gcd(ax, bx)
			gy := gcd(ay, by)
			fmt.Printf("X: %2d na + %2d nb = %2d\n", ax, bx, px)
			fmt.Printf("X: %2d na + %2d nb = %2d\n", ax/gx, bx/gx, px/gx)
			fmt.Printf("X: %2d na + %2d nb = %2d (mod %2d * %2d)\n", ax/gx, bx/gx, px/gx, ax/gx, bx/gx)
			fmt.Printf("X: %2d na + %2d nb = %2d (mod %2d)\n", ax/gx, bx/gx, px/gx, ax/gx*bx/gx)
			fmt.Printf("X: %2d na + %2d nb = %2d (mod %2d)\n", ax/gx, bx/gx, px/gx%(ax/gx*bx/gx), ax/gx*bx/gx)
			fmt.Printf("Y: %2d na + %2d nb = %2d (mod %2d)\n", ay/gy, by/gy, py/gy%(ay/gy*by/gy), ay/gy*by/gy)
			fmt.Println("nb*", ax*by-ay*bx, "=", ax*py-ay*px)

			(func() {
				if ax*by-ay*bx == 0 {
					panic("yolo")
				}
				if (ax*py-ay*px)%(ax*by-ay*bx) != 0 {
					fmt.Println("no solution")
					return
				}
				nb := (ax*py - ay*px) / (ax*by - ay*bx)

				if (px-bx*nb)%ax != 0 {
					fmt.Println("no solution2")
					return
				}
				na := (px - bx*nb) / ax
				fmt.Println("na=", na, "nb=", nb)

				if na < 0 || nb < 0 {
					panic("below 0")
				}

				if na <= 100 && nb <= 100 {
					tokens1b += 3*na + nb
				}
			})()

			// 875318608908 too low
			{
				offset := 10000000000000
				px = px + offset
				py = py + offset

				if y%b == 0 {
					panic("iiiih")
				}

				if y%a == 0 {
					panic("aaah")
				}

				if ax*by-ay*bx == 0 {
					panic("yolo")
				}
				if ax*py < 0 {
					panic("negative")
				}
				if ay*px < 0 {
					panic("negative")
				}
				if (ax*py-ay*px)%(ax*by-ay*bx) != 0 {
					fmt.Println("no solution")
					continue
				}
				nb := (ax*py - ay*px) / (ax*by - ay*bx)

				if (px-bx*nb)%ax != 0 {
					fmt.Println("no solution2")
					continue
				}
				na := (px - bx*nb) / ax
				fmt.Println("na=", na, "nb=", nb)

				if na < 0 || nb < 0 {
					panic("below 0")
				}

				if na <= 100 && nb <= 100 {
					tokens1b += 3*na + nb
				}

				fmt.Println(tokens2)
				tokens2 += 3*na + nb
			}
		}

		// gcd := mathutil.GCD(int64(a), int64(b))

		// fmt.Println(mathutil.GCD(int64(a), int64(b)))
	}
	fmt.Println("Part 1a:", tokens1a)
	fmt.Println("Part 1b:", tokens1b)
	fmt.Println("Part 2:", tokens2)

	return tokens1b, tokens2
}
