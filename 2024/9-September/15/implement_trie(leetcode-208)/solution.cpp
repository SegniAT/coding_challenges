#include <bits/stdc++.h>

using namespace std;

class TrieNode {
  bool end;
  vector<TrieNode *> children;

public:
  TrieNode() {
    this->end = false;
    this->children.resize(26, nullptr);
  }
  void setNode(char c) { children[c - 'a'] = new TrieNode(); }
  void setEnd() { this->end = true; }
  TrieNode *getNode(char c) { return children[c - 'a']; }
  bool getEnd() { return this->end; };
};

class Trie {
  TrieNode *root;

public:
  Trie() { this->root = new TrieNode(); }

  void insert(string word) {
    TrieNode *iter = root;
    for (auto c : word) {
      if (iter->getNode(c)) {
        iter = iter->getNode(c);
      } else {
        iter->setNode(c);
        iter = iter->getNode(c);
      }
    }
    iter->setEnd();
  }

  bool search(string word) {
    TrieNode *iter = root;
    for (auto c : word) {
      if (!iter->getNode(c)) {
        return false;
      }
      iter = iter->getNode(c);
    }
    return iter->getEnd();
  }

  bool startsWith(string prefix) {
    TrieNode *iter = root;
    for (auto c : prefix) {
      if (!iter->getNode(c)) {
        return false;
      }
      iter = iter->getNode(c);
    }
    return true;
  }
};

/**
 * Your Trie object will be instantiated and called as such:
 * Trie* obj = new Trie();
 * obj->insert(word);
 * bool param_2 = obj->search(word);
 * bool param_3 = obj->startsWith(prefix);
 */
