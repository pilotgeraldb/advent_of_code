const { isDigit } = require('../shared/characters')
const { radixSort } = require('../shared/sorting')

const process = (str) => {
    let total = 0
    let buf = ""
    let isRule = false
    let isData = false
    const rules = {}
    let lRule = -1

    let numbers = [];

    const reset = () => {
        buf = ""
        isRule = false
        lRule = -1
    }

    for (let i = 0; i < str.length; i++) {
        let c = str[i];
        if (!isData && isDigit(c)) {
            buf += c;
            continue;
        } else if (!isData && c === '|'){
            lRule = Number(buf)
            buf = ""
            if (!rules[lRule]) {
                rules[lRule] = []
            }
            isRule = true;
            continue;
        } else if (!isData && (c === '\n' || c === '\r\n')){
            if (isRule){
                let x = Number(buf)
                if (!rules[lRule]){
                    rules[lRule] = []
                }
                rules[lRule].push(x)
                reset()
                continue;
            } else {
                isData = true
                continue;
            }
        } else if (isData){
            if(isDigit(c)) {
                buf += c;
                continue;
            } else if (c === ','){
                numbers.push(Number(buf))
                buf = ""
                continue;
            }
            else if (c === '\n' || c === '\r\n'){
                numbers.push(Number(buf))
                let invalid = false
                //console.log(`numbers: ${JSON.stringify(numbers)}`)
                let facts = { }
                let p = []
                for (let i = numbers.length - 1; i >= 0; i--) {
                    let n = numbers[i]
                    p.push(n)
                    if (i === numbers.length - 1) {
                        continue
                    }
                    for (let j = 0; j < p.length; j++) {
                        if(n === p[j]) {
                            continue;
                        }
                        facts[`${n}|${p[j]}`] = true
                    }
                }

                for (let i = 0; i < numbers.length; i++) {
                    let n = numbers[i]
                    let r = rules[n] 
                    if(!r) {
                        continue
                    }   
                    for(let j = 0; j < r.length; j++) {
                        if (numbers.includes(r[j]) && !facts[`${n}|${r[j]}`]) {
                            invalid = true
                            break;
                        }
                    }
                }

                if (!invalid) {
                    let middle = numbers[(numbers.length - 1) / 2]
                    total = total + middle
                }
                buf = ""
                numbers = []
                continue
            }
        }
        buf += c;
    }

    return total;
}

module.exports = { process };