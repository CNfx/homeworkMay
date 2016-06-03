package learn.course;

public class RingQueue<T> {

    private T[] array;

    private int putIndex;
    private int takeIndex;

    public RingQueue(int capacity) {
        this.array = (T[]) new Object[capacity];
    }

    public void enqueue(T element) {
        final int capacity = this.array.length;
        
        if (this.putIndex - this.takeIndex == capacity) {
            throw new RuntimeException("overflow");
        }
        
        this.array[this.putIndex % capacity] = element;
        this.putIndex++;
    }

    public T dequeue() {
        final int capacity = this.array.length;

        if (this.putIndex == this.takeIndex) {
            throw new RuntimeException("underflow");
        }

        T element = this.array[this.takeIndex % capacity];
        this.array[this.takeIndex % capacity] = null;

        this.takeIndex++;

        return element;
    }

    public int size() {
        return this.putIndex - this.takeIndex;
    }

}
