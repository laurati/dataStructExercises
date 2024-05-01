A one-day-long training session will be conducted twice during the next 10 days. There are N employees (numbered from 0 to N-1) willing to attend it.Each employee has provided a list of which of the next 10 days they are able to participate in the training. The employees' preferences are represented as an array of strings. E[K] is a string consisting of digits ('0'-'9') representing the days the K-th employee is able to attend the training.

The dates during which the training will take place are yet to be scheduled. What is the maximum number of employees that can attend during at least one of the two scheduled days?

Write a function:
func Solution(E []string) int

that, given an array E consisting of N strings denoting the available days for each employee, will return the maximum number of employees that can attend during at least one of the two scheduled days.

Given E = {"039", "4", "14", "32", "", "34", "7"}, the answer is 5. It can be achieved for example by running training on days 3 and 4. This way employees number 0,1,2,3 and 5 will attend the training.