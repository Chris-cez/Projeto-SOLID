# Projeto-SOLID

No presente momento desse primeiro commit do projeto, ele não obedece nenhum dos principios SOLID, Demeter ou composição ao invés de herança. Ele é um código de implementação simples no entanto de dificil evolução. No presente momento ele apenas separa funções CRUD de main e mais nada.

Princípio da Responsabilidade Única (SRP)
O princípio da Responsabilidade Única (SRP) é um dos cinco princípios SOLID, que são um conjunto de práticas de design de software que visam melhorar a qualidade do código e torná-lo mais fácil de manter e testar.

O SRP afirma que uma classe ou módulo deve ter apenas uma única responsabilidade. Isso significa que a classe deve ter um único propósito bem definido e não deve ser responsável por muitas tarefas diferentes.

Princípio Aberto/Fechado (OCP)
O Princípio Aberto/Fechado (OCP) é um dos cinco princípios SOLID, que são um conjunto de práticas de design de software que visam melhorar a qualidade do código e torná-lo mais fácil de manter e testar.

O OCP afirma que as entidades de software (classes, módulos, funções, etc.) devem estar abertas para extensão, mas fechadas para modificação. Isso significa que você deve ser capaz de adicionar novas funcionalidades ao seu código sem precisar modificar o código existente.

Princípio da Substituição de Liskov (LSP)
O Princípio da Substituição de Liskov (LSP) é um dos cinco princípios SOLID, que são um conjunto de práticas de design de software que visam melhorar a qualidade do código e torná-lo mais fácil de manter e testar.

O LSP afirma que objetos de uma subclasse devem ser substituíveis por objetos de sua superclasse sem que isso cause problemas no comportamento do programa. Isso significa que, se você tem uma função que espera um objeto de uma superclasse, você deve ser capaz de passar um objeto de uma subclasse para essa função sem que a função se comporte de forma inesperada.

Princípio da Segregação de Interfaces (ISP)
O Princípio da Segregação de Interfaces (ISP) é um dos cinco princípios SOLID, que são um conjunto de práticas de design de software que visam melhorar a qualidade do código e torná-lo mais fácil de manter e testar.

O ISP afirma que as interfaces devem ser pequenas e específicas, em vez de grandes e genéricas. Isso significa que você deve criar interfaces separadas para diferentes funcionalidades, em vez de criar uma única interface grande que abrange todas as funcionalidades.

Princípio da Inversão de Dependência (DIP)
O Princípio da Inversão de Dependência (DIP) é um dos cinco princípios SOLID, que são um conjunto de práticas de design de software que visam melhorar a qualidade do código e torná-lo mais fácil de manter e testar.

O DIP afirma que as classes de alto nível não devem depender de classes de baixo nível. Em vez disso, ambas as classes devem depender de abstrações. As abstrações não devem depender de detalhes; os detalhes devem depender das abstrações.

Princípio de Demeter (LoD)
O Princípio de Demeter (LoD), também conhecido como Lei de Demeter, é um princípio de design de software que afirma que um objeto deve ter conhecimento limitado sobre outros objetos. Em outras palavras, um objeto só deve interagir diretamente com seus objetos amigos (objetos que ele contém ou que estão no mesmo nível de abstração) e não deve interagir diretamente com objetos distantes (objetos que estão em um nível de abstração diferente).

Como esse projeto fere cada um desses principios? 

S (Single Responsibility Principle): A função main tem muitas responsabilidades, incluindo abrir o arquivo, ler os dados, converter os dados, imprimir os dados e realizar operações CRUD.

O (Open/Closed Principle): O código é difícil de estender para suportar novos formatos de arquivos ou novas operações de CRUD. Por exemplo, se você quiser adicionar uma nova operação de CRUD, como "delete", precisará modificar várias funções.

L (Liskov Substitution Principle): O código não permite que subclasses de Person sejam usadas no lugar de Person. Por exemplo, se você criar uma subclasse de Person chamada Employee com campos adicionais, como "cargo", não poderá usar objetos Employee nas funções CRUD.

I (Interface Segregation Principle): O código não separa as interfaces para diferentes operações de CRUD. Por exemplo, a função readPeople é usada para ler todos os registros e também para ler um registro específico.

D (Dependency Inversion Principle): O código depende diretamente de arquivos CSV, em vez de depender de abstrações. Por exemplo, a função readPeople abre o arquivo CSV diretamente, em vez de usar uma interface para abstrair a operação de leitura de arquivos.

Princípio Demeter: O código viola o princípio Demeter porque as funções CRUD acessam diretamente os campos de objetos Person. Por exemplo, a função createPerson acessa os campos ID, Name e Age do objeto person.
