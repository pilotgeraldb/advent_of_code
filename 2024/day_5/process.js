const isLetter = (char) => {
    const code = char.charCodeAt(0);
    return (code >= 65 && code <= 90) || (code >= 97 && code <= 122);
};

const isDigit = (char) => {
    const code = char.charCodeAt(0);
    return code >= 48 && code <= 57;
};

const process = (str) => {
    let buf = ""
    let isRule = false
    let isData = false
    const rules = {}
    let lRule = -1

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
                console.log(buf)
                continue;
            } else if (c === ','){

            }
        }
        buf += c;
    }
}

module.exports = { process };