package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* STRUTTURE DATI*/

//struttura dati per rappresentare un punto su un piano cartesiano
type Punto struct { 
	x int	//colonna
	y int 	//riga
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
//il piano è composto da piastrelle, che sono disposte in una matrice n*n, 
type Piano struct {
	piastrelle [][]Piastrella
	regole []Regola
}

//struttura dati per rappresentare una regola di propagazione,
//con un insieme di alfa, che sono i colori delle piastrelle che circondano la
//piastrella a cui applicare la regola, e un beta, che è il colore che verrà propagato
type Regola struct {
	alfa map[string]int
	beta string
	usato int
}



func main() {
	piano := creaPiano(6)

	colora(piano, 3, 5, "x", 96)
	colora(piano, 1, 3, "h", 90)
	colora(piano, 2, 1, "s", 44)
	colora(piano, 2, 3, "o", 78)
	colora(piano, 0, 5, "f", 6)
	colora(piano, 0, 2, "a", 46)
	colora(piano, 5, 1, "i", 45)
	colora(piano, 4, 1, "i", 27)
	colora(piano, 0, 1, "f", 18)
	colora(piano, 2, 2, "d", 59)
	colora(piano, 5, 5, "u", 69)
	colora(piano, 0, 3, "t", 13)
	colora(piano, 5, 2, "a", 10)
	

	
	stampaPiano(piano)

	piano = regola(piano, "a 1 s 1 o")
	piano = regola(piano, "b 1 h 1 d")
	piano = regola(piano, "c 1 t")

	
	
	propagaBlocco(piano, 2, 2)
	stampaPiano(piano)
}


/* FUNZIONI DI STAMPA*/

//funzione di utilità che stampa lo stato del piano, con il colore delle piastrelle
//attualmente accese (non l'intensità)
func stampaPiano(piano Piano) {
	fmt.Println("  ^")
	fmt.Println("  |")
	str := "|"
	for i:=0; i<len(piano.piastrelle); i++ {
		str += "---|"
	}
	fmt.Println("  "+ str)

	for i := len(piano.piastrelle)-1; i >=0 ; i-- {
		fmt.Print("  |")
		for j := 0; j < len(piano.piastrelle[i]); j++ {
			fmt.Print(" " + piano.piastrelle[j][i].colore + " |")	//lettere sono invertite
		}													//perchè la matrice è invertita
		fmt.Println()
		if i == 0 {
			fmt.Println(i, str+"->")
		} else {
			fmt.Println(i, str)
		}
	}
	fmt.Print("  ")
	for i:= 0; i < len(piano.piastrelle); i++ {
		fmt.Print(i, "   ")
	}
	fmt.Println()
}

//funzione di stampa che stampa tutte le caratteristiche di una piastrella
func stampaPiastrella(piast Piastrella) {
	str := "Piastrella di punti:"

	for _, v := range piast.punti {
		str += fmt.Sprintf(" (%d, %d)", v.x, v.y)
	}

	str += ", colore: '" + piast.colore + "' e intensità: " + fmt.Sprint(piast.intenisita)
	fmt.Println(str)
}




/* FUNZIONI DI COSTRUZIONE */

//funzione costruttrice che crea un piano di piastrelle, con n piastrelle per lato
func creaPiano(n int) Piano{
	piano := make([][]Piastrella, n)
	for i := 0; i < n; i++ {
		piano[i] = make([]Piastrella, n)
		for j := 0; j < n; j++ {
			piano[i][j] = creaPiastrella(i, j)
		}
	}

	return Piano{creaCirconvicini(piano, n), make([]Regola, 0)}
}

//funzione costruttrice che crea i circonvicini di ogni piastrella del piano
//l'array conterrà le piastrelle circonvicine ordinate a partire dalla 
//posizione (x-1, y+1) [per intenderci la posizione NO] e proseguendo in 
//ordine di riga e colonna, ovvero dato un punto (x,y) al centro del piano con
//le posizioni vicine disposte come segue:
//      x-1, y+1 | x, y+1 | x+1, y+1
//      x-1, y   | x, y   | x+1, y
//      x-1, y-1 | x, y-1 | x+1, y-1
//la posizione dei circonvicini sarà disposta come segue:
//		  [0]	 |   [1]  |    [2] 
//        [3]	 |   [4]  |    [5]
//        [6]	 |   [7]  |    [8]
func creaCirconvicini(piano [][]Piastrella, n int) [][]Piastrella{
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			circonvicini := make([]*Piastrella, 0, 9)
			//x-1, y+1
			if i > 0 && j < n-1 {
				circonvicini = append(circonvicini, &piano[i-1][j+1])
			} else {
				circonvicini = append(circonvicini, nil)
			}

			//x, y+1
			if j < n-1 {
				circonvicini = append(circonvicini, &piano[i][j+1])
			} else {
				circonvicini = append(circonvicini, nil)
			}

			//x+1, y+1
			if i < n-1 && j < n-1 {
				circonvicini = append(circonvicini, &piano[i+1][j+1])
			} else {
				circonvicini = append(circonvicini, nil)
			}

			//x-1, y
			if i > 0 {
				circonvicini = append(circonvicini, &piano[i-1][j])
			} else {
				circonvicini = append(circonvicini, nil)
			}

			//x, y
			circonvicini = append(circonvicini, &piano[i][j])

			//x+1, y
			if i < n-1 {
				circonvicini = append(circonvicini, &piano[i+1][j])
			} else {
				circonvicini = append(circonvicini, nil)
			}

			//x-1, y-1
			if i > 0 && j > 0 {
				circonvicini = append(circonvicini, &piano[i-1][j-1])
			} else {
				circonvicini = append(circonvicini, nil)
			}

			//x, y-1
			if j > 0 {
				circonvicini = append(circonvicini, &piano[i][j-1])
			} else {
				circonvicini = append(circonvicini, nil)
			}

			//x+1, y-1
			if i < n-1 && j > 0 {
				circonvicini = append(circonvicini, &piano[i+1][j-1])
			} else {
				circonvicini = append(circonvicini, nil)
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




/*  FUNZIONI DI UTILITA  */

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

//trova il blocco, ovvero la regione di ampiezza massima, di piastrelle adiacenti alla piastrella in posizione x, y
func trovaBlocco(piano Piano, x, y int) []Piastrella {
	visiting, piast := make([]Piastrella, 0), make([]Piastrella, 0)
	visiting = append(visiting, piano.piastrelle[x][y])
	piast = append(piast, piano.piastrelle[x][y])
	for len(visiting) > 0 {
		u := visiting[0]
		visiting = visiting[1:]
		for i := 0; i < len(u.circonvicini); i++ {
			if u.circonvicini[i] == nil {
				continue
			}
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

//trova il blocco omogeneo, ovvero la regione di ampiezza massima, di 
//piastrelle adiacenti alla piastrella in posizione x, y con lo stesso 
//colore della piastrella in posizione x, y
func trovaBloccoOmogeneo(piano Piano, x, y int) []Piastrella {
	visiting, piast := make([]Piastrella, 0), make([]Piastrella, 0)
	colore := piano.piastrelle[x][y].colore
	visiting = append(visiting, piano.piastrelle[x][y])
	piast = append(piast, piano.piastrelle[x][y])
	for len(visiting) > 0 {
		u := visiting[0]
		visiting = visiting[1:]
		for i := 0; i < len(u.circonvicini); i++ {
			if u.circonvicini[i] == nil {
				continue
			}
			pcontrollo := *u.circonvicini[i]
			if Accesa(pcontrollo) && pcontrollo.colore == colore{
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




/* FUNZIONI RICHIESTE DAL PROGETTO */

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
func regola(piano Piano, regola string) Piano{
	reg := strings.Split(regola, " ")
	beta := reg[0]
	reg = reg[1:]
	alfa := make(map[string]int)
	for i := 0; i < len(reg); i = i+2 {
		k, _ := strconv.Atoi(reg[i])
		a := reg[i+1]
		alfa[a] = k
	}

	piano.regole = append(piano.regole, Regola{alfa, beta, 0})
	return piano
}
//Stampa e restituisce il colore e l’intensità di Piastrella(x, y). 
//Se Piastrella(x, y) è spenta, non stampa nulla.
func stato(piano Piano, x, y int) {
	fmt.Println("Colore ed intensità: ", piano.piastrelle[x][y].colore, " ", piano.piastrelle[x][y].intenisita)
}

//Stampa l’elenco delle regole di propagazione, nell’ordine attuale.
func stampa(piano Piano) {
	for _, v := range piano.regole {
		//giallo: 1 rosso 1 blu
		str:=v.beta + ": "

		for k, v := range v.alfa {
			str += fmt.Sprintf("%d %s ", v, k)
		}

		fmt.Println(str)
	}
}

//Calcola e stampa la somma delle intensità delle piastrelle contenute 
//nel blocco di appartenenza di Piastrella(x, y).
//Se Piastrella(x, y) è spenta, restituisce 0.
func blocco(piano Piano, x, y int) int{
	if !Accesa(piano.piastrelle[x][y]) {
		return 0
	}
	pias := trovaBlocco(piano, x, y)
	somma := 0
	for i := 0; i < len(pias); i++ {
		somma += pias[i].intenisita
	}
	fmt.Println("Somma intensità blocco: ", somma)
	return somma
}

//Calcola e stampa la somma delle intensit`a delle piastrelle contenute 
//nel blocco omogeneo di appartenenza di Piastrella(x, y). Se Piastrella(x, y) 
//`e spenta, restituisce 0.
func bloccoOmog(piano Piano, x, y int) int{
	if !Accesa(piano.piastrelle[x][y]) {
		return 0
	}
	pias := trovaBloccoOmogeneo(piano, x, y)
	somma := 0
	for i := 0; i < len(pias); i++ {
		somma += pias[i].intenisita
	}
	fmt.Println("Somma intensità blocco omogeneo: ", somma)
	return somma
}

//Applica a Piastrella(x, y) la prima regola di propagazione applicabile
//dell’elenco, ricolorando la piastrella. Se nessuna regola è applicabile,
//non viene eseguita alcuna operazione.
func propaga(piano Piano, x, y int) {
	intorno := piano.piastrelle[x][y].circonvicini
	moment := intorno[5:]
	intorno = intorno[:4]
	intorno = append(intorno, moment...)
	
	var reg Regola
	regvalida := false

	mappaIntorno := make(map[string]int)
	for i := 0; i < len(intorno); i++ {
		if intorno[i] != nil {
			mappaIntorno[intorno[i].colore]++
		}
	}

	for _, v := range piano.regole {
		trovato := false
		for k, u := range v.alfa {
			if mappaIntorno[k] < u {
				trovato = false
				break
			} else {
				trovato = true
			}
		}

		if trovato {
			reg = v
			regvalida = true
			v.usato++
			break
		}
	}

	if !regvalida {
		return
	}

	piano.piastrelle[x][y].colore = reg.beta
	if !Accesa(piano.piastrelle[x][y]) {
		piano.piastrelle[x][y].intenisita = 1
	}
}

//Propaga il colore sul blocco di appartenenza di Piastrella(x, y).
func propagaBlocco(piano Piano, x, y int) {
	blocco := trovaBlocco(piano, x, y)
	copia := make([]Piastrella, len(blocco))
	copy(copia, blocco)

	for i := 0; i < len(blocco); i++ {
		fmt.Print(blocco[i].colore, " ")
	}

	for i := 0; i < len(blocco); i++ {
		intorno := blocco[i].circonvicini
		moment := intorno[5:]
		intorno = intorno[:4]
		intorno = append(intorno, moment...)
	
		var reg Regola
		regvalida := false

		mappaIntorno := make(map[string]int)
		for i := 0; i < len(intorno); i++ {
			if intorno[i] != nil {
				mappaIntorno[intorno[i].colore]++
			}
		}

		for _, v := range piano.regole {
			trovato := false
			for k, u := range v.alfa {
				if mappaIntorno[k] < u {
					trovato = false
					break
				} else {
					trovato = true
				}
			}

			if trovato {
				reg = v
				regvalida = true
				v.usato++
				break
			}
		}

		if regvalida {
			copia[i].colore = reg.beta
		}
	}
	fmt.Println()
	for i := 0; i < len(copia); i++ {
		fmt.Print(copia[i].colore, " ")
	}

	for i := 0; i < len(copia); i++ {
		if copia[i].colore != blocco[i].colore {
			piano.piastrelle[blocco[i].punti[0].x][blocco[i].punti[0].y].colore = copia[i].colore
		}
	}
}

//Ordina l’elenco delle regole di propagazione in base al consumo delle 
//regole stesse: la regola con consumo maggiore diventa l’ultima dell’elenco. 
//Se due regole hanno consumo uguale mantengono il loro ordine relativo.
func ordina(piano Piano) {
	
}

//Stampa la pista che parte da Piastrella(x, y) e segue la sequenza di 
//direzioni s, se tale pista è definita. Altrimenti non stampa nulla.
func pista(piano Piano, x, y int, s string) {
	indicazioni := make([]string, 0)
	for i := 0; i < len(s); i += 2 {
		indicazioni = append(indicazioni, s[i : i+2])
	}

	piastControllata := piano.piastrelle[x][y]
	if !Accesa(piastControllata) {
		return
	}
	piastrelle := make([]Piastrella, 0)
	piastrelle = append(piastrelle, piastControllata)
	for i := 0; i < len(indicazioni); i++ {
		switch indicazioni[i] {
		case "NO":	//x-1, y+1
			ppros:= piastControllata.circonvicini[0]
			if ppros != nil {
				if Accesa(*ppros) {
					piastControllata = *ppros
					piastrelle = append(piastrelle, piastControllata)
				} else {return}
			} else {return}
			break
		case "NN":	//x, y+1
			ppros:= piastControllata.circonvicini[1]
			if ppros != nil {
				if Accesa(*ppros) {
					piastControllata = *ppros
					piastrelle = append(piastrelle, piastControllata)
				}  else {return}
			} else {return}
			break
		case "NE":	//x+1, y+1
			ppros:= piastControllata.circonvicini[2]
			if ppros != nil {
				if Accesa(*ppros) {
					piastControllata = *ppros
					piastrelle = append(piastrelle, piastControllata)
				} else {return}
			} else {return}
			break
		case "OO":	//x-1, y
			ppros:= piastControllata.circonvicini[3]
			if ppros != nil {
				if Accesa(*ppros) {
					piastControllata = *ppros
					piastrelle = append(piastrelle, piastControllata)
				} else {return}
			} else {return}
			break
		case "EE":	//x+1, y
			ppros:= piastControllata.circonvicini[5]
			if ppros != nil {
				if Accesa(*ppros) {
					piastControllata = *ppros
					piastrelle = append(piastrelle, piastControllata)
				} else {return}
			} else {return}
			break
		case "SO":	//x-1, y-1
			ppros:= piastControllata.circonvicini[6]
			if ppros != nil {
				if Accesa(*ppros) {
					piastControllata = *ppros
					piastrelle = append(piastrelle, piastControllata)
				} else {return}
			} else {return}
			break
		case "SS":	//x, y-1
			ppros:= piastControllata.circonvicini[7]
			if ppros != nil {
				if Accesa(*ppros) {
					piastControllata = *ppros
					piastrelle = append(piastrelle, piastControllata)
				} else {return}
			} else {return}
			break
		case "SE":	//x+1, y-1
			ppros:= piastControllata.circonvicini[8]
			if ppros != nil {
				if Accesa(*ppros) {
					piastControllata = *ppros
					piastrelle = append(piastrelle, piastControllata)
				} else {return}
			} else {return}
			break		
		}
	}

	for i := 0; i < len(piastrelle); i++ {
		stampaPiastrella(piastrelle[i])
	}
}

//Determina la lunghezza della pista più breve che parte da Piastrella(x1, y1)
//e arriva in Piastrella(x2, y2). Altrimenti non stampa nulla.
func lung(piano Piano, x1, y1, x2, y2 int) {
	
}