import array

class Pool:
    data = None
    cap = 0
    
    def __init__(self, n):
        self.data = array.array('l', [0 for i in range(n + 1)])
        self.cap = n
        for i in range(n - 1):
            self.data[i] = i + 1
        self.data[n] = 0

    def Push(self, v):
        n = self.data[0]

        if n == 0:
            return 0, "full"

        new_head = self.data[self.data[0]]

        self.data[self.data[0]] = v
        self.data[0] = new_head

        self.cap -= 1

        return v


if __name__ == "__main__":
    st = Queue(3)
    str_push = "+ %d"
    str_pop = "- %d"
    print str_push%st.Push(1)
    print str_push%st.Push(2)
    print str_push%st.Push(3)
    print st.Push(4)

    print str_pop%st.Pop()
    print str_pop%st.Pop()
    print str_pop%st.Pop()
    print st.Pop()

    print str_push%st.Push(5)
    print str_push%st.Push(6)
    print str_pop%st.Pop()
    print str_push%st.Push(7)

