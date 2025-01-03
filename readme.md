## README – Reconhecimento de Cervejas

Esse projeto foi criado para reconhecer qual é a marca de cerveja em uma imagem que você enviar. Ele tem duas partes:
1. **backend-go** (Go) – Recebe a imagem, fala com o OCR (em Python) e salva no MongoDB.
2. **ocr-service-python** (Flask) – Extrai o texto da lata usando a biblioteca **easyocr** e retorna a marca da cerveja.

---

## Como rodar com Docker
1. **Clonar** ou baixar este repositório.
2. Na pasta raiz, usando o `docker-compose.yml`, é só rodar:
   ```bash
   docker-compose up --build
   ```
   - Ele vai subir:
     - **MongoDB** (para persistência),
     - **backend-go** (porta 8081) 
     - **ocr-service-python** (porta 5001).

Se quiser verificar:
- **Go** rodando: [http://localhost:8081](http://localhost:8081)
- **OCR** rodando: [http://localhost:5001/health](http://localhost:5001/health)

---

## Exemplos de Requisições

1. **Upload de imagem** (no Go):
   ```
   POST http://localhost:8081/upload
   Content-Type: multipart/form-data
   Body (form-data):
   - image (arquivo .jpg ou .png)
   ```
   - Resposta 200 (OK):  
     ```
     Marca identificada: <nome_da_cerveja>
     ```
   - Resposta 404 (Not Found):  
     ```
     Nenhuma marca foi identificada pelo OCR
     ```

2. **OCR Service** (em Python):
   ```
   POST http://localhost:5001/process-image
   Content-Type: multipart/form-data
   Body (form-data):
   - image (arquivo .jpg ou .png)
   ```
   - Resposta 200 (OK) com JSON:  
     ```
     {
       "brand": "AlgumaMarca"
     }
     ```
   - Resposta 404 (Not Found):  
     ```
     {
       "error": "Nenhuma marca identificada"
     }
     ```

---

## Breve Descrição da Arquitetura
- **Go** faz todo o fluxo principal: recebe a imagem, transforma em form-data, chama o **OCR** em Python e salva tudo no **MongoDB**.
- **Python (Flask)** com **easyocr** lê o texto na imagem e tenta mapear para alguma marca conhecida, usando técnicas de similaridade (por ex., difflib).
- **MongoDB** guarda a imagem e o nome da marca, pra saber o que foi reconhecido ao longo do tempo.

As principais tecnologias são:
- **Go**: rápido e ideal para serviços web.
- **MongoDB**: banco NoSQL por ser ideal para armazenar binários e JSON.
- **Python + EasyOCR**: biblioteca de OCR leve, rápida e superior ao a outras libs para esse uso em específico.
- **Docker**: orquestração e conteinerização.

---
