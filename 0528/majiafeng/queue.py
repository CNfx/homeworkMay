import array

class Queue:
    data = None
    rPos = 0
    wPos = 0
    cap = 0
    
    def __init__(self, n):
        self.data = array.array('l', [0 for i in range(n)])
        self.cap = n

    def __repr__(self):
        return str(self.data.tolist())

    def Push(self, v):
        n = self.wPos - self.rPos

        if n == len(self.data):
            return "full"

        self.data[self.wPos % self.cap] = v
        self.wPos += 1

        return None

    def Pop(self):
        if self.rPos == self.wPos:
            self.rPos = self.wPos = 0
            return 0, "empty"

        v = self.data[self.rPos % self.cap]
        self.rPos += 1

        return v


if __name__ == "__main__":
    st = Queue(3)
    print st.Push(1), st
    print st.Push(2), st
    print st.Push(3), st
    print st.Push(4), st

    print st.Pop(), st
    print st.Pop(), st
    print st.Pop(), st
    print st.Pop(), st

    print st.Push(5), st
    print st.Push(6), st
    print st.Pop(), st
    print st.Push(7), st

