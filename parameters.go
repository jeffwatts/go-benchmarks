package go_benchmarks

//import "fmt"

type Inspector interface {
	Inspect()
}

type ValueContainer struct {
	BigString1 string
	BigString2 string
	BigString3 string
	BigString4 string
}

type PointerContainer struct {
	BigString1 string
	BigString2 string
	BigString3 string
	BigString4 string
}

func (v ValueContainer) Inspect() {
	//fmt.Println("Inspect from ValueContainer")
}

func (p *PointerContainer) Inspect() {
	//fmt.Println("Inspect from PointerContainer")
}

func PingInspector(times int, inspector Inspector) {
	for i := 0; i < times; i++ {
		PongInspector(inspector)
	}
}

func PongInspector(inspector Inspector) {
	inspector.Inspect()
}
