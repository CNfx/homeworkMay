import array

class Queue:
    data = None
    rPos = 0
    wPos = 0
    cap = 0
    
    def __init__(self, n):
        self.data = array.array('l', [0 for i in range(n)])
        self.cap = n

    def Push(self, v):
        n = self.wPos - self.rPos

        if n == len(self.data):
            return 0, "full"

        self.data[self.wPos % self.cap] = v
        self.wPos += 1

        return v

    def Pop(self):
        if self.rPos == self.wPos:
            self.rPos = self.wPos = 0
            return 0, "empty"

        v = self.data[self.rPos % self.cap]
        self.rPos += 1

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

