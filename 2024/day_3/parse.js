const isLetter = (char) => {
  const code = char.charCodeAt(0);
  return (code >= 65 && code <= 90) || (code >= 97 && code <= 122);
};

const isDigit = (char) => {
  const code = char.charCodeAt(0);
  return code >= 48 && code <= 57;
};

const parse = (str) => {
  let total = 0
  let isMul = false
  let buf = ""
  let l = ""
  let r = ""
  let foundL = false
  
  let reset = () => {
    buf = ""
    isMul = false
    l = ""
    r = ""
    foundL = false
  }

  for (let i = 0; i < str.length; i++) {
    let c = str[i]
    buf += c
    if(buf.endsWith("mul")){
      while(c !== ')')
      {
        i++
        c = str[i]

        if(c === '('){
          continue
        }
        if(c === ','){
          foundL = true
          continue
        }
        if(c === ')'){
          total += (Number(l) * Number(r))
          reset()
          break
        }
        if (!isDigit(c)){
          reset()
          break
        }
        if (!foundL){
          l += c
        }else {
          r += c
        }
      }
    }
  }
  return total
}

module.exports = { parse }