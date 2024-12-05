const { checkSortOrder, checkAdjacentDistances } = require('../safe')

/*
7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
*/

test('check sort order', () => {
  arr = []
  result = checkSortOrder(arr)
  expect(result).toBe(true)
});

test('check adjacent distances', () => {
  arr = []
  result = checkAdjacentDistances(arr)
  expect(result).toBe(true)
});