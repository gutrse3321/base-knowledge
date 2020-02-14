package ru.reimu.stack;

/**
 * @Author: Tomonori
 * @Date: 2020/2/14 17:58
 * @Title:
 * @Desc: ↓ ↓ ↓ ↓ ↓
 * -----
 */
public interface Stack<T> {

    int getSize();

    boolean isEmpty();

    void push(T t);

    T pop();

    T peek();
}
