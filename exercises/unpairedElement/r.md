A non-empty array A consisting of N integers is given. The array contains an odd number of elements, and each element of the array can be paired with another element that has the same value, except for one element that is left unpaired.

For example, in array A such that:

  A[0] = 9  A[1] = 3  A[2] = 9
  A[3] = 3  A[4] = 9  A[5] = 7
  A[6] = 9
the elements at indexes 0 and 2 have value 9,
the elements at indexes 1 and 3 have value 3,
the elements at indexes 4 and 6 have value 9,
the element at index 5 has value 7 and is unpaired.
Write a function:

func Solution(A []int) int

that, given an array A consisting of N integers fulfilling the above conditions, returns the value of the unpaired element.

For example, given array A such that:

  A[0] = 9  A[1] = 3  A[2] = 9
  A[3] = 3  A[4] = 9  A[5] = 7
  A[6] = 9
the function should return 7, as explained in the example above.


=============================================================

XOR
A operação XOR, ou exclusivo OR, é uma operação bit a bit que retorna verdadeiro (1) se os bits comparados forem diferentes e falso (0) se os bits forem iguais. Quando aplicada a números inteiros, ela compara os bits de cada número e retorna um novo número com os bits resultantes da comparação.

No caso do exemplo dado:
9 ^ 3 ^ 9 ^ 3 ^ 9 ^ 7 ^ 9

Vamos explicar passo a passo:

9 ^ 3:
9 em binário é 1001
3 em binário é 0011
Ao aplicar XOR bit a bit, obtemos 1010, que em decimal é 10.

10 ^ 9:
10 em binário é 1010
9 em binário é 1001
Ao aplicar XOR bit a bit, obtemos 0011, que em decimal é 3.

3 ^ 3:
Como qualquer número XOR consigo mesmo resulta em zero, 3 ^ 3 será zero.

0 ^ 9:
Zero XOR qualquer número resultará nesse mesmo número.

9 ^ 7:
9 em binário é 1001
7 em binário é 0111
Ao aplicar XOR bit a bit, obtemos 1110, que em decimal é 14.

14 ^ 9:
14 em binário é 1110
9 em binário é 1001
Ao aplicar XOR bit a bit, obtemos 0111, que em decimal é 7.

Portanto, o resultado final do XOR de todos os elementos do array é 7. Isso significa que 7 é o elemento sem par no array fornecido.