package main

func main() {

	/* { fmt.Println("1-misol")   }
	var a int
	fmt.Print("3 xonali sonni kiriting: " ); fmt.Scan(&a)
	fmt.Println((a/100)*100 + (a%10)*10 + (a/10)%10)
	*/

	/* { fmt.Println("2-misol")   }
	var a int
	fmt.Print("4 xonali sonni kiriting: " ); fmt.Scan(&a)
	fmt.Println((a/100)%10) */

	/* { fmt.Println("3-misol")   }
	var a int
	fmt.Print("4 xonali sonni kiriting: " ); fmt.Scan(&a)
	fmt.Println(a/1000) */

	/* { fmt.Println("4-misol")   }
	var a int
	fmt.Print("sekuntni kiriting: " ); fmt.Scan(&a)
	fmt.Println("To`la minut: ", a/60) */

	/* { fmt.Println("5-misol")   }
	var a int
	fmt.Print("sekuntni kiriting: " ); fmt.Scan(&a)
	fmt.Println("To`la soat: ", a/3600) */

	/* { fmt.Println("6-misol")   }
	var a int
	fmt.Print("sekuntni kiriting: " ); fmt.Scan(&a)
	fmt.Println("To`la minut: ", a/60, " Sekunt: ", a%60) */

	/* { fmt.Println("7-misol")   }
	var a int
	fmt.Print("sekuntni kiriting: " ); fmt.Scan(&a)
	fmt.Println("To`la soat: ", a/3600, " Sekunt: ", a%60) */

	/* { fmt.Println("8-misol")   }
	var a int
	fmt.Print("sekuntni kiriting: " ); fmt.Scan(&a)
	fmt.Println("To`la soat: ", a/3600, "To`la minut: ", a%3600/60, " Sekunt: ", a%60) */

	/* {
		fmt.Println("9-misol")
	}
	var a int
	fmt.Print("4 xonali sonni kiriting: ")
	fmt.Scan(&a)
	fmt.Println(a / 1000) */

	// 2-dars uyga vazida

	// if

	/* {
		fmt.Println("1-misol")
	}
	var a int
	var b int
	fmt.Print("enter A: ")
	fmt.Scan(&a)
	fmt.Print("enter B: ")
	fmt.Scan(&b)
	if a < b {
		fmt.Println("a:", a, "b:", b)
	} else if a > b {
		a = a + b
		b = a - b
		a = a - b
		fmt.Println("a:", a, "b:", b)
	} else {
		fmt.Println("a=b")
	}
	*/

	/* {
		fmt.Println("2-misol")
	}
	var a int
	var b int
	var c int
	fmt.Print("enter A: ")
	fmt.Scan(&a)
	fmt.Print("enter B: ")
	fmt.Scan(&b)
	if a != b {
		c = a + b; a = c; b = c;
		fmt.Println("a:", a)
		fmt.Println("b:", b)
	} else {
		c = 0; a = c; b = c;
		fmt.Println("a:", c)
		fmt.Println("b:", c)
	} */

	/* {
		fmt.Println("3-misol")
	}
	var a int
	var b int
	fmt.Print("enter A: ")
	fmt.Scan(&a)
	fmt.Print("enter B: ")
	fmt.Scan(&b)
	if a != b && a > b {
		fmt.Println("a:", a)
		b = a
		fmt.Println("b:", b)
	} else if a != b && a < b {
		a = b
		fmt.Println("a:", a)
		fmt.Println("b:", b)
	} else {
		a = 0
		b = 0
		fmt.Println("a:", a)
		fmt.Println("b:", b)
	} */

	/* {
		fmt.Println("4-misol")
	}
	var a int
	var b int
	var c int
	fmt.Print("enter A: ")
	fmt.Scan(&a)
	fmt.Print("enter B: ")
	fmt.Scan(&b)
	fmt.Print("enter C: ")
	fmt.Scan(&c)
	if a < b && a < c {
		fmt.Println("Kichigi:", a)
	} else if b < a && b < c {
		fmt.Println("Kichigi:", b)
	} else if c < b && c < a {
		fmt.Println("Kichigi:", c)
	} */

	/* {
		fmt.Println("5-misol")
	}
	var a int
	var b int
	var c int
	fmt.Print("enter A: ")
	fmt.Scan(&a)
	fmt.Print("enter B: ")
	fmt.Scan(&b)
	fmt.Print("enter C: ")
	fmt.Scan(&c)
	if a < b && a < c && b > c {
		fmt.Println("O`rtachasi:", c)
	} else if a < b && a < c && b < c {
		fmt.Println("O`rtachasi:", b)
	} else if b < a && b < c && a < c {
		fmt.Println("O`rtachasi:", a)
	} else if b < a && b < c && a > c {
		fmt.Println("O`rtachasi:", c)
	} else if c < b && c < a && a > b {
		fmt.Println("O`rtachasi:", b)
	} else if c < b && c < a && a < b {
		fmt.Println("O`rtachasi:", a)
	} */

	/* {
		fmt.Println("6-misol")
	}
	var a int
	var b int
	var c int
	fmt.Print("enter A: ")
	fmt.Scan(&a)
	fmt.Print("enter B: ")
	fmt.Scan(&b)
	fmt.Print("enter C: ")
	fmt.Scan(&c)
	if a < b && a < c && b > c {
		fmt.Println("Kichigi:", a, "Kattasi:", b)
	} else if a < b && a < c && b < c {
		fmt.Println("Kichigi:", a, "Kattasi:", c)
	} else if b < a && b < c && a < c {
		fmt.Println("Kichigi:", b, "Kattasi:", c)
	} else if b < a && b < c && a > c {
		fmt.Println("Kichigi:", b, "Kattasi:", a)
	} else if c < b && c < a && a > b {
		fmt.Println("Kichigi:", c, "Kattasi:", a)
	} else if c < b && c < a && a < b {
		fmt.Println("Kichigi:", c, "Kattasi:", b)
	} */

	/* {
		fmt.Println("7-misol")
	}
	var a int
	var b int
	var c int
	fmt.Print("enter A: ")
	fmt.Scan(&a)
	fmt.Print("enter B: ")
	fmt.Scan(&b)
	fmt.Print("enter C: ")
	fmt.Scan(&c)
	if a < b && a < c && b > c {
		fmt.Println("O`rtachasi:", c, "Kattasi:", b)
	} else if a < b && a < c && b < c {
		fmt.Println("O`rtachasi:", b, "Kattasi:", c)
	} else if b < a && b < c && a < c {
		fmt.Println("O`rtachasi:", a, "Kattasi:", c)
	} else if b < a && b < c && a > c {
		fmt.Println("O`rtachasi:", c, "Kattasi:", a)
	} else if c < b && c < a && a > b {
		fmt.Println("O`rtachasi:", b, "Kattasi:", a)
	} else if c < b && c < a && a < b {
		fmt.Println("O`rtachasi:", a, "Kattasi:", b)
	} */

	/* {
		fmt.Println("8-misol")
	}
	var a int
	var b int
	var c int
	fmt.Print("enter A: ")
	fmt.Scan(&a)
	fmt.Print("enter B: ")
	fmt.Scan(&b)
	fmt.Print("enter C: ")
	fmt.Scan(&c)

	if a < b && b < c {
		a = a * 2
		b = b * 2
		c = c * 2
		fmt.Println("a:", a, "b:", b, "c:", c)
	} else if a > b && b > c {
		a = a * -1
		b = b * -1
		c = c * -1
		fmt.Println("a:", a, "b:", b, "c:", c)
	} else {
		fmt.Println("Aralash tartibda")
	} */

	/* {
		fmt.Println("9-misol")
	}
	var a int
	var b int
	var c int
	fmt.Print("enter A: ")
	fmt.Scan(&a)
	fmt.Print("enter B: ")
	fmt.Scan(&b)
	fmt.Print("enter C: ")
	fmt.Scan(&c)

	if a < b && b < c || a > b && b > c {
		a = a * 2
		b = b * 2
		c = c * 2
		fmt.Println("a:", a, "b:", b, "c:", c)
	} else {
		a = a * -1
		b = b * -1
		c = c * -1
		fmt.Println("a:", a, "b:", b, "c:", c)
	} */

	/* {
		fmt.Println("10-misol")
	}
	var a int
	var b int
	var c int
	fmt.Print("enter A: ")
	fmt.Scan(&a)
	fmt.Print("enter B: ")
	fmt.Scan(&b)
	fmt.Print("enter C: ")
	fmt.Scan(&c)
	if a==b {
		fmt.Println("Qolgan bittasi 3-si:", c)
	} else if a==c {
		fmt.Println("Qolgan bittasi 2-si:", b)
	} else if b==c {
		fmt.Println("Qolgan bittasi 1-si:", a)
	} */

	/* {
		fmt.Println("11-misol")
	}
	var a int
	var b int
	var c int
	var d int
	fmt.Print("enter A: ")
	fmt.Scan(&a)
	fmt.Print("enter B: ")
	fmt.Scan(&b)
	fmt.Print("enter C: ")
	fmt.Scan(&c)
	fmt.Print("enter D: ")
	fmt.Scan(&d)
	if a == b && b == c {
		fmt.Println("Qolgan bittasi 4-si:", d)
	} else if a == c && c == d {
		fmt.Println("Qolgan bittasi 2-si:", b)
	} else if b == c && c == d {
		fmt.Println("Qolgan bittasi 1-si:", a)
	} else if a == b && b == d {
		fmt.Println("Qolgan bittasi 3-si:", c)
	} */

	// for

	/* {
		fmt.Println("1-misol")
	}
	var n int
	fmt.Print("Enter n:")
	fmt.Scan(&n)
	for i := n; i > 0; i-- {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}

	} */

	/* {
		fmt.Println("2-misol")
	}
	var n int
	fmt.Print("Enter n:")
	fmt.Scan(&n)
	for i := n; i > 0; i-- {
		if i%2 == 1 {
			fmt.Print(i, " ")
		}

	} */

	/* {
		fmt.Println("3-misol")
	}
	var counter int
	for i := 10; i < 100; i++ {
		counter = 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				counter++
			}
		}
		if counter == 2 {
			fmt.Print(i, " ")
		}
	} */

	/* {
		fmt.Println("4-misol")

	}
	for i := 100; i < 1000; i++ {
		if i/100 == i/10%10 || i/100 == i%10 || i/10%10 == i%10 {
			fmt.Print(i, " ")
		}
	} */

	/* {
		fmt.Println("5-misol")
	}
	var a int
	var b int
	var count int
	b = a
	fmt.Print("Son:")
	fmt.Scan(&a)
	for i := a; i > 0; {
		fmt.Print()
		a=a/10
	} */

	/* {
		fmt.Println("6-misol")
	} */

	/* {
		fmt.Println("7-misol")
	}
	for i := 0; i < 5; i++ {
		fmt.Println()
		for j := 0; j < 5; j++ {
			if i==0 || j==0 || i==4  || j==4{
				fmt.Print("* ")
			}

		}

	} */

	// 2ta swichni shartiga tushunmadim

}
