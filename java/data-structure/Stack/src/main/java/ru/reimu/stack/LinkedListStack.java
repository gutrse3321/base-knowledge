package ru.reimu.stack;

import ru.reimu.linkedList.LinkedList;

/**
 * @Author: Tomonori
 * @Date: 2020/2/20 14:37
 * @Title:
 * @Desc: ↓ ↓ ↓ ↓ ↓
 * -----
 */
public class LinkedListStack<T> implements Stack<T> {

    private LinkedList<T> linkedList;

    public LinkedListStack() {
        this.linkedList = new LinkedList<>();
    }

    @Override
    public int getSize() {
        return linkedList.getSize();
    }

    @Override
    public boolean isEmpty() {
        return linkedList.isEmpty();
    }

    @Override
    public void push(T t) {
        linkedList.addFirst(t);
    }

    @Override
    public T pop() {
        return linkedList.removeFirst();
    }

    @Override
    public T peek() {
        return linkedList.getFirst();
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append("Stack: top ");
        sb.append(linkedList);

        return sb.toString();
    }
}
