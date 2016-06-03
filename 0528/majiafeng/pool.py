import array

class Pool:
    data = None
    cap = 0
    
    def __init__(self, n):
        self.data = array.array('l', [0 for i in range(n + 1)])
        self.cap = n
        for i in range(n):
            self.data[i] = i + 1
        self.data[n] = 0

    def __repr__(self):
        return str(self.data.tolist())

    def Push(self, v):
        n = self.data[0]

        if n == 0:
            return "full"

        new_head = self.data[self.data[0]]

        self.data[self.data[0]] = v
        self.data[0] = new_head

        self.cap -= 1

        return v


if __name__ == "__main__":
    st = Pool(3)
    print st
    print st.Push(1), st
    print st.Push(2), st
    print st.Push(3), st
    print st.Push(4), st
