const { calculate } = require('../calc');

test('finds total distance', () => {
  let l = [3,4,2,1,3,3]
  let r = [4,3,5,3,9,3]
  let result = calculate(l, r)
  expect(result).toBe(11);
});