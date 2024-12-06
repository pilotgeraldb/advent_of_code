/* 
  Red-Nosed Reports 
  https://adventofcode.com/2024/day/2
*/

const fs = require('fs');
const { performance } = require('perf_hooks');
const { checkSafe } = require('./safe')

const data = fs.readFileSync('input.txt', 'utf8');

const lines = data.split('\r\n');

console.log(`proccessing ${lines.length} line(s)...`)

const start = performance.now();
let safe = 0
lines.forEach(line => {
    let arr = line.split(" ");
    if(checkSafe(arr))
      safe++
});
const end = performance.now();

console.log(`safe:`, safe, `time: ${end - start} ms`)