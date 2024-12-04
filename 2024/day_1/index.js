/* 
  historian_hysteria 
  https://adventofcode.com/2024/day/1
*/
const { calculate } = require("./calc")
const { performance } = require('perf_hooks');

const numberOfListItems = 500000

const list = (length, min, max) => {
    return Array.from({ length }, () => 
        Math.floor(Math.random() * (max - min + 1)) + min
      );
}

const shuffle = (arr) => {
    for (let i = arr.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [arr[i], arr[j]] = [arr[j], arr[i]];
    }
    return arr;
}

//generate two random unordered lists
const l = shuffle(list(numberOfListItems, 1, 99))
const r = shuffle(list(numberOfListItems, 1, 99))

const start = performance.now();
const result = calculate(l, r);
const end = performance.now();
    
console.log(`result:`, d, `time: ${end - start} ms`);