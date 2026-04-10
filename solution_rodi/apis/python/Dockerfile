FROM python:3.12-alpine
WORKDIR /app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY calculator.py ./
EXPOSE 4000
CMD ["python", "calculator.py"]

