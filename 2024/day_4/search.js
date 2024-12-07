const DIR = {
  N: 0,
  NE: 45,
  E: 90,
  SE: 135,
  S: 180,
  SW: 225,
  W: 270,
  NW: 315,
}

const search = (str, xSize, ySize, eolStrLen) => {
  let total = 0
  let found = {}

  if(!eolStrLen){
    eolStrLen = 2
  }

  // nextWordIdx is the zero-based index of the next char in XMAS;
  // where X = 0, M = 1, A = 2, S = 3
  const getChar = (dir, pos, nextWordIdx) => {
    //add 2 for the newline at the end of a row
    let rowOffset = ((xSize + eolStrLen) * nextWordIdx)
    let rowIdx = pos % (xSize + eolStrLen)
    let rowNum = (pos <= xSize) ? 0 : Math.ceil(xSize / pos)
    let idx = -1;
    let charInfo = {
      idx: -1,
      val: ''
    }

    if(dir === DIR.SE){
      if(rowIdx > (xSize - 3)) {
        return charInfo
      }
      idx = pos + rowOffset + nextWordIdx
    }
    else if (dir === DIR.S) {
      if(rowNum > (ySize - 3)) {
        return charInfo
      }
      idx = pos + rowOffset
    }
    else if (dir === DIR.SW) {
      if(rowIdx < 3) {
        return charInfo
      }
      idx = pos + rowOffset - nextWordIdx
    }
    else if (dir === DIR.W) {
      if(rowIdx < 3) {
        return charInfo
      }
      idx = pos - nextWordIdx
    }
    else if (dir === DIR.NW) {
      if(rowIdx < 3) {
        return charInfo
      }
      idx = pos - rowOffset - nextWordIdx
    }
    else if (dir === DIR.N) {
      if(rowIdx < 3) {
        return charInfo
      }
      idx = pos - rowOffset
    }
    else if (dir === DIR.NE) {
      if(rowIdx > (xSize - 3)) {
        return charInfo
      }
      idx = pos - rowOffset + nextWordIdx
    }
    else if (dir === DIR.E) {
      if(rowIdx > (xSize - 3)) {
        return charInfo
      }
      idx = pos + nextWordIdx
    }

    if(idx >= 0){
      let r = str[idx]
      if (r === "\n" || r === "\r")
        return charInfo
      else
      {
        charInfo.idx = idx
        charInfo.val = r
        return charInfo
      }
    } else {
      return charInfo
    }
  }

  const check = (pos) => {
    let t = 0
    let u = null
    //SE
    let buf = str[pos]
    buf += getChar(DIR.SE, pos, 1).val
    buf += getChar(DIR.SE, pos, 2).val
    u = getChar(DIR.SE, pos, 3)
    buf += u.val
    if (buf == "XMAS"){
      if(!found[`${pos}:${u.idx}`] && !found[`${u.idx}:${pos}`]){
        found[`${pos}:${u.idx}`] = true
        found[`${u.idx}:${pos}`] = true
        t++
      }
    }
    if (buf == "SAMX"){
      if(!found[`${pos}:${u.idx}`] && !found[`${u.idx}:${pos}`]){
        found[`${pos}:${u.idx}`] = true
        found[`${u.idx}:${pos}`] = true
        t++
      }
    }
    //S
    buf = str[pos]
    buf += getChar(DIR.S, pos, 1).val
    buf += getChar(DIR.S, pos, 2).val
    u = getChar(DIR.S, pos, 3)
    buf += u.val
    if (buf == "XMAS"){
      if(!found[`${pos}:${u.idx}`] && !found[`${u.idx}:${pos}`]){
        found[`${pos}:${u.idx}`] = true
        found[`${u.idx}:${pos}`] = true
        t++
      }
    }
    if (buf == "SAMX"){
      if(!found[`${pos}:${u.idx}`] && !found[`${u.idx}:${pos}`]){
        found[`${pos}:${u.idx}`] = true
        found[`${u.idx}:${pos}`] = true
        t++
      }
    }
    //SW
    buf = str[pos]
    buf += getChar(DIR.SW, pos, 1).val
    buf += getChar(DIR.SW, pos, 2).val
    u = getChar(DIR.SW, pos, 3)
    buf += u.val
    if (buf == "XMAS"){
      if(!found[`${pos}:${u.idx}`] && !found[`${u.idx}:${pos}`]){
        found[`${pos}:${u.idx}`] = true
        found[`${u.idx}:${pos}`] = true
        t++
      }
    }
    if (buf == "SAMX"){
      if(!found[`${pos}:${u.idx}`] && !found[`${u.idx}:${pos}`]){
        found[`${pos}:${u.idx}`] = true
        found[`${u.idx}:${pos}`] = true
        t++
      }
    }
    //W
    buf = str[pos]
    buf += getChar(DIR.W, pos, 1).val
    buf += getChar(DIR.W, pos, 2).val
    u = getChar(DIR.W, pos, 3)
    buf += u.val
    if (buf == "XMAS"){
      if(!found[`${pos}:${u.idx}`] && !found[`${u.idx}:${pos}`]){
        found[`${pos}:${u.idx}`] = true
        found[`${u.idx}:${pos}`] = true
        t++
      }
    }
    if (buf == "SAMX"){
      if(!found[`${pos}:${u.idx}`] && !found[`${u.idx}:${pos}`]){
        found[`${pos}:${u.idx}`] = true
        found[`${u.idx}:${pos}`] = true
        t++
      }
    }
    //NW
    buf = str[pos]
    buf += getChar(DIR.NW, pos, 1).val
    buf += getChar(DIR.NW, pos, 2).val
    u = getChar(DIR.NW, pos, 3)
    buf += u.val
    if (buf == "XMAS"){
      if(!found[`${pos}:${u.idx}`] && !found[`${u.idx}:${pos}`]){
        found[`${pos}:${u.idx}`] = true
        found[`${u.idx}:${pos}`] = true
        t++
      }
    }
    if (buf == "SAMX"){
      if(!found[`${pos}:${u.idx}`] && !found[`${u.idx}:${pos}`]){
        found[`${pos}:${u.idx}`] = true
        found[`${u.idx}:${pos}`] = true
        t++
      }
    }
    //N
    buf = str[pos]
    buf += getChar(DIR.N, pos, 1).val
    buf += getChar(DIR.N, pos, 2).val
    u = getChar(DIR.N, pos, 3)
    buf += u.val
    if (buf == "XMAS"){
      if(!found[`${pos}:${u.idx}`] && !found[`${u.idx}:${pos}`]){
        found[`${pos}:${u.idx}`] = true
        found[`${u.idx}:${pos}`] = true
        t++
      }
    }
    if (buf == "SAMX"){
      if(!found[`${pos}:${u.idx}`] && !found[`${u.idx}:${pos}`]){
        found[`${pos}:${u.idx}`] = true
        found[`${u.idx}:${pos}`] = true
        t++
      }
    }
    //NE
    buf = str[pos]
    buf += getChar(DIR.NE, pos, 1).val
    buf += getChar(DIR.NE, pos, 2).val
    u = getChar(DIR.NE, pos, 3)
    buf += u.val
    if (buf == "XMAS"){
      if(!found[`${pos}:${u.idx}`] && !found[`${u.idx}:${pos}`]){
        found[`${pos}:${u.idx}`] = true
        found[`${u.idx}:${pos}`] = true
        t++
      }
    }
    if (buf == "SAMX"){
      if(!found[`${pos}:${u.idx}`] && !found[`${u.idx}:${pos}`]){
        found[`${pos}:${u.idx}`] = true
        found[`${u.idx}:${pos}`] = true
        t++
      }
    }
    //E
    buf = str[pos]
    buf += getChar(DIR.E, pos, 1).val
    buf += getChar(DIR.E, pos, 2).val
    u = getChar(DIR.E, pos, 3)
    buf += u.val
    if (buf == "XMAS"){
      if(!found[`${pos}:${u.idx}`] && !found[`${u.idx}:${pos}`]){
        found[`${pos}:${u.idx}`] = true
        found[`${u.idx}:${pos}`] = true
        t++
      }
    }
    if (buf == "SAMX"){
      if(!found[`${pos}:${u.idx}`] && !found[`${u.idx}:${pos}`]){
        found[`${pos}:${u.idx}`] = true
        found[`${u.idx}:${pos}`] = true
        t++
      }
    }

    return t
  }

  for(let x = 0; x < str.length; x++){
      let c = str[x]

      if (c === 'X')
      {
        total = total + check(x)
      }
      if (c === 'S')
      {
        total = total + check(x)
      }
  }

  return total
}

module.exports = { search }