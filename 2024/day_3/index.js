/* 
  Mull It Over
  https://adventofcode.com/2024/day/3
*/

const fs = require('fs');
const { performance } = require('perf_hooks');
const { parse } = require('./parse')

const data = fs.readFileSync('input.txt', 'utf8');

const start = performance.now();

let result = parse(data)

const end = performance.now();

console.log(`result:`, result, `time: ${end - start} ms`)