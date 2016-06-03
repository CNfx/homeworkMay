import array

class LinkedList:
    data = None
    cap = 0

    def __init__(self, n):
        self.data = array.array('l', [0 for i in range(2*n + 2)])
        for i in range(1, n + 1):
            self.data[i] = i + 1

        self.cap = n

        # head
        self.data[0] = -1

        # last one
        self.data[n+1] = -1

    def __repr__(self):
        return str(self.data.tolist())

    def Push(self, v):
        if self.data[1] == -1:
            return "full"
        
        # get a empty node from write list
        new_node = self.data[1]
        # remove the node from write list
        self.data[1] = self.data[new_node]

        # save data
        self.data[new_node + self.cap] = v


        # add te new node into read list
        self.data[new_node] = self.data[0]
        self.data[0] = new_node
        
        self.Print("push %s"%v)
        return v

    def Remove(self, v):
        found = False
        node = self.data[0]
        pre_node = 0 
        while 1:
            if node == -1:
                break

            if self.data[node + self.cap] == v:
                found = True
                break
            else:
                pre_node = node
                node = self.data[node]

        if not found:
            return "Not found"

        self.data[pre_node] = self.data[node]

        # add removed node to writeable list

        self.data[node] = self.data[1]
        self.data[1] = node

        self.Print("remove %s"%v)


    # for trace the changes
    def Print(self, msg):
        print '-----------------------------'
        print msg
        data_list = "Data list: head => "

        node = self.data[0]
        
        while node != -1:
            data_list += "%s(%s) => "%(node, self.data[node + self.cap])
            node = self.data[node]
        
        print data_list
            
        empty_list = "Writeable list: head => "

        node = self.data[1]
        while node != -1:
            empty_list += "%s => "%(node)
            node = self.data[node]
        
        print empty_list



if __name__ == "__main__":
    st = LinkedList(5)
    print st
    st.Push(11)
    st.Push(22)
    st.Push(33)
    st.Push(44)
    st.Push(55)
    st.Push(66)
    
    st.Remove(33)
    st.Remove(33)
    st.Remove(55)


    st.Push(77)
    st.Push(88)

