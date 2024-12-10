const { isDigit } = require('../../shared/numbers')

const process = (str) => {
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
                rules[lRule] = -1
            }
            isRule = true;
            continue;
        } else if (!isData && c === '\r' || c === '\n'){
            if (isRule){
                let x = Number(buf)
                if (!rules[lRule] || rules[lRule] === -1){
                    rules[lRule] = x
                }
                else if (rules[lRule] !== -1 && x < rules[lRule]){
                    rules[lRule] = x
                }
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
            else if (c === '\r' || c === '\n'){
                //determine validation
                for (let i = 0; i < numbers.length; i++){
                    let n1 = numbers[i]
                    let n2 = numbers[i+1]
                    let rule = rules[n1]
                    
                }
                continue;
            }
        }
        buf += c;
    }
}

module.exports = { process };