type Dog struct{}

func (d *Dog) Bark() {
	fmt.Println("Woof!")
}

type GuideDog struct {
	*Dog
}

func (gd *GuideDog) Help(h *Human) {
	fmt.Printf("Please help me (%s)", gd.Name())
}
