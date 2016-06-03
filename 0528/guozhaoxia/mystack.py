class FullException(Exception):
    print 'stack is full.'


class EmptyExcpetion(Exception):
    print 'stack is empty.'


class MyStack(object):

    _list = []
    _size = 0

    def __init__(self, cap=-1):
        self.cap = cap

    def put(self, item):
        if self.cap and self.size() >= self.cap:
            raise FullException
        self._list.append(item)
        self._size += 1

    def get(self):
        if not self.size():
            raise EmptyExcpetion

        item = self._list[-1]
        self._list.pop(0)
        self._size -= 1

        return item

    def size(self):
        return self._size
