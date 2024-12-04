/* 
  historian_hysteria 
  https://adventofcode.com/2024/day/1
*/
const { calculate } = require("./calc")
const fs = require('fs');
const { performance } = require('perf_hooks');

const l = []
const r = []

const data = fs.readFileSync('input.txt', 'utf8');
    
const lines = data.split('\n');

lines.forEach(line => {
    const [num1, num2] = line.split(/\s+/);
    if (num1 && num2) {
    l.push(Number(num1));
    r.push(Number(num2));
    }
});

const start = performance.now();
const result = calculate(l, r);
const end = performance.now();
    
console.log(`result:`, result, `time: ${end - start} ms`);