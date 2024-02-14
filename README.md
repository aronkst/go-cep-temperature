# Go CEP Temperature

## Visão Geral

O projeto oferece uma solução avançada e prática para acessar informações climáticas detalhadas utilizando Códigos de Endereçamento Postal (CEPs) como base para a consulta. Simplesmente ao fornecer um CEP através da URL /?cep=CEP, os usuários recebem uma resposta rápida em formato JSON, que inclui as temperaturas atuais expressas nas três principais escalas termométricas: Celsius, Fahrenheit e Kelvin.

## Características

- **Consulta Direta por CEP**: Permite o acesso rápido a informações de temperatura específicas de uma localização, usando apenas o Código de Endereçamento Postal (CEP) como referência.
- **Validação Rigorosa de CEP**: Implementa uma validação para garantir que o CEP fornecido esteja no formato correto, consistindo apenas de números e contendo exatamente 8 caracteres, assegurando a precisão das consultas.
- **Autenticação Livre**: O sistema foi projetado para ser acessível sem a necessidade de tokens de API ou qualquer forma de autenticação, simplificando o acesso às informações climáticas.
- **Respostas em Formato JSON**: Todas as respostas são fornecidas em formato JSON, facilitando a integração com outras aplicações e a manipulação dos dados.
- **Suporte a Múltiplas Unidades de Temperatura**: Fornece temperaturas em Celsius, Fahrenheit e Kelvin, oferecendo flexibilidade para atender às preferências e necessidades específicas dos usuários.
- **Atualizações Climáticas em Tempo Real**: Integra-se com serviços de meteorologia confiáveis para fornecer informações atualizadas e precisas sobre o clima.

## Exemplo de Uso

Para consultar informações climáticas através da linha de comando, você pode usar o `curl`, uma ferramenta poderosa e disponível na maioria dos sistemas operacionais para fazer requisições HTTP. Abaixo estão exemplos práticos de como usar o curl para obter a temperatura com base em um CEP específico.

### Realizando uma Consulta

Para fazer uma consulta, simplesmente substitua CEP pelo código postal desejado na URL. Aqui estão alguns exemplos:

```bash
curl "http://localhost:8080/?cep=01001000"
```

Retorno esperado:

```json
{"temp_C":27.9,"temp_F":82.22,"temp_K":301.05}
```

Neste exemplo, a requisição retorna a temperatura para o CEP 01001000 (um CEP de São Paulo), mostrando a temperatura em Celsius (temp_C), Fahrenheit (temp_F), e Kelvin (temp_K).

```bash
curl "http://localhost:8080/?cep=80210090"
```

Retorno esperado:

```json
{"temp_C":25.3,"temp_F":77.54,"temp_K":298.45}
```

Aqui, a temperatura é retornada para o CEP 80210090, que corresponde a uma localização em Curitiba.

### Como os Dados são Retornados

Os dados são retornados em formato JSON. Cada campo no JSON representa uma medida diferente de temperatura:

- `temp_C`: Temperatura em graus Celsius.
- `temp_F`: Temperatura em graus Fahrenheit.
- `temp_K`: Temperatura em Kelvin.

## Desenvolvimento

Desenvolvi o projeto centrando-me no uso coordenado de várias APIs externas para entregar informações climáticas precisas baseadas em um Código de Endereçamento Postal (CEP) fornecido. O processo para obter essas informações segue uma sequência lógica de passos, onde cada um faz uso de uma API específica para alcançar o objetivo final. A seguir, detalho cada etapa e como cada API é empregada:

### Busca de Endereço pelo CEP com viacep.com.br

O ponto de partida envolve coletar informações detalhadas sobre o endereço associado ao CEP fornecido. Utilizo a API do ViaCEP para essa finalidade. Ao receber um CEP válido, faço uma requisição à API do ViaCEP, que me retorna dados como logradouro, bairro, cidade e estado correspondente ao CEP. Esses dados são essenciais para determinar a localização geográfica exata a ser usada nas consultas climáticas que se seguem.

### Busca de Longitude e Latitude com nominatim.openstreetmap.org

Tendo em mãos os dados do endereço, o próximo passo é convertê-los em coordenadas geográficas (latitude e longitude). Para isso, recorro à API do Nominatim, parte do projeto OpenStreetMap. Esta API permite que eu envie detalhes de localização, como cidade e estado, e receba em retorno as coordenadas geográficas precisas dessa localização. Esta conversão é vital para assegurar a acurácia das consultas climáticas baseadas em coordenadas.

### Busca de Temperatura com api.open-meteo.com ou wttr.in

Dispondo das coordenadas geográficas, avanço para a fase de consulta das condições climáticas atuais. Neste ponto, o processo se divide, dependendo da disponibilidade das coordenadas:

- Com dados de longitude e latitude disponíveis, recorro à API Open-Meteo. Esta API possibilita consultas climáticas detalhadas baseadas nas coordenadas geográficas, proporcionando dados de temperatura exatos para a localização desejada.
- Sem dados de longitude e latitude, utilizo a API do wttr.in. Esta API fornece informações climáticas baseadas em nomes de localização (como cidades), derivados dos dados obtidos via ViaCEP. Embora este método possa não ser tão preciso quanto a consulta por coordenadas, ainda assim oferece uma estimativa útil das condições climáticas.

## Tratamentos de Erros

Implementei tratamentos de erros em cada etapa para assegurar que o sistema possa lidar de forma adequada com cenários como CEPs inválidos, falhas na obtenção de coordenadas ou erros nas respostas das APIs climáticas.

## Testes Unitários

Uma parte integral do desenvolvimento deste projeto envolve a implementação de testes unitários abrangentes, garantindo a confiabilidade e a robustez de cada funcionalidade oferecida pela aplicação. A abordagem adotada para os testes segue as melhores práticas de desenvolvimento de software, focando na validação de cada componente isoladamente para assegurar seu correto funcionamento em diversos cenários.

### Cobertura dos Testes

Os testes unitários cobrem uma ampla gama de casos de uso e cenários de erro, incluindo, mas não se limitando a:

- Validação de CEPs: Testes para assegurar que apenas CEPs válidos e no formato correto são aceitos, e que as mensagens de erro adequadas são retornadas para CEPs inválidos ou formatados incorretamente.
- Consulta a APIs Externas: Testes para verificar a interação correta com as APIs externas usadas para obter informações de endereço, coordenadas geográficas e dados climáticos. Isso inclui simular respostas das APIs para testar o manejo adequado de dados e erros.
- Conversão de Unidades de Temperatura: Testes que validam a precisão das conversões de temperatura entre Celsius, Fahrenheit e Kelvin, garantindo que os cálculos estejam corretos.
- Tratamento de Erros: Testes específicos para verificar a robustez do sistema ao enfrentar erros durante a consulta de informações, incluindo falhas de rede, erros nas APIs externas e dados inesperados.

## Ambiente de Desenvolvimento

### Pré-requisitos

Antes de começar, certifique-se de que você tem o Docker e o Docker Compose instalados em sua máquina. Caso não tenha, você pode baixar e instalar a partir dos seguintes links:

- Docker: https://docs.docker.com/get-docker/

### Passo a Passo

1. Clonar o Repositório

Primeiro, clone o repositório do projeto para a sua máquina local. Abra um terminal e execute o comando:

```bash
git clone https://github.com/aronkst/go-cep-temperature.git
```

2. Navegar até o Diretório do Projeto

Após clonar o repositório, navegue até o diretório do projeto utilizando o comando cd:

```bash
cd go-cep-temperature
```

3. Construir e Executar o Projeto com Docker Compose

No diretório do projeto, execute o seguinte comando para construir e iniciar o projeto utilizando o Docker Compose:

```bash
docker compose up --build
```

Este comando irá construir a imagem Docker do projeto e iniciar o container. O parâmetro `--build` garante que a imagem seja reconstruída caso haja mudanças no Dockerfile ou nas dependências do projeto.

4. Acessar o Projeto

Com o container rodando, você pode acessar o projeto através do navegador ou utilizando ferramentas como curl, apontando para http://localhost:8080/?cep=CEP, substituindo CEP pelo código postal desejado.

### Exemplo de Comando curl

Para testar se o projeto está rodando corretamente, você pode usar o seguinte comando curl em um novo terminal:

```bash
curl "http://localhost:8080/?cep=01001000"
```

Você deverá receber uma resposta em JSON com as temperaturas em Celsius, Fahrenheit e Kelvin.

### Encerrando o Projeto

Para encerrar o projeto e parar o container do Docker, volte ao terminal onde o Docker Compose está rodando e pressione Ctrl+C. Para remover os containers criados pelo Docker Compose, execute:

```bash
docker compose down
```

Esses passos garantem que você possa executar o projeto localmente de maneira fácil e rápida, utilizando o Docker e Docker Compose para gerenciar o ambiente de execução.

## Nota Sobre Não Fazer o Deploy no Google Cloud Run

Decidi não realizar o deploy do meu projeto no Google Cloud Run e gostaria de compartilhar os motivos por trás dessa decisão:

### Excedi o Limite do Free Tier do Google Cloud

A minha intenção era aproveitar as vantagens do Free Tier oferecido pelo Google Cloud Run. Contudo, já havia ultrapassado esses limites de uso gratuito no passado. Isso significaria que qualquer atividade futura de hospedagem resultaria em custos adicionais, algo que não estava nos meus planos iniciais.

### Preocupações com os Custos Operacionais

O Google Cloud Run é, sem dúvida, uma plataforma excelente que oferece escalabilidade e robustez. No entanto, a possibilidade de custos operacionais crescentes, especialmente quando o uso excede os limites do Free Tier, me levou a reavaliar essa opção.

### Priorização do Desenvolvimento e Teste Local

Optei por focar meus esforços no desenvolvimento local, o que me oferece maior controle e flexibilidade durante a fase de desenvolvimento e testes do projeto.
