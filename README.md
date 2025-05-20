
 ███╗   ███╗███████╗████████╗ █████╗ ████████╗██████╗  ██████╗ ███╗   ██╗
 ████╗ ████║██╔════╝╚══██╔══╝██╔══██╗╚══██╔══╝██╔══██╗██╔═══██╗████╗  ██║
 ██╔████╔██║█████╗     ██║   ███████║   ██║   ██████╔╝██║   ██║██╔██╗ ██║
 ██║╚██╔╝██║██╔══╝     ██║   ██╔══██║   ██║   ██╔══██╗██║   ██║██║╚██╗██║
 ██║ ╚═╝ ██║███████╗   ██║   ██║  ██║   ██║   ██║  ██║╚██████╔╝██║ ╚████║
 ╚═╝     ╚═╝╚══════╝   ╚═╝   ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═══╝

Este projeto apresenta um script automatizado desenvolvido em Go para reconhecimento de subdomínios, verificação de disponibilidade HTTP, coleta de URLs e detecção automatizada de vulnerabilidades do tipo Cross-Site Scripting (XSS).

## 📌 Funcionalidades:
- Enumerar subdomínios utilizando o **subfinder**
- Verificar a disponibilidade HTTP com **httpx**
- Extrair parâmetros de URLs com **katana**
- Descobrir URLs adicionais com **urlfinder**
- Detecção automatizada de vetores XSS com **dalfox**

## 🛠️ Requisitos
- Go instalado ([Tutorial oficial](https://golang.org/doc/install))
- Subfinder, Httpx, Katana, Urlfinder e Dalfox ([Ferramentas do ProjectDiscovery](https://github.com/projectdiscovery))

## 🚀 Como usar?
Clone o repositório:

```bash
git clone https://github.com/AlexSilva-SNO/metatron.git
cd metatron
```

Instale as dependências necessárias:

```bash
go mod init metatron
go mod tidy
```

Execute o script com o domínio desejado:

```bash
go run metatron.go --domain exemplo.com
```
## 📋 Exemplo de resultados:
Os resultados serão salvos na pasta results/exemplo.com, incluindo:

Subdomínios encontrados

URLs analisadas

Vetores XSS detectados

Logs completos da execução

## 📚 Referências Acadêmicas:
Este projeto foi desenvolvido como parte do meu Trabalho de Conclusão de Curso (TCC), com foco na segurança web e automação.

## 📝 Licença:
Este projeto é licenciado sob a licença MIT - veja o arquivo LICENSE para detalhes.

## 📞 Contato:
Alex Patrik da Silva – alex.silva1@unemat.br
