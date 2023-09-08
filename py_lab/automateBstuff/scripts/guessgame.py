import random

print("Hi there, what's your name? ")
name = input()
print(f"hi {name}, guess a number between 1 and 20")
secretnumber = random.randint(1, 19)

for numb in range(5):
    print("take a guess")
    guessnumber = int(input())
    if guessnumber > secretnumber:
        print("guess is larger")
    elif guessnumber < secretnumber:
        print("guess is smaller")
    else:
        break

if guessnumber == secretnumber:
    print("JACKPOT...you got it correct")
else:
    print("ACCOUNT LOCKED: Unsuccessful guess attempts made")
