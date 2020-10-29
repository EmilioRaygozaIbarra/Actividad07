package main

import(
	"fmt"
	"sort"
	"os"
)

type ByCadenaMenor []string

func (a ByCadenaMenor) Len() int           { return len(a) }
func (a ByCadenaMenor) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCadenaMenor) Less(i, j int) bool { return a[i] < a[j] }

type ByCadenaMayor []string

func (a ByCadenaMayor) Len() int           { return len(a) }
func (a ByCadenaMayor) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCadenaMayor) Less(i, j int) bool { return a[i] > a[j] }

func main(){
	asc, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer asc.Close()

	des, err := os.Create("desendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer des.Close()

	str := []string{}
	var aux string
	opc:=1
	
	for opc!=2{
		fmt.Println("Ingrese una string: ")
		fmt.Scan(&aux)
		str=append(str,aux)
		fmt.Println("Quiere ingresar otro string?[1] Si [2] No: ")
		fmt.Scan(&opc)
	}
	sort.Sort(ByCadenaMenor(str))
	for _,cadena := range str{
		asc.WriteString(cadena + "\n")
	}
	sort.Sort(ByCadenaMayor(str))
	for _,cadena := range str{
		des.WriteString(cadena + "\n")
	}
	//fmt.Println(str)
}