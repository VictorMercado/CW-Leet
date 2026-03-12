function DNAStrand(dna: string) : string {
    const complements = {
        A: "T",
        T: "A",
        C: "G",
        G: "C"
    }
    return `${(dna.split("").map((item) => { return `${complements[item]}`; }))}`.replace(/,/g, "");
}
