const { search } = require('../search')

test.each([
  [`MMMSXXMASM
  MSAMXMSMSA
  AMXSXMAAMM
  MSAMASMSMX
  XMASAMXAMM
  XXAMMXXAMA
  SMSMSASXSS
  SAXAMASAAA
  MAMMMXMMMM
  MXMXAXMASX`, 18],
])('parse', (str, expected) => {
  result = search(str)
  expect(result).toBe(expected)
});