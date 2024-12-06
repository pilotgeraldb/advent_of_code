/* 
  Ceres Search
  https://adventofcode.com/2024/day/4
*/

const fs = require('fs');
const { performance } = require('perf_hooks');

const data = fs.readFileSync('input.txt', 'utf8');

const start = performance.now();

/* write solution code here */

const end = performance.now();

console.log(`result:`, result, `time: ${end - start} ms`)