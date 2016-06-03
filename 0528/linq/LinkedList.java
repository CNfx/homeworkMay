package learn.course;

import java.util.List;

public class LinkedList<T> {

    private int head;
    private int tail;
    private int count;

    private NodeEntry<T>[] array;

    public LinkedList(int capacity) {
        this.array = (NodeEntry<T>[]) new Object[capacity];
        this.head = this.tail = -1;
    }

    public void add(T element) {
        this.addFirst(element);
    }

    public void add(T element, int index) {
        if (this.count == this.array.length) {
            throw new RuntimeException("overflow");
        }

        if (0 == index) {
            this.addFirst(element);
        } else if (index == this.count) {
            this.addLast(element);
        } else if (index > 0 && index < this.count) {
            this.addMiddle(element, index);
        } else {
            throw new RuntimeException("invalid index");
        }
    }

    public T get(int index) {
        if (index >= this.count) {
            throw new RuntimeException("overflow");
        }

        if (this.count == 0) {
            throw new RuntimeException("underflow");
        }

        int position = this.head;
        for (int i = 1; i <= index; i++) {
            position = this.array[position].next;
        }

        return this.array[position].data;
    }

    public T remove(int index) {
        if (index == 0) {
            return removeFirst();
        } else if (index == this.count - 1) {
            return this.removeLast();
        } else if (index > 0 && index < this.count - 1) {
            return this.removeMiddle(index);
        } else {
            throw new RuntimeException("invalid index");
        }
    }

    private T removeFirst() {
        int oldHead = this.head;
        T data = this.array[this.head].data;

        this.head = this.array[this.head].next;

        this.array[oldHead] = null;

        this.count--;

        if (this.count == 0) {
            this.head = this.tail = -1;
        }

        return data;
    }

    private T removeLast() {
        int oldTail = this.tail;

        T data = this.array[this.tail].data;

        if (this.count == 1) {
            this.head = this.tail = -1;
        } else {
            int newTail = this.findPreNode(oldTail);
            this.tail = newTail;
        }

        this.array[oldTail] = null;
        this.count--;

        return data;
    }

    private T removeMiddle(int index) {
        int preIndex = this.findPreNode(index);

        T data = this.array[index].data;

        this.array[preIndex].next = this.array[index].next;

        this.array[index] = null;
        this.count--;

        return data;
    }

    private void addFirst(T element) {
        boolean isEmpty = this.isEmpty();
        NodeEntry<T> newHead = new NodeEntry<T>(element, this.head);
        int freeBlockIndex = this.findFreeBlock();
        this.array[freeBlockIndex] = newHead;
        this.head = freeBlockIndex;

        this.count++;

        if (isEmpty) {
            this.tail = this.head;
        }
    }

    private void addLast(T element) {
        boolean isEmpty = this.isEmpty();
        NodeEntry<T> newTail = new NodeEntry<T>(element, -1);
        int freeBlockIndex = this.findFreeBlock();
        this.array[freeBlockIndex] = newTail;

        if (!isEmpty) {
            this.array[this.tail].next = freeBlockIndex;
        }
        this.tail = freeBlockIndex;

        this.count++;

        if (isEmpty) {
            this.head = this.tail;
        }
    }

    private void addMiddle(T element, int index) {
        NodeEntry<T> newMiddle = new NodeEntry<T>(element, -1);

        int freeBlockIndex = this.findFreeBlock();
        this.array[freeBlockIndex] = newMiddle;

        int preIndex = findPreNode(index);
        int postIndex = this.array[preIndex].next;

        this.array[preIndex].next = freeBlockIndex;
        this.array[freeBlockIndex].next = postIndex;

        this.count++;
    }

    private int findPreNode(int index) {
        int position = this.head;
        for (int i = 1; i < index; i++) {
            position = this.array[position].next;
        }

        return position;
    }

    private boolean isEmpty() {
        return this.count == 0;
    }

    private int findFreeBlock() {
        int index = -1;

        for (int i = 0; i < this.array.length; i++) {
            if (null == this.array[i]) {
                index = i;
                break;
            }
        }

        return index;
    }

    class NodeEntry<T> {
        private int next;
        private T data;

        NodeEntry(T data, int next) {
            this.data = data;
            this.next = next;
        }
    }
}
