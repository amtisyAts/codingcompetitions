N = int(input())

for i in range(1, N + 1):
    house, budget = map(int, input().split())
    prices = list(map(int, input().split()))
    prices.sort()
    j = 0
    for price in prices:
        if price > budget:
            break
        budget -= price
        j += 1
    print("Case #", i, ": ", j, sep='')
