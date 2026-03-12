const romanToInt = (s: string): number => {
    let sum: number = 0;
    const romanNums: Record<string, number> = {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000
    };
    for (let i = 0; i < s.length; i++) {
        //item is the key of the romanNums object
        if (i < s.length - 1 && romanNums[s[i]] < romanNums[s[i + 1]]) {
            sum += romanNums[s[i + 1]] - romanNums[s[i]];
            i++;
        }
        else {
            sum += romanNums[s[i]];
        }
    }

    return sum;
};

export default romanToInt;