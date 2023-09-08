passwordFile = open('SecretPasswordFile.txt')
secretpassword = passwordFile.read()
print('Enter your password.')
typedpassword = input()
if typedpassword == secretpassword:
    print('Access granted')
elif typedpassword =='12345':
    print('That password is one that an idiot puts on their luggage.')
else:
    print('Access denied')