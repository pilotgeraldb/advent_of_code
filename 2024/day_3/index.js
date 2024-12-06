/* 
  Mull It Over
  https://adventofcode.com/2024/day/3
*/

const fs = require('fs');
const { performance } = require('perf_hooks');

const data = fs.readFileSync('input.txt', 'utf8');

const start = performance.now();

// write solution code here

const end = performance.now();

console.log(`result:`, 0, `time: ${end - start} ms`)