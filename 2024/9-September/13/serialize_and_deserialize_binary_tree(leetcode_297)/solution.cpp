#include <bits/stdc++.h>
#include <string>

using namespace std;

struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Codec {
  char NULL_NODE = 'N';
  char SEPARATOR = ',';

public:
  void serializeHelper(TreeNode *node, string &s) {
    if (node == nullptr) {
      s += NULL_NODE;
      s += SEPARATOR;
      return;
    }

    s += to_string(node->val);
    s += SEPARATOR;
    serializeHelper(node->left, s);
    serializeHelper(node->right, s);
  }

  // Encodes a tree to a single string.
  string serialize(TreeNode *root) {
    string s;
    serializeHelper(root, s);
    if (!s.empty() && s.back() == SEPARATOR) {
      s.pop_back();
    }
    return s;
  }

  TreeNode *deserializeHelper(deque<string> &nodes) {
    string front = nodes.front();
    nodes.pop_front();

    string t = "";
    t += NULL_NODE;
    if (front == t) {
      return nullptr;
    }

    TreeNode *root = new TreeNode(stoi(front));
    root->left = deserializeHelper(nodes);
    root->right = deserializeHelper(nodes);

    return root;
  }

  // Decodes your encoded data to tree.
  TreeNode *deserialize(string data) {
    deque<string> nodes;
    stringstream ss(data);
    string node;
    while (getline(ss, node, SEPARATOR)) {
      nodes.push_back(node);
    }

    return deserializeHelper(nodes);
  }
};

// Your Codec object will be instantiated and called as such:
// Codec ser, deser;
// TreeNode* ans = deser.deserialize(ser.serialize(root));
