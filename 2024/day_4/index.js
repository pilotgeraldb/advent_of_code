/* 
  Ceres Search
  https://adventofcode.com/2024/day/4
*/

const fs = require('fs');
const { performance } = require('perf_hooks');
const { search } = require('./search')

//const data = fs.readFileSync('input.txt', 'utf8');
const data = fs.readFileSync('test.txt', 'utf8');

const start = performance.now();

//var result = search(data, 140, 140)
var result = search(data, 10, 10)

const end = performance.now();

console.log(`result:`, result, `time: ${end - start} ms`)