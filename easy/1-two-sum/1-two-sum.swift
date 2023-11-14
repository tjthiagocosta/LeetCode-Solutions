class Solution {
    func twoSum(_ nums: [Int], _ target: Int) -> [Int] {
        var dict = [Int: Int]()
        
        for (index, value) in nums.enumerated() {
            if let addent = dict[value] {
                return [addent, index]
            } else {
                dict[target - value] = index
            }
        }
        
        return []
    }
}