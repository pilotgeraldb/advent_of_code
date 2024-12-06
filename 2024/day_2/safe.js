
const distance = (a, b) => {
  a = Number(a)
  b = Number(b)
  if (a >= b)
      return a - b
  else
      return b - a
}

// determines if the levels are either all increasing or all decreasing.
const checkSafe = (arr) => {
  /*
    1. The levels are either all increasing or all decreasing.
    2. Any two adjacent levels differ by at least one and at most three.
  */
  if (!arr || arr.length <= 1) {
    throw Error("invalid array or array length")
  }
  
  //-1 = undetermined; 0 = decreasing; 1 = increasing
  let increasing = -1 

  for (i = 1; i < arr.length; i++) {
    let l = Number(arr[i-1])
    let r = Number(arr[i])

    if (l == r){
      return false
    }

    if(l < r){
      if (increasing == 0) {
        return false
      }
      else if (increasing === -1)
        increasing = 1
    }
    else if(l > r){
      if (increasing == 1) {
        return false
      }
      else if (increasing === -1)
        increasing = 0
    }

    let d = distance(l, r)

    if (d < 0) {
      throw Error(`d should not be less than 0: d = ${d} arr:${arr} l|r: ${l}|${r}` )
    }

    if(d > 3) {
      return false
    }
  }

  return true
}

module.exports = { checkSafe }