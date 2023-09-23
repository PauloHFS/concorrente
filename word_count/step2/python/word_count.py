import os
import sys
from threading import *
count = 0
lock = Lock()

#Recebe um conteudo, fatia seus paragrafos/frases e conta os pedaços/palavras,retorna o total de palavras.
def wc(content):
    return len(content.split())

#Recebe uma arquivo, tenta abrir um arquivo e se sucessedivel chama o método acima para o conteudo desse arquivo, retorna o total de palavras.
def wc_file(filename):
    try:
        with open(filename, 'r', encoding='latin-1') as f:
            file_content = f.read()
        return wc(file_content)
    except FileNotFoundError:
        return 0

#Recebe um diretorio, passa por todos os elementos dentro desse direitorio,se o elemento é um diretorio:recursão, se o elemento é um arquivo, chama o método acima.
def wc_dir(dir_path):
    global count
    global lock
    countI = 0
    for filename in os.listdir(dir_path):
        filepath = os.path.join(dir_path, filename)
        if os.path.isfile(filepath):
            countI += wc_file(filepath)
    with lock:
        count += countI
        #elif os.path.isdir(filepath):
            #count += wc_dir(filepath)  # Chamada recursiva para diretórios

def pick_dir(dir_path):
    diretorios = []
    for filename in os.listdir(dir_path):
        filepath = os.path.join(dir_path, filename)
        if os.path.isdir(filepath):
            diretorios.append(filepath)
    return diretorios
# Minha ideia primeiramente é que para cada diretorio encontrado ele crie uma task para ser executado por uma thread mas estou vendo que por arquivo deve ser menos invasivo ao código original,mudei de ideia, vamos de diretorios mesmo



#Situa uma main que recebe como argumento com um diretorio, e retorna o wc_dir desse direotrio executando todos os métodos acima.
def main():
    global count
    if len(sys.argv) != 2:
        print("Usage: python", sys.argv[0], "root_directory_path")
        return
    root_path = os.path.abspath(sys.argv[1])
    threads = [Thread(target=wc_dir,args=(diretorio,)) for diretorio in pick_dir(root_path)]
    for thread in threads:
        thread.start()
    for thread in threads:
        thread.join()
    print(count)

if __name__ == "__main__":
    main()
