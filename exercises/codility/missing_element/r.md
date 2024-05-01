O problema nos fornece um array A contendo N diferentes inteiros no intervalo de [1..(𝑁+1)], o que significa que exatamente um elemento está faltando.

Quando temos uma sequência de inteiros consecutivos de 1 a N+1, podemos calcular a soma desses inteiros usando a fórmula da soma de uma progressão aritmética.

A soma dos N primeiros números inteiros é dada por N×(N+1)/2. Então, a soma esperada dos números em A seria (N+1)×(N+2)/2, porque estamos incluindo o número faltante.

Se subtrairmos a soma real dos números em A da soma esperada, obteremos o número que está faltando.
