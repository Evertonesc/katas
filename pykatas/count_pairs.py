from typing import List


class Solution:
    def count_pairs(self, nums: List[int], target: int) -> int:
        left, right = 0, 1
        result = 0

        while left < right:
            if right > len(nums) - 1:
                left += 1
                right = left + 1

                if right >= len(nums):
                    break

            if nums[left] + nums[right] < target:
                result += 1

            right += 1

        return result


sol = Solution()
nums = [-1, 1, 2, 3, 1]
target = 2
print(sol.count_pairs(nums, target))
