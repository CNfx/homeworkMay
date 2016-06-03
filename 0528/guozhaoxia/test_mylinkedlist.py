import unittest

from mylinkedlist import MyLinkedList


class MyLinkedListTest(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        cls.ll = MyLinkedList(5)

    @classmethod
    def tearDownClass(cls):
        pass

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def runTest(self):
        pass

    def test_add(self):
        self.ll.add('1')
        self.ll.add('2')

    def test_remove(self):
        self.ll.remove('1')

    def test_search(self):
        print self.ll.search('1')


if __name__ == '__main__':
    case_list = [
        'test_mylinkedlist.MyLinkedListTest.test_add',
        'test_mylinkedlist.MyLinkedListTest.test_remove',
        'test_mylinkedlist.MyLinkedListTest.test_search'
    ]
    suite = unittest.TestLoader().loadTestsFromNames(case_list)
    unittest.TextTestRunner(verbosity=2).run(suite)
