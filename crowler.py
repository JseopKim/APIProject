import time
#? Selenium 웹 드라이버를 사용하기 위해 webdriver 모듈을 가져오기
from selenium import webdriver
#? Chrome 웹 브라우저 옵션을 설정하기 위해 Options 클래스를 가져오기
from selenium.webdriver.chrome.options import Options

#* Chrome 웹 브라우저의 옵션을 설정하기 위한 Options 객체를 생성
options = Options()

#* 메서드를 사용하여 "detach" 옵션을 설정합니다. 이 옵션은 Chrome 브라우저 창이 자동으로 종료되지 않고 열려있는 상태로 유지되도록 합니다.
options.add_experimental_option("detach", True)

#* 생성자를 호출하여 Chrome 웹 드라이버를 인스턴스화합니다. 앞서 설정한 옵션을 options 매개변수로 전달하여 Chrome 브라우저를 열 때 이 옵션을 적용합니다.
driver = webdriver.Chrome(options=options)

url = "https://naver.com"

driver.get(url)