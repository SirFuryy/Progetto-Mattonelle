package main

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

func creaPiastrella(i, j int) Piastrella {
	return Piastrella{
		[]Punto{{i,j}, {i,j+1}, {i+1,j+1}, {i+1,j}},
		"",
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