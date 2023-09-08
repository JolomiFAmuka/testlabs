


####################################################

# Number Guess Game 

##import random
##
##print("Hi there, what's your name? ")
##name = input()
##print(f"hi {name}, guess a number between 1 and 20")
##secretnumber = random.randint(1, 19)
##
##for numb in range(1,5):
##    print("take a guess")
##    guessnumber = int(input())
##    if guessnumber > secretnumber:
##        print("guess is larger")
##    elif guessnumber < secretnumber:
##        print("guess is smaller")
##    else:
##        break
##
##if guessnumber == secretnumber:
##    print(f"JACKPOT...you got it correct in {numb} tries")
##else:
##    print(f"ACCOUNT LOCKED: Unsuccessful...the number was actually {secretnumber}")


##############################

##print("How many cats do you have?")
##
##numCats = input()
##try:
##    if int(numCats) >= 4:
##        print("That's a lot of cats")
##    elif int(numCats) < 0:
##        print("you can't have -ve numbers")
##    else:
##        print("That's not many cats")
##except ValueError:
##    print("you entered a wrong value")

##def divideBy42(numb):
##    try:
##        return 42/numb
##    except ZeroDivisionError:
##        print("Error detected: you tried to divide by 0")
##
##print(divideBy42(2))
##print(divideBy42(10))
##print(divideBy42(0.5))
##print(divideBy42(1))

########################

##import random
##
##def rolldice():
##    print(random.randint(1,6))
##
##rolldice()
##rolldice()
##rolldice()
##rolldice()

################################
##total = 0
##for num in range(101):
##    total += num
##print(total)

##user_name = input("what is your name ")
##print(f"your name is {user_name} and it has {len(user_name)} cters")

