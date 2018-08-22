var twoSum = function(nums, target) {
    let i = 0,
        len = nums.length,
        pool = {}
    for (; i < len; i++) {
        const candy = target - nums[i]
        if (pool.hasOwnProperty(candy)) {
            return [pool[candy], i]
        } else {
            pool[nums[i]] = i
        }
    }
}

module.exports = twoSum
