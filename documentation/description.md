- Race condition acontece quando: Race conditions acontecem quando threads dependem de qualquer estado que esteja compartilhado da aplicação, podendo então gerar resultados inesperados como resultado
- No contexto de multitarefa preemptiva, podemos afirmar que: Que o processo que está sendo executado é gerenciado por um "scheduler" que determina até quando ele poderá ser manter em execução. 
- Qual o principal objetivo do "Mutex"?  Bloquear o estado de alguma informação compartilhada durante sua execução, evitando assim race conditions. 
- O que acontece quando uma thread deixa de ser executada dando espaço a outra?  Context switch 
- Qual a diferença entre concorrência e paralelismo? Quando tarefas estão rodando em paralelo significa que realmente elas estão sendo executadas simultaneamente (por diferentes núcleos computacionais), já tarefas concorrentes significam que cada parte de uma tarefa é executada enquanto a outra aguarda. 
- O que é dead-lock?É a situação onde dois ou mais processos ficam impedidos de continuar suas execuções - ou seja, ficam bloqueados, esperando uns pelos outros. 
- No contexto de multitarefa cooperativa, podemos afirmar que: Que o processo que está sendo executado pode se manter em execução impedindo que outro processo seja iniciado.
- Qual o objetivo de um channel na linguagem go?Realizar a comunicação entre go routines 


Go.mod -> Arquivo que gerencia as dependencias do projeto                  