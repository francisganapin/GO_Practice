
fruits = {'apple','banana','cherry'}


exist = False 
for fruit in range(fruits):
    if fruit == "banana":
        exist = True
        break

if exist == True:
    print('banana exist')
else:
    print('no banana')