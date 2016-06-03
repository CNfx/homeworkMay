import unittest

from mystack import MyStack


class MyStackTest(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        cls.s = MyStack(5)

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
        self.s.put('1')

    def test_get(self):
        self.s.get()

    def test_full(self):
        for i in xrange(6):
            self.s.put(i)

    def test_empty(self):
        for i in xrange(10):
            print self.s.get()


if __name__ == '__main__':
    case_list = [
        # 'test_mystack.MyStackTest.test_put',
        # 'test_mystack.MyStackTest.test_get',
        'test_mystack.MyStackTest.test_full',
        'test_mystack.MyStackTest.test_empty',
    ]
    suite = unittest.TestLoader().loadTestsFromNames(case_list)
    unittest.TextTestRunner(verbosity=2).run(suite)
