Given the head of a singly linked list and an integer k, split the linked list into k consecutive linked list parts.

The length of each part should be as equal as possible: no two parts should have a size differing by more than one. This may lead to some parts being null.

The parts should be in the order of occurrence in the input list, and parts occurring earlier should always have a size greater than or equal to parts occurring later.

Return an array of the k parts.

Solution:
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    vector<ListNode*> splitListToParts(ListNode* head, int k) {
        int len = 0;
        ListNode* tem = head;
        while(tem != NULL){
            tem = tem->next;
            len++;
        }
        int v = len / k;
        int ex = len % k;
        vector<ListNode*> a(k);
        tem = head;
        int i = 0;
        int kk = k;
        while(kk--){
            int val1 = v;
            if(ex){
                ex--;
                val1++;
            }
            ListNode* node = new ListNode();
            ListNode* list = node;
            while(val1--){
                node->next = tem;
                tem = tem->next;
                node=node->next;
            }
            node->next = NULL;
            a[i++] = list->next;
        }
        return a;
    }
};

