### README.md

# Math Skills

This repository contains **Math Skills**, a program that calculates essential statistical measures: **average**, **median**, **variance**, and **standard deviation**. The program reads numerical data from a file and outputs the computed statistics.  

This project was independently implemented by me and served as an opportunity to improve my programming and mathematical skills.

---

## Features

The program computes the following statistical measures:  

1. **Average**  
   The arithmetic mean of the dataset.  
   - Example: For values `1, 2, 3, 4, 5`, the average is `3`.

2. **Median**  
   The middle value of the dataset when sorted.  
   - Example: For `1, 3, 5, 7, 9`, the median is `5`.  
   If the dataset has an even number of elements, the median is the average of the two middle values.  

3. **Variance**  
   The measure of dispersion, calculated as the average of the squared differences from the mean.  

4. **Standard Deviation**  
   The square root of the variance, indicating how much the values deviate from the mean.  

---

## Usage

### Input File Format
- The input file must contain one numeric value per line.  
  Example (`data.txt`):  
  ```
  189
  113
  121
  114
  145
  110
  ```

### Running the Program
To execute the program, provide the path to the data file as an argument.  
```bash
$ go run math-skills.go <data_file>
```

### Example Output
For a file (`data.txt`) containing:  
```
10
20
30
40
50
```

The output will be:  
```
Average: 30
Median: 30
Variance: 200
Standard Deviation: 14
```

---

## What I Learned
- **File Operations**: Reading data from files and handling potential errors.  
- **Mathematical Concepts**: Implementing core statistical functions programmatically.  
- **Sorting Algorithms**: Using sorting to determine the median efficiently.  
- **Precision and Rounding**: Managing numerical computations and rounding results to integers.  

---

## Future Enhancements
- Add support for floating-point precision outputs.  
- Implement additional statistical measures like mode and range.  
- Handle datasets with missing or invalid data gracefully.  

---

## Acknowledgments
This project was my first step into blending programming with statistical analysis. It was a rewarding challenge that enhanced both my coding and mathematical problem-solving abilities.
