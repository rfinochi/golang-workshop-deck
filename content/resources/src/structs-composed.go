type Dog struct{}

func (d *Dog) Bark() {
	fmt.Println("Woof!")
}

type GuideDog struct {
	*Dog
}

func (gd *GuideDog) Help(h *Human) {
	fmt.Printf("Hey human, grab %s’s leash!\n", gd.Name())
}
