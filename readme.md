## 🍺 Projeto – Reconhecimento de Cervejas (Beer Label Recognition)

Esse projeto tem como objetivo identificar a marca de cervejas a partir de imagens enviadas. Ele é composto por dois microsserviços (um em Go e outro em Python) e um banco de dados MongoDB para armazenar os resultados.

---

## 🚀 Como Rodar o Projeto com Docker

### 1. Clone ou baixe este repositório:
```bash
git clone https://github.com/ftangerino/beer_label_recognizer
cd beer_label_recognizer
```

---

### 2. Build e Execução dos Contêineres:
Execute o seguinte comando na raiz do projeto para construir e subir todos os serviços:
```bash
docker-compose up --build
```

Esse comando irá:
- **Construir** e iniciar:
  - `backend-go` (porta 8081)
  - `ocr-service-python` (porta 5001)
  - `mongo` (porta 27017)
  
Para verificar:
- **Backend Go**: [http://localhost:8081](http://localhost:8081)  
- **OCR Flask**: [http://localhost:5001/health](http://localhost:5001/health)  

---

## 🛠️ Configurando o Docker Buildx (opcional)
Se o buildx não estiver habilitado, crie-o com:
```bash
docker buildx create --name mybuilder --use
```
Depois, rode:
```bash
docker buildx build --platform linux/amd64 -t beer_label_recognizer-backend-go ./backend-go --load
docker buildx build --platform linux/amd64 -t beer_label_recognizer-ocr-service-python ./ocr-service-python --load
```

---

## 🏗️ Inicializando Banco de Dados MongoDB
O MongoDB será automaticamente configurado com um banco `beerdb` e uma coleção `beer_recognition` ao subir o container.  
Caso precise reiniciar manualmente:
```bash
docker exec -it mongodb mongosh beerdb
show collections
```

Arquivo `init.js` usado para inicializar o banco:
```javascript
db = db.getSiblingDB('beerdb');

db.createCollection('beer_recognition');

db.beer_recognition.insertOne({
  brand_name: "Example Beer",
  image: null,
  created_at: new Date()
});
```

---

## 🔄 Fluxo de Trabalho (Arquitetura)

### 1. Microsserviço 1: API em Go
- **Recebe imagens** enviadas para o endpoint `/upload`.
- A imagem é transformada em `multipart/form-data` e enviada para o **serviço OCR** em Flask.
- O resultado (marca identificada) é salvo no **MongoDB**.

### 2. Microsserviço 2: API OCR em Python
- Recebe a imagem e aplica **OCR** com `EasyOCR`.
- Compara o texto extraído com uma lista de marcas de cerveja usando similaridade (`difflib`).
- Retorna a marca encontrada para o Go ou indica que não houve correspondência.

---

## 📡 Endpoints Disponíveis

### 1. Upload de Imagem (Go)
```
POST http://localhost:8081/upload
Content-Type: multipart/form-data
Body (form-data):
- file: <arquivo .jpg ou .png>
```
**Resposta:**
```json
{
  "match": "Lata de Heineken"
}
```

---

### 2. OCR Direto (Flask Python)
```
POST http://localhost:5001/ocr
Content-Type: multipart/form-data
Body (form-data):
- file: <arquivo .jpg ou .png>
```
**Resposta:**
```json
{
  "match": "Lata de Skol"
}
```
---

## 📦 Principais Tecnologias Utilizadas
- **Go** – Para construir uma API performática e robusta.
- **Python (Flask)** – Serviço OCR leve e rápido.
- **PaddleOCR** – Biblioteca de reconhecimento de texto eficiente.
- **MongoDB** – Banco de dados NoSQL usado para persistir resultados.
- **Docker/Docker Compose** – Orquestração de microsserviços e banco de dados.

---