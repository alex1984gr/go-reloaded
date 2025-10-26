# 🧠 Go Reloaded
 
 ## 📖 Overview
 
 Το **Go Reloaded** είναι ένα εργαλείο αυτόματης επεξεργασίας και μορφοποίησης κειμένου, γραμμένο εξ ολοκλήρου στη γλώσσα **Go**. 
 Διαβάζει ένα αρχείο εισόδου, εντοπίζει ειδικές ενδείξεις (tags) και στίξη, και παράγει ένα καθαρό, διορθωμένο αρχείο εξόδου. 
 
 Η αρχιτεκτονική του βασίζεται σε **modular pipeline**, όπου κάθε στάδιο της επεξεργασίας είναι ανεξάρτητο και μπορεί να δοκιμαστεί ή να επεκταθεί μεμονωμένα. 
 Το project έχει σχεδιαστεί να συνεργάζεται με **AI Agents** (Codex, Copilot, Claude, ChatGPT) που ακολουθούν το πρωτόκολλο που περιγράφεται στο AGENTS.md.
 
 ---
 
 ## 📂 Project Structure
bash
📁 go-reloaded/
├── main.go
│   # Κεντρικό σημείο εκκίνησης του προγράμματος
│
├── pipeline/
│   ├── readInput.go              # Διαβάζει το αρχείο εισόδου
│   ├── tokenize.go               # Διαχωρίζει tokens
│   ├── replaceHex.go             # Μετατρέπει hex σε δεκαδικό
│   ├── replaceBin.go             # Μετατρέπει binary σε δεκαδικό
│   ├── applyCaseTransform.go     # Εφαρμόζει (up), (low), (cap)
│   ├── formatPunctuation.go      # Διορθώνει στίξη
│   ├── fixQuotes.go              # Τοποθετεί σωστά quotes
│   ├── fixArticles.go            # Αντικαθιστά “a” με “an”
│   ├── applyTransformations.go   # Ενοποιεί όλα τα βήματα
│   └── writeOutput.go            # Γράφει το αποτέλεσμα στο αρχείο εξόδου
│
├── tests/
│   ├── readInput_test.go
│   ├── tokenize_test.go
│   ├── replaceHex_test.go
│   ├── replaceBin_test.go
│   ├── applyCaseTransform_test.go
│   ├── formatPunctuation_test.go
│   ├── fixQuotes_test.go
│   ├── fixArticles_test.go
│   ├── applyTransformations_test.go
│   └── writeOutput_test.go
│   # Όλα τα test αρχεία είναι βασισμένα στο TDD pipeline (Test Driven Development)
│
├── docs/
│   ├── architecture.md           # Αναλύει το data flow και την εσωτερική αρχιτεκτονική
│   ├── coding_standards.md       # Κανόνες γραφής Go κώδικα και ονοματοδοσίας
│   ├── how_to_work.md            # Οδηγίες για developers και AI agents
│   ├── blueprint-index.md        # Κεντρικός πίνακας αναφοράς των tasks
│   └── glossary.md               # Ορολογία για το project
│
├── tasks/
│   ├── TASK-A1.md                # Παράδειγμα: replaceHex() implementation
│   ├── TASK-A2.md                # Παράδειγμα: replaceBin() implementation
│   └── ...                       # Επιπλέον task αρχεία για AI Agents
│
├── .github/
│   ├── workflows/
│   │   └── ci.yml                # GitHub Actions για αυτόματο testing και QA
│   └── .actrc                    # Config για local CI testing με act
│
├── AGENTS.md                     # Πρωτόκολλο εκτέλεσης για AI Agents
└── README.md                     # Κεντρική τεκμηρίωση του project

⚙️ Functional Pipeline 

readInput
↓
tokenize
↓
applyTransformations
↓
formatPunctuation
↓
fixQuotes
↓
fixArticles
↓
writeOutput

🧩 Core Features

| Εντολή      | Περιγραφή                                                     | Παράδειγμα                                                |
| ----------- | ------------------------------------------------------------- | --------------------------------------------------------- |
| `(hex)`     | Μετατρέπει αριθμό hex σε δεκαδικό                             | `"1E (hex)" → "30"`                                       |
| `(bin)`     | Μετατρέπει αριθμό binary σε δεκαδικό                          | `"10 (bin)" → "2"`                                        |
| `(up)`      | Κάνει τη λέξη πριν κεφαλαία                                   | `"go (up)" → "GO"`                                        |
| `(low)`     | Κάνει τη λέξη πριν πεζά                                       | `"STOP (low)" → "stop"`                                   |
| `(cap)`     | Κάνει τη λέξη πριν κεφαλαιογράμματη                           | `"bridge (cap)" → "Bridge"`                               |
| `(up, n)`   | Μετατρέπει τις **n** προηγούμενες λέξεις σε κεφαλαία          | `"so exciting (up,2)" → "SO EXCITING"`                    |
| `(low, n)`  | Μετατρέπει τις **n** προηγούμενες λέξεις σε πεζά              | `"WOW THAT'S COOL (low,3)" → "wow that's cool"`           |
| `(cap, n)`  | Μετατρέπει τις **n** προηγούμενες λέξεις σε κεφαλαιογράμματες | `"brooklyn bridge park (cap,3)" → "Brooklyn Bridge Park"` |
| Punctuation | Διορθώνει τη στίξη με σωστά κενά                              | `"Hello , world !" → "Hello, world!"`                     |
| `'quotes'`  | Τοποθετεί σωστά τα μονά εισαγωγικά                            | `" ' awesome ' " → "'awesome'"`                           |
| `a → an`    | Αντικαθιστά το “a” με “an” πριν από φωνήεν ή “h”              | `"a apple" → "an apple"`                                  |


🧱 Function Breakdown 

| Function                 | Περιγραφή                                |
| ------------------------ | ---------------------------------------- |
| `readInput()`            | Διαβάζει το αρχείο εισόδου               |
| `tokenize()`             | Διαχωρίζει λέξεις, στίξη και tags        |
| `replaceHex()`           | Μετατρέπει hex σε δεκαδικό               |
| `replaceBin()`           | Μετατρέπει binary σε δεκαδικό            |
| `applyCaseTransform()`   | Εφαρμόζει αλλαγές κεφαλαίων/πεζών        |
| `formatPunctuation()`    | Διορθώνει τα κενά γύρω από σημεία στίξης |
| `fixQuotes()`            | Τοποθετεί σωστά τα quotes                |
| `fixArticles()`          | Ελέγχει το “a/an”                        |
| `applyTransformations()` | Συνδυάζει όλες τις μετατροπές            |
| `writeOutput()`          | Γράφει το τελικό αποτέλεσμα              |

🧪 Testing & Quality Assurance

Όλες οι συναρτήσεις ακολουθούν τη φιλοσοφία Test Driven Development (TDD).
Τα tests βρίσκονται στον φάκελο /tests και εκτελούνται με:

go test ./tests/...


// Για local CI execution:

act -j build -W .github/workflows/ci.yml

🤖 AI Integration (AGENTS.md)

Το Go Reloaded είναι σχεδιασμένο ώστε να συνεργάζεται με AI Agents.
Ο φάκελος tasks/ περιέχει markdown αρχεία όπου κάθε AI Agent εκτελεί τα βήματα:

Analyze & Confirm

Generate the Tests

Generate the Code

QA & Mark Complete

Οι οδηγίες χρήσης και οι κανόνες για Agents περιγράφονται αναλυτικά στο AGENTS.md. Δες [AGENTS.md](./AGENTS.md) για πλήρεις οδηγίες εκτέλεσης.

Επιπλέον, στο docs/how_to_work.md υπάρχουν τα workflows που κάθε agent ακολουθεί:

analyze → ask for operator confirmation → implement tests → implement code → QA

🧭 Developer Docs

Αν είσαι νέος developer ή agent, ξεκίνα από εδώ:

Αρχείο	Σκοπός
docs/architecture.md	Αναλύει το pipeline και το data flow
docs/how_to_work.md	Οδηγεί developers & agents στα βήματα εργασίας
docs/coding_standards.md	Οδηγίες για καθαρό, συνεπή Go κώδικα
docs/blueprint-index.md	Επισκόπηση όλων των tasks και progress
AGENTS.md	Το πρωτόκολλο εκτέλεσης για AI Agents

🚀 Execution

Τρέξε το πρόγραμμα με:

go run main.go input.txt output.txt


Το αποτέλεσμα θα αποθηκευτεί στο αρχείο output.txt.