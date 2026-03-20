import boto3
import json
import psycopg2
import os

def handler(event, context):

    client = boto3.client('secretsmanager')

    secret = client.get_secret_value(
        SecretId="DBSecret"
    )

    secret_dict = json.loads(secret['SecretString'])

    conn = psycopg2.connect(
        host=os.environ['DB_HOST'],
        user=secret_dict['username'],
        password=secret_dict['password'],
        database="postgres"
    )

    print("DB connection successful")
