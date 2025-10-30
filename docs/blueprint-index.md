# 🧭 Go Reloaded — Blueprint Index

## 🎯 Σκοπός
Το αρχείο αυτό χρησιμεύει ως **κεντρικός πίνακας αναφοράς** για όλο το project **Go Reloaded**.  
Παρακολουθεί την **πρόοδο**, τις **αναθέσεις**, και τις **ενημερώσεις** σε κάθε component, ώστε να διασφαλίζεται συνοχή ανάμεσα σε developers, AI agents και auditors.

---

## 📁 Δομή Project

| Κατηγορία | Φάκελος | Περιγραφή |
|------------|----------|------------|
| Pipeline Functions | `pipeline/` | Περιέχει όλα τα modules επεξεργασίας κειμένου |
| Tests | `tests/` | Unit & integration tests για κάθε module |
| Documentation | `docs/` | Όλη η τεκμηρίωση (architecture, how_to_work, glossary κ.ά.) |
| Main Program | `cmd/` | Κεντρική είσοδος της εφαρμογής |
| Data | `data/` | Input & output αρχεία για δοκιμές |

---

## 📜 Task Tracker

| # | Περιγραφή Εργασίας | Αρχείο / Module | Assigned To | Κατάσταση | Τελευταία Ενημέρωση |
|---|---------------------|------------------|--------------|------------|----------------------|
| 1 | Ανάγνωση αρχείου εισόδου | `pipeline/readInput.go` | Developer | ✅ Done | 2025-10-26 |
| 2 | Tokenization (διάσπαση σε λέξεις) | `pipeline/tokenize.go` | Developer | ✅ Done | 2025-10-26 |
| 3 | Αντικατάσταση binary → decimal | `pipeline/replaceBin.go` | Developer | ✅ Done | 2025-10-26 |
| 4 | Αντικατάσταση hex → decimal | `pipeline/replaceHex.go` | Developer | ✅ Done | 2025-10-26 |
| 5 | Εφαρμογή μετατροπών κεφαλαίων/πεζών | `pipeline/applyCaseTransform.go` | AI Agent | ✅ Done | 2025-10-26 |
| 6 | Διόρθωση σημείων στίξης | `pipeline/formatPunctuation.go` | Developer | 🧩 In Progress | — |
| 7 | Διόρθωση εισαγωγικών | `pipeline/fixQuotes.go` | AI Agent | 🧩 In Progress | — |
| 8 | Διόρθωση άρθρων (a/an) | `pipeline/fixArticles.go` | AI Agent | ⏳ Pending | — |
| 9 | Ενοποίηση όλων των μετασχηματισμών | `pipeline/applyTransformations.go` | Developer | ⏳ Pending | — |
| 10 | Γράψιμο αποτελέσματος στο αρχείο εξόδου | `pipeline/writeOutput.go` | Developer | ⏳ Pending | — |
| 11 | Δημιουργία tests για κάθε module | `tests/` | Auditor | 🧩 In Progress | — |
| 12 | Δημιουργία documentation αρχιτεκτονικής | `docs/architecture.md` | AI Agent | ✅ Done | 2025-10-26 |
| 13 | Δημιουργία coding standards | `docs/coding_standards.md` | AI Agent | ✅ Done | 2025-10-26 |
| 14 | Δημιουργία εγχειριδίου συνεργασίας | `docs/how_to_work.md` | AI Agent | ✅ Done | 2025-10-26 |
| 15 | Δημιουργία γλωσσαρίου όρων | `docs/glossary.md` | AI Agent | ⏳ Pending | — |

---

## 🧠 Rules για Ανανεώσεις

Κάθε φορά που ένα task αλλάζει κατάσταση:

- Ενημερώνεται η στήλη **Κατάσταση** (`Pending`, `In Progress`, `Done`, `Verified`).
- Προστίθεται **ημερομηνία ενημέρωσης**.
- Αν ένα task εγκριθεί από auditor, σημειώνεται ως **Verified ✅**.

Παράδειγμα ενημέρωσης:

| # | Περιγραφή Εργασίας | Αρχείο / Module | Assigned To | Κατάσταση | Τελευταία Ενημέρωση |
|---|---------------------|------------------|--------------|------------|----------------------|
| 7 | Διόρθωση εισαγωγικών | `pipeline/fixQuotes.go` | AI Agent | ✅ Verified | 2025-10-28 |

---

## 🔍 Κατάσταση Συνολικά

| Σύνολο Tasks | Ολοκληρωμένα | Σε εξέλιξη | Εκκρεμή | Ποσοστό Ολοκλήρωσης |
|---------------|--------------|------------|----------|----------------------|
| 15 | 6 | 3 | 6 | 40% ✅ |

---

## 💬 Notes

- Όλες οι αλλαγές πρέπει να συνοδεύονται από test case.
- Αν αλλάξει το data flow → ενημέρωση `architecture.md`.
- Ο auditor ενημερώνει καθημερινά το blueprint-index για πρόοδο.

---
