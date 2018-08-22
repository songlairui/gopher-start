const twoSumF = require('./index-for')
const twoSumR = require('./index-reduce')
const twoSumA = require('./index-advance')

const nums = [99, 23, 2, 7, 11, 15],
    target = 9

console.warn(twoSumF(nums, 9))
console.warn(twoSumR(nums, 9))
console.warn(twoSumA(nums, 9))
