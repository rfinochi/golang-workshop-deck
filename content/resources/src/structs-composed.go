type Dog struct{}

func (d *Dog) Bark() string {
	return "Woof!"
}

type GuideDog struct {
	*Dog
}

func (gd *GuideDog) Help() {
	fmt.Printf("Please help me (%s)", gd.Bark())
}
