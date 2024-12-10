// Radix Sort is a special sorting algorithm that works on lists of numbers. 
// It never makes comparisons between elements! It exploits the fact that information 
// about the size of a number is encoded in the number of digits. 
// More digits means a bigger number!
const radixSort = (arr) => {
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

module.exports = { radixSort };