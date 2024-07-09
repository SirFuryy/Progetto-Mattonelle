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
//con un numero di piastrelle n*m, con n,m > 0
//
//il piano è composto da piastrelle e da regole di propagazione 
type piano struct {
	piastrelle map[Punto]*Piastrella
	regole []Regola
}

//struttura dati per rappresentare una regola di propagazione,
//con un insieme di alfa, che sono i colori delle piastrelle che circondano la
//piastrella a cui applicare la regola, e un beta, che è il colore che verrà propagato
type Regola struct {
	alfa map[string]int
	beta string
	ordine []string		//mantiene l'ordine con cui sono inseriti gli alfa perchè le mappe non mantengono l'ordine
	usato int
}

var eseguito bool = true

func main() {
	var input string
	piano := piano{make(map[Punto]*Piastrella), make([]Regola, 0)}

	scanner := bufio.NewScanner(os.Stdin)
	for eseguito {
		if scanner.Scan() {
			input = scanner.Text()
		}
		piano = esegui(piano, input)
	}
}

//funzione che esegue il comando passato in input
func esegui(piano piano, input string) piano{
	inp := strings.Split(input, " ")
	comando := inp[0]
	inp = inp[1:]
	switch comando {
	case "C":
		x, _ := strconv.Atoi(inp[0])
		y, _ := strconv.Atoi(inp[1])
		alpha := inp[2]
		i, _ := strconv.Atoi(inp[3])
		colora(piano, x, y, alpha, i)
		break
	case "S":
		x, _ := strconv.Atoi(inp[0])
		y, _ := strconv.Atoi(inp[1])
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
		stato(piano, x, y)
		break
	case "s":
		stampa(piano)
		break
	case "b":
		x, _ := strconv.Atoi(inp[0])
		y, _ := strconv.Atoi(inp[1])
		blocco(&piano, x, y)
		break
	case "B":
		x, _ := strconv.Atoi(inp[0])
		y, _ := strconv.Atoi(inp[1])
		bloccoOmog(&piano, x, y)
		break
	case "p":
		x, _ := strconv.Atoi(inp[0])
		y, _ := strconv.Atoi(inp[1])
		propaga(&piano, x, y)
		break
	case "P":
		x, _ := strconv.Atoi(inp[0])
		y, _ := strconv.Atoi(inp[1])
		propagaBlocco(&piano, x, y)
		break
	case "o":
		piano.regole = mergeSort(piano.regole)
		break
	case "q":
		eseguito = false
		fmt.Println()
		break
	case "t":
		x, _ := strconv.Atoi(inp[0])
		y, _ := strconv.Atoi(inp[1])
		s := inp[2]
		pista(&piano, x, y, s)
		break
	case "L":
		x1, _ := strconv.Atoi(inp[0])
		y1, _ := strconv.Atoi(inp[1])
		x2, _ := strconv.Atoi(inp[2])
		y2, _ := strconv.Atoi(inp[3])
		lung(&piano, x1, y1, x2, y2)
		break
	default:
		break
	}
	return piano
}

/*  FUNZIONI DI UTILITA  */

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
func creaCirconvicini(piano piano, piast *Piastrella, i, j int) {
	circonvicini := make([]*Piastrella, 0, 9)
	//x-1, y+1
	if val, exist := piano.piastrelle[Punto{i-1, j+1}]; exist {
		circonvicini = append(circonvicini, val)
		val.circonvicini[8] = piast
	} else {
		circonvicini = append(circonvicini, nil)
	}

	//x, y+1
	if val, exist := piano.piastrelle[Punto{i, j+1}]; exist {
		circonvicini = append(circonvicini, val)
		val.circonvicini[7] = piast
	} else {
		circonvicini = append(circonvicini, nil)
	}

	//x+1, y+1
	if val, exist := piano.piastrelle[Punto{i+1, j+1}]; exist {
		circonvicini = append(circonvicini, val)
		val.circonvicini[6] = piast
	} else {
		circonvicini = append(circonvicini, nil)
	}

	//x-1, y
	if val, exist := piano.piastrelle[Punto{i-1, j}]; exist {
		circonvicini = append(circonvicini, val)
		val.circonvicini[5] = piast
	} else {
		circonvicini = append(circonvicini, nil)
	}

	//x, y
	circonvicini = append(circonvicini, piast)

	//x+1, y
	if val, exist := piano.piastrelle[Punto{i+1, j}]; exist {
		circonvicini = append(circonvicini, val)
		val.circonvicini[3] = piast
	} else {
		circonvicini = append(circonvicini, nil)
	}

	//x-1, y-1
	if val, exist := piano.piastrelle[Punto{i-1, j-1}]; exist {
		circonvicini = append(circonvicini, val)
		val.circonvicini[2] = piast
	} else {
		circonvicini = append(circonvicini, nil)
	}

	//x, y-1
	if val, exist := piano.piastrelle[Punto{i, j-1}]; exist {
		circonvicini = append(circonvicini, val)
		val.circonvicini[1] = piast
	} else {
		circonvicini = append(circonvicini, nil)
	}

	//x+1, y-1
	if val, exist := piano.piastrelle[Punto{i+1, j-1}]; exist {
		circonvicini = append(circonvicini, val)
		val.circonvicini[0] = piast	
	} else {
		circonvicini = append(circonvicini, nil)
	}

	piast.circonvicini = circonvicini
}

//funzione di utilità che restituisce true se la piastrella è accessa, 
//false altrimenti
func Accesa(piano *piano, x,y int) bool {
	if val, exist := piano.piastrelle[Punto{x, y}]; exist {
		return val.intenisita > 0
	}
	return false
}

//trova il blocco, ovvero la regione di ampiezza massima, di piastrelle 
//adiacenti alla piastrella in posizione x, y
func trovaBlocco(piano *piano, x, y int) []*Piastrella {
	if !Accesa(piano, x, y) {
		return nil
	}

	visitato := make(map[Punto]bool)
	coda := []*Piastrella{piano.piastrelle[Punto{x, y}]}	//creo e inserisco il primo elemento
	blocco := []*Piastrella{piano.piastrelle[Punto{x, y}]}	//creo e inserisco il primo elemento
	visitato[Punto{x, y}] = true

	for len(coda) > 0 {
		curr := coda[0]
		coda = coda[1:]

		for _, circonvicino := range curr.circonvicini {
			if circonvicino != nil && !visitato[circonvicino.punti[0]] && circonvicino.intenisita > 0 {
				coda = append(coda, circonvicino)
				blocco = append(blocco, circonvicino)
				visitato[circonvicino.punti[0]] = true
			}
		}
	}

	return blocco
}

//trova il blocco omogeneo, ovvero la regione di ampiezza massima, di 
//piastrelle adiacenti alla piastrella in posizione x, y con lo stesso 
//colore della piastrella in posizione x, y
func trovaBloccoOmogeneo(piano *piano, x, y int) []*Piastrella {
	if !Accesa(piano, x, y) {
		return nil
	}

	visitato := make(map[Punto]bool)
	colore := piano.piastrelle[Punto{x, y}].colore
	coda := []*Piastrella{piano.piastrelle[Punto{x, y}]}
	blocco := []*Piastrella{piano.piastrelle[Punto{x, y}]}
	visitato[Punto{x, y}] = true

	for len(coda) > 0 {
		curr := coda[0]
		coda = coda[1:]

		for _, circonvicino := range curr.circonvicini {
			if circonvicino != nil && !visitato[circonvicino.punti[0]] && circonvicino.intenisita > 0 && circonvicino.colore == colore{
				coda = append(coda, circonvicino)
				blocco = append(blocco, circonvicino)
				visitato[circonvicino.punti[0]] = true
			}
		}
	}

	return blocco
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

//funzione che colora la piastrella in posizione x, y con il colore 
//e l'intensità passati, a prescindere dallo stato precedente della piastrella.
//Se la piastrella non esiste, la crea
func colora(p piano, x int, y int, alpha string, i int) {
	punt := Punto{x, y}
	if _, exist := p.piastrelle[punt]; !exist {
		piast := Piastrella{[]Punto{{x, y}, {x,y+1}, {x+1,y+1}, {x+1,y}}, alpha, i, nil}
		p.piastrelle[punt] = &piast
		creaCirconvicini(p, &piast, x, y)
		return
	} 
	p.piastrelle[punt].colore = alpha
	p.piastrelle[punt].intenisita = i
}

//Spegne Piastrella(x, y). Se Piastrella(x, y) è già spenta, 
//non fa nulla.
func spegni(p piano, x int, y int) {
	if !Accesa(&p, x, y) {
		return
	}
	p.piastrelle[Punto{x,y}].colore = ""
	p.piastrelle[Punto{x,y}].intenisita = 0
}

//Trasforma la regola di propagazione in formato β k1 α1 k2 α2 ... kn αn e 
//la inserisce in fondo all’elenco delle regole.
func regola(p piano, r string) {
	reg := strings.Split(r, " ")
	beta := reg[0]
	reg = reg[1:]
	alfa := make(map[string]int)
	ordine := make([]string, 0)
	for i := 0; i < len(reg); i = i+2 {
		k, _ := strconv.Atoi(reg[i])
		a := reg[i+1]
		alfa[a] = k
		ordine = append(ordine, a)
	}

	p.regole[len(p.regole)-1] = Regola{alfa, beta, ordine, 0}
}

//Stampa e restituisce il colore e l’intensità di Piastrella(x, y). 
//Se Piastrella(x, y) è spenta, non stampa nulla.
func stato(p piano, x int, y int) (string, int){
	if !Accesa(&p, x, y) {
		return "", 0
	}
	fmt.Println(p.piastrelle[Punto{x,y}].colore, p.piastrelle[Punto{x,y}].intenisita)
	return p.piastrelle[Punto{x,y}].colore, p.piastrelle[Punto{x,y}].intenisita
}

//Stampa l’elenco delle regole di propagazione, nell’ordine attuale.
func stampa(p piano) {
	fmt.Println("(")
	for _, v := range p.regole {
		//giallo: 1 rosso 1 blu
		str:=v.beta + ": "

		for _, alpha := range v.ordine {
			str += fmt.Sprintf("%d %s ", v.alfa[alpha], alpha)
		}

		fmt.Println(str)
	}
	fmt.Println(")")
}

//Calcola e stampa la somma delle intensità delle piastrelle contenute 
//nel blocco di appartenenza di Piastrella(x, y).
//Se Piastrella(x, y) è spenta, restituisce 0.
func blocco(piano *piano, x, y int) {
	if !Accesa(piano, x, y) {
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
func bloccoOmog(piano *piano, x, y int) {
	if !Accesa(piano, x, y) {
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
func propaga(piano *piano, x, y int) {
	if _, exist := piano.piastrelle[Punto{x, y}]; !exist {
		colora(*piano, x, y, "", 0)		//la coloro, dunque la creo, ma come se fosse spenta
	}
	intorno := piano.piastrelle[Punto{x,y}].circonvicini[:4]
	intorno = append(intorno, piano.piastrelle[Punto{x,y}].circonvicini[5:]...)

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
			piano.piastrelle[Punto{x,y}].colore = v.beta
			if piano.piastrelle[Punto{x,y}].intenisita == 0 {
				piano.piastrelle[Punto{x,y}].intenisita = 1
			}
			piano.regole[i].usato++
			return
		}
	}

}

//Propaga il colore sul blocco di appartenenza di Piastrella(x, y).
func propagaBlocco(piano *piano, x, y int) {
	if !Accesa(piano, x, y) {
		return
	}
	blocco := trovaBlocco(piano, x, y)
	copia := make([]Piastrella, len(blocco))
	for i := 0; i < len(blocco); i++ {
		copia[i] = *blocco[i]
	}

	for i := 0; i < len(blocco); i++ {
		intorno := blocco[i].circonvicini[:4]
		intorno = append(intorno, blocco[i].circonvicini[5:]...)

		mappaIntorno := make(map[string]int)
		for j := 0; j < len(intorno); j++ {
			if intorno[j] != nil {
				mappaIntorno[intorno[j].colore]++
			}
		}
		
		for j, v := range piano.regole {
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
				piano.regole[j].usato++
				copia[i].colore = v.beta
				break
			}
		}
	}

	for i := 0; i < len(copia); i++ {
		if copia[i].colore != blocco[i].colore {
			piano.piastrelle[Punto{blocco[i].punti[0].x,blocco[i].punti[0].y}].colore = copia[i].colore
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
//s sarà una stringa nel formatto NN,NO,NE,...
func pista(piano *piano, x, y int, s string) {
	if !Accesa(piano, x, y) {
		return
	}
	indicazioni := make([]string, 0)
	for i := 0; i < len(s); i += 3 {
		indicazioni = append(indicazioni, s[i : i+2])
	}

	piastControllata := piano.piastrelle[Punto{x,y}]
	
	piastrelle := make([]*Piastrella, 0)
	piastrelle = append(piastrelle, piastControllata)
	for i := 0; i < len(indicazioni); i++ {
		switch indicazioni[i] {
		case "NO":	//x-1, y+1
			ppros:= piastControllata.circonvicini[0]
			if ppros != nil && ppros.intenisita > 0 { 
					piastControllata = ppros
					piastrelle = append(piastrelle, piastControllata)
			} else {return}
			break
		case "NN":	//x, y+1
			ppros:= piastControllata.circonvicini[1]
			if ppros != nil && ppros.intenisita > 0 { 
				piastControllata = ppros
				piastrelle = append(piastrelle, piastControllata)
			} else {return}
			break
		case "NE":	//x+1, y+1
			ppros:= piastControllata.circonvicini[2]
			if ppros != nil && ppros.intenisita > 0 { 
				piastControllata = ppros
				piastrelle = append(piastrelle, piastControllata)
			} else {return}
			break
		case "OO":	//x-1, y
			ppros:= piastControllata.circonvicini[3]
			if ppros != nil && ppros.intenisita > 0 { 
				piastControllata = ppros
				piastrelle = append(piastrelle, piastControllata)
			} else {return}
			break
		case "EE":	//x+1, y
			ppros:= piastControllata.circonvicini[5]
			if ppros != nil && ppros.intenisita > 0 { 
				piastControllata = ppros
				piastrelle = append(piastrelle, piastControllata)
			} else {return}
			break
		case "SO":	//x-1, y-1
			ppros:= piastControllata.circonvicini[6]
			if ppros != nil && ppros.intenisita > 0 { 
				piastControllata = ppros
				piastrelle = append(piastrelle, piastControllata)
			} else {return}
			break
		case "SS":	//x, y-1
			ppros:= piastControllata.circonvicini[7]
			if ppros != nil && ppros.intenisita > 0 { 
				piastControllata = ppros
				piastrelle = append(piastrelle, piastControllata)
			} else {return}
			break
		case "SE":	//x+1, y-1
			ppros:= piastControllata.circonvicini[8]
			if ppros != nil && ppros.intenisita > 0 { 
				piastControllata = ppros
				piastrelle = append(piastrelle, piastControllata)
			} else {return}
			break		
		}
	}

	fmt.Println("[")
	for i := 0; i < len(piastrelle); i++ {
		fmt.Println(piastrelle[i].punti[0].x, piastrelle[i].punti[0].y, piastrelle[i].colore, piastrelle[i].intenisita)
	}
	fmt.Println("]")
}

//Determina la lunghezza della pista più breve che parte da Piastrella(x1, y1)
//e arriva in Piastrella(x2, y2). Altrimenti non stampa nulla.
func lung(piano *piano, x1, y1, x2, y2 int) {
	if !Accesa(piano, x1, y1) || !Accesa(piano, x2, y2) {
		return
	}

	coda := make([]*Piastrella, 0)
	visitato := make(map[Punto]bool)

	coda = append(coda, piano.piastrelle[Punto{x1, y1}])
	visitato[Punto{x1, y1}] = true
	dist := 1

	for len(coda) > 0 {
		numPiastrelle := len(coda)
		for i := 0; i < numPiastrelle; i++ {
			piastrellaCorrente := coda[0]
			coda = coda[1:]
			if piastrellaCorrente.punti[0].x == x2 && piastrellaCorrente.punti[0].y == y2 {
				fmt.Println(dist)
				return
			}
			for _, circonvicino := range piastrellaCorrente.circonvicini {
				if circonvicino != nil && !visitato[circonvicino.punti[0]] {
					coda = append(coda, circonvicino)
					visitato[circonvicino.punti[0]] = true
				}
			}
		}
		dist++
	}
}

