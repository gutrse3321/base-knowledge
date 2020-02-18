package ru.reimu.queue;

/**
 * @Author: Tomonori
 * @Date: 2020/2/18 11:38
 * @Title:
 * @Desc: ↓ ↓ ↓ ↓ ↓
 * -----
 */
public class LoopQueue<T> implements IQueue<T> {

    private T[] data;
    private int front;
    private int tail;
    private int size;

    public LoopQueue() {
        this(10);
    }

    public LoopQueue(int cap) {
        this.data = (T[]) new Object[cap + 1];
        this.front = 0;
        this.tail = 0;
        this.size = 0;
    }

    public int getCap() {
        return data.length - 1;
    }

    @Override
    public int getSize() {
        return size;
    }

    @Override
    public boolean isEmpty() {
        return front == tail;
    }

    @Override
    public void enqueue(T t) {
        if ((tail + 1) % data.length == front) {
            resize(getCap() * 2);
        }
        data[tail] = t;
        tail = (tail + 1) % data.length;
        size++;
    }

    @Override
    public T dequeue() {
        if (isEmpty()) {
            throw new IllegalArgumentException("cannot dequeue from an empty queue.");
        }

        T item = data[front];
        data[front] = null;
        front = (front + 1) % data.length;
        size--;

        if (size == getCap() / 4 && getCap() / 2 != 0) {
            resize(getCap() / 2);
        }

        return item;
    }

    @Override
    public T getFront() {
        if (isEmpty()) {
            throw new IllegalArgumentException("cannot getFront from an empty queue.");
        }

        return data[front];
    }

    private void resize(int newCap) {
        T[] newData = (T[]) new Object[newCap + 1];

        for (int i = 0; i < size; i++) {
            newData[i] = data[(front + i) % data.length];
        }

        data = newData;
        front = 0;
        tail = size;
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append(String.format("Queue: size = %d, capacity = %d\n", size, getCap()));
        sb.append("front [");

        for (int i = front; i != tail; i = (i + 1) % data.length) {
            sb.append(data[i]);
            if ((i + 1) % data.length != tail) {
                sb.append(", ");
            }
        }

        sb.append("] tail");

        return sb.toString();
    }
}
