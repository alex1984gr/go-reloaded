🧱 docs/how_to_work.md

# ⚙️ Go Reloaded — How to Work

## 🎯 Σκοπός
Αυτό το έγγραφο περιγράφει **τη ροή εργασίας**, **τον ρόλο των developers & AI agents**, και **τους κανόνες συνεργασίας** μέσα στο project Go Reloaded.  
Στόχος είναι η **πλήρης συνέπεια**, η **αναπαραγωγιμότητα** και η **καθαρή επικοινωνία** σε κάθε βήμα ανάπτυξης.

---

## 🧠 Ρόλοι

| Ρόλος | Περιγραφή |
|-------|------------|
| **Developer** | Γράφει και δοκιμάζει τον κώδικα, δημιουργεί νέα modules, φροντίζει για την τεκμηρίωση. |
| **AI Agent** | Εκτελεί αυτόματα tasks, εντοπίζει προβλήματα consistency, προτείνει βελτιώσεις και παράγει documentation. |
| **Auditor** | Ελέγχει commits άλλων, αξιολογεί tests, δίνει feedback και εγκρίνει merges. |

---

## 🧩 Ροή Εργασίας

1. **Δημιουργία Task**  
   - Κάθε νέα λειτουργία ή διόρθωση καταγράφεται στο `docs/blueprint-index.md`.  
   - Περιλαμβάνει περιγραφή, αρμόδιο developer/agent, ημερομηνία και κατάσταση (Pending / Done / Verified).

2. **Ανάπτυξη (Development)**  
   - Ο developer ή ο AI agent δημιουργεί branch:  
     ```bash
     git checkout -b feature/fixQuotes
     ```
   - Υλοποιεί το feature στο αντίστοιχο αρχείο του pipeline.  
   - Δημιουργεί tests στο φάκελο `tests/`.

3. **Έλεγχος (Testing)**  
   - Εκτελούνται tests:  
     ```bash
     go test ./...
     ```
   - Όλα τα tests πρέπει να περνούν πριν γίνει commit.  
   - Αν κάτι αποτύχει, ενημερώνεται ο auditor στο blueprint-index.md.

4. **Commit & Push**  
   - Format και lint πριν το commit:
     ```bash
     go fmt ./...
     go vet ./...
     ```
   - Έπειτα:
     ```bash
     git add .
     git commit -m "Implement replaceHex and add tests"
     git push
     ```

5. **Αναθεώρηση (Review)**  
   - Ο auditor διαβάζει το diff, τα σχόλια και τα tests.  
   - Αν εγκρίνει, σημειώνει το task ως `Verified ✅` στο blueprint-index.md.

---

## 🧱 Κανόνες Συνεργασίας

1. **Ποτέ δεν αλλάζουμε core αρχεία χωρίς document update.**  
2. **Όλες οι αλλαγές περνούν από testing.**  
3. **Κάθε commit πρέπει να έχει καθαρό, περιγραφικό μήνυμα.**  
4. **Κάθε module πρέπει να έχει docstring και test.**  
5. **Οποιαδήποτε αλλαγή σε pipeline → ενημέρωση σε architecture.md.**

---

## 🧪 Δοκιμές (Testing Workflow)

| Είδος Test | Περιγραφή | Παράδειγμα |
|-------------|------------|-------------|
| **Unit Test** | Ελέγχει μεμονωμένες συναρτήσεις | `TestReplaceHex_ValidHexToDecimal` |
| **Integration Test** | Ελέγχει ροή δεδομένων μεταξύ modules | `TestApplyTransformations_FullFlow` |
| **Behavioral Test** | Ελέγχει αν η έξοδος ταιριάζει με το επιθυμητό αποτέλεσμα | `TestFixQuotes_ComplexText` |

---

## 📘 Κανόνες Ενημέρωσης Documentation

Μετά από κάθε αλλαγή:

| Αλλαγή | Ενημερώνεται αρχείο |
|--------|----------------------|
| Νέα λειτουργία | `blueprint-index.md` |
| Αλλαγή αρχιτεκτονικής | `architecture.md` |
| Κανόνας κώδικα | `coding_standards.md` |
| Νέος όρος ή συντομογραφία | `glossary.md` |

---

## 🔐 Checklist πριν από κάθε Commit

- [ ] Όλα τα tests περνούν (`go test ./...`)
- [ ] Δεν υπάρχουν unused imports (`go vet ./...`)
- [ ] Ο κώδικας είναι formatted (`go fmt ./...`)
- [ ] Η τεκμηρίωση είναι ενημερωμένη
- [ ] Το blueprint-index.md έχει ενημερωθεί
- [ ] Ο auditor έχει ενημερωθεί

---

## 🧭 Παράδειγμα Ροής (End-to-End)

```bash
# 1. Δημιουργία νέου feature
git checkout -b feature/fixPunctuation

# 2. Ανάπτυξη και test
vim pipeline/fixPunctuation.go
vim tests/fixPunctuation_test.go
go test ./...

# 3. Format & commit
go fmt ./...
git add .
git commit -m "Implement fixPunctuation and tests"

# 4. Push & ενημέρωση
git push
vim docs/blueprint-index.md

🤝 Στόχος

Να υπάρχει τέλεια συνεργασία ανθρώπου και μηχανής.
Κάθε commit είναι ένα βήμα προς καθαρότερο, πιο συντονισμένο και διαφανές project.