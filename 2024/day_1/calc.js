const distance = (a, b) => {
    if (a >= b)
        return a - b
    else
        return b - a
}

function radixSort(arr) {
    const getDigit = (num, place) => Math.floor(Math.abs(num) / Math.pow(10, place)) % 10;

    const digitCount = (num) => (num === 0 ? 1 : Math.floor(Math.log10(Math.abs(num))) + 1);

    const mostDigits = (nums) => nums.reduce((maxDigits, num) => Math.max(maxDigits, digitCount(num)), 0);

    const maxDigitCount = mostDigits(arr);

    for (let k = 0; k < maxDigitCount; k++) {
        const buckets = Array.from({ length: 10 }, () => []);

        for (let i = 0; i < arr.length; i++) {
            const digit = getDigit(arr[i], k);
            buckets[digit].push(arr[i]);
        }

        arr = [].concat(...buckets);
    }

    return arr;
}

const calculate = (l, r) => {
    let numberOfListItems = l.length

    const oL = radixSort(l)
    const oR = radixSort(r)
    
    d = 0
    for (i = 0; i < numberOfListItems; i++){
        d += distance(oL[i], oR[i])
    }

    return d
}

module.exports = { calculate }