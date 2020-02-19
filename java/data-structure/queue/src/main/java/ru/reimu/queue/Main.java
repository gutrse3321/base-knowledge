package ru.reimu.queue;

import java.util.Random;

/**
 * @Author: Tomonori
 * @Date: 2020/2/17 15:08
 * @Title:
 * @Desc: ↓ ↓ ↓ ↓ ↓
 * -----
 */
public class Main {

    public static void main(String[] args) {
        System.out.println(Integer.MAX_VALUE);

        int count = 100000;

        ArrayQueue<Integer> arrayQueue = new ArrayQueue<>();
        System.out.println(testQueue(arrayQueue, count) + "s");

        LoopQueue<Integer> loopQueue = new LoopQueue<>();
        System.out.println(testQueue(loopQueue, count) + "s");
    }

    private static double testQueue(IQueue<Integer> queue, int opCount) {
        long startTime = System.nanoTime();
        Random random = new Random();

        for (int i = 0; i < opCount; i++) {
            queue.enqueue(random.nextInt(Integer.MAX_VALUE));
        }

        for (int i = 0; i < opCount; i++) {
            queue.dequeue();
        }

        long endTime = System.nanoTime();

        return (endTime - startTime) / 1000000000.0;
    }
}
