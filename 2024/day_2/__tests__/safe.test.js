const { checkSafe } = require('../safe')

/*
  1. The levels are either all increasing or all decreasing.
  2. Any two adjacent levels differ by at least one and at most three.
*/
/*
  7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
  1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
  9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
  1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
  8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
  1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
*/

test.each([
  [[7, 6, 4, 2, 1], true],              //Safe because the levels are all decreasing by 1 or 2.
  [[6, 7, 8, 9, 10], true],             //Safe because the levels are all increasing
  [[15, 14, 13, 12, 11], true],         //Safe because the levels are all decreasing
  [[6, 8, 10, 12, 14], true],           //Safe because the levels are all increasing (d = 2)
  [[14, 12, 10, 8, 6], true],           //Safe because the levels are all decreasing (d = 2)
  [[6, 10, 14, 18, 22], false],         //Unsafe increasing, but because d = 4
  [[22, 18, 14, 10, 6], false],         //Unsafe decreasing but because d = 4
  [[1, 2, 7, 8, 9], false],             //Unsafe because 2 7 is an increase of 5.
  [[9, 7, 6, 2, 1], false],             //Unsafe because 6 2 is a decrease of 4.
  [[1, 3, 2, 4, 5], false],             //Unsafe because 1 3 is increasing but 3 2 is decreasing.
  [[8, 6, 4, 4, 1], false],             //Unsafe because 4 4 is neither an increase or a decrease.
  [[1, 3, 6, 7, 9], true],              //Safe because the levels are all increasing by 1, 2, or 3.
  [[12, 9, 7, 4, 1], true],             //Safe because all decreasing
  [[21, 18, 15, 13, 10, 7, 6, 4], true], //Safe all decreasing by 1, 2, or 3 
  [["12", "9", "7", "4", "1"], true],   //Safe because all decreasing (string check)
])('check safe', (arr, expected) => {
  result = checkSafe(arr)
  expect(result).toBe(expected)
});