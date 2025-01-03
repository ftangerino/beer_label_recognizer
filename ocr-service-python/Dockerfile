# Build the application and install dependencies
FROM python:3.9-slim AS builder

# Instala apenas o essencial (reduzindo a camada de pacotes desnecessários)
RUN apt-get update && apt-get install -y --no-install-recommends \
    libglib2.0-0 \
    libsm6 \
    libxrender1 \
    libxext6 \
    curl \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY requirements.txt .

# Instala pacotes de forma isolada para não levar o cache para a imagem final
RUN pip install --prefix=/install -r requirements.txt \
    --extra-index-url https://download.pytorch.org/whl/cpu \
    torch==2.5.1+cpu \
 && rm -rf ~/.cache/pip  # Remover cache do pip para economizar espaço

# Final runtime image
FROM python:3.9-slim

WORKDIR /app
COPY --from=builder /install /usr/local
COPY . .

EXPOSE 5001
ENV FLASK_ENV=production
ENV PYTHONPATH=/app

CMD ["python", "app/app.py"]
