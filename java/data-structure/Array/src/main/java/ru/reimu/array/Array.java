package ru.reimu.array;

/**
 * @Author: Tomonori
 * @Date: 2020/2/13 14:49
 * @Title:
 * @Desc: ↓ ↓ ↓ ↓ ↓
 * -----
 */
public class Array<T> {

    private T[] data;
    private int   size;

    public Array() {
        this(10);
    }

    /**
     * 传入数组的容量capacity构造Array
     * @param capacity
     */
    public Array(int capacity) {
        this.data = (T[]) new Object[capacity];
        this.size = 0;
    }

    /**
     *
     * @return 元素个数
     */
    public int getSize() {
        return size;
    }

    /**
     *
     * @return 数组的容量
     */
    public int getCap() {
        return data.length;
    }

    public T get(int index) {
        if (index < 0 || index >= size) {
            throw new IllegalArgumentException("Get failed, Require index >= 0 and index <= size");
        }
        return data[index];
    }

    public T getFirst() {
        return get(0);
    }

    public T getLast() {
        return get(size - 1);
    }

    public void set(int index, T e) {
        if (index < 0 || index >= size) {
            throw new IllegalArgumentException("Set failed, Require index >= 0 and index <= size");
        }
        data[index] = e;
    }

    public boolean isEmpty() {
        return size == 0;
    }

    /**
     * 查找数组中是否有元素e
     * @param e
     * @return
     */
    public boolean contains(T e) {
        for (T ele : data) {
            //值比较
            if (ele.equals(e)) {
                return true;
            }
        }
        return false;
    }

    /**
     * 查找数组中的元素索引，-1没有找到
     * @param e
     * @return
     */
    public int find(T e) {
        for (int i = 0; i < size; i++) {
            if (data[i].equals(e)) {
                return i;
            }
        }
        return -1;
    }

    /**
     * 所有元素钱添加元素
     * @param e
     */
    public void addFirst(T e) {
        add(0, e);
    }

    /**
     * 所有元素后添加元素
     * @param e
     */
    public void addLast(T e) {
        add(size, e);
    }

    /**
     * 在第index的位置插入一个新元素
     * @param index
     * @param e
     */
    public void add(int index, T e) {
        if (index < 0 || index > size) {
            throw new IllegalArgumentException("Add failed, Require index >= 0 and index <= size");
        }

        if (size == data.length) {
            resize(2 * data.length);
        }

        for (int i = size - 1; i >= index; i--) {
            data[i + 1] = data[i];
        }

        data[index] = e;
        size++;
    }

    private void resize(int newCapacity) {
        T[] newArr = (T[]) new Object[newCapacity];

        for (int i = 0; i < size; i++) {
            newArr[i] = data[i];
        }
        data = newArr;
    }

    /**
     * 从数组删除index位置的元素，返回删除的元素
     * @param index
     * @return
     */
    public T remove(int index) {
        if (index < 0 || index >= size) {
            throw new IllegalArgumentException("Remove failed, index is illegal");
        }

        T ele = data[index];
        for (int i = index + 1; i < size; i++) {
            data[i - 1] = data[i];
        }
        size--;
        data[size] = null;

        //防止减容resize过于着急(Eager)，在1/4的时候才减容
        //!=0防止数组中只有一个的情况
        if (size == data.length / 4 && data.length / 2 != 0) {
            resize(data.length / 2);
        }

        return ele;
    }

    public T removeFirst() {
        return remove(0);
    }

    public T removeLast() {
        return remove(size - 1);
    }

    /**
     * 删除数组中的某个元素e
     * @param e
     */
    public void removeElement(T e) {
        int index = find(e);

        if (index != -1) {
            remove(index);
        }
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append(String.format("Array: size = %d, capacity = %d\n", size, data.length));
        sb.append('[');

        for (int i = 0; i < size; i++) {
            sb.append(data[i]);
            if (i != size - 1) {
                sb.append(", ");
            }
        }

        sb.append(']');

        return sb.toString();
    }
}
