from selenium import webdriver
from webdriver_manager.chrome import ChromeDriverManager
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By

browser = webdriver.Chrome(ChromeDriverManager().install())
browser.get('https://forms.gle/TNNBtfCdgkpWCZA16')
browser.implicitly_wait(120)
angkatan = '2019'
insert = browser.find_element(By.CSS_SELECTOR, ".AgroKb .zHQkBf")
insert.send_keys(angkatan)
# browser.find_element(By.CSS_SELECTOR, 'div[data-value="Pernah"]')