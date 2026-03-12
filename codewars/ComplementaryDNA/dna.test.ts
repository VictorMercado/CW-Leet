import { expect, test, describe } from "bun:test";

// The function to test
function DNAStrand(dna: string): string {
  const complements: Record<string, string> = {
    A: "T",
    T: "A",
    C: "G",
    G: "C"
  };
  return `${(dna.split("").map((item) => { return `${complements[item]}`; }))}`.replace(/,/g, "");
}

describe("DNA Strand Complements", () => {
  test("Basic strands", () => {
    expect(DNAStrand("AAAA")).toBe("TTTT");
    expect(DNAStrand("ATTGC")).toBe("TAACG");
    expect(DNAStrand("GTAT")).toBe("CATA");
  });

  test("Longer sequences", () => {
    expect(DNAStrand("CGCG")).toBe("GCGC");
    expect(DNAStrand("ATTGCGCG")).toBe("TAACGCGC");
  });

  test("Single character strands", () => {
    expect(DNAStrand("A")).toBe("T");
    expect(DNAStrand("G")).toBe("C");
  });
});