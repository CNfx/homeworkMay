import uuid


class Node(object):
    def __init__(self):
        self._id = uuid.uuid4().hex
        self._use = 0

    def get(self):
        return self._id

    def use(self):
        self._use += 1

    def rel(self):
        self._use -= 1


class MyPool(object):

    _list = []
    _size = 0

    def __init__(self, cap=-1):
        self.cap = cap

    def create(self):
        node = Node()
        self._list.append(node)
        self._size += 1

        return node

    def get(self):
        if not self.size():
            node = self.create()
        else:
            s = self.size()
            node = self._list[0]
            i = 0
            while i < (s - 1):
                next = self._list[i + 1]
                if node._use > next._use:
                    node = next
                i += 1

            if node._use >= 1 and s < self.cap:
                node = self.create()

        node.use()

        return node

    def back(self, node):
        node.rel()

    def size(self):
        return self._size
