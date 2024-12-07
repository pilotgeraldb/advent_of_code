const { search } = require('../search')

test.each([
  [`....XXMAS.\r\n.SAMXMS...\r\n...S..A...\r\n..A.A.MS.X\r\nXMASAMX.MM\r\nX.....XA.A\r\nS.S.S.S.SS\r\n.A.A.A.A.A\r\n..M.M.M.MM\r\n.X.X.XMASX`, 18],
])('parse', (str, expected) => {
  result = search(str, 10, 10)
  expect(result).toBe(expected)
});

//horizontal overlap
test.each([
  [`XMASAMX`, 2],
])('parse', (str, expected) => {
  result = search(str, 7, 1)
  expect(result).toBe(expected)
});

//vertical overlap
test.each([
  [`...X\r\n...M\r\n...A\r\n...S\r\n...A\r\n...M\r\n...X`, 2],
])('parse', (str, expected) => {
  result = search(str, 4, 7)
  expect(result).toBe(expected)
});

//E
test.each([
  [`XMAS`, 1],
])('parse', (str, expected) => {
  result = search(str, 4, 1)
  expect(result).toBe(expected)
});

//W
test.each([
  [`SAMX`, 1],
])('parse', (str, expected) => {
  result = search(str, 4, 1)
  expect(result).toBe(expected)
});

//NW
test.each([
  [`S...\r\n.A..\r\n..M.\r\n...X`, 1],
])('parse', (str, expected) => {
  result = search(str, 4, 4)
  expect(result).toBe(expected)
});

//SE
test.each([
  [`X...\r\n.M..\r\n..A.\r\n...S`, 1],
])('parse', (str, expected) => {
  result = search(str, 4, 4)
  expect(result).toBe(expected)
});

//SW
test.each([
  [`...X\r\n..M.\r\n.A..\r\nS...`, 1],
])('parse', (str, expected) => {
  result = search(str, 4, 4)
  expect(result).toBe(expected)
});

//S
test.each([
  [`X\r\nM\r\nA\r\nS`, 1],
])('parse', (str, expected) => {
  result = search(str, 1, 4)
  expect(result).toBe(expected)
});

//N
test.each([
  [`S\r\nA\r\nM\r\nX`, 1],
])('parse', (str, expected) => {
  result = search(str, 1, 4)
  expect(result).toBe(expected)
});

//NE
test.each([
  [`...S\r\n..A.\r\n.M..\r\nX...`, 1],
])('parse', (str, expected) => {
  result = search(str, 4, 4)
  expect(result).toBe(expected)
});