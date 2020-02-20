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

        int count = 10000000;

        //time-out O(n)
        ArrayQueue<Integer> arrayQueue = new ArrayQueue<>();
        System.out.println(testQueue(arrayQueue, count) + "s");

        //1.1471862s O(1)
        LoopQueue<Integer> loopQueue = new LoopQueue<>();
        System.out.println(testQueue(loopQueue, count) + "s");

        //19.7045847s O(1)
        LinkedListQueue<Integer> linkedListQueue = new LinkedListQueue<>();
        System.out.println(testQueue(linkedListQueue, count) + "s");
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
