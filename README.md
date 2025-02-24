# 🏗️ Progetto Mattonelle  

**Laboratorio di Algoritmi e Basi di Dati – Sessione Estiva**  

Questo progetto implementa un sistema per la gestione e l'elaborazione di insiemi di **piastrelle digitali** su un piano bidimensionale.  
L'obiettivo è analizzare le configurazioni delle piastrelle, applicare regole di propagazione e studiare l'influenza reciproca tra le piastrelle sulla base del loro stato.  

---

## ✨ Funzionalità  

Il programma supporta le seguenti operazioni:  

### 🎨 Manipolazione delle Piastrelle  

- **`colora(x, y, α, i)`** – Assegna alla piastrella `(x, y)` il colore `α` con intensità `i`, indipendentemente dal suo stato precedente.  
- **`spegni(x, y)`** – Spegne la piastrella `(x, y)`. Se è già spenta, l'operazione non ha effetto.  
- **`stato(x, y)`** – Restituisce il colore e l'intensità della piastrella `(x, y)`. Se è spenta, restituisce una stringa vuota e `0`.  
- **`blocco(x, y)`** – Calcola e restituisce la somma delle intensità delle piastrelle nel blocco di `(x, y)`. Se la piastrella è spenta, restituisce `0`.  
- **`bloccoOmog(x, y)`** – Calcola la somma delle intensità delle piastrelle appartenenti al **blocco omogeneo** di `(x, y)`. Se la piastrella è spenta, restituisce `0`.  

### 🔥 Propagazione e Regole  

- **`regola(k1, α1, k2, α2, ..., kn, αn, β)`** – Definisce una regola di propagazione `k1α1 + k2α2 + ... + knαn → β` e la aggiunge all’elenco delle regole.  
- **`stampa`** – Stampa l’elenco delle regole di propagazione nell’ordine attuale.  
- **`propaga(x, y)`** – Applica alla piastrella `(x, y)` la prima regola applicabile dell’elenco. Se nessuna regola è applicabile, l'operazione viene ignorata.  
- **`propagaBlocco(x, y)`** – Propaga il colore sull’intero blocco di appartenenza della piastrella `(x, y)`.  
- **`ordina`** – Ordina l’elenco delle regole di propagazione in base al **consumo** delle regole: la più utilizzata diventa l'ultima. A parità di consumo, l'ordine relativo viene mantenuto.  

### 🏁 Percorsi e Analisi  

- **`pista(x, y, s)`** – Stampa il percorso che parte dalla piastrella `(x, y)` e segue la sequenza di direzioni `s`, se definito.  
- **`lung(x1, y1, x2, y2)`** – Determina la lunghezza del percorso più breve tra `(x1, y1)` e `(x2, y2)`. Se non esiste, non restituisce nulla.  

---

## 🧪 Testing  

Il progetto include due suite di test:  

1. **Test I/O** – Una serie di file di test in formato `.txt` contenenti input/output predefiniti.  
2. **Test Go** – Una collezione di test unitari scritti in **Go**, realizzati in collaborazione con [Luca Carone](https://github.com/lochy54/test.git).  

### 🔧 Esecuzione dei Test  

Assicurati di aggiornare il valore della variabile **`prog`** nel file `Luglio_test.go` con il nome del tuo programma prima di eseguire i test:  

```sh
cd test/
go test Luglio_test.go
```

---

## 📜 Licenza
Questo progetto è distribuito sotto licenza GPL-3.0.
Consulta i dettagli qui: [GNU GPL v3.0.](https://www.gnu.org/licenses/gpl-3.0.en.html)
