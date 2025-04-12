

def testcount(x):
    if x == 11:
        return 0
    else:
        print(x)
        return testcount(x + 1)
    

testcount(1)