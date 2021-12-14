package com.example.sxy.leetcode.day2;

import java.util.HashMap;
import java.util.Map;
import java.util.Stack;

public class Solution {
    // 06
    public int[] reversePrint(ListNode head) {
        //处理临界情况
        if (head == null) {
            return new int[0];
        }
        // 创建临时实例，后续使用不改变head
        ListNode temp = head;
        Stack<ListNode> stack = new Stack<>();
        while (temp != null) {
            stack.push(temp);
            temp = temp.next;
        }
        int[] rev = new int[stack.size()];
        int i = 0;
        while (!stack.empty()) {
            rev[i++] = stack.pop().val;
        }
        return rev;
    }

    // 24
    public ListNode reverseList(ListNode head) {
        //处理临界情况
        if (head == null) {
            return null;
        }
        // 创建临时实例，后续使用不改变head
        ListNode temp = head;
        Stack<Integer> stack = new Stack<>();
        while (temp != null) {
            stack.push(temp.val);
            temp = temp.next;
        }
        ListNode rev = new ListNode(stack.pop());
        ListNode temp1 = rev;
        while (!stack.empty()) {
            ListNode rev1 = new ListNode(stack.pop());
            temp1.next = rev1;
            temp1 = rev1;
        }
        return rev;
    }

    // 35. 复杂链表的复制
    public Node copyRandomList(Node head) {
        if (head == null) {
            return null;
        }
        Map<Node, Node> map = new HashMap<>();
        Node curNode = head;
        while (curNode != null) {
            map.put(curNode, new Node(curNode.val));
            curNode =  curNode.next;
        }
        curNode = head;
        while (curNode != null) {
            map.get(curNode).next = map.get(curNode.next);
            map.get(curNode).random = map.get(curNode.random);
            curNode = curNode.next;
        }
        return map.get(head);
    }
}
