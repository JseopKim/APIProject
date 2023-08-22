from bs4 import BeautifulSoup
import requests

url = "https://kream.co.kr/search"
response = requests.get(url)
html = response.text

soup = BeautifulSoup(html, "html.parser")

content = soup.find("div", class_ = "product_card")

for a in content:
  print(a.text)
