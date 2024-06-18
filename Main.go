package main

import (
	"fmt"
	"math/rand"
	"time"
)

//struttura dati per rappresentare un punto su un piano cartesiano
type Punto struct { 
	x int
	y int 
}

// struttura dati per rappresentare una piastrella, centrata all'interno 
//di un quadrato di punti di lato 1, con un colore basato sull'alfabeto {a..z} 
//e un'intensità numerica da 1 a infinito, con 0 che rappresenta la piastrella 
//"spenta". La piastrella ha un massimo di 9 circonvicini, che sono le 
//piastrelle adiacenti ad essa sui punti cardinali e i loro mezzi, compresa se 
//stessa. Come definizione i circonvicini sono le piastrelle che hanno almeno un 
//punto in comune con la piastrella in questione.
type Piastrella struct {
	punti []Punto
	colore string
	intenisita int
	circonvicini []*Piastrella 
}

//struttura dati per rappresentare un piano di piastrelle, 
//con un numero di piastrelle n*n, con n > 0
//
//il piano è composto da piastrelle, che sono disposte in una matrice n*n
type Piano struct {
	piastrelle [][]Piastrella
	regole []string
}

func main() {
	piano := creaPiano(6)

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	valgen:= rnd.Intn(26)
	//genera un numero casuale tra 10 e 26
	for valgen < 10 {
		valgen = rnd.Intn(26)
	}

	for i := 0; i < valgen; i++ {
		num1 := rand.Intn(6)
		num2 := rand.Intn(6)
		color := string(rune(97 + rand.Intn(26)))
		intens := rand.Int()

		colora(piano, num1, num2, color, intens)
	}
	

	stampaPiano(piano)
}

//funzione che colora la piastrella in posizione num1, num2 con il colore 
//e l'intensità passati, a prescindere dallo stato precedente
func colora(piano Piano, num1, num2 int, color string, intens int) {
	piano.piastrelle[num1][num2].colore = color
	piano.piastrelle[num1][num2].intenisita = intens
}

//funzione di utilità che stampa lo stato del piano, con il colore delle piastrelle
//attualmente accese (non l'intensità)
func stampaPiano(piano Piano) {
	fmt.Println("^")
	fmt.Println("|")
	str := "|"
	for i:=0; i<len(piano.piastrelle); i++ {
		str += "---|"
	}
	fmt.Println(str)

	for i := 0; i < len(piano.piastrelle); i++ {
		fmt.Print("|")
		for j := 0; j < len(piano.piastrelle[i]); j++ {
			fmt.Print(" " + piano.piastrelle[i][j].colore + " |")
		}
		fmt.Println()
		if i == len(piano.piastrelle)-1 {
			fmt.Println(str+"->")
		} else {
			fmt.Println(str)
		}
	}
}


//funzione che crea un piano di piastrelle, con n piastrelle per lato
func creaPiano(n int) Piano{
	piano := make([][]Piastrella, n)
	for i := 0; i < n; i++ {
		piano[i] = make([]Piastrella, n)
		for j := 0; j < n; j++ {
			piano[i][j] = creaPiastrella(i, j)
		}
	}

	return Piano{creaCirconvicini(piano, n), nil}
}

//funzione che crea i circonvicini di ogni piastrella del piano
func creaCirconvicini(piano [][]Piastrella, n int) [][]Piastrella{
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			circonvicini := make([]*Piastrella, 0)
			if i > 0 {
				circonvicini = append(circonvicini, &piano[i-1][j])
				if j > 0 {
					circonvicini = append(circonvicini, &piano[i-1][j-1])
				}
				if j < n-1 {
					circonvicini = append(circonvicini, &piano[i-1][j+1])
				}
			}
			if i < n-1 {
				circonvicini = append(circonvicini, &piano[i+1][j])
				if j > 0 {
					circonvicini = append(circonvicini, &piano[i+1][j-1])
				}
				if j < n-1 {
					circonvicini = append(circonvicini, &piano[i+1][j+1])
				}
			}
			if j > 0 {
				circonvicini = append(circonvicini, &piano[i][j-1])
			}
			if j < n-1 {
				circonvicini = append(circonvicini, &piano[i][j+1])
			}
			piano[i][j].circonvicini = circonvicini
		}
	}
	return piano
}

//funzione che crea una piastrella con i punti passati
func creaPiastrella(i, j int) Piastrella {
	return Piastrella{
		[]Punto{{i,j}, {i,j+1}, {i+1,j+1}, {i+1,j}},
		" ",
		0, 
		nil}
}

//funzione che restituisce true se la piastrella è accessa, false altrimenti
func Accesa(piastrella Piastrella) bool {
	return piastrella.intenisita > 0
}

//restituisce true se la piastrella X è circonvicina alla piastrella y, false altrimenti
func circonvicini(piastX, piastY Piastrella) bool {
	for i := 0; i < len(piastX.circonvicini); i++ {
		if piastX.circonvicini[i].punti[0] == piastY.punti[0] {
			return true
		}
	}
	return false
}