#include <algorithm>
#include <bits/stdc++.h>
#include <deque>
#include <unordered_map>
#include <vector>

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

/*
 * HOW???
 * Time complexity: O(n^2) ???
 * Space complexity: O(n^2) ???
 *
class Solution {
public:
  TreeNode *buildTree(vector<int> &preorder, vector<int> &inorder) {
    deque<int> preorderQ(preorder.begin(), preorder.end());
    return build(preorderQ, inorder);
  }
  TreeNode *build(deque<int> &preorderQ, vector<int> &inorder) {
    if (preorderQ.empty() || inorder.empty()) {
      return nullptr;
    }

    int val = preorderQ.front();
    preorderQ.pop_front();
    auto it = find(inorder.begin(), inorder.end(), val);
    int ind = it - inorder.begin();

    TreeNode *root = new TreeNode(val);
    vector<int> lInorder(inorder.begin(), inorder.begin() + ind);
    vector<int> rInorder(inorder.begin() + ind + 1, inorder.end());

    root->left = build(preorderQ, lInorder);
    root->right = build(preorderQ, rInorder);

    return root;
  }


};
*/

// better solution
class Solution {
  int preorderInd;
  unordered_map<int, int> inorderIndMap;

public:
  TreeNode *buildTree(vector<int> &preorder, vector<int> &inorder) {
    inorderIndMap.clear();
    for (int i = 0; i < inorder.size(); i++) {
      inorderIndMap[inorder[i]] = i;
    }

    return build(preorder, 0, inorder.size() - 1);
  }

  TreeNode *build(vector<int> &preorder, int s, int e) {
    if (s > e) {
      return nullptr;
    }

    int val = preorder[preorderInd++];
    TreeNode *root = new TreeNode(val);
    int mid = inorderIndMap[val];
    root->left = build(preorder, s, mid - 1);
    root->right = build(preorder, mid + 1, e);
    return root;
  }
};
