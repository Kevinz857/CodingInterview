package com.example.sxy.leetcode.day1;

import java.util.Stack;

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(x);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.min();
 */
public class MinStack {

    Stack stack1;
    Stack stack2;

    /** initialize your data structure here. */
    public MinStack() {
        stack1 = new Stack();
        stack2 = new Stack();
    }

    //保证stack2中是有序的
    public void push(int x) {
        stack1.push(x);
        if (stack2.empty()) {
            stack2.push(x);
            return;
        }
        if (x <= (int)stack2.peek()) {
            stack2.push(x);
        }
    }

    public void pop() {
        int a = (int)stack1.pop();
        if (a == (int)stack2.peek()) {
            stack2.pop();
        }
    }

    public int top() {
        return (int)stack1.peek();

    }

    public int min() {
        if (stack1.empty() || stack2.empty()) {
            return -1;
        }
        return (int)stack2.peek();
    }
}


