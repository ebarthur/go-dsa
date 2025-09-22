The **order of complexities** refers to the hierarchy of common time complexities in terms of their growth rates. This helps you understand which complexity dominates another as the input size (`n`) grows larger.

Here’s the order of common time complexities from **fastest (best)** to **slowest (worst)**:

---

### Order of Complexities:
1. **O(1)** (Constant Time):
   - Operations that take the same amount of time regardless of the input size.
   - Example: Accessing an element in an array.

2. **O(log n)** (Logarithmic Time):
   - Operations that reduce the problem size by a factor (e.g., binary search).
   - Example: Searching in a sorted array using binary search.

3. **O(n)** (Linear Time):
   - Operations that process each element of the input once.
   - Example: Iterating through an array.

4. **O(n log n)** (Linearithmic Time):
   - Operations that involve dividing the input and processing each division (e.g., sorting algorithms like merge sort or quicksort).
   - Example: Sorting an array.

5. **O(n²)** (Quadratic Time):
   - Operations that involve nested loops, where each loop iterates over the input.
   - Example: Comparing all pairs of elements in an array.

6. **O(n³)** (Cubic Time):
   - Operations with three nested loops, where each loop iterates over the input.
   - Example: Comparing all triplets in an array.

7. **O(2ⁿ)** (Exponential Time):
   - Operations that grow exponentially with the input size.
   - Example: Generating all subsets of a set.

8. **O(n!)** (Factorial Time):
   - Operations that involve permutations of the input.
   - Example: Solving the traveling salesman problem using brute force.

---

### What Dominates What:
- **Lower complexities dominate higher complexities**:
  - **O(1) < O(log n) < O(n) < O(n log n) < O(n²) < O(n³) < O(2ⁿ) < O(n!)**
- For example:
  - **O(n log n)** dominates **O(n)** because the additional logarithmic factor grows faster as `n` increases.
  - **O(n²)** dominates **O(n log n)** because the quadratic growth rate is faster than linearithmic growth.
