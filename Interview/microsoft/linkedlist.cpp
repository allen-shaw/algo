#include <stdio.h>
#include <unordered_map>

class Node {
public:
  int val_;
  Node *next;
  Node *random;

  Node(int val) : val_(val), next(nullptr), random(nullptr) {}
};

Node *deepcopy(Node *head) {
  if (head == nullptr) {
    return nullptr;
  }

  std::unordered_map<Node *, Node *> nodeMap;
  Node *curr = head;
  while (curr != nullptr) {
    nodeMap[curr] = new Node(curr->val_);
    curr = curr->next;
  }

  // 1 -> 2 -> 3
  // map 1->1' 2->2' 3->3' 
  //   1'-> 2'

  // 1 -> 1' -> 2 -> 2' -> 3 -> 3'
  // curr -> next -> random = curr -> random -> next

  curr = head;
  while (curr != nullptr) {
    nodeMap[curr]->next = nodeMap[curr->next];
    nodeMap[curr]->random = nodeMap[curr->random];
    curr = curr->next;
  }

  return nodeMap[head];
}

// nullptr, 1 node
// random = nullptr
// random -> self
// random -> same_node
// 超长的输入