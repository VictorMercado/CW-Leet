#include<iostream>

struct ListNode {
  int val;
  ListNode *next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
      long long int x;
    }
};

int main() {
  ListNode* val3 = new ListNode(3);
  ListNode* val2 = new ListNode(4, val3);
  ListNode* val1 = new ListNode(2, val2);
  
  ListNode* val6 = new ListNode(4);
  ListNode* val5 = new ListNode(6, val6);
  ListNode* val4 = new ListNode(5, val5);

  Solution sol = Solution();
  std::cout << sol.addTwoNumbers(val1, val4) << std::endl;
}