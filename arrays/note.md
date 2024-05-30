#Array and slices

Slices in Go are dynamic collections of elements of the same type, providing a more flexible alternative to arrays.

## Key points

- **Declaration:**  
    `mySlice := []int{1, 2, 3}` (No need to specify size)
- **Slicing:**`mySlice[low:high]` (Extracts a portion, omitting values defaults to start/end)
- **Making slices:**`make([]int, length, capacity)` (Creates slice with defined length and optional capacity)
- **Appending:**`append(slice, element)` (Adds to the end, creating new slice if needed)
- **Length and Capacity:**
    - `len(slice)` (Number of elements)
    - `cap(slice)` (Maximum capacity)

## Important Note

Modifying a slice can affect the underlying array due to how they are implemented. To prevent unintended changes, create a copy of the slice before modifying it.