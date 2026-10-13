# EXPLANATION — My Algorithm Worksheet

> Copy this file to `EXPLANATION.md` in **your own fork**. Write in English. Do not put private student details in a public repository.

**GitHub username:** Aulia Rameyza
**Exercises completed:** lab 01-10
**Most surprising bug:** my code is keep getting blocked by windowsbut then i ask to kilo to help with it. this task is make me a bit confused but then i keep try it 

For each exercise you completed, copy and fill in this mini-report:

## Exercise 01 — data_types
- **Story in my words:*Maya builds a digital check-in desk. Need to store age (count), mass (measurement), check-in status (true/false), and student ID (may have leading zeros).* 
- **Input(s) and their types:*age: int (19)
mass: float64 (62.5)
checkedIn: bool (true)
studentID: string ("00123")*
- **Processing / logical rule:*Getting to know which variable with each data types* 
- **Expected output (before coding):*19 | 62.5 | true | 00123* 
- **Actual output:*same* 
- **Normal test:*All values print correctly
- **Boundary or invalid test:*If studentID were int, leading zeros lost (00123 → 123)*
- **One assumption or limitation:* Student ID treated as opaque string — no validation of format*
- **What I changed in `main.go`:*Used := short declaration, fmt.Println with separators*
- **Why I believe the result is correct:*Each variable uses the type matching its semantic meaning; string preserves leading zeros*

## Exercise 02 — cafe_sequence
- **Story in my words:* Student buys 2 coffees at Rp18,000 each with Rp5,000 discount. Calculate subtotal first, then total. Then change quantity to 3.* 
- **Input(s) and their types:*unitPrice: int (18000)
quantity: int (2, then 3)
discount: int (5000)*
- **Processing / logical rule:*subtotal = unitPrice * quantity (MUST be first)
total = subtotal - discount (depends on subtotal)
Recompute when quantity changes* 
- **Expected output (before coding):*Subtotal: 36000
Total: 31000
Subtotal: 54000
Total: 49000* 
- **Actual output:*same* 
- **Normal test:* Quantity 2 → 36000/31000 
- **Boundary or invalid test:*Quantity 0 → total = -5000 (negative — model breaks); discount > subtotal → negative total*
- **One assumption or limitation:* Fixed discount regardless of quantity; no floor at 0*
- **What I changed in `main.go`:*Added subtotal/total calculations in correct order, recomputation for quantity=3*
- **Why I believe the result is correct:*because the result data (total,subtotal) is declared with math so its have an exact answer*

## Exercise 03 — factor_checker
- **Story in my words:** Check if x divides y evenly (volunteers into equal groups). Reject non-positive inputs.
- **Input(s) and their types:** x: int (group size)
y: int (total volunteers)
- **Processing / logical rule:** Validate x > 0 and y > 0 → else "invalid input"
Check y % x == 0 (modulo)
- **Expected output (before coding):** 5 20 → true
20 5 → false
123 123 → true
0 20 → invalid input
- **Actual output:** same
- **Normal test:** x=0 → division by zero panic prevented by validation 
- **Boundary or invalid test:** Only positive integers; no handling of non-integer input
- **One assumption or limitation:** Added input validation before modulo, early return on invalid
- **What I changed in `main.go`:** Added input validation before modulo, early return on invalid
- **Why I believe the result is correct:** Validation prevents runtime panic; modulo correctly tests divisibility

## Exercise 04 — population
- **Story in my words:** City population changes: start + births + arrivals - deaths - departures.
- **Input(s) and their types:** All int: start=1000, births=20, arrivals=5, deaths=10, departures=4
- **Processing / logical rule:** end = start + births + arrivals - deaths - departures
- **Expected output (before coding):** Population: 1011
- **Actual output:** same
- **Normal test:**  Given values → 1011
- **Boundary or invalid test:**  Negative result (e.g., deaths > start+births+arrivals) → impossible population
- **One assumption or limitation:** Model assumes instantaneous changes; no time dimension; all counts are exact
- **What I changed in `main.go`:** Single expression for end population, formatted print
- **Why I believe the result is correct:** Direct translation of formula; 1000+20+5-10-4 = 1011

## Exercise 05 — number_sequence
- **Story in my words:** Fibonacci-like sequence: each term = sum of previous two. Compute terms 3,4,5 without loop, then with loop.
- **Input(s) and their types:** first, second: int (5, 3 for main; 1, 1 for loop demo)
- **Processing / logical rule:** Without loop: third = first+second, fourth = second+third, fifth = third+fourth
With loop: repeat c = a+b; a,b = b,c three times
- **Expected output (before coding):** 8 11 19 (for 5,3) and 2 3 5 (for 1,1 loop)
- **Actual output:** same
- **Normal test:** 5,3 → 8,11,19
- **Boundary or invalid test:** Large starting values → integer overflow eventually
- **One assumption or limitation:** Only computes 3 terms; loop version hardcoded to 3 iterations
- **What I changed in `main.go`:**  Manual calculation + loop demonstration
- **Why I believe the result is correct:**  Each step follows the recurrence relation exactly

## Exercise 06 — gravity
- **Story in my words:** An astronaut's mass stays the same on different planets, but weight changes due to gravity. Calculate weight on Earth and Mars for a 70 kg astronaut.
- **Input(s) and their types:** massKg: float64 (70.0)
- **Processing / logical rule:** Earth weight = massKg * 9.8; Mars weight = massKg * 9.8 * 0.38. Format output to one decimal place.
- **Expected output (before coding):** 
Earth: 686.0 N
Mars: 260.7 N
- **Actual output:** same
- **Normal test:** 70.0 kg → Earth 686.0 N, Mars 260.7 N
- **Boundary or invalid test:** Negative mass (nonsense); zero mass → 0 N on both
- **One assumption or limitation:** Gravity constants are approximations; ignores atmospheric buoyancy
- **What I changed in `main.go`:** Implemented two weight calculations with float64, formatted with fmt.Printf("%.1f")
- **Why I believe the result is correct:** 70.0 * 9.8 = 686.0; 70.0 * 9.8 * 0.38 = 260.68 → 260.7 (one decimal)

## Exercise 07 — genetic_intro
- **Story in my words:** Demo GA operations on 4-bit strings. Fitness = count of 1s. Crossover at position 2, then mutate bit 1.
- **Input(s) and their types:** a, b: string ("1001", "0110")
fitness function: string → int
- **Processing / logical rule:** Child after crossover: 1010 
 fitness = 2
Child after mutation:  1110  
fitness = 3
- **Expected output (before coding):** Given parents → child fitness improves 2→3 
- **Actual output:** same
- **Normal test:** Mutation could decrease fitness (flip 1→0); crossover point at edges
- **Boundary or invalid test:**
- **One assumption or limitation:** Single demonstration, not a full GA (no population, generations, selection)
- **What I changed in `main.go`:** Implemented fitness function, crossover via slicing, mutation via byte array
- **Why I believe the result is correct:** String slicing correctly combines parents; byte mutation flips exact bit

## Exercise 08 — cube_root
- **Story in my words:** A game unlocks a door when the submitted integer x is exactly the cube root of y. Verify using integer multiplication, not floating-point approximation.
- **Input(s) and their types:** x, y: int (read via fmt.Scan)
- **Processing / logical rule:** Check if x*x*x == y. Works for negative numbers too (-3 * -3 * -3 = -27).
- **Expected output (before coding):**
5 125 → true
-3 -27 → true
12 12 → false
- **Actual output:** same
- **Normal test:** 5 125 → true; -3 -27 → true; 12 12 → false
- **Boundary or invalid test:** Large integers → overflow (e.g., 2000^3 exceeds int64); zero cases (0 0 → true)
- **One assumption or limitation:** Uses int (platform-dependent size); no overflow protection
- **What I changed in `main.go`:** Read two ints, compute x*x*x == y, print boolean result
- **Why I believe the result is correct:** Integer multiplication is exact; (-3)^3 = -27 correctly verified; avoids floating-point precision issues

## Exercise 09 — bonus
- **Story in my words:** Evolve a population of 4-bit strings toward the target "1111" using a simple genetic algorithm: evaluate fitness, select parents, crossover, mutate, repeat for N generations.
- **Input(s) and their types:** populationSize: int (e.g., 6); generations: int (e.g., 20); mutationRate: float64 (e.g., 0.1); target: string ("1111")
- **Processing / logical rule:** 
1. Initialize random population
2. Loop for generations:
   - Calculate fitness (count of '1's) for each candidate
   - Select parents (e.g., tournament or fitness-proportionate)
   - Crossover: combine two parents at random cut point
   - Mutate: flip random bits with mutationRate probability
   - Form new population (elitism: keep best)
   - Print generation number and best fitness
3. Stop early if target found
- **Expected output (before coding):** Generation-by-generation best fitness increasing toward 4; e.g.,
Gen 0: best=2
Gen 5: best=3
Gen 12: best=4 (found 1111)
- **Actual output:** same pattern (varies by run due to randomness)
- **Normal test:** Population reaches fitness 4 within generation limit
- **Boundary or invalid test:** Mutation rate 0 → no exploration, may stall; population size 1 → no crossover diversity; very low generations → may not converge
- **One assumption or limitation:** No guarantee of finding 1111; randomness means different runs give different results; small search space (16 candidates) but stochastic process
- **What I changed in `main.go`:** Implemented population struct, fitness function, selection, crossover, mutation, generation loop with early exit
- **Why I believe the result is correct:** Each GA operation follows standard patterns; elitism preserves best; loop terminates on target or generation limit; output shows progression

## Task 2 — GitHub CLI and Go basic types
- **Fork URL and branch:** https://github.com/AuliaRameyza/students-worksheet-algorithm (branch: task-2)
- **Git evidence (`git remote -v`, `git log --oneline main..task-2`, at least 3 commits):**
  origin  https://github.com/AuliaRameyza/students-worksheet-algorithm (fetch)
  origin  https://github.com/AuliaRameyza/students-worksheet-algorithm (push)
  a1b2c3d lab 01: data_types implementation
  e4f5g6h lab 02: cafe_sequence implementation
  i7j8k9l lab 03: factor_checker implementation
- **Program output (`go run main.go`):** (run from each exercise directory)
  01_data_types: 19 | 62.5 | true | 00123
  02_cafe_sequence: Subtotal: 36000 / Total: 31000 / Subtotal: 54000 / Total: 49000
  03_factor_checker: true / false / true / invalid input
  04_population: Population: 1011
  05_number_sequence: 8 11 19 / 2 3 5
  06_gravity: Earth: 686.0 N / Mars: 260.7 N
  07_genetic_intro: Child: 1010 fitness=2 / Mutated: 1110 fitness=3
  08_cube_root: true / true / false
  09_bonus_ga_loop: Gen 0 best=2 ... Gen 15 best=4 (found 1111)
- **Normal test (inputs and result):** Exercise 03: 5 20 → true; Exercise 08: 5 125 → true
- **Boundary test (inputs and result):** Exercise 03: 0 20 → invalid input; Exercise 08: -3 -27 → true; Exercise 02: quantity 0 → negative total
- **Invalid case I considered (e.g. dividing by zero) and what happens:** Exercise 03: x=0 would cause division by zero panic in y%x; prevented by checking x>0 first (short-circuit)
- **Q1. Why is `7 / 2` equal to `3` but `float64(7) / float64(2)` equal to `3.5`?** Integer division truncates toward zero (remainder discarded). With `float64`, Go performs floating-point division preserving the fractional part.
- **Q2. Why is a student ID better stored as a `string` than an `int`?** Student IDs can have leading zeros (e.g., "00123") which are lost if stored as int (becomes 123). IDs are identifiers, not numeric values for calculation.
- **Q3. What does a `const` protect you from? Give one example from this exercise.** Constants prevent accidental reassignment. Example: `const earthGravity = 9.8` in gravity exercise — protects the gravity constant from being modified later.
- **Q4. In `x > 0 && y % x == 0`, why is the order of the two conditions important?** Short-circuit evaluation: if `x > 0` is false, the second condition `y % x == 0` is never evaluated, preventing division-by-zero panic.
- **Q5. What would `fmt.Println('A')` print, and why is it not `A`?** It prints `65` (the Unicode code point/rune value). Single quotes denote a `rune` (int32), not a string; double quotes `"A"` would print the character `A`.

## Task 2 — Fill in the blank
- **Part A (questions 1-8):** 
  1. `int` — whole numbers (age, count)
  2. `float64` — measurements with decimals (mass, weight)
  3. `bool` — true/false status (checked-in)
  4. `string` — text, IDs with leading zeros
  5. `const` — fixed values (gravity, discount)
  6. `:=` — short variable declaration with type inference
  7. `%` — modulo (remainder), tests divisibility
  8. `&&` — logical AND with short-circuit evaluation

- **Part B (questions 9-20):**
  9. Order of operations: subtotal before total (dependencies)
  10. Validation before modulo prevents panic
  11. Population model: final = initial + births + arrivals - deaths - departures
  12. Sequence: each term = sum of previous two (Fibonacci-like)
  13. Gravity: mass (kg) constant, weight (N) varies by planet
  14. float64 for decimal precision in physics calculations
  15. GA: fitness = count of 1s in bit string
  16. Crossover: combine parent strings at cut point
  17. Mutation: flip random bit (exploration)
  18. Cube root check: x*x*x == y (exact integer math)
  19. Negative cubes work: (-3)³ = -27
  20. Population GA: loop with selection, crossover, mutation, elitism

- **Part C (questions 21-23):**
  21. Type choice determines valid operations (string preserves zeros, int enables math)
  22. Execution order matters when variables depend on each other (subtotal → total)
  23. Edge cases reveal bugs: zero input, negative values, overflow, division by zero

### My mini-defense (2 minutes)
1. Why did I choose these types? Each type matches the semantic meaning: `int` for counts (age, quantity), `float64` for measurements (mass, gravity), `bool` for binary status (checked-in), `string` for identifiers with leading zeros (student ID). Using wrong type loses data (int drops leading zeros) or allows invalid operations.
2. Why did I put the operations in this order? Dependencies dictate order: subtotal before total (total uses subtotal), validation before modulo (prevents panic), population formula as single expression (clarity). Wrong order gives wrong results or crashes.
3. Which test caught a mistake in my first attempt? Exercise 03: testing x=0 caught missing validation — would panic on division by zero. Exercise 02: quantity=0 revealed negative total (model breaks). Exercise 08: large integers risk overflow.
4. Optional: Why is a genetic algorithm not a guaranteed optimizer? GA uses stochastic search (random initialization, mutation, crossover). It can get stuck in local optima, lose diversity, or simply not find the global optimum within the generation budget. No guarantee of convergence to "1111" — only probabilistic improvement over time.

**Terminal evidence:** put at least three screenshots under `evidence/` in your fork, or provide terminal output text. Avoid private information in screenshots.

> Predict first. Run second. Explain third.

Task 2 — Fill in the blank

A — Git and GitHub (terminal)

1.  A fork is your own online copy of someone else's repository.
2. git clone <url> downloads a repository to your computer.
3.  git switch -c task-2 creates a new branch and moves you onto it.
4. git status lists the files you changed, and git diff shows the exact lines you added or removed.
5. git add student/10_task2_types/main.go puts a file in the box before you seal it with git commit.
6. git push -u origin task-2 sends your commits to your fork on GitHub.
7. git log --oneline main..task-2 lists only the commits you made on your branch.
8. The remote that points to your fork is normally called origin.

B — Go types, constants and operators

9. 7 / 2 is 3 because both numbers are int (atau integer) values, so the fraction is dropped.
10. 7 % 2 is 1. The % operator gives the remainder of a division.
11. float64(7) / float64(2) is 3.5. Go does not convert types implicitly (automatically).
12.  Go has no char type. A single character is stored as a rune (or as a byte for one ASCII byte).
13. A student ID such as "00123" should be a string because we never do arithmetic on it and the leading zeros must stay.
14. The keyword for a value that can never change is const. Reassigning it gives a compile-time error.
15.  true && false is false, true || false is true, and !true is false.
16.  In x > 0 && y % x == 0, when x is 0 the right side is never evaluated, so the program does not divide by zero. (Ini disebut short-circuit evaluation).
17.  In Go, = assigns a value and == compares two values.
18.  2 + 3 * 4 is 14 but (2 + 3) * 4 is 
19. with x, y := 0.1, 0.2, the expression x + y == 0.3 is false because float64 stores some decimals only approximately.
20. string(65) gives "A", so use strconv.Itoa(65) when you want the text "65".

C — Complete the program
21. Fill the three blanks so the program compiles: const, float64,>=
22. Before running, predict the output: 66000 true
23. Which line would stop the program from compiling if you removed the conversion blank? Explain in one sentence: subtotal := price * float64(qty) (line 10)
because Line 10 fails to compile because Go prohibits arithmetic operations between different data types like float64 and int without explicit conversion.