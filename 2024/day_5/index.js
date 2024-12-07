/* 
  Print Queue
  https://adventofcode.com/2024/day/5
*/

const fs = require('fs');
const { performance } = require('perf_hooks');
const { process } = require('./process');

const data = fs.readFileSync('test.txt', 'utf8');

const start = performance.now();

const result = process(data);

const end = performance.now();

console.log(`result:`, result, `time: ${end - start} ms`)