/* 
  Red-Nosed Reports 
  https://adventofcode.com/2024/day/2
*/

const fs = require('fs');
const { performance } = require('perf_hooks');
const { checkSortOrder, checkAdjacentDistances } = require('./safe')

const data = fs.readFileSync('input.txt', 'utf8');

const lines = data.split('\n');

const start = performance.now();
let safe = 0
lines.forEach(line => {
    // do something
    arr = []

    if(checkAdjacentDistances(arr) && checkSortOrder(arr))
      safe++
});
const end = performance.now();

console.log(`safe:`, safe, `time: ${end - start} ms`)