/* 
  Ceres Search
  https://adventofcode.com/2024/day/4
*/

const fs = require('fs');
const { performance } = require('perf_hooks');
const { search } = require('./search')

const data = fs.readFileSync('input.txt', 'utf8');

const start = performance.now();

//change the eol string length based on your OS \r\n vs \n
// windows: 2
// linux/macos: 1
var result = search(data, 140, 140, 1)

const end = performance.now();

console.log(`result:`, result, `time: ${end - start} ms`)