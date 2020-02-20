package ru.reimu.queue;

/**
 * @Author: Tomonori
 * @Date: 2020/2/20 18:15
 * @Title:
 * @Desc: ↓ ↓ ↓ ↓ ↓
 * -----
 */
public class LinkedListQueue<T> implements IQueue<T> {

    private Node head;
    private Node tail;
    private int size;

    public LinkedListQueue() {
        this.head = null;
        this.tail = null;
        this.size = 0;
    }

    @Override
    public int getSize() {
        return size;
    }

    @Override
    public boolean isEmpty() {
        return size == 0;
    }

    @Override
    public void enqueue(T t) {
        if (tail == null) {
            tail = new Node(t);
            head = tail;
        } else {
            tail.next = new Node(t);
            tail = tail.next;
        }
        size++;
    }

    @Override
    public T dequeue() {
        if (isEmpty()) {
            throw new IllegalArgumentException("Cannot dequeue from an empty queue.");
        }

        Node deNode = head;
        head = head.next;
        deNode.next = null;

        if (head == null) {
            tail = null;
        }
        size--;

        return deNode.t;
    }

    @Override
    public T getFront() {
        if (isEmpty()) {
            throw new IllegalArgumentException("Cannot dequeue from an empty queue.");
        }
        return head.t;
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append("Queue: head ");

        Node current = head;
        while (current != null) {
            sb.append(current + " -> ");
            current = current.next;
        }

        sb.append("null tail");

        return sb.toString();
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
