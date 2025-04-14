i = 1

while i <= 3:
    print(i)
    i += 1


for j in range(3):
    print(j)
    
for i in range(3):
    print("range",i)

while True:
    print("loop")
    break

for n in range(6):
    if n % 2 ==0:
        continue
    print(n)
    