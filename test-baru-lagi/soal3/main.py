def removeSpaces(inputStr):
    return input.replace(" ","")

def changetolowercase(input):
    return input.lower()

def countCharacters(inputStr):
    pecahtext = {}
    for char in inputStr:
        if char in pecahtext:
            pecahtext[char] += 1
        else:
            pecahtext[char] = 1
    return pecahtext
     
def printCharacterCounts(inputStr):
    counts = countCharacters(inputStr)
    outputStr = ""
    for char in counts:
        count = counts[char]
        if count == 1:
            outputStr += char
        else:
            outputStr += str(count) + char
    return outputStr

def main(input):
    return printCharacterCounts(removeSpaces(changetolowercase(input)))


input = "dani Maulana"
print("text yang anda masukkan adalah: "+ input);
print(countCharacters(input))
output = main(input)
print("Hasil setelah di kelola: "+ output);
print("")