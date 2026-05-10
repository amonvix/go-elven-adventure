# ⚔️ Phase 01 — Village of Awakening

## Quest 3: Types of Destiny

The Hall of Records has opened its ancient gates.

Every being in Eldora must be registered by true nature.
A number must be a number.
A word must be a word.
Truth itself must be declared as true or false.

The Guild Master speaks:

> "Many travelers fail because they confuse appearance with essence.
> Learn the Types of Destiny, and your code shall become strong."

---

## Objective

Understand how Go stores different kinds of values and how type inference works.

---

## Concepts Learned

- string
- int
- bool
- float64
- short declaration :=
- explicit declaration with var
- zero values (preview)

---

## Your Task

Create a new file:

types.go

Build a runnable Go program containing the following variables:

- Hero name (string)
- Hero level (int)
- Hero mana (float64)
- Hero alive status (bool)

Then print all values.

Example inspiration:

Hero: Amon
Level: 1
Mana: 75.5
Alive: true

---

## Side Quest

Create one extra variable using var with explicit type.

Example:

var kingdom string

Assign a value later and print it.

Suggested output:

Kingdom: Eldora

---

## Mini Boss Fight

The Royal Clerk made a terrible mistake.

He wrote the hero level as text.

Correct the registry.

Create:

- one wrong example as string: "41"
- one correct example as int: 41

Print both and identify which one is valid for calculations.

Example:

Wrong Level: 41
Correct Level: 41
Only one of them can be used in math.

Bonus challenge:

Print the result of:

Correct Level + 1

Expected:

42

---

## Important Guidance

Use fmt.Println() to display your values.

Use := when declaring and assigning quickly.

Use var when declaring type explicitly.

Remember:

"41" = text
41 = number

75.5 = decimal number

true / false = boolean values

---

## Victory Condition

Create the file.
Run the program successfully.
Complete the Mini Boss.
Return to the Guild Master.

---

## Reward Unlocked

- You understand primitive types in Go.
- You can choose the correct data type.
- You are ready for calculations and logic.

---

## Next Phase →

move.forward()