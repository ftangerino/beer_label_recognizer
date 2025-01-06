FROM python:3.10-slim

RUN apt update && apt install -y \
    libgl1 \
    libgomp1 \
    libglib2.0-0 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY . .

RUN pip install --no-cache-dir \
    paddleocr \
    paddlepaddle \
    opencv-python-headless \
    flask pillow

RUN rm -rf ~/.cache/pip

# Apontar corretamente o FLASK_APP
ENV FLASK_APP=app.beer_ocr

EXPOSE 5000
CMD ["flask", "run", "--host=0.0.0.0", "--port=5000"]
