package com.example.sxy.leetcode.day1;

import java.util.Stack;

/**
 * Your CQueue object will be instantiated and called as such:
 * CQueue obj = new CQueue();
 * obj.appendTail(value);
 * int param_2 = obj.deleteHead();
 */

public class CQueue {

    // 添加元素
    Stack stack1;
    // 取出元素
    Stack stack2;

    // 构造器
    public CQueue() {
        stack1 = new Stack();
        stack2 = new Stack();
    }
    // 在队列里添加元素 所有的元素都维护在stack1中
    public void appendTail(int value) {
        // 往stack1的栈顶添加元素
        stack1.push(value);
    }

    public int deleteHead() {
        // 如果stack1是空，则当前无元素
        if (stack1.empty()) {
            return -1;
        }
        // 把stack1的元素放在stack2中，顺序刚好相反
        while (!stack1.empty()) {
            stack2.push(stack1.pop());
        }
        // stack2栈顶即是队头元素
        int head = (int) stack2.pop();
        // 把元素再放回stack1中
        while (!stack2.empty()) {
            stack1.push(stack2.pop());
        }
        return head;
    }


}
