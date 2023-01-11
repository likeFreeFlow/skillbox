package main
import "fmt" 
type Ibarker interface {
	Bark()string
}

type DogName string
func (dn DogName) Bark() string{
	return string(dn)+" aw aw aw "
}

func main() [
	dn:+dogname("whisper")
	fmt.Println(dn.Bark())

	var Idn IBarker=dn
	Idn.Barker()
	

	]