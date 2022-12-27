# katas
Collection of coding examples, problems generally sourced from online coding problems

# the wave (codewars)
In this simple Kata your task is to create a function that turns a string into the Wave. You will be passed a string and you must return that string in an array where an uppercase letter is a person standing up.
    - Example: `wave("hello") => ["Hello", "hEllo", "heLlo", "helLo", "hellO"]`
    -  - Example: `wave("hello me") => ["Hello", "hEllo", "heLlo", "helLo", "hellO", "hello Me", "hello mE]`

# the thief
A stealthy burglar is causing a big headache for people in a town. He is known to be swift in getting in the house unnoticed and steal all the diamonds from the locker. He holds a bag in one hand and picks the diamonds from other hand, while stealing.

Your objective - if choose to accept - is to calculate and return an integer representing the minimum number of repetitions required for the burglar to pick all the diamonds from the locker.
    - Example locker array:
        ['*.*.*.*.*.',
        '...*..**..',
        '**.**...*.',
        '**..**..**'] -> 13 repetitions
Rules
The burglar steals a single diamond every time but can steal two if diamonds are horizontally adjacent to each other.

The burglar cannot steal more than two diamonds at once.

A row in the locker can have 0 diamonds.

The locker will have at least 1 diamond in it.
All the inputs are valid.