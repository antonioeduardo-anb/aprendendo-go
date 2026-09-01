# Aprendendo Go

Repositório destinado aos meus estudos da linguagem de programação **Go (Golang)**.

## Estrutura

```text
.
├── ambiente-de-estudos
│   └── Containerfile
└── README.md
```

## Ambiente de estudos

A pasta `ambiente-de-estudos` contém o `Containerfile` utilizado para criar a imagem do ambiente de estudos em Go.

O ambiente utiliza a imagem oficial `golang:1.25` e adiciona algumas ferramentas úteis para os estudos, como Git e ferramentas de compilação.

### Criar a imagem

Entre na pasta do ambiente:

```bash
cd ambiente-de-estudos
```

Construa a imagem:

```bash
podman build -t localhost/go-estudos:1.0 .
```

### Criar um container temporário

Para iniciar um ambiente temporário baseado nessa imagem:

```bash
podman run --rm -it localhost/go-estudos:1.0
```

O parâmetro `--rm` faz com que o container seja removido automaticamente quando ele for encerrado.

O parâmetro `-it` permite utilizar o terminal interativamente.

Dentro do container, é possível verificar o ambiente:

```bash
go version
git --version
```

### Trabalhar com os arquivos do projeto

O container é temporário, portanto os arquivos criados dentro dele serão perdidos quando o container for removido.

Para estudar utilizando os arquivos do repositório no computador, monte a pasta atual no container:

```bash
podman run --rm -it \
    --userns=keep-id \
    -v "$PWD:/workspace:Z" \
    localhost/go-estudos:1.0
```

Dessa forma, `/workspace` dentro do container corresponde à pasta `ambiente-de-estudos` no sistema hospedeiro.

## Objetivo

Este repositório acompanha meu aprendizado de Go, contendo exemplos, exercícios e pequenos projetos desenvolvidos durante os estudos. 🚀
