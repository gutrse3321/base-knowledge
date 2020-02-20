package ru.reimu.stack;

import java.util.Random;

/**
 * @Author: Tomonori
 * @Date: 2020/2/14 17:55
 * @Title:
 * @Desc: ↓ ↓ ↓ ↓ ↓
 * -----
 */
public class main {

    public static void main(String[] args) {
        int count = 10000000;

        ArrayStack<Integer> arrayStack = new ArrayStack<>();
        System.out.println(testStack(arrayStack, count) + "s");

        LinkedListStack<Integer> linkedListStack = new LinkedListStack<>();
        System.out.println(testStack(linkedListStack, count) + "s");
    }


    private static double testStack(Stack<Integer> stack, int opCount) {
        long startTime = System.nanoTime();
        Random random = new Random();

        for (int i = 0; i < opCount; i++) {
            stack.push(random.nextInt(Integer.MAX_VALUE));
        }

        for (int i = 0; i < opCount; i++) {
            stack.pop();
        }

        long endTime = System.nanoTime();

        return (endTime - startTime) / 1000000000.0;
    }
}
