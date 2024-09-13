#include <bits/stdc++.h>

using namespace std;

struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode() : val(0), left(nullptr), right(nullptr) {}
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
  TreeNode(int x, TreeNode *left, TreeNode *right)
      : val(x), left(left), right(right) {}
};

class Solution {
public:
  int maxPathSum(TreeNode *root) {
    int res = INT_MIN;
    helper(root, res);
    return res;
  }

  int helper(TreeNode *node, int &res) {
    if (node == nullptr) {
      return 0;
    }
    int left = max(0, helper(node->left, res));
    int right = max(0, helper(node->right, res));

    res = max(res, left + right + node->val);

    return node->val + max(left, right);
  }
};
