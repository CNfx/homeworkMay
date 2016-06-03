import array

class Stack:
    data = None
    
    def __init__(self, n):
        self.data = array.array('l', [0 for i in range(n + 1)])
        self.data[0] = n

    def __str__(self):
        return str(self.data.tolist())

    def Push(self, v):
        n = self.data[0]
        if n == 0:
            return "full"

        self.data[n] = v
        self.data[0] = n-1

        return None

    def Pop(self):
        n = self.data[0] + 1
        if n == len(self.data):
            return 0, "empty"

        v = self.data[n]
        self.data[0] = n

        return v


if __name__ == "__main__":
    st = Stack(3)
    print st.Push(1), st
    print st.Push(2), st
    print st.Push(3), st
    print st.Push(4), st

    print st.Pop(), st
    print st.Pop(), st
    print st.Pop(), st
    print st.Pop(), st


