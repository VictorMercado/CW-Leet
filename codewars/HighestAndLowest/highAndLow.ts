function highAndLow(input: string) : string {
    let max: number = -Infinity;
    let min: number = Infinity;
    input.split(" ").forEach((item) => {
        const num = parseInt(item);
        if (num >= max) 
        {
            max = num;
        }
        if (num < min) 
        {
            min = num;
        }
    });

    return `${max} ${min}`;
}

export default highAndLow;