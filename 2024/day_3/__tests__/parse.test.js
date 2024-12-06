const { parse } = require('../parse')

test.each([
  ["xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", 161],
])('parse', (str, expected) => {
  result = parse(str)
  expect(result).toBe(expected)
});