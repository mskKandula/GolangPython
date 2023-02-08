FROM python:3.10.5

# Create app directory
WORKDIR /app/winevarieties/python

# Install app dependencies
COPY ./python/requirements.txt .

RUN pip install -r requirements.txt

# Copy code files
COPY ./python/ .

COPY ./data/ /app/winevarieties/data

EXPOSE 9000

CMD [ "python", "server.py" ]