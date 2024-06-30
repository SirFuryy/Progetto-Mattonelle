package main

import (
	"bufio"
	"fmt"
	"os"
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

var eseguito bool = true

func main() {
	var input string
	var piano Piano
	piano = creaPiano(10, 10)

	/*
	scanner := bufio.NewScanner(os.Stdin)
	for eseguito {
		if scanner.Scan() {
			input = scanner.Text()
		}
		esegui(piano, input)
	}*/


	file, err := os.Open("input_2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = scanner.Text()
		piano = esegui(piano, input)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

//funzione che esegue il comando passato in input
func esegui(piano Piano, input string) Piano{
	inp := strings.Split(input, " ")
	comando := inp[0]
	inp = inp[1:]
	switch comando {
	case "C":
		x, _ := strconv.Atoi(inp[0])
		y, _ := strconv.Atoi(inp[1])
		alpha := inp[2]
		i, _ := strconv.Atoi(inp[3])
		checkPiano(&piano, x, y)
		colora(piano, x, y, alpha, i)
		break
	case "S":
		x, _ := strconv.Atoi(inp[0])
		y, _ := strconv.Atoi(inp[1])
		checkPiano(&piano, x, y)
		spegni(piano, x, y)
		break
	case "r":
		reg := strings.Join(inp, " ")
		piano.regole = append(piano.regole, Regola{})
		regola(piano, reg)
		break
	case "?":
		x, _ := strconv.Atoi(inp[0])
		y, _ := strconv.Atoi(inp[1])
		checkPiano(&piano, x, y)
		stato(piano, x, y)
		break
	case "s":
		stampa(piano)
		break
	case "b":
		x, _ := strconv.Atoi(inp[0])
		y, _ := strconv.Atoi(inp[1])
		checkPiano(&piano, x, y)
		blocco(piano, x, y)
		break
	case "B":
		x, _ := strconv.Atoi(inp[0])
		y, _ := strconv.Atoi(inp[1])
		checkPiano(&piano, x, y)
		bloccoOmog(piano, x, y)
		break
	case "p":
		x, _ := strconv.Atoi(inp[0])
		y, _ := strconv.Atoi(inp[1])
		checkPiano(&piano, x, y)
		propaga(piano, x, y)
		break
	case "P":
		x, _ := strconv.Atoi(inp[0])
		y, _ := strconv.Atoi(inp[1])
		checkPiano(&piano, x, y)
		propagaBlocco(piano, x, y)
		break
	case "o":
		piano.regole = mergeSort(piano.regole)
		break
	case "q":
		eseguito = false
		break
	case "t":
		x, _ := strconv.Atoi(inp[0])
		y, _ := strconv.Atoi(inp[1])
		s := inp[2]
		checkPiano(&piano, x, y)
		pista(piano, x, y, s)
		break
	case "L":
		x1, _ := strconv.Atoi(inp[0])
		y1, _ := strconv.Atoi(inp[1])
		x2, _ := strconv.Atoi(inp[2])
		y2, _ := strconv.Atoi(inp[3])
		checkPiano(&piano, x1, y1)
		checkPiano(&piano, x2, y2)
		lung(piano, x1, y1, x2, y2)
		break
	case "a":
		stampaPiano(piano)
		break
	default:
		break
	}
	return piano
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

	for j := len(piano.piastrelle[0])-1; j >= 0; j-- {
		fmt.Print("  |")
		for i := 0; i < len(piano.piastrelle) ; i++ {
			fmt.Print(" " + piano.piastrelle[i][j].colore + " |")	//lettere sono invertite
		}													//perchè la matrice è invertita
		fmt.Println()
		if j == 0 {
			fmt.Println(j, str+"->")
		} else {
			fmt.Println(j, str)
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

//funzione di utilità che controlla se la piastrella in posizione x, y è
//all'interno del piano
func checkPiano(piano *Piano, x, y int) {
	lenX := len(piano.piastrelle)
	lenY := len(piano.piastrelle[0])

	if x >= lenX && y >= lenY {
		piano.piastrelle = append(piano.piastrelle, make([][]Piastrella, x+1-lenX)...)
		for i := lenX; i <= x; i++ {		//inserisco le nuove x
			piano.piastrelle[i] = make([]Piastrella, lenY)
			for j := 0; j < lenY; j++ {
				piano.piastrelle[i][j] = creaPiastrella(i, j)
			}
		}
		for i := 0; i <= x; i++ {		//inserisco le nuove y e l'incorcio tra le nuove x e e le nuove y
			piano.piastrelle[i] = append(piano.piastrelle[i], make([]Piastrella, y+1-lenY)...)
			for j := lenY; j <= y; j++ {
				piano.piastrelle[i][j] = creaPiastrella(i, j)
			}
		}

		for i:= lenX-1; i <= x; i++ {		//calcolo i circonvicini delle nuove piastrelle dalla parte di x
			for j := 0; j < lenY; j++ {
				creaCirconvicini(*piano, i, j)
			}
		}

		for i := 0; i <= x; i++ {		//calcolo i circonvicini delle nuove piastrelle dalla parte di y e l'incrocio
			for j := lenY-1; j <= y; j++ {
				creaCirconvicini(*piano, i, j)
			}
		}
	} else 	if x >= lenX {
		piano.piastrelle = append(piano.piastrelle, make([][]Piastrella, x+1-lenX)...)
		for i := lenX; i <= x; i++ {
			piano.piastrelle[i] = make([]Piastrella, lenY)
			for j := 0; j < lenY; j++ {
				piano.piastrelle[i][j] = creaPiastrella(i, j)
			}
		}

		for i:= lenX-1; i <= x; i++ {
			for j := 0; j < lenY; j++ {
				creaCirconvicini(*piano, i, j)
			}
		}
	} else if y >= lenY {
		for i := 0; i < lenX; i++ {
			piano.piastrelle[i] = append(piano.piastrelle[i], make([]Piastrella, y+1-lenY)...)
			for j:= lenY; j <= y; j++ {
				piano.piastrelle[i][j] = creaPiastrella(i, j)
			}
		}

		for i := 0; i < lenX; i++ {
			for j := lenY-1; j <= y; j++ {
				creaCirconvicini(*piano, i, j)
			}
		}
	} 
}

//funzione costruttrice che crea un piano di piastrelle, con n piastrelle per lato
func creaPiano(n, m int) Piano{
	piano := make([][]Piastrella, n)
	for i := 0; i < n; i++ {
		piano[i] = make([]Piastrella, m)
		for j := 0; j < m; j++ {
			piano[i][j] = creaPiastrella(i, j)
		}
	}
	p := Piano{piano, make([]Regola, 0)}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			creaCirconvicini(p, i, j)
		}
	}
	return p
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
func creaCirconvicini(piano Piano, i, j int) {
	circonvicini := make([]*Piastrella, 0, 9)
	lenX := len(piano.piastrelle)
	lenY := len(piano.piastrelle[0])
	//x-1, y+1
	if i > 0 && j < lenY-1 {
		circonvicini = append(circonvicini, &piano.piastrelle[i-1][j+1])
	} else {
		circonvicini = append(circonvicini, nil)
	}

	//x, y+1
	if j < lenY-1 {
		circonvicini = append(circonvicini, &piano.piastrelle[i][j+1])
	} else {
		circonvicini = append(circonvicini, nil)
	}

	//x+1, y+1
	if i < lenX-1 && j < lenY-1 {
		circonvicini = append(circonvicini, &piano.piastrelle[i+1][j+1])
	} else {
		circonvicini = append(circonvicini, nil)
	}

	//x-1, y
	if i > 0 {
		circonvicini = append(circonvicini, &piano.piastrelle[i-1][j])
	} else {
		circonvicini = append(circonvicini, nil)
	}

	//x, y
	circonvicini = append(circonvicini, &piano.piastrelle[i][j])

	//x+1, y
	if i < lenX-1 {
		circonvicini = append(circonvicini, &piano.piastrelle[i+1][j])
	} else {
		circonvicini = append(circonvicini, nil)
	}

	//x-1, y-1
	if i > 0 && j > 0 {
		circonvicini = append(circonvicini, &piano.piastrelle[i-1][j-1])
	} else {
		circonvicini = append(circonvicini, nil)
	}

	//x, y-1
	if j > 0 {
		circonvicini = append(circonvicini, &piano.piastrelle[i][j-1])
	} else {
		circonvicini = append(circonvicini, nil)
	}

	//x+1, y-1
	if i < lenX-1 && j > 0 {
		circonvicini = append(circonvicini, &piano.piastrelle[i+1][j-1])
	} else {
		circonvicini = append(circonvicini, nil)
	}

	piano.piastrelle[i][j].circonvicini = circonvicini
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

//trova il blocco, ovvero la regione di ampiezza massima, di piastrelle 
//adiacenti alla piastrella in posizione x, y
func trovaBlocco(piano Piano, x, y int) []Piastrella {
	if !Accesa(piano.piastrelle[x][y]) {
		return nil
	}
	piast := make([]Piastrella, 0)
	piast = append(piast, piano.piastrelle[x][y])
	for i:=0; i<len(piast); i++ {
		u := piast[i]
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
	if !Accesa(piano.piastrelle[x][y]) {
		return nil
	}
	piast := make([]Piastrella, 0)
	colore := piano.piastrelle[x][y].colore
	piast = append(piast, piano.piastrelle[x][y])
	for i:=0; i<len(piast); i++{
		u := piast[i]
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
					piast = append(piast, pcontrollo)
				}
			}
		}
	}

	return piast
}

//Funzione di utilità che svolge il merge per combinare due array 
//ordinati di regole
func merge(sinistro, destro []Regola) []Regola {
	risultato := make([]Regola, 0)
	i, j := 0, 0

	for i < len(sinistro) && j < len(destro) {
		if sinistro[i].usato <= destro[j].usato {
			risultato = append(risultato, sinistro[i])
			i++
		} else {
			risultato = append(risultato, destro[j])
			j++
		}
	}

	for i < len(sinistro) {
		risultato = append(risultato, sinistro[i])
		i++
	}

	for j < len(destro) {
		risultato = append(risultato, destro[j])
		j++
	}

	return risultato
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
func regola(piano Piano, regola string) {
	reg := strings.Split(regola, " ")
	beta := reg[0]
	reg = reg[1:]
	alfa := make(map[string]int)
	for i := 0; i < len(reg); i = i+2 {
		k, _ := strconv.Atoi(reg[i])
		a := reg[i+1]
		alfa[a] = k
	}

	piano.regole[len(piano.regole)-1] = Regola{alfa, beta, 0}
}

//Stampa e restituisce il colore e l’intensità di Piastrella(x, y). 
//Se Piastrella(x, y) è spenta, non stampa nulla.
func stato(piano Piano, x, y int) (string, int){
	if !Accesa(piano.piastrelle[x][y]) {
		return "", 0
	}
	fmt.Println(piano.piastrelle[x][y].colore, piano.piastrelle[x][y].intenisita)
	return piano.piastrelle[x][y].colore, piano.piastrelle[x][y].intenisita
}

//Stampa l’elenco delle regole di propagazione, nell’ordine attuale.
func stampa(piano Piano) {
	fmt.Println("(")
	for _, v := range piano.regole {
		//giallo: 1 rosso 1 blu
		str:=v.beta + ": "

		for k, v := range v.alfa {
			str += fmt.Sprintf("%d %s ", v, k)
		}

		fmt.Println(str, v.usato)
	}
	fmt.Println(")")
}

//Calcola e stampa la somma delle intensità delle piastrelle contenute 
//nel blocco di appartenenza di Piastrella(x, y).
//Se Piastrella(x, y) è spenta, restituisce 0.
func blocco(piano Piano, x, y int) {
	if !Accesa(piano.piastrelle[x][y]) {
		fmt.Println("0")
		return
	}
	pias := trovaBlocco(piano, x, y)
	somma := 0
	for i := 0; i < len(pias); i++ {
		somma += pias[i].intenisita
	}
	fmt.Println(somma)
}

//Calcola e stampa la somma delle intensit`a delle piastrelle contenute 
//nel blocco omogeneo di appartenenza di Piastrella(x, y). Se Piastrella(x, y) 
//`e spenta, restituisce 0.
func bloccoOmog(piano Piano, x, y int) {
	if !Accesa(piano.piastrelle[x][y]) {
		fmt.Println("0")
		return
	}
	pias := trovaBloccoOmogeneo(piano, x, y)
	somma := 0
	for i := 0; i < len(pias); i++ {
		somma += pias[i].intenisita
	}
	fmt.Println(somma)
}

//Applica a Piastrella(x, y) la prima regola di propagazione applicabile
//dell’elenco, ricolorando la piastrella. Se nessuna regola è applicabile,
//non viene eseguita alcuna operazione.
func propaga(piano Piano, x, y int) {
	intorno := piano.piastrelle[x][y].circonvicini[:4]
	intorno = append(intorno, piano.piastrelle[x][y].circonvicini[5:]...)

	mappaIntorno := make(map[string]int)
	for i := 0; i < len(intorno); i++ {
		if intorno[i] != nil {
			mappaIntorno[intorno[i].colore]++
		}
	}

	for i, v := range piano.regole {
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
			if !Accesa(piano.piastrelle[x][y]) {
				colora(piano, x, y, v.beta, 1)
			} else {
				piano.piastrelle[x][y].colore = v.beta
			}
			piano.regole[i].usato++
			return
		}
	}
}

//Propaga il colore sul blocco di appartenenza di Piastrella(x, y).
func propagaBlocco(piano Piano, x, y int) {
	if !Accesa(piano.piastrelle[x][y]) {
		return
	}
	blocco := trovaBlocco(piano, x, y)
	copia := make([]Piastrella, len(blocco))
	copy(copia, blocco)

	for i := 0; i < len(blocco); i++ {
		intorno := blocco[i].circonvicini[:4]
		intorno = append(intorno, blocco[i].circonvicini[5:]...)
	
		var reg Regola
		regvalida := false

		mappaIntorno := make(map[string]int)
		for i := 0; i < len(intorno); i++ {
			if intorno[i] != nil {
				mappaIntorno[intorno[i].colore]++
			}
		}
		for i, v := range piano.regole {
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
				piano.regole[i].usato++
				break
			}
		}
		
		if regvalida {
			copia[i].colore = reg.beta
		}
	}

	for i := 0; i < len(copia); i++ {
		if copia[i].colore != blocco[i].colore {
			piano.piastrelle[blocco[i].punti[0].x][blocco[i].punti[0].y].colore = copia[i].colore
		}
	}
}

//Ordina l’elenco delle regole di propagazione in base al consumo delle 
//regole stesse: la regola con consumo maggiore diventa l’ultima dell’elenco. 
//Se due regole hanno consumo uguale mantengono il loro ordine relativo. Lo fa
//implementando l'algoritmo di merge sort
func mergeSort(regole []Regola) []Regola {
	if len(regole) <= 1 {
		return regole
	}

	meta := len(regole) / 2
	sinistro := mergeSort(regole[:meta])
	destro := mergeSort(regole[meta:])

	return merge(sinistro, destro)
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
	if !Accesa(piano.piastrelle[x1][y1]) || !Accesa(piano.piastrelle[x2][y2]) {
		return
	}
	coda := make([]Piastrella, 0)
	visitato := make([][]bool, len(piano.piastrelle))
	for i := 0; i < len(piano.piastrelle); i++ {
		visitato[i] = make([]bool, len(piano.piastrelle[0]))
	}

	coda = append(coda, piano.piastrelle[x1][y1])
	visitato[x1][y1] = true
	dist := 1

	for len(coda) > 0 {
		size := len(coda)
		dist++

		for i := 0; i < size; i++ {
			piastrella := coda[0]
			coda = coda[1:]

			if piastrella.punti[0].x == x2 && piastrella.punti[0].y == y2 {
				fmt.Println("Lunghezza della pista più breve:", dist)
				return
			}

			for _, circonvicino := range piastrella.circonvicini {
				if circonvicino == nil {
					continue
				}
				
				if !visitato[circonvicino.punti[0].x][circonvicino.punti[0].y] {
					coda = append(coda, *circonvicino)
					visitato[circonvicino.punti[0].x][circonvicino.punti[0].y] = true
				}
			}
		}

		
	}

	fmt.Println("Lunghezza della pista più breve: ", dist)
}