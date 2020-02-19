package ru.reimu.linkedList;

/**
 * @Author: Tomonori
 * @Date: 2020/2/19 15:27
 * @Title:
 * @Desc: ↓ ↓ ↓ ↓ ↓
 * -----
 */
public class LinkedList<T> {

    private Node dummyHead;
    private int size;

    public LinkedList() {
        this.dummyHead = new Node(null, null);
        this.size = 0;
    }

    public int getSize() {
        return size;
    }

    public boolean isEmpty() {
        return size == 0;
    }

    /**
     * TODO 根据索引位置添加节点(不常用)
     * 在链表的index位置添加新的元素
     * @param index
     * @param t
     */
    public void add(int index, T t) {
        if (index < 0 || index > size) {
            throw new IllegalArgumentException("Add failed. Illegal index");
        }

        Node prev = dummyHead;
        for (int i = 0; i < index; i++) {
            prev = prev.next;
        }

        //找到index - 1的节点元素，构造一个新的节点，并将index - 1之前的
        //下一个节点赋给新的节点的next，最后将新的节点赋值index - 1的next
        prev.next = new Node(t, prev.next);
        size++;
    }

    /**
     * 在链表头部添加新的节点，并将头部改成这个新加的
     * @param t
     */
    public void addFirst(T t) {
        add(0, t);
    }

    public void addLast(T t) {
        add(size, t);
    }

    private class Node {

        public T t;

        public Node next;

        public Node() {
            this(null, null);
        }

        public Node(T t) {
            this(t, null);
        }

        public Node(T t, Node next) {
            this.t = t;
            this.next = next;
        }

        @Override
        public String toString() {
            return t.toString();
        }
    }
}
