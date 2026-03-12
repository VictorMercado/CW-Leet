#include <iostream>
#include <vector>
#include <chrono>
#include <cstdlib>

class Solution {
public:
  bool isPalindrome(long x) {
    if(x < 0) {
      return false;
    }
    long temp = x;
    long temp2 = x;
    int length = 0;
    while(temp != 0) {
      length++;
      temp = temp / 10;
    }
    int arr[length];
    int index = 0;
    while(temp2 != 0) {
      arr[(length - 1) - index] = temp2 % 10;
      temp2 = temp2 / 10;
      index++;
    }
    long val = x;
    bool isEvenDigits = length % 2; 
    for(int i=0; i<length; i++) {
      int curr = val % 10;
      val = val / 10;
      if (arr[i] == curr) {
        continue;
      } else {
        return false;
      }
    }
    return true;
  }
  bool isPalindrome2(long x) {
    if (x < 0) return false; // Negative numbers are not palindromes

    std::vector<int> digits; // Use vector instead of VLA
    long temp = x;

    // Extract digits and store them in a vector
    while (temp != 0) {
        digits.push_back(temp % 10);
        temp /= 10;
    }

    // Compare digits from front and back
    int n = digits.size();
    for (int i = 0; i < n / 2; ++i) {
        if (digits[i] != digits[n - i - 1]) {
            return false;
        }
    }
    return true;
  }
  bool isPalindrome3(long x) {
    if (x < 0) return false; // Negative numbers are not palindromes
    if (x != 0 && x % 10 == 0) return false; // Numbers ending with 0 (except 0) are not palindromes

    long reversed = 0;
    long original = x;

    // Reverse the number
    while (x > 0) {
        reversed = reversed * 10 + x % 10;
        x /= 10;
    }

    // Check if the original number is equal to its reversed version
    return original == reversed;
  }
};

int main() {
  // 4294967296
  long num = 64295559246;
  Solution sol = Solution();
  auto start = std::chrono::high_resolution_clock::now();
  std::cout << (sol.isPalindrome(num) ? "true": "false") << std::endl;
  auto end = std::chrono::high_resolution_clock::now();
  auto duration = std::chrono::duration_cast<std::chrono::microseconds>(end - start);
  std::cout << "Execution time: " << duration.count() << " microseconds" << std::endl;

  start = std::chrono::high_resolution_clock::now();
  std::cout << (sol.isPalindrome2(num) ? "true": "false") << std::endl;
  end = std::chrono::high_resolution_clock::now();
  duration = std::chrono::duration_cast<std::chrono::microseconds>(end - start);
  std::cout << "Execution time: " << duration.count() << " microseconds" << std::endl;

  start = std::chrono::high_resolution_clock::now();
  std::cout << (sol.isPalindrome3(num) ? "true": "false") << std::endl;
  end = std::chrono::high_resolution_clock::now();
  duration = std::chrono::duration_cast<std::chrono::microseconds>(end - start);
  std::cout << "Execution time: " << duration.count() << " microseconds" << std::endl;


  exit(0);
}