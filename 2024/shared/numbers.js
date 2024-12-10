// distance Returns the absolute value distance between two numbers
const distance = (a, b) => {
    if (a >= b)
        return a - b
    else
        return b - a
}

module.exports = { distance }