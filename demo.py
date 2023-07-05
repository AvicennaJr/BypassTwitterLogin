from fastapi import FastAPI
from fastapi.responses import HTMLResponse
import requests

app = FastAPI()


@app.get("/{url:path}", response_class=HTMLResponse)
def proxy(url: str):
    response = requests.get(f'https://twitter.com/{url}', headers={'User-Agent': 'Googlebot'})
    return response.text
