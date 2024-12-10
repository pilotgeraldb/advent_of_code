const { distance } = require('../shared/numbers')
const { radixSort } = require('../shared/sorting')

const calculate = (l, r) => {
    let numberOfListItems = l.length

    const oL = radixSort(l)
    const oR = radixSort(r)

    d = 0
    for (i = 0; i < numberOfListItems; i++){
        d += distance(oL[i], oR[i])
    }

    return d
}

module.exports = { calculate }