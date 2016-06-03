import array
import itertools

class Stack:
    data = None
    
    def __init__(self, n):
        self.data = array.array('l', list(itertools.repeat(0, n+1)))
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

def test():
    st = Stack(3)
    print st.Push(1), st
    print st.Push(2), st
    print st.Push(3), st
    print st.Push(4), st

    print st.Pop(), st
    print st.Pop(), st
    print st.Pop(), st
    print st.Pop(), st

def test2(n):
    stack = []
    for i in range(2*n):
        stack.append(i)

def test3(n):
    st = Stack(n)
    for i in range(n):
        st.Push(i)


if __name__ == "__main__":
    #test()
    import timeit

    print timeit.timeit("test2(10000000)", setup="from __main__ import test2, test3", number=1)
    print timeit.timeit("test3(10000000)", setup="from __main__ import test2, test3", number=1)

