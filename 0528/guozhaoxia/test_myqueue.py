import unittest

from myqueue import MyQueue


class MyQueueTest(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        cls.q = MyQueue(5)

    @classmethod
    def tearDownClass(cls):
        pass

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def runTest(self):
        pass

    def test_put(self):
        self.q.put('1')

    def test_get(self):
        self.q.get()

    def test_full(self):
        for i in xrange(6):
            self.q.put(i)

    def test_empty(self):
        for i in xrange(10):
            self.q.get()


if __name__ == '__main__':
    case_list = [
        # 'test_myqueue.MyQueueTest.test_put',
        # 'test_myqueue.MyQueueTest.test_get',
        # 'test_myqueue.MyQueueTest.test_full',
        'test_myqueue.MyQueueTest.test_empty',
    ]
    suite = unittest.TestLoader().loadTestsFromNames(case_list)
    unittest.TextTestRunner(verbosity=2).run(suite)
