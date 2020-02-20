package ru.reimu.linkedList;

import java.util.Random;

/**
 * @Author: Tomonori
 * @Date: 2020/2/19 15:27
 * @Title:
 * @Desc: ↓ ↓ ↓ ↓ ↓
 * -----
 */
public class Main {

    public static void main(String[] args) {
        LinkedList<Integer> linkedList = new LinkedList<>();

        for (int i = 0; i < 5; i++) {
            linkedList.addFirst(i);
            System.out.println(linkedList);
        }

        linkedList.add(3, 100);
        System.out.println(linkedList);

        linkedList.set(2, 50);
        System.out.println(linkedList);

        System.out.println(linkedList.contains(1));

        System.out.println(linkedList.get(5));

        linkedList.remove(5);
        System.out.println(linkedList);

        linkedList.removeLast();
        System.out.println(linkedList);
    }
}
