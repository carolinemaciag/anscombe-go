<<<<<<< HEAD
# anscombe-go
=======
# Anscombe's Quartet Regression Analysis

This project uses **Anscombe's Quartet** dataset to perform a linear regression analysis- utilizing both **Go**,and **Python**. The goal is using both languages to compare the results of the linear regression, including the slopes and intercepts for the datasets.

## Overview

This repository includes code that performs linear regression analysis on Anscombe's Quartet datasets. Anscombe’s Quartet contains four datasets that have nearly identical simple descriptive statistics but very different distributions and appearances. The regression is performed using:

- **Go** for the main application logic.
- **Python** for comparison and additional analysis.

The the following python program example was used in the comparison, with only a few adjustments made to the code: 

miller-mtpa-chapter-1-program.py

The linear regression models calculate the **slope** and **intercept** for each dataset, and execution times are also deployed as a comparison.

## Installation

To install and use the project, clone the repository to your local machine:

```bash
git clone https://github.com/carolinemaciag/anscombe-go.git
```

### Go Dependencies

For Go, make sure you have Go installed on your system (version 1.16 or above).

- Download Go from the official website: [https://golang.org/dl/](https://golang.org/dl/)

### Python Dependencies

For Python, you will need the following libraries:

- **pandas**
- **numpy**
- **matplotlib**
- **statsmodels**

You can install them using `pip`:

```bash
pip install pandas numpy matplotlib statsmodels
```

### Running the Program

To run the program for **Go**, use the following:

```bash
go run main.go
```

For the **Python regression**, it is easiest to run the program in a Jupyter notebook.
>>>>>>> d5e29a6 (Initial commit of Go regression project)
