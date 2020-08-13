# -*- coding: utf-8 -*-  
"""
Create on 04-26 11:05 2020
@Author ywx 
@File 1.py
"""


class Solution:
    def twoSum(self, nums: list, target: int) -> list:
        d = {}
        t = 0
        for i in range(0,len(nums)):
            t = target - nums[i]
            if t in d:
                return [d[t],i]
            d[nums[i]] = t
