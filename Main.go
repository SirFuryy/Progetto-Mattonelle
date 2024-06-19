package main

import (
	"fmt"
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

	colora(piano,1, 2, "f", 48)
	colora(piano,3, 5, "o", 66)
	colora(piano,4, 3, "r", 18)
	colora(piano,5, 1, "j", 5)	
	colora(piano,2, 0, "g", 26)
	colora(piano,4, 0, "x", 45)
	colora(piano,4, 1, "j", 6)
	colora(piano,0, 0, "d", 77)
	colora(piano,1, 5, "m", 79)
	colora(piano,0, 2, "e", 71)
	colora(piano,4, 4, "g", 6)
	colora(piano,0, 1, "j", 47)
	colora(piano,1, 3, "o", 43)
	colora(piano,3, 1, "h", 68)
	colora(piano,4, 5, "c", 50)
	colora(piano,1, 5, "e", 66)
	colora(piano,2, 4, "m", 49)
	colora(piano,2, 2, "g", 96)
	colora(piano,5, 0, "h", 28)
	colora(piano,3, 2, "f", 56)
	

	stampaPiano(piano)

	fmt.Println("Accesa: ", Accesa(piano.piastrelle[1][2]))
	fmt.Println("Accesa: ", Accesa(piano.piastrelle[1][4]))
	fmt.Println("Circonvicini: ", circonvicini(piano.piastrelle[1][2], piano.piastrelle[1][3]))
	fmt.Println("Circonvicini: ", circonvicini(piano.piastrelle[1][2], piano.piastrelle[1][4]))
}

//funzione di utilità che stampa lo stato del piano, con il colore delle piastrelle
//attualmente accese (non l'intensità)
func stampaPiano(piano Piano) {
	fmt.Println("  ^")
	fmt.Println("  |")
	str := "  |"
	for i:=0; i<len(piano.piastrelle); i++ {
		str += "---|"
	}
	fmt.Println(str)

	for i := len(piano.piastrelle)-1; i >=0 ; i-- {
		fmt.Printf("%d |", i)
		for j := 0; j < len(piano.piastrelle[i]); j++ {
			fmt.Print(" " + piano.piastrelle[i][j].colore + " |")
		}
		fmt.Println()
		if i == 0 {
			fmt.Println(str+"->")
		} else {
			fmt.Println(str)
		}
	}
	fmt.Print("    ")
	for i:= 0; i < len(piano.piastrelle); i++ {
		fmt.Print(i, "   ")
	}
	fmt.Println()
}


//funzione costruttrice che crea un piano di piastrelle, con n piastrelle per lato
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

//funzione costruttrice che crea i circonvicini di ogni piastrella del piano
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

//funzione costruttrice che crea una piastrella con i punti passati
func creaPiastrella(i, j int) Piastrella {
	return Piastrella{
		[]Punto{{i,j}, {i,j+1}, {i+1,j+1}, {i+1,j}},
		" ",
		0, 
		nil}
}

//funzione di utilità che restituisce true se la piastrella è accessa, 
//false altrimenti
func Accesa(piastrella Piastrella) bool {
	return piastrella.intenisita > 0
}

//funzione di utilità che restituisce true se la piastrella X è circonvicina 
//alla piastrella y, false altrimenti
func circonvicini(piastX, piastY Piastrella) bool {
	for i := 0; i < len(piastX.circonvicini); i++ {
		if piastX.circonvicini[i].punti[0] == piastY.punti[0] {
			return true
		}
	}
	return false
}

//trova la regione di piastrelle adiacenti alla piastrella in posizione x, y
func trovaRegione(piano Piano, x, y int) []Piastrella {
	visiting, piast := make([]Piastrella, 0), make([]Piastrella, 0)
	visiting = append(visiting, piano.piastrelle[x][y])
	piast = append(piast, piano.piastrelle[x][y])
	for len(visiting) > 0 {
		u := visiting[0]
		visiting = visiting[1:]
		for i := 0; i < len(u.circonvicini); i++ {
			pcontrollo := *u.circonvicini[i]
			if Accesa(pcontrollo) {
				nonvisit := true
				for j := 0; j < len(piast); j++ {
					if piast[j].punti[0] == pcontrollo.punti[0] {
						nonvisit = false
						break
					}
				}
				if nonvisit {
					visiting = append(visiting, pcontrollo)
					piast = append(piast, pcontrollo)
				}
			}
		}
	}

	return piast
}



//funzione che colora la piastrella in posizione num1, num2 con il colore 
//e l'intensità passati, a prescindere dallo stato precedente
func colora(piano Piano, num1, num2 int, color string, intens int) {
	piano.piastrelle[num1][num2].colore = color
	piano.piastrelle[num1][num2].intenisita = intens
}

//Spegne Piastrella(x, y). Se Piastrella(x, y) è già spenta, 
//non fa nulla.
func spegni(piano Piano, x, y int) {
	piano.piastrelle[x][y].colore = " "
	piano.piastrelle[x][y].intenisita = 0
}

//Definisce la regola di propagazione k1α1 + k2α2 + · · · + knαn → β e 
//la inserisce in fondo all’elenco delle regole.
func regola(piano Piano, regola []string) {
	
}

//Stampa e restituisce il colore e l’intensità di Piastrella(x, y). 
//Se Piastrella(x, y) è spenta, non stampa nulla.
func stato(piano Piano, x, y int) {
	fmt.Println("Colore ed intensità: ", piano.piastrelle[x][y].colore, " ", piano.piastrelle[x][y].intenisita)
}

//Stampa l’elenco delle regole di propagazione, nell’ordine attuale.
func stampa(piano Piano) {
	
}

//Calcola e stampa la somma delle intensità delle piastrelle contenute 
//nel blocco di appartenenza di Piastrella(x, y).
//Se Piastrella(x, y) è spenta, restituisce 0.
func blocco(piano Piano, x, y int) {
	
}

//Calcola e stampa la somma delle intensit`a delle piastrelle contenute 
//nel blocco omogeneo di appartenenza di Piastrella(x, y). Se Piastrella(x, y) 
//`e spenta, restituisce 0.
func bloccoOmog(piano Piano, x, y int) {
	
}

//Applica a Piastrella(x, y) la prima regola di propagazione applicabile
//dell’elenco, ricolorando la piastrella. Se nessuna regola è applicabile,
//non viene eseguita alcuna operazione.
func propaga(piano Piano, x, y int) {
	
}

//Propaga il colore sul blocco di appartenenza di Piastrella(x, y).
func propagaBlocco(piano Piano, x, y int) {
	
}

//Ordina l’elenco delle regole di propagazione in base al consumo delle 
//regole stesse: la regola con consumo maggiore diventa l’ultima dell’elenco. 
//Se due regole hanno consumo uguale mantengono il loro ordine relativo.
func ordina(piano Piano) {
	
}

//Stampa la pista che parte da Piastrella(x, y) e segue la sequenza di 
//direzioni s, se tale pista è definita. Altrimenti non stampa nulla.
func pista(piano Piano, x, y int, s string) {
	
}

//Determina la lunghezza della pista più breve che parte da Piastrella(x1, y1)
//e arriva in Piastrella(x2, y2). Altrimenti non stampa nulla.
func lung(piano Piano, x1, y1, x2, y2 int) {
	
}