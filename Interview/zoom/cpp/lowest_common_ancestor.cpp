#include <unordered_set>
#include <vector>

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

TreeNode* lowestCommonAncestor4(TreeNode* root, std::vector<TreeNode*>& nodes) {
    std::unordered_set<TreeNode*> nodeSet;
    for (TreeNode* node : nodes) {
        nodeSet.insert(node);
    }
    
    std::function<TreeNode*(TreeNode*)> dfs = [&](TreeNode* root) -> TreeNode* {
        if (root == nullptr) {
            return nullptr;
        }
        if (nodeSet.find(root) != nodeSet.end()) {
            return root;
        }

        TreeNode* left = dfs(root->left);
        TreeNode* right = dfs(root->right);

        if (left && right) {
            return root;
        }
        return left ? left : right;
    };

    return dfs(root);
}