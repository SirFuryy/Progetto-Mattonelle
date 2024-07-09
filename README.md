# Progetto-Mattonelle
Progetto di Algoritmi e basi di dati laboratorio sessione estiva

Il progetto si propone di mantenere e risolvere diverse operazioni di insiemi di piastrelle digitali disposte su un piano bidimensionale, di analizzare le configurazioni e di studiare l'influenza che queste configurazioni esercitano sulle piastrelle circostanti sulla base del loro stato.

## Operazioni
Il programma permette di eseguire le seguenti operazioni:

* **colora(x, y, α, i)**  
Colora Piastrella(x, y) di colore α e intensit`a i, qualunque sia lo stato di Piastrella(x, y) prima
dell’operazione.
* **spegni(x, y)**  
Spegne Piastrella(x, y). Se Piastrella(x, y) è già spenta, non fa nulla.
* **regola(k1, α1, k2, α2, . . . , kn, αn, β)**  
Definisce la regola di propagazione k1α1 + k2α2 + · · · + knαn → β e la inserisce in fondo all’elenco
delle regole.
* **stato(x, y)**  
Stampa e restituisce il colore e l’intensità di Piastrella(x, y). Se Piastrella(x, y) è spenta, non stampa
nulla e restituisce la stringa vuota e l’intero 0.
* **stampa**  
Stampa l’elenco delle regole di propagazione, nell’ordine attuale.
* **blocco(x, y)**  
Calcola le stampa a somma delle intensità delle piastrelle contenute nel blocco di appartenenza di
Piastrella(x, y). Se Piastrella(x, y) è spenta, stampa 0.
* **colora(x, y, α, i)**  
Colora Piastrella(x, y) di colore α e intensit`a i, qualunque sia lo stato di Piastrella(x, y) prima
dell’operazione.
* **spegni(x, y)**  
Spegne Piastrella(x, y). Se Piastrella(x, y) è già spenta, non fa nulla.
* **regola(k1, α1, k2, α2, . . . , kn, αn, β)**  
Definisce la regola di propagazione k1α1 + k2α2 + · · · + knαn → β e la inserisce in fondo all’elenco
delle regole.
* **stato(x, y)**  
Stampa e restituisce il colore e l’intensità di Piastrella(x, y). Se Piastrella(x, y) è spenta, non stampa
nulla e restituisce la stringa vuota e l’intero 0.
* **stampa**  
Stampa l’elenco delle regole di propagazione, nell’ordine attuale.
* **blocco(x, y)**  
Calcola le stampa a somma delle intensit`a delle piastrelle contenute nel blocco di appartenenza di
Piastrella(x, y). Se Piastrella(x, y) è spenta, stampa 0.

* **bloccoOmog(x, y)**  
Calcola e stampa la somma delle intensità delle piastrelle contenute nel blocco omogeneo di appartenenza di Piastrella(x, y). Se Piastrella(x, y) è spenta, restituisce 0.
* **propaga(x, y)**  
Applica a Piastrella(x, y) la prima regola di propagazione applicabile dell’elenco, ricolorando la
piastrella. Se nessuna regola è applicabile, non viene eseguita alcuna operazione.
* **propagaBlocco(x, y)**  
Propaga il colore sul blocco di appartenenza di Piastrella(x, y).
* **ordina**  
Ordina l’elenco delle regole di propagazione in base al consumo delle regole stesse: la regola con
consumo maggiore diventa l’ultima dell’elenco. Se due regole hanno consumo uguale mantengono
il loro ordine relativo.
* **pista(x, y, s)**  
Stampa la pista che parte da Piastrella(x, y) e segue la sequenza di direzioni s, se tale pista è
definita. Altrimenti non stampa nulla.
* **lung(x1, y1, x2, y2)**  
Determina la lunghezza della pista più breve che parte da Piastrella(x1, y1) e arriva in Piastrella(x2, y2).
Altrimenti non stampa nulla.

## Test
Nel progetto sono presenti due cartelle di test, una contenente dei test input output in formato .txt e una contenente una serie di test in go creati in collaborazione con [Luca Carone](https://github.com/lochy54/test.git), di cui la cartella è un effettivo collegamento al suo progetto git a riguardo, che possono essere runnati nel seguente modo, dopo aver prestato attenzione al modificare il valore della variabile **prog** nel file **Luglio_test.go** con il nome del proprio programma da testare

```bash
  cd .\test\

  go test .\Luglio_test.go
```
## License
[GPL-3.0](https://www.gnu.org/licenses/gpl-3.0.en.html)