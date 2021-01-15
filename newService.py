# -*- coding: UTF-8 -*-
"""
@author: xxsCoder
@contact: xxscoder@gmail.com
@file: newService.py
@time: 2021/1/13 10:25
"""
import os
import winreg

startText = '\t本工具用于生成go micro的项目基本结构，项目生成后只需要编写自定义proto文件,并通过protoc来生成go还有micro文件\n' \
            '项目的数据库文件需要在domain/model里面定义，本项目需要的gorm来操作数据库，具体的数据库模型tag定义需要参考gorm模型定义\n' \
            '具体的服务器与客户端的业务逻辑调用需要在handler中编写，最后在main中注册即可，启动本项目需要安装并启动redis,mysql,etcd,\njaeger\n\n\n'


def showTree(path, depth):
    if depth == 0:
        print("项目创建目录:[" + path + "]")

    for item in os.listdir(path):
        if '.git' not in item:
            print("|      " * depth + "+--" + item)
            newItem = path + '/' + item
            if os.path.isdir(newItem):
                showTree(newItem, depth + 1)


class NewMicro(object):
    def __init__(self, path: str, microName: str, newPort: str):
        self.path: str = path
        self.newMicroName: str = microName
        self.newPort = newPort
        self.pathList: list = []
        self.fileNameList: list = []
        self.newPathList: list = []
        self.ReadList: list = []
        self.winDesName: str = ""
        self.GetWinDesPath()
        self.newProjectName = os.path.join(self.winDesName, self.newMicroName)
        try:
            os.mkdir(self.newProjectName)
        except FileExistsError:
            print("项目名称已经存在或者项目已创建，无需重复创建")

    def CreatePath(self):
        for p in self.newPathList:
            isExists = os.path.exists(p)
            if not isExists:
                # 如果不存在，则创建目录（多层）
                if "Xxxx" in p:
                    try:
                        newPathName = p.replace("Xxxx", self.newMicroName, 5)
                        os.makedirs(newPathName)
                    except FileExistsError:
                        pass
                else:
                    os.makedirs(p)

    def GetWinDesPath(self):
        key = winreg.OpenKey(winreg.HKEY_CURRENT_USER,
                             r'Software\Microsoft\Windows\CurrentVersion\Explorer\Shell Folders')
        self.winDesName = str(winreg.QueryValueEx(key, "Desktop")[0])

    def GetFilePath(self):
        for path, dirList, files in os.walk(self.path):
            for i in files:
                filePath = os.path.join(path, i)
                if ".idea" in filePath:
                    continue
                if ".git" in filePath:
                    continue
                if ".py" in filePath:
                    continue
                self.fileNameList.append(filePath)
            pathList = path.split(r"microServerTemplate")
            if ".idea" in pathList[1]:
                continue
            if ".git" in pathList[1]:
                continue
            if ".py" in pathList[1]:
                continue
            self.pathList.append(pathList[1])
            for p in self.pathList:
                newPath = self.newProjectName + p
                if p == "":
                    continue
                self.newPathList.append(newPath)
            # print("path List:", self.pathList)
            # print("file list:", self.fileNameList)
            # print("newPath:", self.newPathList)

    def ReadOldFile(self):
        for pathName in self.fileNameList:
            print(self.newMicroName)
            path = str(pathName).replace("microServerTemplate", self.newMicroName)
            if "Xxxx" in path:
                newPath = path.replace("Xxxx", self.newMicroName)
                self.OpenFileRead(pathName)
                self.OpenFileWrite(newPath)
                self.ReadList.clear()
            else:
                print(path)
                newPath = path.replace("Xxxx", self.newMicroName)
                self.OpenFileRead(pathName)
                self.OpenFileWrite(newPath)
                self.ReadList.clear()

    def OpenFileRead(self, pathName: str):
        with open(pathName, mode="r", encoding="UTF-8") as rf:
            while True:
                read = rf.readline()
                if "Xxxx" in read:
                    newRead = read.replace("Xxxx", self.newMicroName)
                    self.ReadList.append(newRead)
                    if not read:
                        break
                elif "port" in read:
                    port = read.replace("8772", self.newPort)
                    self.ReadList.append(port)
                    if not read:
                        break
                elif "microServerTemplate" in read:
                    modelName = read.replace("microServerTemplate", self.newMicroName)
                    self.ReadList.append(modelName)
                    if not read:
                        break
                else:
                    self.ReadList.append(read)
                    if not read:
                        break

    def OpenFileWrite(self, pathName: str):
        with open(pathName, mode="w", encoding="UTF-8") as wf:
            for i in self.ReadList:
                wf.write(i)

    def __del__(self):
        print("====================")
        print("=====项目创建完成=====")
        print("端口为：", self.newPort)
        print('\t')
        print('\t')
        print("=======项目的结构=====")
        showTree(self.newProjectName, 0)


if __name__ == '__main__':
    print("=" * 100)
    print("欢迎使用go-micro服务端项目基本结构生成工具".center(50))
    print("=" * 100)
    print(startText)
    newProjectName = input("输入新的项目名称:").capitalize()
    newPort = input("输入新的端口(默认端口：8772)：")
    if newPort == "":
        newPort = "8772"
    absPathName = os.path.abspath('.')
    newService = NewMicro(absPathName, newProjectName, newPort)
    newService.GetFilePath()
    newService.CreatePath()
    newService.ReadOldFile()
