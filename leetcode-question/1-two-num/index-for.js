var twoSum = function(nums, target) {
    let i = 0,
        len = nums.length
    for (; i < len; i++) {
        for (let j = 0; j < i; j++) {
            if (nums[j] + nums[i] === target) {
                return [j, i]
            }
        }
    }
}

module.exports = twoSum
