package learn.course;

public class Queue<T> {
    
    private T[] array;
    
    private int putIndex;
    private int takeIndex;
    
    private int count;
    
    public Queue(int capacity) {
        this.array = (T[])new Object[capacity];
    }
    
    public void enqueue(T element) {
        if (this.count == this.array.length) {
            throw new RuntimeException("overflow");
        }
        
        this.array[this.putIndex] = element;
        this.putIndex++;
        this.count++;
        if (this.putIndex == this.array.length) {
            this.putIndex = 0;
        }
    }
    
    public T dequeue() {
        if (this.count == 0) {
            throw new RuntimeException("underflow");
        }
        
        T element = this.array[this.takeIndex];
        this.array[this.takeIndex] = null;
        
        this.count--;
        this.takeIndex++;
        if (this.takeIndex == this.array.length) {
            this.takeIndex = 0;
        }
        
        return element;
    }
    
    public int size() {
        return this.count;
    }

}
