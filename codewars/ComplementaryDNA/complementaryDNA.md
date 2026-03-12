https://www.codewars.com/kata/554e4a2f232cdd87d9000038
# Codewars: Complementary DNA

## Problem Description
Deoxyribonucleic acid (DNA) is a chemical found in the nucleus of cells and carries the "instructions" for the development and functioning of living organisms.

In DNA strings, symbols "A" and "T" are complements of each other, as "C" and "G". You receive one side of the DNA; you need to return the other complementary side.

### Complement Rules
* **A** ↔ **T**
* **C** ↔ **G**

---

## Examples (Input → Output)
* `"ATTGC"` → `"TAACG"`
* `"GTAT"` → `"CATA"`
* `"AAAA"` → `"TTTT"`
* `"CGCG"` → `"GCGC"`

---

## Constraints
* The DNA strand is never empty (except for specific language-specific edge cases like Haskell).
* Input will consist of capital letters `A`, `T`, `C`, and `G`.