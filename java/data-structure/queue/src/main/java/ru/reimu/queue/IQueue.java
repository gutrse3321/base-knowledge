package ru.reimu.queue;

/**
 * @Author: Tomonori
 * @Date: 2020/2/17 15:09
 * @Title:
 * @Desc: ↓ ↓ ↓ ↓ ↓
 * -----
 */
public interface IQueue<T> {

    int getSize();

    boolean isEmpty();

    void enqueue(T t);

    T dequeue();

    T getFront();
}
