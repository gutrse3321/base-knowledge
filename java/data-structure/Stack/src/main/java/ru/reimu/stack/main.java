package ru.reimu.stack;

/**
 * @Author: Tomonori
 * @Date: 2020/2/14 17:55
 * @Title:
 * @Desc: ↓ ↓ ↓ ↓ ↓
 * -----
 */
public class main {

    public static void main(String[] args) {
        ArrayStack<Integer> stack = new ArrayStack<>();

        for (int i = 0; i < 5; i++) {
            stack.push(i);
            System.out.println(stack);
        }

        stack.pop();
        System.out.println(stack);
    }
}
