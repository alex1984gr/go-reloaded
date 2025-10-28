# 🧩 System Architecture

## 🎯 Σκοπός
Αυτό το έγγραφο περιγράφει τη συνολική αρχιτεκτονική του project και τον τρόπο που αλληλεπιδρούν μεταξύ τους τα επιμέρους components.  
Στόχος είναι να παρέχει στον **developer** και στον **AI Agent** μια ξεκάθαρη εικόνα της ροής δεδομένων, των σημείων ελέγχου και των κανόνων συντονισμού.

---

## 🏗️ Βασική Δομή

Το σύστημα αποτελείται από τέσσερα κύρια επίπεδα:

1. **Input Layer** — δέχεται εντολές και δεδομένα (tasks, αρχεία, user prompts).  
2. **Processing Layer** — αναλύει, οργανώνει και κατανέμει το έργο στους agents ή modules.  
3. **Execution Layer** — εκτελεί τις λειτουργίες (π.χ. test generation, code implementation).  
4. **Output Layer** — επιστρέφει αποτελέσματα, reports ή commits στον developer.

Κάθε επίπεδο επικοινωνεί μόνο με το αμέσως επόμενο, για να διατηρείται καθαρή ιεραρχία και modular σχεδίαση.

---

## 🔁 Data Flow Overview

```mermaid
flowchart TD
    A[User / Operator] --> B[Agent Interface]
    B --> C[Task Analyzer]
    C --> D[Implementation Engine]
    D --> E[Testing & QA Module]
    E --> F[Output / Reports / Commits]
