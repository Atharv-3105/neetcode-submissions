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
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        ListNode* tmp1 = list1;
        ListNode* tmp2 = list2;
        ListNode* dummy = new ListNode();
        ListNode* result = dummy;

        while(tmp1 && tmp2)
        {
            if(tmp1->val <= tmp2->val){
                result->next = tmp1;
                tmp1 = tmp1->next;
                result = result->next;
            }
            else{
                result->next = tmp2;
                tmp2 = tmp2->next;
                result = result->next;
            }
        }
        if(tmp1)
        {
            result->next = tmp1;
        }
        if(tmp2)
        {
            result->next = tmp2;
        }

        return dummy->next;
    }
};
