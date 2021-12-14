/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
    pre:=new(ListNode)
    //反转链表
    for head!=nil{
        pre,head,head.Next=head,head.Next,pre
    }
    //遍历链表
    nums:=make([]int,0)//申请一个数组
    for pre.Next!=nil{
        nums=append(nums[:],pre.Val)
        pre=pre.Next
    }
    return nums
}

