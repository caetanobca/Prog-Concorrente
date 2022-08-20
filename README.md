# Prog-Concorrente
Exercícios realizadas para a disciplina de programação concorrente

## 5. Fork-sleep-join
    Crie um programa que recebe um número inteiro n como argumento e cria n goroutines. Cada uma dessas goroutines deve dormir por um tempo aleatório de no máximo 5 segundos. A main-goroutine deve esperar todas as goroutines filhas terminarem de executar para em seguida escrever na saída padrão o valor de n.

## 6. Two-phase sleep 
    Crie um programa que recebe um número inteiro n como argumento e cria n goroutines. Cada uma dessas goroutines deve dormir por um tempo aleatório de no máximo 5 segundos. Depois que  acordar, cada thread deve sortear um outro número aleatório s (entre 0 e 10).  Somente depois de todas as n goroutines terminarem suas escolhas (ou seja, ao fim da primeira fase), começamos a segunda fase. Nesta segunda fase, a n-ésima goroutine criada deve dormir pelo tempo s escolhido pela goroutine n - 1 (faça a contagem de maneira modular, ou seja, a primeira goroutine dorme conforme o número sorteado pela última).

## 7. Pipeline 
    Crie um programa organizado como um pipeline de goroutines. Esse programa deve receber como argumento um caminho absoluto para um diretório. Uma goroutine deve navegar na árvore que tem como raiz o diretório passado como argumento. Essa goroutine deve passar para uma próxima goroutine do pipeline o nome dos arquivos encontrados na busca dos diretórios, ou seja, ignore os diretórios. Esta segunda goroutine deve ler o primeiro byte de conteúdo de cada um desses arquivos e escrever na saída padrão o nome dos arquivos que tem esse valor do primeiro byte sendo par.

## 8. Bid
    Implemente a função handle seguindo a assinatura abaixo (e qualquer outra função utilitária que julgue necessárias, incluindo a main).

    func handle(nServers int) chan Bid;

    Sua implementação da função handle deve usar a API com as funções itemsStream e bid, considerando as seguintes restrições:
    Você deve criar um número de nServers goroutines para executar as funções bid; 
    A execução da função bid pode demorar muito tempo;
    Um Item recebido no canal retornado pela função itemsStream deve ser enviado somente para um servidor. Ou seja, você deve executar somente uma vez a função bid para um determinado Item;
    Você deve evitar ociosidade dos servidores de bid. Ou seja, se um Item estiver disponível no canal itemsStream, deve ser enviado para processamento por um servidor de bid;
    Na medida em que novos valores de Bid seja retornados, esse valores devem ser enviados no canal de resposta da função handle;
    Após o canal retornado por itemsStream ter sido fechado e não houver mais conteúdo a ser consumido, as goroutines que executam bid devem terminar;
    Após todo o conteúdo de bid ter sido gerado e enviado para o canal de resposta da função handle, esse canal de resposta deve ser fechado.

## 9. Bid timeout
    Implemente uma variação da função handle descrita acima, seguindo a nova assinatura abaixo:

    func handle(nServers int, timeoutSecs int) chan Bid;

    Nesta nova implementação, você não quer esperar mais do que timeoutSecs segundos executando a função bid. Em caso desse tempo máximo para execução ter sido ultrapassado, um bid especial deve ser enviado para o canal retornado pela função handle. Esse bid pode ser criado assim:
