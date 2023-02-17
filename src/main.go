package main

//"strconv"

/*
type Doctor struct {
	id           int
	name         string
	designastion string
	address      []string
}

type Animal struct {
	name    string
	origion string
}

type Bird struct {
	Animal
	speed  float32
	canFly bool
}
*/

func main() {
	/*
			fmt.Println("Hello GO !")
			var i int = 10
			var s string
			fmt.Printf("vaue of i is %v and is type %T \n", i, i)
			s = strconv.Itoa(i)
			fmt.Printf("vaue of s is %v and is type %T \n", s, s)
			//array
			am := make([]int, 3)
			am[0] = 47
			arr := [...]int{1, 5, 6}
			fmt.Printf("vaue of arr is %v and is type %T \n", arr[0], arr)
			fmt.Printf("vaue of am is %v and is type %T \n", am[0], am)
			//slice
			sl := []int{}
			sl = append(sl, 1)
			fmt.Printf("vaue of slice is %v and is type %T \n", sl[0], sl)

		a := [...]int{1, 2, 3}
		c := a
		fmt.Print("a", a, "\n")
		b := append(c[:1], c[2:]...)
		fmt.Print("b", b, "\n")
		fmt.Print("a", a, "\n")
	*/

	//locations := make(map[int]string)
	/*
		locations := map[int]string{
			1:  "US",
			91: "India",
		}
		fmt.Print("locations  : ", locations[1], "\n")
		newLocations := locations
		newLocations[1] = "USA"
		fmt.Print("locations  : ", locations[1], "\n")

		if pop, ok := locations[1]; ok {
			fmt.Print("You are good ", pop, "\n")
		}

		for k, v := range locations {
			fmt.Println(k, v)
		}
		ranger := []int{1, 4, 5, 6}
		for i := range ranger {
			fmt.Println(i)
		}

			for i := 1; i <= 5; i++ {
				fmt.Println(i)
			}


				doctor := Doctor{
					id:           2,
					name:         "Abhi",
					designastion: "Surjeon",
					address: []string{
						"2",
						"Renfrew steet",
						"Glasgow",
					},
				}

				dummy := &doctor
				dummy.id = 3
				fmt.Print("doctor is : ", doctor, "\n")
				fmt.Print("dummy doctor is : ", dummy, "\n")
				fmt.Print("doctor is : ", doctor, "\n")

	*/

	/*
		bird := Bird{}
		bird.name = "Emu"
		bird.origion = "India"
		bird.speed = 48
		bird.canFly = true

		secondBird := Bird{
			Animal: Animal{name: "Momu",
				origion: "India"},
			speed:  68,
			canFly: false,
		}

		fmt.Print("Bird  is : ", bird, "\n")
		fmt.Print("Second Bird  is : ", secondBird, "\n")
	*/

}
