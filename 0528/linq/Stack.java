package learn.course;

public class Stack<T> {
    private T[] array;
    
    /**
     * The position available to push the element.
     */
    private int position;
    
    public Stack(int capacity) {
        this.array = (T[]) new Object[capacity];
    } 
    
    public void push(T element) {
        if (this.position >= this.array.length) {
            throw new RuntimeException("overflow");
        }
        
        this.array[this.position] = element;
        this.position++;
    }
    
    public T pop() {
        if (this.position == 0) {
            throw new RuntimeException("underflow");
        }
        this.position--;
        T element = this.array[this.position];
        this.array[this.position] = null;
        
        return element;
    }
    
    public int size() {
        return this.position;
    }
}
