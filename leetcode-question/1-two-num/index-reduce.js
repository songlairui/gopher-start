var twoSum = function(nums, target) {
    return nums.reduce((result, current, idx) => {
        if (result) return result
        const prevIdx = nums
            .slice(0, idx)
            .findIndex(item => item + current === target)
        if (prevIdx != -1) {
            result = [prevIdx, idx]
        }
        return result
    }, null)
}

module.exports = twoSum
