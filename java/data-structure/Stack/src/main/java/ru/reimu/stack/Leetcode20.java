package ru.reimu.stack;

import java.util.Stack;

/**
 * @Author: Tomonori
 * @Date: 2020/2/17 10:54
 * @Title:
 * @Desc: ↓ ↓ ↓ ↓ ↓
 * -----
 */
public class Leetcode20 {

    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<>();
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (c == '(' || c == '[' || c == '{') {
                stack.push(c);
            } else {
                if (stack.empty()) {
                    return false;
                }

                char topItem = stack.pop();

                if (c == ')' && topItem != '(') {
                    return false;
                }

                if (c == ']' && topItem != '[') {
                    return false;
                }

                if (c == '}' && topItem != '{') {
                    return false;
                }
            }
        }

        return stack.isEmpty();
    }
}
