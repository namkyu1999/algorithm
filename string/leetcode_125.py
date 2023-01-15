from collections import deque


class Solution:
    def isPalindrome(self, s: str) -> bool:
        q = deque()
        for word in s:
            if word.isalnum():
                q.append(word.lower())
        while True:
            if len(q) <= 1:
                break
            if q.popleft() != q.pop():
                return False
        return True


sol = Solution()
inputValue = "A man, a plan, a canal: Panama"
print(sol.isPalindrome(inputValue))
