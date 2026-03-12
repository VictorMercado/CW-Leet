#include <iostream>
#include <vector>
#include <chrono>

class Solution {
  public:
    bool isPalindrome(long x ) {
      if (x < 0) return false;
      if (x != 0 && x % 10 == 0 ) return false;
      
      long reverse = 0;
      long value = x;
      
      while (value != 0) {
        reverse *= 10;
        reverse += value % 10;
        value /= 10;
      }
      return reverse == x;
    } 
};

int main() {
  Solution sol = Solution();
  std::cout << (sol.isPalindrome(0) ? "true" : "false") << std::endl;
}