import unittest

from mypool import MyPool


class MyPoolTest(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        cls.p = MyPool(5)

    @classmethod
    def tearDownClass(cls):
        pass

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def runTest(self):
        pass

    def test_get(self):
        node = self.p.get()
        print 'get: ', node._id, node._use

    def test_back(self):
        node = self.p.get()
        print 'back1: ', node._id, node._use
        self.p.back(node)
        print 'back2: ', node._id, node._use


if __name__ == '__main__':
    case_list = [
        'test_mypool.MyPoolTest.test_get',
        'test_mypool.MyPoolTest.test_back'
    ]
    suite = unittest.TestLoader().loadTestsFromNames(case_list)
    unittest.TextTestRunner(verbosity=2).run(suite)
