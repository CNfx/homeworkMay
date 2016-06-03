class FullException(Exception):
    print 'list is full.'


class EmptyExcpetion(Exception):
    print 'list is empty.'


class Node(object):
    def __init__(self, item):
        self.__data = item
        self.__next = None

    def get(self):
        return self.__data

    def set(self, item):
        self.__data = item

    def next(self):
        return self.__next

    def setNext(self, next):
        self.__next = next

    def getNext(self):
        return self.__next


class MyLinkedList(object):

    _list = []
    _head = None
    _next = None
    _size = 0

    def __init__(self, cap=-1):
        self.cap = cap
        self._head = Node(None)
        self._head.setNext(self._head)

    def add(self, item):
        if self.cap and self.size() >= self.cap:
            raise FullException

        node = Node(item)
        self._head.setNext(node)
        node.setNext(self._head)

        self._size += 1

    def remove(self, item):
        if not self.size():
            raise EmptyExcpetion

        prev = self._head
        while prev.getNext() != self._head:
            cur = prev.getNext()
            if cur.get() == item:
                prev.setNext(cur.getNext())
                break
            prev = cur

        self._size -= 1

    def search(self, item):
        if not self.size():
            raise EmptyExcpetion

        cur = self._head.getNext()
        while cur != self._head:
            if cur.get() == item:
                return True
            cur = cur.getNext()

        return False

    def size(self):
        return self._size
