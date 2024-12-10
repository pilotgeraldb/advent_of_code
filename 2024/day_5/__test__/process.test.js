const { process } = require('../process')


test.each([
    [`12|56\r\n12, 56, 1, 34`, true],
  ])('parse', (str, expected) => {
    result = process(str)
    expect(result).toBe(expected)
  });