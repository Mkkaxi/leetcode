node_modules 不应该跟踪，太大
npm i  会修改package.json


ls  查看当前目录
ls -al  
cd 进入目录
pwd 当前文件绝对路径
du -hs node_modules/  显示当前文件目录下文件大小
cat 文件  输出文件


git add .
git commit -m ''
git push origin master

git pull origin master
yarn  安装依赖

1. 删除node.modules 
  git rm  
2. .gitignore
3. 提交删除
